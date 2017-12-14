/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceClusterNodeDbFailedEventSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"failure_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceIpAddrSchema()},
			"node_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
