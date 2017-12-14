/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSamlSettingsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"idp": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSamlIdentityProviderSettingsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"sp": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSamlServiceProviderSettingsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
