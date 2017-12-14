/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceCloudInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cca_props": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCC_AgentPropertiesSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"controller_props": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceControllerPropertiesSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"flavor_props": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCloudFlavorSchema(),
			},
			"flavor_regex_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"htypes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vtype": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
