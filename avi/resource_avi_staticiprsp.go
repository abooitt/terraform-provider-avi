/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceStaticIPRspSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip_alloc": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceStaticIPInfoSchema()},
			"rsp_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceStaticIPRspStatusSchema()},
		},
	}
}
