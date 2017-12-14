/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceConnectionItemSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"connection": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceConnectionEntrySchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"reuse_cnt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
