/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHTTPResponseRuleSchema() *schema.Resource {
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
			"loc_hdr_action": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPRewriteLocHdrActionSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"log": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"match": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceResponseMatchTargetSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
