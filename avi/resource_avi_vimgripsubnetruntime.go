/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVIMgrIPSubnetRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"fip_available": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"fip_subnet_uuids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"floatingip_subnets": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceFloatingIpSubnetSchema(),
			},
			"ip_subnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"prefix": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceIpAddrPrefixSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"primary": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ref_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"se_ref_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
