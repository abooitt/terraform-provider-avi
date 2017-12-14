/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHealthMonitorHttpSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"exact_http_request": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"http_request": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "GET / HTTP/1.0",
			},
			"http_response": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_response_code": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"maintenance_code": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"maintenance_response": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_attributes": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHealthMonitorSSLAttributesSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
