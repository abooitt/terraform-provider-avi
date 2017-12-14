/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceDnsRequestSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"additional_records_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"answer_records_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"authentic_data": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"checking_disabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"client_location": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGeoLocationSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"identifier": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"nameserver_records_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"opcode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"opt_record": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDnsOptRecordSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"query_or_response": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"question_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"recursion_desired": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}
