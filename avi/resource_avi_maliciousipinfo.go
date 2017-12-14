/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceMaliciousIpInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"count": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"dos_dimension": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DOS_IP",
			},
			"dst_ip": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"meta_data": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAttackMetaDataSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"src_ip": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
