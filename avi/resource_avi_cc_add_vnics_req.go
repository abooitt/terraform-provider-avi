/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Resourcecc_add_vnics_reqSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "cloud-0"},
			"cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_vm_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"vnics": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCC_VnicInfoSchema()},
			"vs_uuids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
		},
	}
}
