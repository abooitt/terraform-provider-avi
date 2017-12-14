/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVirtualServiceDbExtensionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"apic_extension": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsApicExtensionSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ipam_dns_records": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsRecordSchema(),
			},
			"lif": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"redis_db": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"redis_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"redis_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tls_ticket_key": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceTLSTicketSchema(),
			},
			"version": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"vh_child_vs_uuid": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vip_runtime": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVipDbExtensionSchema(),
			},
		},
	}
}
