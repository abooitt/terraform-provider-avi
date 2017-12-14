/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHSMThalesRFSSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  9004,
			},
		},
	}
}
