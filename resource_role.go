package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func roleResourceServer() *schema.Resource {
	return &schema.Resource{
		Create: roleResourceServerCreate,
		Read:   roleResourceServerRead,
		Update: roleResourceServerUpdate,
		Delete: roleResourceServerDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db": {
				Type:     schema.TypeString,
				Required: true,
			},
			"role":      roleRefSet(),
			"privilege": privilegeSet(),
		},
	}
}

func roleResourceServerCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	if err := client.CreateRole(roleInfo(d)); err != nil {
		return err
	}

	d.SetId(d.Get("name").(string))

	return nil
}

func roleResourceServerRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	obj, err := client.GetRole(d.Get("db").(string), d.Id())
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

	if err := d.Set("privilege", flattenPrivileges(obj.Privileges)); err != nil {
		return err
	}

	return nil
}

func roleResourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	if err := client.UpdateRole(roleInfo(d)); err != nil {
		return err
	}

	return nil
}

func roleResourceServerDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	if err := client.DeleteRole(roleInfo(d)); err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func roleInfo(d *schema.ResourceData) Role {
	return Role{
		Role:       d.Get("name").(string),
		Db:         d.Get("db").(string),
		Roles:      expandRoleRefs(d.Get("role").(*schema.Set)),
		Privileges: expandPrivileges(d.Get("privilege").(*schema.Set)),
	}
}
