/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHostAttributesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attr_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"attr_val": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
