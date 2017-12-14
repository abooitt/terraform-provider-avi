/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeResourceVnicSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"connected": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"del_retries": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"in_use": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_avi_internal_network": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"last_vnic_op_ticks": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"lif": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"linux_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mac_addr": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"marked_for_del": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"port_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vips": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"virtual_network_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vnic_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "RM_VNIC_FRONTEND",
			},
			"vrf_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
