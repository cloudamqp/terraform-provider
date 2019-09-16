package main

import (
	"github.com/84codes/go-api/api"
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"apikey": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLOUDAMQP_APIKEY", nil),
				Description: "Key used to authentication to the CloudAMQP Customer API",
			},
			"baseurl": &schema.Schema{
				Type:        schema.TypeString,
				Default:     "https://customer.cloudamqp.com",
				Optional:    true,
				Description: "Base URL to CloudAMQP Customer website",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"cloudamqp_instance" : resourceInstance(),
			"cloudamqp_notification" : resourceNotification(),
			"cloudamqp_alarm" : resourceAlarm(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return api.New(d.Get("baseurl").(string), d.Get("apikey").(string)), nil
}
