/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeAssignReqSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"con_info": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceSeResourceFindReqSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"curr_se_info": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCurSeInfoSchema(),
			},
			"host_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"new_se": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"new_vcpus": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
