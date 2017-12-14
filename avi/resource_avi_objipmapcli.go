/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceObjIpMapCliSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"obj_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
