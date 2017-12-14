/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceIpAddrPrefixSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip_addr": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Set: func(v interface{}) int { return 0 }, Elem: ResourceIpAddrSchema()},
			"mask": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
		},
	}
}
