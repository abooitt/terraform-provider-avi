/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceGslbPerDnsStateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"geo_download": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbDownloadStatusSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"gslb_download": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbDownloadStatusSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"oper_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOperationalStatusSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"placement_rules": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbSubDomainPlacementRuntimeSchema(),
			},
			"se_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"valid_dns_vs": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}
