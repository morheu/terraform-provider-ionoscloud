package ionoscloud

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	ionoscloud "github.com/ionos-cloud/sdk-go/v6"
)

func resourceLanIPFailover() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceLanIPFailoverCreate,
		ReadContext:   resourceLanIPFailoverRead,
		UpdateContext: resourceLanIPFailoverUpdate,
		DeleteContext: resourceLanIPFailoverDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceIpFailoverImporter,
		},
		Schema: map[string]*schema.Schema{
			"ip": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.All(validation.IsIPAddress),
			},
			"nicuuid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.All(validation.IsUUID),
			},
			"lan_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.All(validation.StringIsNotWhiteSpace),
			},
			"datacenter_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.All(validation.IsUUID),
			},
		},
		Timeouts: &resourceDefaultTimeouts,
	}
}

func resourceLanIPFailoverCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(SdkBundle).CloudApiClient
	dcid := d.Get("datacenter_id").(string)
	lanid := d.Get("lan_id").(string)
	if lanid == "" {
		diags := diag.FromErr(fmt.Errorf("'lan_id' is missing, please provide a valid lan ID "))
		return diags
	}
	ip := d.Get("ip").(string)
	nicUuid := d.Get("nicuuid").(string)

	properties := &ionoscloud.LanProperties{}

	properties.IpFailover = &[]ionoscloud.IPFailover{
		{
			Ip:      &ip,
			NicUuid: &nicUuid,
		}}

	lan, apiResponse, err := client.LANsApi.DatacentersLansPatch(ctx, dcid, lanid).Lan(*properties).Execute()
	logApiRequestTime(apiResponse)
	if err != nil {
		diags := diag.FromErr(fmt.Errorf("an error occured while patching a lans failover group  %s %w", lanid, err))
		return diags
	}

	// Wait, catching any errors
	_, errState := getStateChangeConf(meta, d, apiResponse.Header.Get("Location"), schema.TimeoutCreate).WaitForStateContext(ctx)
	if errState != nil {
		diags := diag.FromErr(errState)
		return diags
	}

	d.SetId(*lan.Id)

	return resourceLanIPFailoverRead(ctx, d, meta)
}

func resourceLanIPFailoverRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(SdkBundle).CloudApiClient

	lan, apiResponse, err := client.LANsApi.DatacentersLansFindById(ctx, d.Get("datacenter_id").(string), d.Id()).Execute()
	logApiRequestTime(apiResponse)

	if err != nil {
		if apiResponse != nil && apiResponse.Response != nil && apiResponse.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		diags := diag.FromErr(fmt.Errorf("an error occured while fetching a lan ID %s %w", d.Id(), err))
		return diags
	}

	if lan.Properties.IpFailover != nil {
		err := d.Set("ip", *(*lan.Properties.IpFailover)[0].Ip)
		if err != nil {
			diags := diag.FromErr(fmt.Errorf("error while setting ip property for IpFailover %s: %w", d.Id(), err))
			return diags
		}
	}

	if lan.Properties.IpFailover != nil {
		err := d.Set("nicuuid", *(*lan.Properties.IpFailover)[0].NicUuid)
		if err != nil {
			diags := diag.FromErr(fmt.Errorf("error while setting nicuuid property for IpFailover %s: %w", d.Id(), err))
			return diags
		}
	}

	if lan.Id != nil {
		err := d.Set("lan_id", *lan.Id)
		if err != nil {
			diags := diag.FromErr(fmt.Errorf("error while setting lan_id property for IpFailover %s: %sw", d.Id(), err))
			return diags
		}
	}

	if err := d.Set("datacenter_id", d.Get("datacenter_id").(string)); err != nil {
		diags := diag.FromErr(err)
		return diags
	}

	return nil
}

func resourceLanIPFailoverUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(SdkBundle).CloudApiClient

	properties := &ionoscloud.LanProperties{}
	dcid := d.Get("datacenter_id").(string)
	lanid := d.Get("lan_id").(string)
	ip := d.Get("ip").(string)
	nicUuid := d.Get("nicuuid").(string)

	properties.IpFailover = &[]ionoscloud.IPFailover{
		{
			Ip:      &ip,
			NicUuid: &nicUuid,
		}}

	_, apiResponse, err := client.LANsApi.DatacentersLansPatch(ctx, dcid, lanid).Lan(*properties).Execute()
	logApiRequestTime(apiResponse)
	if err != nil {
		diags := diag.FromErr(fmt.Errorf("an error occured while patching a lan ID %s %w", d.Id(), err))
		return diags
	}

	// Wait, catching any errors
	_, errState := getStateChangeConf(meta, d, apiResponse.Header.Get("Location"), schema.TimeoutUpdate).WaitForStateContext(ctx)
	if errState != nil {
		diags := diag.FromErr(errState)
		return diags
	}

	return resourceLanIPFailoverRead(ctx, d, meta)
}

func resourceLanIPFailoverDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(SdkBundle).CloudApiClient

	dcid := d.Get("datacenter_id").(string)
	lanid := d.Get("lan_id").(string)

	//remove the failover group
	properties := &ionoscloud.LanProperties{
		IpFailover: &[]ionoscloud.IPFailover{},
	}

	_, apiResponse, err := client.LANsApi.DatacentersLansPatch(ctx, dcid, lanid).Lan(*properties).Execute()
	logApiRequestTime(apiResponse)
	if err != nil {
		/*
						//try again in 90 seconds
						time.Sleep(90 * time.Second)
						_, apiResponse, err = client.LANsApi.DatacentersLansPatch(ctx, dcid, lanid).Lan(*properties).Execute()
			logApiRequestTime(apiResponse)

						if err != nil && (apiResponse == nil || apiResponse.StatusCode != 404) {
							return fmt.Errorf("an error occured while removing a lans ipfailover groups dcId %s ID %s %s", d.Get("datacenter_id").(string), d.Id(), err)
						}
		*/
		diags := diag.FromErr(fmt.Errorf("an error occured while removing a lans ipfailover groups dcId %s ID %s %w", d.Get("datacenter_id").(string), d.Id(), err))
		return diags
	}

	// Wait, catching any errors
	_, errState := getStateChangeConf(meta, d, apiResponse.Header.Get("Location"), schema.TimeoutDelete).WaitForStateContext(ctx)
	if errState != nil {
		diags := diag.FromErr(errState)
		return diags
	}

	d.SetId("")
	return nil
}
