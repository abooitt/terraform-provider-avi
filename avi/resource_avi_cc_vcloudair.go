/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceCC_VCloudAirSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"access_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cfg": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcevCloudAirConfigurationSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"mgmt_nw_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
