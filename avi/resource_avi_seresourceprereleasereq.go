/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeResourcePreReleaseReqSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"admin_down": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"curr_primary_se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"migrate_primary": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"new_primary_se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"notify_failure": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"num_se": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"num_stby_se": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"rel_se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"res_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsResProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_alloc_info": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeAllocInfoSchema(),
			},
			"svc_id": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceSeMgrSvcIdSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"use_res_profile": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}
