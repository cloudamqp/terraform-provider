package main

import (
	"github.com/84codes/go-api/api"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceCreate,
		Read:   resourceRead,
		Update: resourceUpdate,
		Delete: resourceDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the instance",
			},
			"plan": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the plan, valid options are: lemur, tiger, bunny, rabbit, panda, ape, hippo, lion",
			},
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the region you want to create your instance in",
			},
			"vpc_subnet": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Dedicated VPC subnet, shouldn't overlap with your current VPC's subnet",
			},
			"nodes": {
				Type:        schema.TypeInt,
				Default:     1,
				Optional:    true,
				Description: "Number of nodes in cluster (plan must support it)",
			},
			"rmq_version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "RabbitMQ version",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "URL of the CloudAMQP instance",
			},
			"apikey": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "API key for the CloudAMQP instance",
			},
			"notifications": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: resourceNotification(),
			},
		},
	}
}

func resourceCreate(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*api.API)
	keys := []string{"name", "plan", "region", "nodes"}
	params := make(map[string]interface{})
	for _, k := range keys {
		if v := d.Get(k); v != nil {
			params[k] = v
		}
	}
	data, err := api.CreateInstance(params)
	if err != nil {
		return err
	}
	d.SetId(data["id"].(string))
	for k, v := range data {
		d.Set(k, v)
	}
	return nil

}

func resourceRead(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*api.API)
	data, err := api.ReadInstance(d.Id())
	if err != nil {
		return err
	}
	for k, v := range data {
		d.Set(k, v)
	}
	return nil
}

func resourceUpdate(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*api.API)
	keys := []string{"name", "plan", "nodes"}
	params := make(map[string]interface{})
	for _, k := range keys {
		params[k] = d.Get(k)
	}
	return api.UpdateInstance(d.Id(), params)
}

func resourceDelete(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*api.API)
	return api.DeleteInstance(d.Id())
}
