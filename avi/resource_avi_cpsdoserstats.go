/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceCpsDoserStatsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cps_doser_entry": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCpsDoserEntrySchema()},
			"num_cps_dosers": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
