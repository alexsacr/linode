// +build !integration

package linode

import (
	"testing"

	"github.com/alexsacr/linode/_third_party/testify/assert"
	"github.com/alexsacr/linode/_third_party/testify/require"
)

func mockDomainCreateOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"DomainID":716074},"ACTION":"domain.create"}`
	params = map[string]string{
		"Description":      `foo`,
		"Domain":           `foo.com`,
		"Expire_sec":       `40`,
		"Refresh_sec":      `20`,
		"Retry_sec":        `30`,
		"SOA_Email":        `foo@foo.com`,
		"TTL_sec":          `50`,
		"Type":             `master`,
		"api_action":       `domain.create`,
		"api_key":          `foo`,
		"lpm_displayGroup": `test`,
		"status":           `1`,
	}
	responses = append(responses, newMockAPIResponse("domain.create", params, output))

	return responses
}

func TestDomainCreateOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockDomainCreateOK()))
	defer ts.Close()

	dco := DomainCreateOpts{
		Description:  String("foo"),
		SOAEmail:     String("foo@foo.com"),
		RefreshSec:   Int(20),
		RetrySec:     Int(30),
		ExpireSec:    Int(40),
		TTLSec:       Int(50),
		DisplayGroup: String("test"),
		Status:       Int(1),
	}

	dID, err := c.DomainCreate("foo.com", "master", dco)
	require.NoError(t, err)
	require.Equal(t, 716074, dID)
}

func mockDomainListOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":[{"DOMAINID":716074,"DESCRIPTION":"foo","EXPIRE_SEC":300,"RETRY_SEC":300,"STATUS":1,"LPM_DISPLAYGROUP":"test","MASTER_IPS":"1;2;3;","REFRESH_SEC":300,"SOA_EMAIL":"foo@foo.com","TTL_SEC":300,"DOMAIN":"foo.com","TYPE":"master","AXFR_IPS":"none"}],"ACTION":"domain.list"}`
	params = map[string]string{
		"DomainID":   `716074`,
		"api_action": `domain.list`,
		"api_key":    `foo`,
	}
	responses = append(responses, newMockAPIResponse("domain.list", params, output))

	return responses
}

func TestDomainListOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockDomainListOK()))
	defer ts.Close()

	domains, err := c.DomainList(Int(716074))
	require.NoError(t, err)
	require.Len(t, domains, 1)
	d := domains[0]

	assert.Equal(t, 716074, d.ID)
	assert.Equal(t, "foo", d.Description)
	assert.Equal(t, 300, d.ExpireSec)
	assert.Equal(t, 300, d.RetrySec)
	assert.Equal(t, 1, d.Status)
	assert.Equal(t, "test", d.DisplayGroup)
	assert.Equal(t, "1;2;3;", d.MasterIPs)
	assert.Equal(t, 300, d.RefreshSec)
	assert.Equal(t, "foo@foo.com", d.SOAEmail)
	assert.Equal(t, 300, d.TTLSec)
	assert.Equal(t, "foo.com", d.Domain)
	assert.Equal(t, "master", d.Type)
	assert.Equal(t, "", d.AXFRIPs)
}

func mockDomainUpdateOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"DomainID":716074},"ACTION":"domain.update"}`
	params = map[string]string{
		"Domain":           `baz.com`,
		"DomainID":         `716074`,
		"Expire_sec":       `3600`,
		"Refresh_sec":      `3600`,
		"Retry_sec":        `3600`,
		"SOA_Email":        `baz@baz.com`,
		"TTL_sec":          `3600`,
		"Type":             `master`,
		"api_action":       `domain.update`,
		"api_key":          `foo`,
		"lpm_displayGroup": `still-testing`,
		"status":           `2`,
	}
	responses = append(responses, newMockAPIResponse("domain.update", params, output))

	return responses
}

func TestDomainUpdateOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockDomainUpdateOK()))
	defer ts.Close()

	duo := DomainUpdateOpts{
		Domain:       String("baz.com"),
		Type:         String("master"),
		SOAEmail:     String("baz@baz.com"),
		RefreshSec:   Int(3600),
		RetrySec:     Int(3600),
		ExpireSec:    Int(3600),
		TTLSec:       Int(3600),
		DisplayGroup: String("still-testing"),
		Status:       Int(2),
	}

	err := c.DomainUpdate(716074, duo)
	require.NoError(t, err)
}

func mockDomainResourceCreateOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"ResourceID":5337468},"ACTION":"domain.resource.create"}`
	params = map[string]string{
		"DomainID":   `716074`,
		"Name":       `_foo._tcp`,
		"Port":       `15`,
		"Priority":   `5`,
		"Protocol":   `bar`,
		"TTL_sec":    `20`,
		"Target":     `bar.baz.com`,
		"Type":       `srv`,
		"Weight":     `10`,
		"api_action": `domain.resource.create`,
		"api_key":    `foo`,
	}
	responses = append(responses, newMockAPIResponse("domain.resource.create", params, output))

	return responses
}

func TestDomainResourceCreateOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockDomainResourceCreateOK()))
	defer ts.Close()

	drco := DomainResourceCreateOpts{
		Name:     String("_foo._tcp"),
		Target:   String("bar.baz.com"),
		Priority: Int(5),
		Weight:   Int(10),
		Port:     Int(15),
		Protocol: String("bar"),
		TTLSec:   Int(20),
	}

	rID, err := c.DomainResourceCreate(716074, "srv", drco)
	require.NoError(t, err)
	require.Equal(t, 5337468, rID)
}

func mockDomainResourceListOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":[{"DOMAINID":716074,"PORT":15,"RESOURCEID":5337468,"NAME":"_foo._tcp","WEIGHT":10,"TTL_SEC":300,"TARGET":"bar.baz.com","PRIORITY":5,"PROTOCOL":"tcp","TYPE":"srv"}],"ACTION":"domain.resource.list"}`
	params = map[string]string{
		"DomainID":   `716074`,
		"ResourceID": `5337468`,
		"api_action": `domain.resource.list`,
		"api_key":    `foo`,
	}
	responses = append(responses, newMockAPIResponse("domain.resource.list", params, output))

	return responses
}

func TestDomainResourceListOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockDomainResourceListOK()))
	defer ts.Close()

	resources, err := c.DomainResourceList(716074, Int(5337468))
	require.NoError(t, err)
	require.Len(t, resources, 1)
	r := resources[0]

	assert.Equal(t, 716074, r.DomainID)
	assert.Equal(t, 15, r.Port)
	assert.Equal(t, 5337468, r.ID)
	assert.Equal(t, "_foo._tcp", r.Name)
	assert.Equal(t, 10, r.Weight)
	assert.Equal(t, 300, r.TTLSec)
	assert.Equal(t, "bar.baz.com", r.Target)
	assert.Equal(t, 5, r.Priority)
	assert.Equal(t, "tcp", r.Protocol)
	assert.Equal(t, "srv", r.Type)
}

func mockDomainResourceUpdateOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"ResourceID":5337468},"ACTION":"domain.resource.update"}`
	params = map[string]string{
		"DomainID":   `716074`,
		"Name":       `_qux._udp`,
		"Port":       `30`,
		"Priority":   `20`,
		"Protocol":   `udp`,
		"ResourceID": `5337468`,
		"TTL_sec":    `301`,
		"Target":     `qux.baz.com`,
		"Weight":     `25`,
		"api_action": `domain.resource.update`,
		"api_key":    `foo`,
	}
	responses = append(responses, newMockAPIResponse("domain.resource.update", params, output))

	return responses
}

func TestDomainResourceUpdateOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockDomainResourceUpdateOK()))
	defer ts.Close()

	druo := DomainResourceUpdateOpts{
		DomainID: Int(716074),
		Name:     String("_qux._udp"),
		Target:   String("qux.baz.com"),
		Priority: Int(20),
		Weight:   Int(25),
		Port:     Int(30),
		Protocol: String("udp"),
		TTLSec:   Int(301),
	}

	err := c.DomainResourceUpdate(5337468, druo)
	require.NoError(t, err)
}

func mockDomainResourceDeleteOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"ResourceID":5337468},"ACTION":"domain.resource.delete"}`
	params = map[string]string{
		"DomainID":   `716074`,
		"ResourceID": `5337468`,
		"api_action": `domain.resource.delete`,
		"api_key":    `foo`,
	}
	responses = append(responses, newMockAPIResponse("domain.resource.delete", params, output))

	return responses
}

func TestDomainResourceDeleteOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockDomainResourceDeleteOK()))
	defer ts.Close()

	err := c.DomainResourceDelete(716074, 5337468)
	require.NoError(t, err)
}

func mockDomainDeleteOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"DomainID":716074},"ACTION":"domain.delete"}`
	params = map[string]string{
		"DomainID":   `716074`,
		"api_action": `domain.delete`,
		"api_key":    `foo`,
	}
	responses = append(responses, newMockAPIResponse("domain.delete", params, output))

	return responses
}

func TestDomainDeleteOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockDomainDeleteOK()))
	defer ts.Close()

	err := c.DomainDelete(716074)
	require.NoError(t, err)
}
