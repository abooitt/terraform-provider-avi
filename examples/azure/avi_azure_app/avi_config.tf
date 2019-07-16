provider "azurerm" {
  subscription_id = var.subscription_id
  client_id       = var.client_id
  client_secret   = var.client_secret
  tenant_id       = var.tenant_id
}

data "azurerm_network_interface" "controller_nic" {
  name                = "${var.project_name}-terraform-network-interface"
  resource_group_name = "${var.project_name}-terraform-resource-group"
  //resource_group_name = "${var.resource_group_name}"
}

data "azurerm_subnet" "terraform_subnet" {
  name                 = "${var.project_name}-terraform-subnet"
  virtual_network_name = "${var.project_name}-terraform-virtual-network"
  resource_group_name  = "${var.project_name}-terraform-resource-group"
}

provider "avi" {
  avi_username   = var.avi_username
  avi_password   = var.avi_password
  avi_controller = data.azurerm_network_interface.controller_nic.private_ip_address
  avi_tenant     = "admin"
  avi_version    = var.avi_version
}

data "avi_tenant" "default_tenant" {
  name = "admin"
}

data "avi_cloud" "azure_cloud_cfg" {
  name = "AZURE"
}

data "avi_vrfcontext" "terraform_vrf" {
  name      = "global"
  cloud_ref = data.avi_cloud.azure_cloud_cfg.id
}

data "avi_healthmonitor" "system_http_healthmonitor" {
  name = "System-HTTP"
}

resource "avi_pool" "azure-pool-v1" {
  name                = "azure_poolv1"
  health_monitor_refs = [data.avi_healthmonitor.system_http_healthmonitor.id]
  tenant_ref          = data.avi_tenant.default_tenant.id
  vrf_ref             = data.avi_vrfcontext.terraform_vrf.id
  cloud_ref           = data.avi_cloud.azure_cloud_cfg.id

  //external_autoscale_groups = ["${azurerm_virtual_machine_scale_set.terraform_scale_set_v1.name}@${var.resource_group_name}"]
  external_autoscale_groups = ["${azurerm_virtual_machine_scale_set.terraform_scale_set_v1.name}@${var.project_name}-terraform-resource-group"]
  fail_action {
    type = "FAIL_ACTION_CLOSE_CONN"
  }
}

resource "avi_pool" "azure-pool-v2" {
  name                = "azure_poolv2"
  health_monitor_refs = [data.avi_healthmonitor.system_http_healthmonitor.id]
  tenant_ref          = data.avi_tenant.default_tenant.id
  vrf_ref             = data.avi_vrfcontext.terraform_vrf.id
  cloud_ref           = data.avi_cloud.azure_cloud_cfg.id

  //external_autoscale_groups = ["${azurerm_virtual_machine_scale_set.terraform_scale_set_v1.name}@${var.resource_group_name}"]
  external_autoscale_groups = ["${azurerm_virtual_machine_scale_set.terraform_scale_set_v2.name}@${var.project_name}-terraform-resource-group"]
  fail_action {
    type = "FAIL_ACTION_CLOSE_CONN"
  }
}

resource "azurerm_virtual_machine_scale_set" "terraform_scale_set_v1" {
  name                = "${var.project_name}_scale_set_v1"
  location            = var.location
  resource_group_name = "${var.project_name}-terraform-resource-group"

  //resource_group_name = "${var.resource_group_name}"
  upgrade_policy_mode = "Manual"

  sku {
    name     = "Standard_A0"
    tier     = "Standard"
    capacity = 1
  }

  storage_profile_image_reference {
    publisher = "Canonical"
    offer     = "UbuntuServer"
    sku       = "16.04-LTS"
    version   = "latest"
  }

  storage_profile_os_disk {
    name              = ""
    caching           = "ReadWrite"
    create_option     = "FromImage"
    managed_disk_type = "Standard_LRS"
  }

  storage_profile_data_disk {
    lun           = 0
    caching       = "ReadWrite"
    create_option = "Empty"
    disk_size_gb  = 10
  }

  os_profile {
    computer_name_prefix = "${var.project_name}testvm"
    admin_username       = "myadmin"
    admin_password       = var.azure_vm_password
  }

  network_profile {
    name    = "${var.project_name}terraformnetworkprofile"
    primary = true
    ip_configuration {
      name      = "${var.project_name}TestIPConfiguration"
      subnet_id = data.azurerm_subnet.terraform_subnet.id
    }
  }

  os_profile_linux_config {
    disable_password_authentication = false
    ssh_keys {
      path     = "/home/myadmin/.ssh/authorized_keys"
      key_data = file("~/.ssh/id_rsa.pub")
    }
  }

  extension {
    name                 = "vmssextension"
    publisher            = "Microsoft.OSTCExtensions"
    type                 = "CustomScriptForLinux"
    type_handler_version = "1.2"
    settings             = <<SETTINGS
    {
        "commandToExecute": "apt-get -y install nginx"
    }
    
SETTINGS

  }

  tags = {
    environment = "${var.project_name}-terraform-${var.project_environment}"
  }
}

output "azure_scale_set1" {
  value = azurerm_virtual_machine_scale_set.terraform_scale_set_v1.name
}

resource "azurerm_virtual_machine_scale_set" "terraform_scale_set_v2" {
  name = "${var.project_name}_scale_set_v2"
  location = var.location
  resource_group_name = "${var.project_name}-terraform-resource-group"

  //resource_group_name = "${var.resource_group_name}"
  upgrade_policy_mode = "Manual"

  sku {
    name = "Standard_A0"
    tier = "Standard"
    capacity = 1
  }

  storage_profile_image_reference {
    publisher = "Canonical"
    offer = "UbuntuServer"
    sku = "16.04-LTS"
    version = "latest"
  }

  storage_profile_os_disk {
    name = ""
    caching = "ReadWrite"
    create_option = "FromImage"
    managed_disk_type = "Standard_LRS"
  }

  storage_profile_data_disk {
    lun = 0
    caching = "ReadWrite"
    create_option = "Empty"
    disk_size_gb = 10
  }

  os_profile {
    computer_name_prefix = "${var.project_name}testvm"
    admin_username = "myadmin"
    admin_password = var.azure_vm_password
  }

  network_profile {
    name = "${var.project_name}terraformnetworkprofile"
    primary = true
    ip_configuration {
      name = "${var.project_name}TestIPConfiguration"
      subnet_id = data.azurerm_subnet.terraform_subnet.id
    }
  }

  os_profile_linux_config {
    disable_password_authentication = false
    ssh_keys {
      path = "/home/myadmin/.ssh/authorized_keys"
      key_data = file("~/.ssh/id_rsa.pub")
    }
  }

  extension {
    name = "vmssextension"
    publisher = "Microsoft.OSTCExtensions"
    type = "CustomScriptForLinux"
    type_handler_version = "1.2"
    settings = <<SETTINGS
    {
        "commandToExecute": "apt-get -y install nginx"
    }
    
SETTINGS

}

tags = {
environment = "${var.project_name}-terraform-${var.project_environment}"
}
}

