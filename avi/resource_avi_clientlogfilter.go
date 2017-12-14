/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceClientLogFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"all_headers": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"client_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"duration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30,
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"uri": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceStringMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
