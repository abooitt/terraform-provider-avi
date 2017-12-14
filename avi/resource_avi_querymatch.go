/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceQueryMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_case": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "INSENSITIVE"},
			"match_criteria": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"match_str": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"string_group_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
		},
	}
}
