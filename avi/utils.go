/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"log"
	"reflect"
	"strings"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/avinetworks/sdk/go/session"
	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
)

func SchemaToAviData(d interface{}, s map[string]*schema.Schema) (interface{}, error) {
	switch d.(type) {
	default:
		log.Printf("[INFO] SchemaToAviData: resource d: %v(%v)", d, reflect.TypeOf(d))
		//return d, nil
	case map[string]interface{}:
		m := make(map[string]interface{})
		for k, v := range d.(map[string]interface{}) {
			if obj, err := SchemaToAviData(v, nil); err == nil && obj != nil && obj != "" {
				m[k] = obj
			} else if err != nil {
				log.Printf("[ERROR] SchemaToAviData %v in parsing k: %v v: %v", err, k, v)
			}
		}
		return m, nil
	case []interface{}:
		var objList []interface{}
		varray := d.([]interface{})
		for i := 0; i < len(varray); i++ {
			obj, err := SchemaToAviData(varray[i], nil)
			if err == nil && obj != nil {
				objList = append(objList, obj)
			}
		}
		if len(objList) == 0 {
			return nil, nil
		}
		return objList, nil

	case *schema.Set:
		if len(d.(*schema.Set).List()) == 0 {
			return nil, nil
		}
		obj, err := SchemaToAviData(d.(*schema.Set).List()[0], nil)
		return obj, err

	case *schema.ResourceData:
		// In this case the top level schema should be present.
		m := make(map[string]interface{})
		r := d.(*schema.ResourceData)
		for k, v := range s {
			if obj, err := SchemaToAviData(r.Get(k), nil); err == nil && obj != nil && obj != "" {
				m[k] = obj
			} else if err != nil {
				log.Printf("[ERROR] SchemaToAviData %v in converting k: %v v: %v", err, k, v)
			}
		}
		return m, nil
	}
	/** Return the same object as there is nothing special about **/
	return d, nil
}

func CommonHash(v interface{}) int {
	return hashcode.String("avi")
}

func ApiDataToSchema(adata interface{}, d *schema.ResourceData, t map[string]*schema.Schema) (interface{}, error) {
	switch adata.(type) {
	default:
	case map[string]interface{}:
		// resolve d interface into a set
		if t == nil {
			var m map[string]interface{}
			m = map[string]interface{}{}
			for k, v := range adata.(map[string]interface{}) {
				if obj, err := ApiDataToSchema(v, nil, nil); err == nil {
					m[k] = obj
				} else if err != nil {
					log.Printf("[ERROR] ApiDataToSchema %v in converting k: %v v: %v", err, k, v)
				}

			}
			//var s schema.Set
			objs := []interface{}{}
			objs = append(objs, m)
			s := schema.NewSet(CommonHash, objs)
			//s.Add(m)
			return s, nil
		} else {
			for k, v := range adata.(map[string]interface{}) {
				if _, ok := t[k]; ok {
					// found in the schema
					if obj, err := ApiDataToSchema(v, nil, nil); err == nil {
						err := d.Set(k, obj)
						if err != nil {
							log.Printf("[ERROR] ApiDataToSchema %v in setting %v", err, obj)
						}
					}
				}
			}
			return d, nil
		}
	case []interface{}:
		var objList []interface{}
		varray := adata.([]interface{})
		for i := 0; i < len(varray); i++ {
			obj, err := ApiDataToSchema(varray[i], nil, nil)
			if err == nil {
				switch obj.(type) {
				default:
					objList = append(objList, obj)
				case *schema.Set:
					objList = append(objList, obj.(*schema.Set).List()[0])
				}
			} else {
				log.Printf("[ERROR] ApiDataToSchema %v", err)
			}
		}
		return objList, nil
		/** Return the same object as there is nothing special about **/
	}
	return adata, nil
}

func ApiCreateOrUpdate(d *schema.ResourceData, meta interface{}, objType string, s map[string]*schema.Schema) error {
	var err error
	client := meta.(*clients.AviClient)
	var robj interface{}
	obj := d

	if data, err := SchemaToAviData(obj, s); err == nil {
		path := "api/" + objType
		if uuid, ok := d.GetOk("uuid"); ok {
			path = path + "/" + uuid.(string)
			err = client.AviSession.Put(path, data, &robj)
		} else if name, ok := d.GetOk("name"); ok {
			var existing_obj interface{}
			if cloudRef, ok := d.GetOk("cloud_ref"); ok && strings.Contains(cloudRef.(string), "api/cloud/") {
				cloudUUID := strings.SplitN(cloudRef.(string), "api/cloud/", 2)[1]
				log.Printf("[INFO] ApiCreateOrUpdate: using cloud %v \n", cloudUUID)

				err = client.AviSession.GetObject(objType, session.SetName(name.(string)),
					session.SetResult(&existing_obj), session.SetCloudUUID(cloudUUID))
			} else {
				err = client.AviSession.GetObject(objType, session.SetName(name.(string)),
					session.SetResult(&existing_obj))
			}
			if err != nil {
				// object not found
				log.Printf("[INFO] ApiCreateOrUpdate: Creating obj type %v schema %v data %v\n", objType, d, data)

				err = client.AviSession.Post(path, data, &robj)
				if err != nil {
					log.Printf("[ERROR] ApiCreateOrUpdate creation failed %v object with name %v\n", err, name)
				}
			} else {
				// found existing object.
				uuid = existing_obj.(map[string]interface{})["uuid"].(string)
				d.Set("uuid", uuid)
				d.SetId(uuid.(string))
				path = path + "/" + uuid.(string)
				err = client.AviSession.Put(path, data, &robj)
			}
		}
		if err != nil {
			d.SetId("")
			log.Printf("[ERROR] ApiCreateOrUpdate: Error %v path %v id %v\n", err, path, d.Id())
			return err
		}
		log.Printf("[DEBUG] ApiCreateOrUpdate: object %v\n", robj)
		url := robj.(map[string]interface{})["url"].(string)
		uuid := robj.(map[string]interface{})["uuid"].(string)
		url = strings.SplitN(url, "#", 2)[0]
		d.SetId(url)
		d.Set("uuid", uuid)
	} else {
		log.Printf("[ERROR] ApiCreateOrUpdate: Error %v", err)
	}
	return err
}

func ApiRead(d *schema.ResourceData, meta interface{}, objType string, s map[string]*schema.Schema) error {
	client := meta.(*clients.AviClient)
	var obj interface{}
	uuid := ""
	log.Printf("[DEBUG] ApiRead reading object with objType %v id %v\n",
		objType, d.Id())

	if d.Id() != "" {
		// extract the uuid from it.
		log.Printf("[DEBUG] ApiRead reading object with objType %v id %v \n", objType, d.Id())
		url_parts := strings.Split(d.Id(), "/")
		uuid = url_parts[len(url_parts)-1]
	} else if u, ok := d.GetOk("uuid"); ok {
		uuid = u.(string)
		log.Printf("[DEBUG] ApiRead reading object with uuid %v \n", uuid)
	}
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		log.Printf("[DEBUG] ApiRead reading object with id %v path %v\n", uuid, path)
		err := client.AviSession.Get(path, &obj)
		if err != nil {
			d.SetId("")
			log.Printf("[ERROR] ApiRead object with uuid %v not found err %v\n", uuid, err)
			return nil
		}
	} else if name, ok := d.GetOk("name"); ok {
		var err error
		if cloudRef, ok := d.GetOk("cloud_ref"); ok && strings.Contains(cloudRef.(string), "api/cloud/") {
			cloudUUID := strings.SplitN(cloudRef.(string), "api/cloud/", 2)[1]
			log.Printf("[DEBUG] ApiRead using cloud %v \n", cloudUUID)
			err = client.AviSession.GetObject(objType, session.SetName(name.(string)),
				session.SetResult(&obj), session.SetCloudUUID(cloudUUID))
		} else {
			log.Printf("[DEBUG] ApiRead using name %v \n", name)
			err = client.AviSession.GetObject(objType, session.SetName(name.(string)),
				session.SetResult(&obj))
		}
		if err != nil {
			d.SetId("")
			log.Printf("[ERROR] ApiRead object with name %v:%v not found err %v\n", objType, name, err)
			return nil
		}
	} else {
		d.SetId("")
		log.Printf("[ERROR] ApiRead not found %v\n", d.Get("uuid"))
		return nil
	}
	if _, err := ApiDataToSchema(obj, d, s); err == nil {
		url := obj.(map[string]interface{})["url"].(string)
		uuid := obj.(map[string]interface{})["uuid"].(string)
		url = strings.SplitN(url, "#", 2)[0]
		d.SetId(url)
		d.Set("uuid", uuid)
	} else {
		log.Printf("[ERROR] ApiRead in setting read object %v\n", err)
	}
	return nil
}

func ResourceImporter(d *schema.ResourceData, meta interface{}, objType string, s map[string]*schema.Schema) ([]*schema.ResourceData, error) {
	if d.Id() != "" {
		// return the ID based import
		return []*schema.ResourceData{d}, nil
	}
	var data interface{}
	client := meta.(*clients.AviClient)
	path := "api/" + objType + "/"
	log.Printf("[DEBUG] ResourceImporter reading object with path %v\n", path)

	err := client.AviSession.Get(path, &data)
	if err != nil {
		log.Printf("[ERROR] ResourceImporter %v in GET of path %v\n", err, path)
		return nil, err
	}
	count := int((data.(map[string]interface{})["count"]).(float64))
	log.Printf("[DEBUG] ResourceImporter read data with path %v -> count %v\n", path, count)
	if count > 0 {
		results := make([]*schema.ResourceData, count)
		apiResults := (data.(map[string]interface{})["results"]).([]interface{})
		for index := 0; index < count; index++ {
			obj := apiResults[index].(map[string]interface{})
			log.Printf("[DEBUG] ResourceImporter processing obj %v results %v\n", obj, results[index])
			result := new(schema.ResourceData)
			if _, err := ApiDataToSchema(obj, result, s); err == nil {
				log.Printf("[DEBUG] ResourceImporter Processing obj %v\n", obj)
				url := obj["url"].(string)
				uuid := obj["uuid"].(string)
				url = strings.SplitN(url, "#", 2)[0]
				result.SetId(url)
				result.Set("uuid", uuid)
				result.SetType("avi_" + objType)
				results[index] = result
			}
		}
		return results, nil
	}
	return nil, nil
}

func ApiDeleteSystemDefaultCheck(d *schema.ResourceData) bool {
	var systemDefault bool
	var sysName string
	if sysdef, ok := d.GetOk("system_default"); ok {
		systemDefault = sysdef.(bool)
	}
	if name, ok := d.GetOk("name"); ok {
		sysName = name.(string)
	}
	if systemDefault || strings.HasPrefix(sysName, "System-") || sysName == "Default-Group" {
		d.SetId("")
		return true
	}
	return false
}
