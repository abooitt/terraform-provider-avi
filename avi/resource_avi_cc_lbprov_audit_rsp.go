/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Resourcecc_lbprov_audit_rspSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"data": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOsLbProvAuditDetailSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ret_status": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ret_string": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
