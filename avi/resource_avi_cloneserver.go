/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceCloneServerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip_address": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_ref": &schema.Schema{
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
		},
	}
}
