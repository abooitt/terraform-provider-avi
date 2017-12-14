/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourcePortRangeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"end": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"start": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
		},
	}
}
