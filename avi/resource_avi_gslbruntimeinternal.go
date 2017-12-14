/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceGslbRuntimeInternalSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbRuntimeSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
