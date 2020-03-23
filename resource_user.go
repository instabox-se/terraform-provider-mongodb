package main

import (
	"crypto/sha512"
	"encoding/base64"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func userResourceServer() *schema.Resource {
	return &schema.Resource{
		Create: userResourceServerCreate,
		Read:   userResourceServerRead,
		Update: userResourceServerUpdate,
		Delete: userResourceServerDelete,

		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
				StateFunc: func(val interface{}) string {
					hasher := sha512.New()
					hasher.Write([]byte(val.(string)))
					return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
				},
			},
			"allow_password_update": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"db": {
				Type:     schema.TypeString,
				Required: true,
			},
			"role": roleRefSet(),
		},
	}
}

func userResourceServerCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	err := client.CreateUser(userInfo(d))

	if err != nil {
		return err
	}

	d.SetId(d.Get("username").(string))

	return nil
}

func userResourceServerRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	database := d.Get("db").(string)

	obj, err := client.GetUser(database, d.Id())
	if err != nil {
		return err
	}

	if obj == nil {
		d.SetId("")
		return nil
	}

	if err := d.Set("role", flattenRoleRefs(obj.Roles)); err != nil {
		return err
	}

	return nil
}

func userResourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	if err := client.UpdateUser(userInfo(d)); err != nil {
		return err
	}

	return nil
}

func userResourceServerDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	if err := client.DeleteUser(userInfo(d)); err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func userInfo(d *schema.ResourceData) User {
	return User{
		Username:            d.Get("username").(string),
		Name:                d.Get("name").(string),
		Password:            d.Get("password").(string),
		AllowPasswordUpdate: d.Get("allow_password_update").(bool),
		Db:                  d.Get("db").(string),
		Roles:               expandRoleRefs(d.Get("role").(*schema.Set)),
	}
}
