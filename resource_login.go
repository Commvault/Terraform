package main

import (
    "github.com/TerraformProvider/handler"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


func resourceLogin() *schema.Resource {
    return &schema.Resource{
        Create: resourceLoginCreate,
        Read:   resourceLoginRead,
        Update: resourceLoginUpdate,
        Delete: resourceLoginDelete,

        Schema: map[string]*schema.Schema{
            "username": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "password": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceLoginCreate(d *schema.ResourceData, m interface{}) error {
    username := d.Get("username").(string)
    password := d.Get("password").(string)
    handler.GenerateAuthToken(username,password)
    d.SetId(username)
    return resourceLoginRead(d,m)
}

func resourceLoginRead(d *schema.ResourceData, m interface{}) error {
    return nil
}

func resourceLoginUpdate(d *schema.ResourceData, m interface{}) error {
    username := d.Get("username").(string)
    password := d.Get("password").(string)
    handler.GenerateAuthToken(username,password)
    d.SetId(username)
    return resourceLoginRead(d,m)
}

func resourceLoginDelete(d *schema.ResourceData, m interface{}) error {
    return nil
}