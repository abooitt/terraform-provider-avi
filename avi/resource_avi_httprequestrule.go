/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHTTPRequestRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"all_headers": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enable": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"hdr_action": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHTTPHdrActionSchema(),
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"log": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"match": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMatchTargetSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"redirect_action": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPRedirectActionSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"rewrite_url_action": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPRewriteURLActionSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"switching_action": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPSwitchingActionSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
