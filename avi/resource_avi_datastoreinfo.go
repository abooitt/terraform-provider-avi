/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceDatastoreInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host_names": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"managed_object_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
