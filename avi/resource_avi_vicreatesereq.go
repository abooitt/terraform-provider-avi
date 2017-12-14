/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVICreateSEReqSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_grp_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_params": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceVISeVmOvaParamsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"tenant_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_admin": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVIAdminCredentialsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vcenter_vnic_portgroups": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}
