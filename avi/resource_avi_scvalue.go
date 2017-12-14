/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSCValueSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			//"pool_state": &schema.Schema{
			//	Type:     schema.TypeSet,
			//	Optional: true,
			//	//Elem:     ResourcePoolRuntimeDetailSchema(),
			//	Set: func(v interface{}) int {
			//		return 0
			//	},
			//},
			//"server_state": &schema.Schema{
			//	Type:     schema.TypeList,
			//	Optional: true,
			//	//Elem:     ResourceServerRuntimeSummarySchema(),
			//},
			"vs_state": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVirtualServiceRuntimeDetailSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
