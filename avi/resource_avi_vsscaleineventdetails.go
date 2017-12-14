/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVsScaleInEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error_message": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rpc_status": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"scale_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceScaleStatusSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_assigned": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVipSeAssignedSchema(),
			},
			"se_requested": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVirtualServiceResourceSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
