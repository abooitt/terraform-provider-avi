/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceFloatingIpSubnetSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"prefix": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceIpAddrPrefixSchema()},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true},
		},
	}
}
