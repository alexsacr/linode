// +build !integration

package linode

import (
	"testing"

	"github.com/alexsacr/linode/_third_party/testify/assert"
	"github.com/alexsacr/linode/_third_party/testify/require"
)

func mockNodeBalancerCreateOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"NodeBalancerID":13128},"ACTION":"nodebalancer.create"}`
	params = map[string]string{
		"ClientConnThrottle": `5`,
		"DatacenterID":       `2`,
		"Label":              `testing`,
		"api_action":         `nodebalancer.create`,
		"api_key":            `foo`,
	}
	responses = append(responses, newMockAPIResponse("nodebalancer.create", params, output))

	return responses
}

func TestNodeBalancerCreateOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockNodeBalancerCreateOK()))
	defer ts.Close()

	nbID, err := c.NodeBalancerCreate(2, String("testing"), Int(5))
	require.NoError(t, err)
	require.Equal(t, 13128, nbID)
}

func mockNodeBalancerListOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":[{"HOSTNAME":"nb-198-58-121-23.dallas.nodebalancer.linode.com","LABEL":"testing","CLIENTCONNTHROTTLE":5,"DATACENTERID":2,"ADDRESS4":"198.58.121.23","NODEBALANCERID":13128,"ADDRESS6":"2600:3c00:1::c63a:7917"}],"ACTION":"nodebalancer.list"}`
	params = map[string]string{
		"NodeBalancerID": `13128`,
		"api_action":     `nodebalancer.list`,
		"api_key":        `foo`,
	}
	responses = append(responses, newMockAPIResponse("nodebalancer.list", params, output))

	return responses
}

func TestNodeBalancerListOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockNodeBalancerListOK()))
	defer ts.Close()

	nbList, err := c.NodeBalancerList(Int(13128))
	require.NoError(t, err)
	require.Len(t, nbList, 1)

	n := nbList[0]
	assert.Equal(t, "nb-198-58-121-23.dallas.nodebalancer.linode.com", n.Hostname)
	assert.Equal(t, "testing", n.Label)
	assert.Equal(t, 5, n.Throttle)
	assert.Equal(t, 2, n.DatacenterID)
	assert.Equal(t, "198.58.121.23", n.IPv4Addr)
	assert.Equal(t, "2600:3c00:1::c63a:7917", n.IPv6Addr)
	assert.Equal(t, 13128, n.ID)
}

func mockNodeBalancerUpdateOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"NodeBalancerID":13128},"ACTION":"nodebalancer.update"}`
	params = map[string]string{
		"ClientConnThrottle": `10`,
		"Label":              `testing-2`,
		"NodeBalancerID":     `13128`,
		"api_action":         `nodebalancer.update`,
		"api_key":            `foo`,
	}
	responses = append(responses, newMockAPIResponse("nodebalancer.update", params, output))

	return responses
}

func TestNodeBalancerUpdateOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockNodeBalancerUpdateOK()))
	defer ts.Close()

	err := c.NodeBalancerUpdate(13128, String("testing-2"), Int(10))
	require.NoError(t, err)
}

func mockNodeBalancerConfigCreateOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"ConfigID":14591},"ACTION":"nodebalancer.config.create"}`
	params = map[string]string{
		"Algorithm":      `roundrobin`,
		"NodeBalancerID": `13128`,
		"Port":           `80`,
		"Protocol":       `http`,
		"Stickiness":     `http_cookie`,
		"api_action":     `nodebalancer.config.create`,
		"api_key":        `foo`,
		"check":          `http`,
		"check_attempts": `15`,
		"check_body":     `bar`,
		"check_interval": `30`,
		"check_passive":  `0`,
		"check_path":     `/foo`,
		"check_timeout":  `29`,
	}
	responses = append(responses, newMockAPIResponse("nodebalancer.config.create", params, output))

	return responses
}

func TestNodeBalancerConfigCreateOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockNodeBalancerConfigCreateOK()))
	defer ts.Close()

	nbcco := NodeBalancerConfigCreateOpts{
		Port:          Int(80),
		Protocol:      String("http"),
		Algorithm:     String("roundrobin"),
		Stickiness:    String("http_cookie"),
		Check:         String("http"),
		CheckInterval: Int(30),
		CheckTimeout:  Int(29),
		CheckAttempts: Int(15),
		CheckPath:     String("/foo"),
		CheckBody:     String("bar"),
		CheckPassive:  Bool(false),
	}

	t.Log("nodebalancer.config.create...")
	confID, err := c.NodeBalancerConfigCreate(13128, nbcco)
	require.NoError(t, err)
	require.Equal(t, 14591, confID)
}

func mockNodeBalancerConfigListOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":[{"STICKINESS":"http_cookie","SSL_COMMONNAME":"","CHECK_PATH":"\/foo","CHECK_BODY":"bar","CHECK_INTERVAL":30,"SSL_FINGERPRINT":"","ALGORITHM":"roundrobin","CONFIGID":14591,"CHECK_ATTEMPTS":15,"NODEBALANCERID":13128,"PORT":80,"CHECK":"http","CHECK_PASSIVE":0,"PROTOCOL":"http","CHECK_TIMEOUT":29}],"ACTION":"nodebalancer.config.list"}`
	params = map[string]string{
		"ConfigID":       `14591`,
		"NodeBalancerID": `13128`,
		"api_action":     `nodebalancer.config.list`,
		"api_key":        `foo`,
	}
	responses = append(responses, newMockAPIResponse("nodebalancer.config.list", params, output))

	return responses
}

func TestNodeBalancerConfigListOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockNodeBalancerConfigListOK()))
	defer ts.Close()

	nbconfs, err := c.NodeBalancerConfigList(13128, Int(14591))
	require.NoError(t, err)
	require.Len(t, nbconfs, 1)

	nbc := nbconfs[0]
	assert.Equal(t, "http_cookie", nbc.Stickiness)
	assert.Equal(t, "", nbc.SSLCommonName)
	assert.Equal(t, "/foo", nbc.CheckPath)
	assert.Equal(t, "bar", nbc.CheckBody)
	assert.Equal(t, 30, nbc.CheckInterval)
	assert.Equal(t, "", nbc.SSLFingerprint)
	assert.Equal(t, "roundrobin", nbc.Algorithm)
	assert.Equal(t, 14591, nbc.ID)
	assert.Equal(t, 15, nbc.CheckAttempts)
	assert.Equal(t, 13128, nbc.NodeBalancerID)
	assert.Equal(t, 80, nbc.Port)
	assert.Equal(t, "http", nbc.Check)
	assert.False(t, nbc.CheckPassive)
	assert.Equal(t, "http", nbc.Protocol)
	assert.Equal(t, 29, nbc.CheckTimeout)
}

func mockNodeBalancerConfigUpdateOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"ConfigID":14591},"ACTION":"nodebalancer.config.update"}`
	params = map[string]string{
		"Algorithm":      `leastconn`,
		"ConfigID":       `14591`,
		"Port":           `90`,
		"Protocol":       `https`,
		"Stickiness":     `table`,
		"api_action":     `nodebalancer.config.update`,
		"api_key":        `foo`,
		"check":          `http_body`,
		"check_attempts": `20`,
		"check_body":     `quux`,
		"check_interval": `24`,
		"check_passive":  `1`,
		"check_path":     `/bar`,
		"check_timeout":  `23`,
		"ssl_cert":       `foo-cert`,
		"ssl_key":        `foo-key`,
	}
	responses = append(responses, newMockAPIResponse("nodebalancer.config.update", params, output))

	return responses
}

func TestNodeBalancerConfigUpdateOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockNodeBalancerConfigUpdateOK()))
	defer ts.Close()

	nbcuo := NodeBalancerConfigUpdateOpts{
		Port:          Int(90),
		Protocol:      String("https"),
		Algorithm:     String("leastconn"),
		Stickiness:    String("table"),
		Check:         String("http_body"),
		CheckInterval: Int(24),
		CheckTimeout:  Int(23),
		CheckAttempts: Int(20),
		CheckPath:     String("/bar"),
		CheckBody:     String("quux"),
		CheckPassive:  Bool(true),
		SSLCert:       String("foo-cert"),
		SSLKey:        String("foo-key"),
	}

	t.Log("nodebalancer.config.update...")
	err := c.NodeBalancerConfigUpdate(14591, nbcuo)
	require.NoError(t, err)
}

func mockNodeBalancerNodeCreateOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"NodeID":140574},"ACTION":"nodebalancer.node.create"}`
	params = map[string]string{
		"Address":    `192.168.202.68:90`,
		"ConfigID":   `14591`,
		"Label":      `test`,
		"Mode":       `accept`,
		"Weight":     `50`,
		"api_action": `nodebalancer.node.create`,
		"api_key":    `foo`,
	}
	responses = append(responses, newMockAPIResponse("nodebalancer.node.create", params, output))

	return responses
}

func TestNodeBalancerNodeCreateOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockNodeBalancerNodeCreateOK()))
	defer ts.Close()

	nbNodeID, err := c.NodeBalancerNodeCreate(14591, "test", "192.168.202.68:90", Int(50),
		String("accept"))
	require.NoError(t, err)
	require.Equal(t, 140574, nbNodeID)
}

func mockNodeBalancerNodeListOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":[{"WEIGHT":50,"ADDRESS":"192.168.202.68:90","LABEL":"test","NODEID":140574,"MODE":"accept","CONFIGID":14591,"STATUS":"Unknown","NODEBALANCERID":13128}],"ACTION":"nodebalancer.node.list"}`
	params = map[string]string{
		"ConfigID":   `14591`,
		"NodeID":     `140574`,
		"api_action": `nodebalancer.node.list`,
		"api_key":    `foo`,
	}
	responses = append(responses, newMockAPIResponse("nodebalancer.node.list", params, output))

	return responses
}

func TestNodeBalancerNodeListOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockNodeBalancerNodeListOK()))
	defer ts.Close()

	nbNodes, err := c.NodeBalancerNodeList(14591, Int(140574))
	require.NoError(t, err)
	require.Len(t, nbNodes, 1)

	node := nbNodes[0]
	assert.Equal(t, 50, node.Weight)
	assert.Equal(t, "192.168.202.68:90", node.Address)
	assert.Equal(t, "test", node.Label)
	assert.Equal(t, 140574, node.ID)
	assert.Equal(t, "accept", node.Mode)
	assert.Equal(t, 14591, node.ConfigID)
	assert.Equal(t, "Unknown", node.Status)
	assert.Equal(t, 13128, node.NodeBalancerID)
}

func mockNodeBalancerNodeUpdateOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"NodeID":140574},"ACTION":"nodebalancer.node.update"}`
	params = map[string]string{
		"Address":    `192.168.202.68:80`,
		"Label":      `test-2`,
		"Mode":       `reject`,
		"NodeID":     `140574`,
		"Weight":     `60`,
		"api_action": `nodebalancer.node.update`,
		"api_key":    `foo`,
	}
	responses = append(responses, newMockAPIResponse("nodebalancer.node.update", params, output))

	return responses
}

func TestNodeBalancerNodeUpdateOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockNodeBalancerNodeUpdateOK()))
	defer ts.Close()

	err := c.NodeBalancerNodeUpdate(140574, String("test-2"), String("192.168.202.68:80"),
		Int(60), String("reject"))
	require.NoError(t, err)
}

func mockNodeBalancerNodeDeleteOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"NodeID":140574},"ACTION":"nodebalancer.node.delete"}`
	params = map[string]string{
		"NodeID":     `140574`,
		"api_action": `nodebalancer.node.delete`,
		"api_key":    `foo`,
	}
	responses = append(responses, newMockAPIResponse("nodebalancer.node.delete", params, output))

	return responses
}

func TestNodeBalancerNodeDeleteOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockNodeBalancerNodeDeleteOK()))
	defer ts.Close()

	err := c.NodeBalancerNodeDelete(140574)
	require.NoError(t, err)
}

func mockNodeBalancerConfigDeleteOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"ConfigID":14591},"ACTION":"nodebalancer.config.delete"}`
	params = map[string]string{
		"ConfigID":       `14591`,
		"NodeBalancerID": `13128`,
		"api_action":     `nodebalancer.config.delete`,
		"api_key":        `foo`,
	}
	responses = append(responses, newMockAPIResponse("nodebalancer.config.delete", params, output))

	return responses
}

func TestNodeBalancerConfigDeleteOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockNodeBalancerConfigDeleteOK()))
	defer ts.Close()

	err := c.NodeBalancerConfigDelete(13128, 14591)
	require.NoError(t, err)
}

func mockNodeBalancerDeleteOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"NodeBalancerID":13128},"ACTION":"nodebalancer.delete"}`
	params = map[string]string{
		"NodeBalancerID": `13128`,
		"api_action":     `nodebalancer.delete`,
		"api_key":        `foo`,
	}
	responses = append(responses, newMockAPIResponse("nodebalancer.delete", params, output))

	return responses
}

func TestNodeBalancerDeleteOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockNodeBalancerDeleteOK()))
	defer ts.Close()

	err := c.NodeBalancerDelete(13128)
	require.NoError(t, err)
}
