/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourcensxVshieldAppConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"excludelistconfiguration": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourcensxExcludeListConfigurationSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
