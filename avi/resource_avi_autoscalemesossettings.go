/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceAutoScaleMesosSettingsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"force": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}
