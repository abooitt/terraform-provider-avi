/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSdbEntrySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"method": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"sdb_key_s": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"sdb_val_s": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
