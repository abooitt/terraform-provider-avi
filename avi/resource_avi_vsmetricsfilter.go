/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVsMetricsFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"obj_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
