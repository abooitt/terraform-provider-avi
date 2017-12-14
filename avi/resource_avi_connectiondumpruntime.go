/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceConnectionDumpRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"connection": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceConnectionEntrySchema()},
			"proc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
