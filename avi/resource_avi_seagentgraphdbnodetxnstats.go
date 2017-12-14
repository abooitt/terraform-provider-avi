/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeAgentGraphDBNodeTxnStatsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"history": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeAgentGraphDBNodeTxnDetailSchema()},
			"longest_txn": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeAgentGraphDBNodeTxnDetailSchema()},
		},
	}
}
