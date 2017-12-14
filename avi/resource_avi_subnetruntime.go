/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSubnetRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"free_ip_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ip_alloced": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAllocInfoSchema(),
			},
			"prefix": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceIpAddrPrefixSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"total_ip_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"used_ip_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
