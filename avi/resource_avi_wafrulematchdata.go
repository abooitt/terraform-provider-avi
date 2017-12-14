/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceWafRuleMatchDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"is_internal": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"match_element": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"match_value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
