/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceRateLimiterActionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"file": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPLocalFileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"redirect": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPRedirectActionSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"status_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "HTTP_LOCAL_RESPONSE_STATUS_CODE_429",
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "RL_ACTION_NONE",
			},
		},
	}
}
