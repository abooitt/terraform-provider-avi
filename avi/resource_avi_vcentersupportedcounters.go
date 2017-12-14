/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVcenterSupportedCountersSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"counters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVmwarePerfCounterInfoSchema()},
			"vcenter": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
