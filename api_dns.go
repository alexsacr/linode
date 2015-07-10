package linode

// DomainCreateOpts contains the optional arguments to DomainCreate().
type DomainCreateOpts struct {
	Description  *string `args:"Description"`
	SOAEmail     *string `args:"SOA_Email"`
	RefreshSec   *int    `args:"Refresh_sec"`
	RetrySec     *int    `args:"Retry_sec"`
	ExpireSec    *int    `args:"Expire_sec"`
	TTLSec       *int    `args:"TTL_sec"`
	DisplayGroup *string `args:"lpm_displayGroup"`
	Status       *int    `args:"status"`
	MasterIPs    *string `args:"master_ips"`
	AXFRIPs      *string `args:"axfr_ips"`
}

// DomainCreate maps to the 'domain.create' call.
//
// https://www.linode.com/api/dns/domain.create
func (c *Client) DomainCreate(domain string, Type string,
	conf DomainCreateOpts) (domainID int, err error) {

	args, err := c.argMarshal(conf)
	if err != nil {
		return 0, err
	}
	args["Domain"] = domain
	args["Type"] = Type

	data, err := c.apiCall("domain.create", args)
	if err != nil {
		return 0, err
	}

	err = unmarshalSingle(data, "DomainID", &domainID)
	if err != nil {
		return 0, err
	}

	return domainID, nil
}

// DomainDelete maps to the 'domain.delete' call.
//
// https://www.linode.com/api/dns/domain.delete
func (c *Client) DomainDelete(domainID int) error {
	args := make(map[string]interface{})
	args["DomainID"] = domainID

	_, err := c.apiCall("domain.delete", args)
	if err != nil {
		return err
	}

	return nil
}

// Domain is the API response to the 'domain.list' call.
type Domain struct {
	ID           int    `mapstructure:"DOMAINID"`
	Description  string `mapstructure:"DESCRIPTION"`
	Type         string `mapstructure:"TYPE"`
	Status       int    `mapstructure:"STATUS"`
	SOAEmail     string `mapstructure:"SOA_EMAIL"`
	Domain       string `mapstructure:"DOMAIN"`
	RetrySec     int    `mapstructure:"RETRY_SEC"`
	MasterIPs    string `mapstructure:"MASTER_IPS"`
	AXFRIPs      string `mapstructure:"AXFR_IPS"`
	ExpireSec    int    `mapstructure:"EXPIRE_SEC"`
	RefreshSec   int    `mapstructure:"REFRESH_SEC"`
	TTLSec       int    `mapstructure:"TTL_SEC"`
	DisplayGroup string `mapstructure:"LPM_DISPLAYGROUP"`
}

// DomainList maps to the 'domain.list' call.
//
// https://www.linode.com/api/dns/domain.list
func (c *Client) DomainList(domainID *int) ([]Domain, error) {
	args := make(map[string]interface{})
	args["DomainID"] = domainID

	data, err := c.apiCall("domain.list", args)
	if err != nil {
		return nil, err
	}

	var domains []Domain
	err = unmarshalMultiMap(data, &domains)
	if err != nil {
		return nil, err
	}

	// fix goofy return value
	for i, d := range domains {
		if d.AXFRIPs == "none" {
			domains[i].AXFRIPs = ""
		}
	}

	return domains, nil
}

// DomainUpdateOpts contains the optional arguments to DomainUpdate().
type DomainUpdateOpts struct {
	Domain       *string `args:"Domain"`
	Type         *string `args:"Type"`
	SOAEmail     *string `args:"SOA_Email"`
	RefreshSec   *int    `args:"Refresh_sec"`
	RetrySec     *int    `args:"Retry_sec"`
	ExpireSec    *int    `args:"Expire_sec"`
	TTLSec       *int    `args:"TTL_sec"`
	DisplayGroup *string `args:"lpm_displayGroup"`
	Status       *int    `args:"status"`
	MasterIPs    *string `args:"master_ips"`
	AXFRIPs      *string `args:"axfr_ips"`
}

// DomainUpdate maps to the 'domain.update' call.
//
// https://www.linode.com/api/dns/domain.update
func (c *Client) DomainUpdate(domainID int, conf DomainUpdateOpts) error {
	args, err := c.argMarshal(conf)
	if err != nil {
		return err
	}
	args["DomainID"] = domainID

	_, err = c.apiCall("domain.update", args)
	if err != nil {
		return err
	}

	return nil
}

// DomainResourceCreateOpts contains the optional arguments to
// DomainResourceCreate().
type DomainResourceCreateOpts struct {
	Name     *string `args:"Name"`
	Target   *string `args:"Target"`
	Priority *int    `args:"Priority"`
	Weight   *int    `args:"Weight"`
	Protocol *string `args:"Protocol"`
	TTLSec   *int    `args:"TTL_sec"`
	Port     *int    `args:"Port"`
}

// DomainResourceCreate maps to the 'domain.resource.create' call.
//
// NOTE: The TTL passed is not respected by the API.  It must be set to one
// of: 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800,
// 1209600, or 2419200.
//
// https://www.linode.com/api/dns/domain.resource.create
func (c *Client) DomainResourceCreate(domainID int, rType string,
	conf DomainResourceCreateOpts) (resourceID int, err error) {

	args, err := c.argMarshal(conf)
	if err != nil {
		return 0, err
	}
	args["DomainID"] = domainID
	args["Type"] = rType

	data, err := c.apiCall("domain.resource.create", args)
	if err != nil {
		return 0, err
	}

	err = unmarshalSingle(data, "ResourceID", &resourceID)
	if err != nil {
		return 0, err
	}

	return resourceID, nil
}

// DomainResourceDelete maps to the "domain.resource.delete" call.
//
// https://www.linode.com/api/dns/domain.resource.delete
func (c *Client) DomainResourceDelete(domainID int, resourceID int) error {
	args := make(map[string]interface{})
	args["DomainID"] = domainID
	args["ResourceID"] = resourceID

	_, err := c.apiCall("domain.resource.delete", args)
	if err != nil {
		return err
	}

	return nil
}

// DomainResource is the API response to the 'domain.resource.list' call.
type DomainResource struct {
	Protocol string `mapstructure:"PROTOCOL"`
	TTLSec   int    `mapstructure:"TTL_SEC"`
	Priority int    `mapstructure:"PRIORITY"`
	Type     string `mapstructure:"TYPE"`
	Target   string `mapstructure:"TARGET"`
	Weight   int    `mapstructure:"WEIGHT"`
	ID       int    `mapstructure:"RESOURCEID"`
	Port     int    `mapstructure:"PORT"`
	DomainID int    `mapstructure:"DOMAINID"`
	Name     string `mapstructure:"NAME"`
}

// DomainResourceList maps to the 'domain.resource.list' call.
//
// https://www.linode.com/api/dns/domain.resource.list
func (c *Client) DomainResourceList(domainID int, resourceID *int) ([]DomainResource, error) {
	args := make(map[string]interface{})
	args["DomainID"] = domainID
	args["ResourceID"] = resourceID

	data, err := c.apiCall("domain.resource.list", args)
	if err != nil {
		return nil, err
	}

	var dr []DomainResource
	err = unmarshalMultiMap(data, &dr)
	if err != nil {
		return nil, err
	}

	return dr, nil
}

// DomainResourceUpdateOpts contains the optional arguments to
// DomainResourceUpdate().
type DomainResourceUpdateOpts struct {
	DomainID *int    `args:"DomainID"`
	Name     *string `args:"Name"`
	Target   *string `args:"Target"`
	Priority *int    `args:"Priority"`
	Weight   *int    `args:"Weight"`
	Port     *int    `args:"Port"`
	Protocol *string `args:"Protocol"`
	TTLSec   *int    `args:"TTL_sec"`
}

// DomainResourceUpdate maps to the 'domain.resource.update' call.
//
// https://www.linode.com/api/dns/domain.resource.update
func (c *Client) DomainResourceUpdate(resourceID int, conf DomainResourceUpdateOpts) error {
	args, err := c.argMarshal(conf)
	if err != nil {
		return err
	}
	args["ResourceID"] = resourceID

	_, err = c.apiCall("domain.resource.update", args)
	if err != nil {
		return err
	}

	return nil
}
