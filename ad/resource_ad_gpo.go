package ad

import (
	"strings"

	"github.com/Sidoran/terraform-provider-ad/ad/internal/config"

	"github.com/Sidoran/terraform-provider-ad/ad/internal/winrmhelper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceADGPO() *schema.Resource {
	return &schema.Resource{
		Description: "`ad_gpo` manages Group Policy Objects (GPOs).",
		Create:      resourceADGPOCreate,
		Read:        resourceADGPORead,
		Update:      resourceADGPOUpdate,
		Delete:      resourceADGPODelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the GPO.",
			},
			"domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Domain of the GPO.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the GPO.",
			},
			"status": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "AllSettingsEnabled",
				ValidateFunc: validation.StringInSlice([]string{"AllSettingsEnabled", "UserSettingsDisabled", "ComputerSettingsDisabled", "AllSettingsDisabled"}, false),
				Description:  "Status of the GPO. Can be one of `AllSettingsEnabled`, `UserSettingsDisabled`, `ComputerSettingsDisabled`, or `AllSettingsDisabled` (case sensitive).",
			},
			"numeric_status": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"dn": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceADGPOCreate(d *schema.ResourceData, meta interface{}) error {
	g := winrmhelper.GetGPOFromResource(d)
	guid, err := g.NewGPO(meta.(*config.ProviderConf))
	if err != nil {
		return err
	}
	d.SetId(guid)
	return resourceADGPORead(d, meta)
}

func resourceADGPORead(d *schema.ResourceData, meta interface{}) error {
	if d.Id() == "" {
		return nil
	}
	g, err := winrmhelper.GetGPOFromHost(meta.(*config.ProviderConf), "", d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "GpoWithNameNotFound") || strings.Contains(err.Error(), "GpoWithIdNotFound") {
			d.SetId("")
			return nil
		}
		return err
	}
	_ = d.Set("domain", g.Domain)
	_ = d.Set("description", g.Description)
	_ = d.Set("status", g.Status)
	_ = d.Set("numeric_status", g.NumericStatus)
	_ = d.Set("name", g.Name)
	return nil
}

func resourceADGPOUpdate(d *schema.ResourceData, meta interface{}) error {
	g := winrmhelper.GetGPOFromResource(d)
	_, err := g.UpdateGPO(meta.(*config.ProviderConf), d)
	if err != nil {
		return err
	}
	return resourceADGPORead(d, meta)
}

func resourceADGPODelete(d *schema.ResourceData, meta interface{}) error {
	g := winrmhelper.GetGPOFromResource(d)
	err := g.DeleteGPO(meta.(*config.ProviderConf))
	if err != nil {
		return err
	}
	return nil
}
