/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceLcEntrySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ip_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"max_queue_length": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"num_open_connections": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
		},
	}
}
