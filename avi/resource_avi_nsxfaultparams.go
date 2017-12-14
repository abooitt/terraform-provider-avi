/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceNsxFaultParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"count": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"delay": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"fault_api": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
