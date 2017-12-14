/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Resourcecc_configure_reqSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_config": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceCloudSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cc_old_config": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
