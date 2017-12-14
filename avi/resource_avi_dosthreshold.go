/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceDosThresholdSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attack": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"max_value": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"min_value": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}
