package dme

import (
	"fmt"
	"log"
	"strings"

	"github.com/DNSMadeEasy/dme-go-client/client"
	"github.com/DNSMadeEasy/dme-go-client/models"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	// "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	// "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagedDNSRecordActions() *schema.Resource {
	return &schema.Resource{
		Create: resourceManagedDNSRecordActionsCreate,
		Update: resourceManagedDNSRecordActionsUpdate,
		Read:   resourceManagedDNSRecordActionsRead,
		Delete: resourceManagedDNSRecordActionsDelete,

		Schema: map[string]*schema.Schema{
			"domain_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"dynamic_dns": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"ttl": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			// "monitor": &schema.Schema{
			// 	Type:     schema.TypeBool,
			// 	Optional: true,
			// 	Computed: true,
			// },

			// "failover": &schema.Schema{
			// 	Type:     schema.TypeBool,
			// 	Optional: true,
			// 	Computed: true,
			// },

			// "failed": &schema.Schema{
			// 	Type:     schema.TypeBool,
			// 	Optional: true,
			// 	Computed: true,
			// },

			"gtd_location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"caa_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"issuer_critical": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"keywords": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"title": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"redirect_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"hardlink": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"mx_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"weight": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"allow_overwrite": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceManagedDNSRecordActionsCreate(d *schema.ResourceData, m interface{}) error {
	dmeClient := m.(*client.Client)

	recordAttr := models.ManagedDNSRecordActions{}

	recordAttr.Name = d.Get("name").(string)
	// if name, ok := d.GetOk("name"); ok {
	// 	recordAttr.Name = name.(string)
	// }

	if value, ok := d.GetOk("value"); ok {
		recordAttr.Value = value.(string)
	}

	if Type, ok := d.GetOk("type"); ok {
		recordAttr.Type = Type.(string)
	}

	if dynamicdns, ok := d.GetOk("dynamic_dns"); ok {
		recordAttr.DynamicDNS = dynamicdns.(string)
	}

	if password, ok := d.GetOk("password"); ok {
		recordAttr.Password = password.(string)
	}

	if ttl, ok := d.GetOk("ttl"); ok {
		recordAttr.Ttl = ttl.(string)
	}

	if gtdlocation, ok := d.GetOk("gtd_location"); ok {
		recordAttr.GtdLocation = gtdlocation.(string)
	}

	if description, ok := d.GetOk("description"); ok {
		recordAttr.Description = description.(string)
	}

	if keywords, ok := d.GetOk("keywords"); ok {
		recordAttr.Keywords = keywords.(string)
	}

	if title, ok := d.GetOk("title"); ok {
		recordAttr.Title = title.(string)
	}

	if redirecttype, ok := d.GetOk("redirect_type"); ok {
		recordAttr.RedirectType = redirecttype.(string)
	}

	if hardlink, ok := d.GetOk("hardlink"); ok {
		recordAttr.HardLink = hardlink.(string)
	}

	if mxlevel, ok := d.GetOk("mx_level"); ok {
		recordAttr.MxLevel = mxlevel.(string)
	}

	if weight, ok := d.GetOk("weight"); ok {
		recordAttr.Weight = weight.(string)
	}

	if priority, ok := d.GetOk("priority"); ok {
		recordAttr.Priority = priority.(string)
	}

	if port, ok := d.GetOk("port"); ok {
		recordAttr.Port = port.(string)
	}

	if caatype, ok := d.GetOk("caa_type"); ok {
		recordAttr.CaaType = caatype.(string)
	}
	if issuer, ok := d.GetOk("issuer_critical"); ok {
		recordAttr.IssuerCritical = issuer.(string)
	}
	log.Println("Value of recordAttr: ", &recordAttr)

	cont, err := dmeClient.Save(&recordAttr, "dns/managed/"+d.Get("domain_id").(string)+"/records/")

	if err != nil {
		if strings.Contains(err.Error(), "already exists") && d.Get("allow_overwrite").(bool) && len(recordAttr.Name) > 0 {
			log.Printf("[DEBUG] DNS Record already exists however we are overwriting it")

			cont, err = dmeClient.GetbyId("dns/managed/" + d.Get("domain_id").(string) + "/records?recordName=" + recordAttr.Name + "&type=" + recordAttr.Type)
			if err != nil {
				return err
			}
			log.Println("Found record: ", cont)
			log.Println("Id valueinside update: ", strings.Trim(cont.S("data").S("id").String(), "[]"))

			_, err := dmeClient.Update(&recordAttr, "dns/managed/"+d.Get("domain_id").(string)+"/records/"+strings.Trim(fmt.Sprintf("%v", cont.S("data").S("id").String()), "[]"))
			if err != nil {
				log.Println("Error returned: ", err)
				return err
			}

			log.Println("Value of container: ", cont)
			idname := strings.Trim(cont.S("data").S("name").String(), "[]")
			if strings.HasPrefix(idname, "\"") && strings.HasSuffix(idname, "\"") {
				idname = strings.TrimSuffix(strings.TrimPrefix(idname, "\""), "\"")
			}
			log.Println("Idname value inside create: ", idname)
			log.Println("Id valueinside create: ", strings.Trim(cont.S("data").S("id").String(), "[]"))
			d.Set("name", strings.Trim(fmt.Sprintf("%v", idname), "[]"))
			d.SetId(strings.Trim(fmt.Sprintf("%v", cont.S("data").S("id")), "[]"))

			return resourceManagedDNSRecordActionsRead(d, m)
		} else {
			log.Println("Error returned: ", err)
			return err
		}
	}

	log.Println("Value of container: ", cont)
	idname := cont.S("name").String()
	if strings.HasPrefix(idname, "\"") && strings.HasSuffix(idname, "\"") {
		idname = strings.TrimSuffix(strings.TrimPrefix(idname, "\""), "\"")
	}
	log.Println("Idname value inside create: ", idname)
	log.Println("Id valueinside create: ", cont.S("id"))
	d.Set("name", fmt.Sprintf("%v", idname))
	d.SetId(fmt.Sprintf("%v", cont.S("id")))

	return resourceManagedDNSRecordActionsRead(d, m)
}

func resourceManagedDNSRecordActionsRead(d *schema.ResourceData, m interface{}) error {
	dmeClient := m.(*client.Client)
	log.Println("andkamdak")
	dnsId := d.Id()
	log.Println("Inside read ID value: ", dnsId)
	con, err := dmeClient.GetbyId("dns/managed/" + d.Get("domain_id").(string) + "/records?recordName=" + d.Get("name").(string) + "&type=" + d.Get("type").(string))
	if err != nil {
		return err
	}
	log.Println("Inside read method: ", con)

	data := con.S("data").Data().([]interface{})
	var count int
	log.Println("data: ", data)

	for _, info := range data {
		val := info.(map[string]interface{})
		s := fmt.Sprintf("%.f", val["id"])
		log.Println("s value: ", s)
		if s == dnsId {
			break
		}
		count = count + 1
	}

	cont1 := con.S("data").Index(count)

	d.SetId(fmt.Sprintf("%v", cont1.S("id").String()))

	log.Println("INSIDE READ ID value: ", cont1.S("id").String())
	d.Set("name", StripQuotes(cont1.S("name").String()))
	log.Println("Inside read ID name value: ", StripQuotes(cont1.S("name").String()))

	str := StripQuotes(cont1.S("value").String())

	if d.Get("type").(string) == "TXT" || d.Get("type").(string) == "SPF" || d.Get("type").(string) == "CAA" {
		str = str[2 : len(str)-2]
	}
	log.Println("After trim: ", str)

	d.Set("value", str)

	d.Set("type", StripQuotes(cont1.S("type").String()))
	d.Set("dynamic_dns", StripQuotes(cont1.S("dynamicDns").String()))
	d.Set("password", StripQuotes(cont1.S("password").String()))
	d.Set("ttl", StripQuotes(cont1.S("ttl").String()))
	d.Set("gtd_location", StripQuotes(cont1.S("gtdLocation").String()))
	d.Set("description", StripQuotes(cont1.S("description").String()))
	d.Set("keywords", StripQuotes(cont1.S("keywords").String()))
	d.Set("title", StripQuotes(cont1.S("title").String()))
	d.Set("redirect_type", StripQuotes(cont1.S("redirectType").String()))
	d.Set("hardlink", StripQuotes(cont1.S("hardLink").String()))
	d.Set("mx_level", StripQuotes(cont1.S("mxLevel").String()))
	d.Set("weight", StripQuotes(cont1.S("weight").String()))
	d.Set("port", StripQuotes(cont1.S("port").String()))
	d.Set("priority", StripQuotes(cont1.S("priority").String()))
	d.Set("caa_type", StripQuotes(cont1.S("caaType").String()))
	d.Set("issuer_critical", StripQuotes(cont1.S("issuerCritical").String()))

	return nil
}

func resourceManagedDNSRecordActionsUpdate(d *schema.ResourceData, m interface{}) error {
	dmeClient := m.(*client.Client)
	recordAttr := models.ManagedDNSRecordActions{}

	if name, ok := d.GetOk("name"); ok {
		recordAttr.Name = name.(string)
	}

	if value, ok := d.GetOk("value"); ok {
		recordAttr.Value = value.(string)
	}

	if Type, ok := d.GetOk("type"); ok {
		recordAttr.Type = Type.(string)
	}

	if dynamicdns, ok := d.GetOk("dynamic_dns"); ok {
		recordAttr.DynamicDNS = dynamicdns.(string)
	}

	if password, ok := d.GetOk("password"); ok {
		recordAttr.Password = password.(string)
	}

	if ttl, ok := d.GetOk("ttl"); ok {
		recordAttr.Ttl = ttl.(string)
	}

	if gtdlocation, ok := d.GetOk("gtd_location"); ok {
		recordAttr.GtdLocation = gtdlocation.(string)
	}

	if description, ok := d.GetOk("description"); ok {
		recordAttr.Description = description.(string)
	}

	if keywords, ok := d.GetOk("keywords"); ok {
		recordAttr.Keywords = keywords.(string)
	}

	if title, ok := d.GetOk("title"); ok {
		recordAttr.Title = title.(string)
	}

	if redirecttype, ok := d.GetOk("redirect_type"); ok {
		recordAttr.RedirectType = redirecttype.(string)
	}

	if hardlink, ok := d.GetOk("hardlink"); ok {
		recordAttr.HardLink = hardlink.(string)
	}

	if mxlevel, ok := d.GetOk("mx_level"); ok {
		recordAttr.MxLevel = mxlevel.(string)
	}

	if weight, ok := d.GetOk("weight"); ok {
		recordAttr.Weight = weight.(string)
	}

	if priority, ok := d.GetOk("priority"); ok {
		recordAttr.Priority = priority.(string)
	}

	if port, ok := d.GetOk("port"); ok {
		recordAttr.Port = port.(string)
	}

	if caatype, ok := d.GetOk("caa_type"); ok {
		recordAttr.CaaType = caatype.(string)
	}

	if issuer, ok := d.GetOk("issuer_critical"); ok {
		recordAttr.IssuerCritical = issuer.(string)
	}

	log.Println("Inside update method: recordattr: ", recordAttr)
	recordId := d.Id()

	recordAttr.IdUpdate = recordId
	_, err := dmeClient.Update(&recordAttr, "dns/managed/"+d.Get("domain_id").(string)+"/records/"+recordId)
	if err != nil {
		return err
	}

	return resourceManagedDNSRecordActionsRead(d, m)
}

func resourceManagedDNSRecordActionsDelete(d *schema.ResourceData, m interface{}) error {
	dmeClient := m.(*client.Client)
	dn := d.Id()

	err := dmeClient.Delete("dns/managed/" + d.Get("domain_id").(string) + "/records/" + dn)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}
