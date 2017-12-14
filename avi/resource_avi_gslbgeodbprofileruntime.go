/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceGslbGeoDbProfileRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"checksum": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"flr_state": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCfgStateSchema(),
			},
			"ldr_state": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCfgStateSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
