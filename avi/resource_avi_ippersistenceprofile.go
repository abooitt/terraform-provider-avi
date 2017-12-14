/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceIPPersistenceProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip_persistent_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
		},
	}
}
