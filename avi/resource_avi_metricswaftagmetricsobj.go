/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceMetricsWafTagMetricsObjSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"avg_flagged": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_matched": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_rejected": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_flagged": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_matched": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_rejected": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"node_obj_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"sum_flagged": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_matched": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_rejected": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
		},
	}
}
