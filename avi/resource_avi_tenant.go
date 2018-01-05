/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strings"
)

func ResourceTenantSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"config_settings": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceTenantConfigurationSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"created_by": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"local": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviTenant() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviTenantCreate,
		Read:   ResourceAviTenantRead,
		Update: resourceAviTenantUpdate,
		Delete: resourceAviTenantDelete,
		Schema: ResourceTenantSchema(),
	}
}

func ResourceAviTenantRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTenantSchema()
	err := ApiRead(d, meta, "tenant", s)
	log.Printf("[DEBUG] data read as %v uuid %v id %v\n", d.Get("name"), d.Get("uuid"), d.Id())
	return err
}

func resourceAviTenantCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTenantSchema()
	err := ApiCreateOrUpdate(d, meta, "tenant", s)
	if err == nil {
		err = ResourceAviTenantRead(d, meta)
	}
	return err
}

func resourceAviTenantUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTenantSchema()
	err := ApiCreateOrUpdate(d, meta, "tenant", s)
	if err == nil {
		err = ResourceAviTenantRead(d, meta)
	}
	return err
}

func resourceAviTenantDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "tenant"
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviTenantDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
