/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSdbValueSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip_persistent_value": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpPersistentValueSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"prst_srv_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ssl_value": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSslValueSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"string_val": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}
