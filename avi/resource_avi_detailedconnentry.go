/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceDetailedConnEntrySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"f_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"f_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"inpcb_fast_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceInpcbFastInfoSchema()},
			"inpcb_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceInpcbInfoSchema()},
			"inpcb_proxy_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceInpcbProxyInfoSchema()},
			"l_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"l_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"socket_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSocketInfoSchema()},
			"tcpcb_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceTcpcbInfoSchema()},
			"tcptw_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceTcptwInfoSchema()},
		},
	}
}
