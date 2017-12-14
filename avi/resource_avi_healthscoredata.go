/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHealthScoreDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"anomaly_penalty": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"anomaly_score": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHealthScoreAnomalyDataSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"dos_attack_level": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
				Default:  "0.0",
			},
			"is_null": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"num_samples": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"performance_score": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHealthScorePerformanceDataSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"performance_value": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"reason_attr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"resources_penalty": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"resources_score": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHealthScoreResourcesDataSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"security_penalty": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"security_threat_level": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHealthScoreSecurityDataSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ssl_score": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
				Default:  "5.0",
			},
			"timestamp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeFloat,
				Required: true,
			},
		},
	}
}
