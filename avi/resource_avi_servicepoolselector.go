/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceServicePoolSelectorSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"service_pool_group_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_pool_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"service_port_range_end": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"service_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
