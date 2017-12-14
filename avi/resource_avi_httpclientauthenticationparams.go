/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHTTPClientAuthenticationParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"auth_profile_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"realm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"request_uri_path": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceStringMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
