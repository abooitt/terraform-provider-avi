/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourcePoolResourcesScoreDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"avg_server_resources_score": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_uptime": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pool_ref": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"reason_attr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_resources_scores": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceServerResourcesScoreSchema(),
			},
		},
	}
}
