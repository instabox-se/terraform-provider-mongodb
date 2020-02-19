package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"connection_string": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A MongoDB connection string",
			},
		},
		ConfigureFunc: providerConfigure,
		ResourcesMap: map[string]*schema.Resource{
			"mongodb_user": userResourceServer(),
			"mongodb_role": roleResourceServer(),
		},
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	connectionString, _ := d.GetOk("connection_string") //dTos("connection_string", d)

	return NewClient(connectionString.(string))
}
