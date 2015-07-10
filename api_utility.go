package linode

import (
	"encoding/json"
	"fmt"
)

// Datacenter is the API response to the 'avail.datacenters' call.
type Datacenter struct {
	ID       int    `json:"DATACENTERID"`
	Location string `json:"LOCATION"`
	Abbr     string `json:"ABBR"`
}

// AvailDatacenters maps to the 'avail.datacenters' call.
//
// https://www.linode.com/api/utility/avail.datacenters
func (c *Client) AvailDatacenters() ([]Datacenter, error) {
	data, err := c.apiCall("avail.datacenters", map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	var ret []Datacenter
	err = json.Unmarshal(data, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// Distribution is the API repsonse to the 'avail.distributions' call.
type Distribution struct {
	Is64Bit       bool   `mapstructure:"IS64BIT"`
	Label         string `mapstructure:"LABEL"`
	MinImageSize  int    `mapstructure:"MINIMAGESIZE"`
	ID            int    `mapstructure:"DISTRIBUTIONID"`
	CreateDT      string `mapstructure:"CREATE_DT"`
	RequiresPVOps bool   `mapstructure:"REQUIRESPVOPSKERNEL"`
}

// AvailDistributions maps to the 'avail.distributions' call.
//
// https://www.linode.com/api/utility/avail.distributions
func (c *Client) AvailDistributions(distributionID *int) ([]Distribution, error) {
	args := make(map[string]interface{})
	args["DistributionID"] = distributionID

	data, err := c.apiCall("avail.distributions", args)
	if err != nil {
		return nil, err
	}

	var ret []Distribution
	err = unmarshalMultiMap(data, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// Kernel is the API response to the 'avail.kernels' call.
type Kernel struct {
	Label   string `mapstructure:"LABEL"`
	IsXen   bool   `mapstructure:"ISXEN"`
	IsPVOps bool   `mapstructure:"ISPVOPS"`
	ID      int    `mapstructure:"KERNELID"`
}

// AvailKernels maps to the 'avail.kernels' call.
//
// https://www.linode.com/api/utility/avail.kernels
func (c *Client) AvailKernels(kernelID *int, isXen *bool) ([]Kernel, error) {
	args := make(map[string]interface{})
	args["KernelID"] = kernelID
	args["isXen"] = isXen

	data, err := c.apiCall("avail.kernels", args)
	if err != nil {
		return nil, err
	}

	var ret []Kernel
	err = unmarshalMultiMap(data, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// LinodePlan is the API response to the 'avail.linodeplans' call.
type LinodePlan struct {
	Cores  int     `json:"CORES"`
	Price  float64 `json:"PRICE"`
	RAM    int     `json:"RAM"`
	Xfer   int     `json:"XFER"`
	ID     int     `json:"PLANID"`
	Label  string  `json:"LABEL"`
	Disk   int     `json:"DISK"`
	Hourly float64 `json:"HOURLY"`
}

// AvailLinodePlans maps to the 'avail.linodeplans' call.
//
// https://www.linode.com/api/utility/avail.linodeplans
func (c *Client) AvailLinodePlans(planID *int) ([]LinodePlan, error) {
	args := make(map[string]interface{})
	args["PlanID"] = planID

	data, err := c.apiCall("avail.linodeplans", args)
	if err != nil {
		return nil, err
	}

	var ret []LinodePlan
	err = json.Unmarshal(data, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// AvailStackScripts maps to the 'avail.stackscripts' call.
//
// https://www.linode.com/api/utility/avail.stackscripts
func (c *Client) AvailStackScripts(distID *int, distVendor *string,
	keywords *string) ([]StackScript, error) {

	args := make(map[string]interface{})
	args["DistributionID"] = distID
	args["DistributionVendor"] = distVendor
	args["keywords"] = keywords

	data, err := c.apiCall("avail.stackscripts", args)
	if err != nil {
		return nil, err
	}

	var ret []StackScript
	err = unmarshalMultiMap(data, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// TestEcho maps to the 'test.echo' call.  It has hardcoded arguments
// (foo:bar).  It can be used to test your API key.
//
// https://www.linode.com/api/utility/test.echo
func (c *Client) TestEcho() error {
	data, err := c.apiCall("test.echo", map[string]interface{}{"foo": "bar"})
	if err != nil {
		return err
	}

	var out string
	err = unmarshalSingle(data, "FOO", &out)
	if err != nil {
		return err
	}

	if out != "bar" {
		return fmt.Errorf("unexpected echo response: '%s'", out)
	}

	return nil
}
