/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceCustomTagSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"tag_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"tag_val": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
