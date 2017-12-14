/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceRPCRequestMetaDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"marker": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"page_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  25,
			},
		},
	}
}
