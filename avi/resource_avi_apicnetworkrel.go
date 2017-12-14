/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceAPICNetworkRelSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"connector": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rel_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_network": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
