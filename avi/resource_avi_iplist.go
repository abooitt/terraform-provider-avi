/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceIpListSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ips": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema()},
			"obj_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Unknown"},
			"obj_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Unknown"},
			"only_conn_subnets": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeBool}},
			"pool_update": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourcePoolUpdateSchema()},
			"vrf_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_update": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVsUpdateSchema()},
		},
	}
}
