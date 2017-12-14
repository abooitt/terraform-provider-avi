/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceRetrieveVcenterDatastoreRspSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ds_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVIDatastoreSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
