package ionoscloud

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	ionoscloud "github.com/ionos-cloud/sdk-go/v6"
	"log"
)

func dataSourceNIC() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNicRead,
		Schema: map[string]*schema.Schema{
			"server_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.All(validation.StringIsNotWhiteSpace),
			},
			"datacenter_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.All(validation.StringIsNotWhiteSpace),
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"lan": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dhcp": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"ips": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
				Optional: true,
			},
			"firewall_active": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"firewall_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"device_number": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"pci_slot": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
		Timeouts: &resourceDefaultTimeouts,
	}
}
func getNicDataSourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"server_id": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.All(validation.StringIsNotWhiteSpace),
		},
		"datacenter_id": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.All(validation.StringIsNotWhiteSpace),
		},
		"id": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"lan": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"dhcp": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"ips": {
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Computed: true,
			Optional: true,
		},
		"firewall_active": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"firewall_type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"mac": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"device_number": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"pci_slot": {
			Type:     schema.TypeInt,
			Computed: true,
		},
	}
}
func dataSourceNicRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(SdkBundle).CloudApiClient

	t, dIdOk := data.GetOk("datacenter_id")
	st, sIdOk := data.GetOk("server_id")
	if !dIdOk || !sIdOk {
		return diag.FromErr(fmt.Errorf("datacenter id and server id must be set"))
	}
	var datacenterId, serverId string

	if dIdOk {
		datacenterId = t.(string)
	}
	if sIdOk {
		serverId = st.(string)
	}
	var name string
	id, idOk := data.GetOk("id")

	t, nameOk := data.GetOk("name")
	if nameOk {
		name = t.(string)
	}
	var nic ionoscloud.Nic
	var err error
	var apiResponse *ionoscloud.APIResponse

	if !idOk && !nameOk {
		return diag.FromErr(fmt.Errorf("either id, or name must be set"))
	}
	if idOk {
		nic, apiResponse, err = client.NetworkInterfacesApi.DatacentersServersNicsFindById(ctx, datacenterId, serverId, id.(string)).Execute()
		logApiRequestTime(apiResponse)
		if err != nil {
			return diag.FromErr(fmt.Errorf("error getting nic with id %s %w", id.(string), err))
		}
		if nameOk {
			if *nic.Properties.Name != name {
				return diag.FromErr(fmt.Errorf("name of nic (UUID=%s, name=%s) does not match expected name: %s",
					*nic.Id, *nic.Properties.Name, name))
			}
		}
	} else {
		nics, apiResponse, err := client.NetworkInterfacesApi.DatacentersServersNicsGet(ctx, datacenterId, serverId).Depth(1).Filter("name", name).Execute()
		logApiRequestTime(apiResponse)

		if err != nil {
			return diag.FromErr(fmt.Errorf("an error occured while fetching nics: %w ", err))
		}

		if nics.Items != nil && len(*nics.Items) > 0 {
			nic = (*nics.Items)[len(*nics.Items)-1]
			log.Printf("[INFO] %v nics found matching the search criteria. Getting the latest nic from the list %v", len(*nics.Items), *nic.Id)
		} else {
			return diag.FromErr(fmt.Errorf("no nic found with the specified name %s", name))
		}
	}

	if err := NicSetData(data, &nic); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
