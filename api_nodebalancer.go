package linode

import "encoding/json"

// NodeBalancerCreate maps to the 'nodebalancer.create' call.
//
// https://www.linode.com/api/nodebalancer/nodebalancer.create
func (c *Client) NodeBalancerCreate(datacenterID int, label *string,
	throttle *int) (nbID int, err error) {

	args := make(map[string]interface{})
	args["DatacenterID"] = datacenterID
	args["Label"] = label
	args["ClientConnThrottle"] = throttle

	data, err := c.apiCall("nodebalancer.create", args)
	if err != nil {
		return 0, err
	}

	err = unmarshalSingle(data, "NodeBalancerID", &nbID)
	if err != nil {
		return 0, err
	}

	return nbID, nil
}

// NodeBalancerDelete maps to the 'nodebalancer.delete' call.
//
// https://www.linode.com/api/nodebalancer/nodebalancer.delete
func (c *Client) NodeBalancerDelete(nbID int) error {
	args := make(map[string]interface{})
	args["NodeBalancerID"] = nbID

	_, err := c.apiCall("nodebalancer.delete", args)
	if err != nil {
		return err
	}

	return nil
}

// NodeBalancer is the API response to the 'nodebalancer.list' call.
type NodeBalancer struct {
	ID           int    `json:"NODEBALANCERID"`
	Label        string `json:"LABEL"`
	DatacenterID int    `json:"DATACENTERID"`
	Hostname     string `json:"HOSTNAME"`
	IPv4Addr     string `json:"ADDRESS4"`
	IPv6Addr     string `json:"ADDRESS6"`
	Throttle     int    `json:"CLIENTCONNTHROTTLE"`
}

// NodeBalancerList maps to the 'nodebalancer.list' call.
//
// https://www.linode.com/api/nodebalancer/nodebalancer.list
func (c *Client) NodeBalancerList(nbID *int) ([]NodeBalancer, error) {
	args := make(map[string]interface{})
	args["NodeBalancerID"] = nbID

	data, err := c.apiCall("nodebalancer.list", args)
	if err != nil {
		return nil, err
	}

	var ret []NodeBalancer
	err = json.Unmarshal(data, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// NodeBalancerUpdate maps to the 'nodebalancer.update' call.
//
// https://www.linode.com/api/nodebalancer/nodebalancer.update
func (c *Client) NodeBalancerUpdate(nbID int, label *string, throttle *int) error {
	args := make(map[string]interface{})
	args["NodeBalancerID"] = nbID
	args["Label"] = label
	args["ClientConnThrottle"] = throttle

	_, err := c.apiCall("nodebalancer.update", args)
	if err != nil {
		return err
	}

	return nil
}

// NodeBalancerConfigCreateOpts contains the optional arguments to
// NodeBalancerConfigCreate().
type NodeBalancerConfigCreateOpts struct {
	Port          *int    `args:"Port"`
	Protocol      *string `args:"Protocol"`
	Algorithm     *string `args:"Algorithm"`
	Stickiness    *string `args:"Stickiness"`
	Check         *string `args:"check"`
	CheckInterval *int    `args:"check_interval"`
	CheckTimeout  *int    `args:"check_timeout"`
	CheckAttempts *int    `args:"check_attempts"`
	CheckPath     *string `args:"check_path"`
	CheckBody     *string `args:"check_body"`
	CheckPassive  *bool   `args:"check_passive,int"`
	SSLCert       *string `args:"ssl_cert"`
	SSLKey        *string `args:"ssl_key"`
}

// NodeBalancerConfigCreate maps to the 'nodebalancer.config.create' call.
//
// https://www.linode.com/api/nodebalancer/nodebalancer.config.create
func (c *Client) NodeBalancerConfigCreate(nbID int,
	conf NodeBalancerConfigCreateOpts) (confID int, err error) {

	args, err := c.argMarshal(conf)
	if err != nil {
		return 0, err
	}
	args["NodeBalancerID"] = nbID

	data, err := c.apiCall("nodebalancer.config.create", args)
	if err != nil {
		return 0, err
	}

	var cID int
	err = unmarshalSingle(data, "ConfigID", &cID)
	if err != nil {
		return 0, err
	}

	return cID, nil
}

// NodeBalancerConfigDelete maps to the 'nodebalancer.config.delete' call.
//
// https://www.linode.com/api/nodebalancer/nodebalancer.config.delete
func (c *Client) NodeBalancerConfigDelete(nbID int, confID int) error {
	args := make(map[string]interface{})
	args["NodeBalancerID"] = nbID
	args["ConfigID"] = confID

	_, err := c.apiCall("nodebalancer.config.delete", args)
	if err != nil {
		return err
	}

	return nil
}

// NodeBalancerConfig is the API response to the 'nodebalancer.config.list' call.
type NodeBalancerConfig struct {
	Stickiness     string `mapstructure:"STICKINESS"`
	CheckPath      string `mapstructure:"CHECK_PATH"`
	Port           int    `mapstructure:"PORT"`
	CheckBody      string `mapstructure:"CHECK_BODY"`
	Check          string `mapstructure:"CHECK"`
	CheckInterval  int    `mapstructure:"CHECK_INTERVAL"`
	Protocol       string `mapstructure:"PROTOCOL"`
	ID             int    `mapstructure:"CONFIGID"`
	Algorithm      string `mapstructure:"ALGORITHM"`
	CheckTimeout   int    `mapstructure:"CHECK_TIMEOUT"`
	NodeBalancerID int    `mapstructure:"NODEBALANCERID"`
	CheckAttempts  int    `mapstructure:"CHECK_ATTEMPTS"`
	CheckPassive   bool   `mapstructure:"CHECK_PASSIVE"`
	SSLFingerprint string `mapstructure:"SSL_FINGERPRINT"`
	SSLCommonName  string `mapstructure:"SSL_COMMONNAME"`
}

// NodeBalancerConfigList maps to the 'nodebalancer.config.list' call.
//
// https://www.linode.com/api/nodebalancer/nodebalancer.config.list
func (c *Client) NodeBalancerConfigList(nbID int, confID *int) ([]NodeBalancerConfig, error) {
	args := make(map[string]interface{})
	args["NodeBalancerID"] = nbID
	args["ConfigID"] = confID

	data, err := c.apiCall("nodebalancer.config.list", args)
	if err != nil {
		return nil, err
	}

	var ret []NodeBalancerConfig
	err = unmarshalMultiMap(data, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// NodeBalancerConfigUpdateOpts contains the optional arguments to
// NodeBalancerConfigUpdate().
type NodeBalancerConfigUpdateOpts struct {
	Port          *int    `args:"Port"`
	Protocol      *string `args:"Protocol"`
	Algorithm     *string `args:"Algorithm"`
	Stickiness    *string `args:"Stickiness"`
	Check         *string `args:"check"`
	CheckInterval *int    `args:"check_interval"`
	CheckTimeout  *int    `args:"check_timeout"`
	CheckAttempts *int    `args:"check_attempts"`
	CheckPath     *string `args:"check_path"`
	CheckBody     *string `args:"check_body"`
	CheckPassive  *bool   `args:"check_passive,int"`
	SSLCert       *string `args:"ssl_cert"`
	SSLKey        *string `args:"ssl_key"`
}

// NodeBalancerConfigUpdate maps to the 'nodebalancer.config.update' call.
//
// https://www.linode.com/api/nodebalancer/nodebalancer.config.update
func (c *Client) NodeBalancerConfigUpdate(confID int, conf NodeBalancerConfigUpdateOpts) error {
	args, err := c.argMarshal(conf)
	if err != nil {
		return err
	}
	args["ConfigID"] = confID

	_, err = c.apiCall("nodebalancer.config.update", args)
	if err != nil {
		return err
	}

	return nil
}

// NodeBalancerNodeCreate maps to the 'nodebalancer.node.create' call.
//
// https://www.linode.com/api/nodebalancer/nodebalancer.node.create
func (c *Client) NodeBalancerNodeCreate(confID int, label string, address string, weight *int,
	mode *string) (nodeID int, err error) {

	args := make(map[string]interface{})
	args["ConfigID"] = confID
	args["Label"] = label
	args["Address"] = address
	args["Weight"] = weight
	args["Mode"] = mode

	data, err := c.apiCall("nodebalancer.node.create", args)
	if err != nil {
		return 0, err
	}

	var nID int
	err = unmarshalSingle(data, "NodeID", &nID)
	if err != nil {
		return 0, err
	}

	return nID, nil
}

// NodeBalancerNodeDelete maps to the 'nodebalancer.node.delete' call.
//
// https://www.linode.com/api/nodebalancer/nodebalancer.node.delete
func (c *Client) NodeBalancerNodeDelete(nodeID int) error {
	args := make(map[string]interface{})
	args["NodeID"] = nodeID

	_, err := c.apiCall("nodebalancer.node.delete", args)
	if err != nil {
		return err
	}

	return nil
}

// NodeBalancerNode is the API response to the 'nodebalancer.node.list' call.
type NodeBalancerNode struct {
	Weight         int    `mapstructure:"WEIGHT"`
	Address        string `mapstructure:"ADDRESS"`
	Label          string `mapstructure:"LABEL"`
	ID             int    `mapstructure:"NODEID"`
	Mode           string `mapstructure:"MODE"`
	Status         string `mapstructure:"STATUS"`
	NodeBalancerID int    `mapstructure:"NODEBALANCERID"`
	ConfigID       int    `mapstructure:"CONFIGID"`
}

// NodeBalancerNodeList maps to the 'nodebalancer.node.list' call.
//
// https://www.linode.com/api/nodebalancer/nodebalancer.node.list
func (c *Client) NodeBalancerNodeList(confID int, nodeID *int) ([]NodeBalancerNode, error) {
	args := make(map[string]interface{})
	args["ConfigID"] = confID
	args["NodeID"] = nodeID

	data, err := c.apiCall("nodebalancer.node.list", args)
	if err != nil {
		return nil, err
	}

	var nbn []NodeBalancerNode
	err = unmarshalMultiMap(data, &nbn)
	if err != nil {
		return nil, err
	}

	return nbn, nil
}

// NodeBalancerNodeUpdate maps to the 'nodebalancer.node.update' call.
//
// https://www.linode.com/api/nodebalancer/nodebalancer.node.update
func (c *Client) NodeBalancerNodeUpdate(nodeID int, label *string, address *string, weight *int,
	mode *string) error {

	args := make(map[string]interface{})
	args["NodeID"] = nodeID
	args["Label"] = label
	args["Address"] = address
	args["Weight"] = weight
	args["Mode"] = mode

	_, err := c.apiCall("nodebalancer.node.update", args)
	if err != nil {
		return err
	}

	return nil
}
