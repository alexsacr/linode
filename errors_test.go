// +build !integration

package linode

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/alexsacr/linode/_third_party/testify/assert"
)

func httpPostError(_ string, _ url.Values) (*http.Response, error) {
	return nil, errors.New("foo")
}

func TestHTTPError(t *testing.T) {
	c := NewClient("foo")
	c.post = httpPostError
	testErrors(t, c, "foo", false)
}

func apiCallerError(_ string, _ map[string]interface{}) (json.RawMessage, error) {
	return nil, errors.New("bar")
}

func TestAPICallerError(t *testing.T) {
	c := NewClient("foo")
	c.apiCall = apiCallerError
	testErrors(t, c, "bar", false)
}

func apiCallerNilJSON(_ string, _ map[string]interface{}) (json.RawMessage, error) {
	return nil, nil
}

func TestAPICallerNilJSON(t *testing.T) {
	c := NewClient("foo")
	c.apiCall = apiCallerNilJSON
	testErrors(t, c, "unexpected end of JSON input", true)
}

func argMarshalerError(_ interface{}) (map[string]interface{}, error) {
	return nil, errors.New("foo")
}

func TestArgMarshalerError(t *testing.T) {
	c := NewClient("foo")
	c.argMarshal = argMarshalerError

	err := c.LinodeUpdate(0, LinodeOpts{})
	assert.Error(t, err)

	_, err = c.LinodeConfigCreate(0, 0, "", "", LinodeConfigCreateOpts{})
	assert.Error(t, err)

	err = c.LinodeConfigUpdate(0, LinodeConfigUpdateOpts{})
	assert.Error(t, err)

	_, err = c.DomainCreate("", "", DomainCreateOpts{})
	assert.Error(t, err)

	err = c.DomainUpdate(0, DomainUpdateOpts{})
	assert.Error(t, err)

	_, err = c.DomainResourceCreate(0, "", DomainResourceCreateOpts{})
	assert.Error(t, err)

	err = c.DomainResourceUpdate(0, DomainResourceUpdateOpts{})
	assert.Error(t, err)

	_, err = c.NodeBalancerConfigCreate(0, NodeBalancerConfigCreateOpts{})
	assert.Error(t, err)

	err = c.NodeBalancerConfigUpdate(0, NodeBalancerConfigUpdateOpts{})
	assert.Error(t, err)
}

func testErrors(t *testing.T, c *Client, msg string, badReturnData bool) {
	errMap := make(map[string]error)

	if !badReturnData {
		errMap["LinodeDelete"] = c.LinodeDelete(0, nil)
		errMap["LinodeResize"] = c.LinodeResize(0, 0)
		errMap["LinodeUpdate"] = c.LinodeUpdate(0, LinodeOpts{})
		errMap["LinodeConfigDelete"] = c.LinodeConfigDelete(0, 0)
		errMap["LinodeConfigUpdate"] = c.LinodeConfigUpdate(0, LinodeConfigUpdateOpts{})
		errMap["LinodeDiskUpdate"] = c.LinodeDiskUpdate(0, 0, nil, nil)
		errMap["LinodeIPSwap"] = c.LinodeIPSwap(0, nil, nil)
		errMap["DomainDelete"] = c.DomainDelete(0)
		errMap["DomainUpdate"] = c.DomainUpdate(0, DomainUpdateOpts{})
		errMap["DomainResourceDelete"] = c.DomainResourceDelete(0, 0)
		errMap["DomainResourceUpdate"] = c.DomainResourceUpdate(0, DomainResourceUpdateOpts{})
		errMap["NodeBalancerDelete"] = c.NodeBalancerDelete(0)
		errMap["NodeBalancerUpdate"] = c.NodeBalancerUpdate(0, nil, nil)
		errMap["NodeBalancerConfigDelete"] = c.NodeBalancerConfigDelete(0, 0)
		errMap["NodeBalancerConfigUpdate"] = c.NodeBalancerConfigUpdate(0, NodeBalancerConfigUpdateOpts{})
		errMap["NodeBalancerNodeDelete"] = c.NodeBalancerNodeDelete(0)
		errMap["NodeBalancerNodeUpdate"] = c.NodeBalancerNodeUpdate(0, nil, nil, nil, nil)
		errMap["ImageDelete"] = c.ImageDelete(0)
		errMap["ImageUpdate"] = c.ImageUpdate(0, nil, nil)
		errMap["StackScriptDelete"] = c.StackScriptDelete(0)
		errMap["StackScriptUpdate"] = c.StackScriptUpdate(0, nil, nil, nil, nil, nil, nil)
	}

	errMap["TestEcho"] = c.TestEcho()
	errMap["WaitForAllJobs"] = c.WaitForAllJobs(0, 1*time.Nanosecond, 1*time.Second)

	_, errMap["AvailDatacenters"] = c.AvailDatacenters()
	_, errMap["AvailDistributions"] = c.AvailDistributions(nil)
	_, errMap["AvailKernels"] = c.AvailKernels(nil, nil)
	_, errMap["AvailLinodePlans"] = c.AvailLinodePlans(nil)
	_, errMap["AvailStackScripts"] = c.AvailStackScripts(nil, nil, nil)
	_, errMap["LinodeBoot"] = c.LinodeBoot(0, nil)
	_, errMap["LinodeClone"] = c.LinodeClone(0, 0, 0, nil, nil)
	_, errMap["LinodeCreate"] = c.LinodeCreate(0, 0, nil)
	_, errMap["LinodeList"] = c.LinodeList(nil)
	_, errMap["LinodeReboot"] = c.LinodeReboot(0, nil)
	_, errMap["LinodeShutdown"] = c.LinodeShutdown(0)
	_, errMap["LinodeConfigCreate"] = c.LinodeConfigCreate(0, 0, "", "", LinodeConfigCreateOpts{})
	_, errMap["LinodeConfigList"] = c.LinodeConfigList(0, nil)
	_, errMap["LinodeDiskDelete"] = c.LinodeDiskDelete(0, 0)
	_, errMap["LinodeDiskList"] = c.LinodeDiskList(0, nil)
	_, errMap["LinodeDiskResize"] = c.LinodeDiskResize(0, 0, 0)
	_, errMap["LinodeIPList"] = c.LinodeIPList(nil, nil)
	_, errMap["LinodeJobList"] = c.LinodeJobList(0, nil, nil)
	_, errMap["WaitForJob"] = c.WaitForJob(0, 0, 1*time.Nanosecond, 1*time.Second)
	_, errMap["DomainCreate"] = c.DomainCreate("", "", DomainCreateOpts{})
	_, errMap["DomainList"] = c.DomainList(nil)
	_, errMap["DomainResourceCreate"] = c.DomainResourceCreate(0, "", DomainResourceCreateOpts{})
	_, errMap["DomainResourceList"] = c.DomainResourceList(0, nil)
	_, errMap["NodeBalancerCreate"] = c.NodeBalancerCreate(0, nil, nil)
	_, errMap["NodeBalancerList"] = c.NodeBalancerList(nil)
	_, errMap["NodeBalancerConfigCreate"] = c.NodeBalancerConfigCreate(0, NodeBalancerConfigCreateOpts{})
	_, errMap["NodeBalancerConfigList"] = c.NodeBalancerConfigList(0, nil)
	_, errMap["NodeBalancerNodeCreate"] = c.NodeBalancerNodeCreate(0, "", "", nil, nil)
	_, errMap["NodeBalancerNodeList"] = c.NodeBalancerNodeList(0, nil)
	_, errMap["AccountEstimateInvoice"] = c.AccountEstimateInvoice("", nil, nil, nil)
	_, errMap["AccountInfo"] = c.AccountInfo()
	_, errMap["UserGetAPIKey"] = c.UserGetAPIKey("", "", nil, nil, nil)
	_, errMap["ImageList"] = c.ImageList(nil, nil)
	_, errMap["StackScriptCreate"] = c.StackScriptCreate("", "", "", nil, nil, nil)
	_, errMap["StackScriptList"] = c.StackScriptList(nil)

	_, _, errMap["LinodeDiskCreate"] = c.LinodeDiskCreate(0, "", "", 0)
	_, _, errMap["LinodeDiskCreateFromDistribution"] = c.LinodeDiskCreateFromDistribution(0, 0, "", 0, "", nil)
	_, _, errMap["LinodeDiskCreateFromImage"] = c.LinodeDiskCreateFromImage(0, 0, "", nil, nil, nil)
	_, _, errMap["LinodeDiskCreateFromStackScript"] = c.LinodeDiskCreateFromStackScript(0, 0, "", 0, "", 0, "", nil)
	_, _, errMap["LinodeDiskDuplicate"] = c.LinodeDiskDuplicate(0, 0)
	_, _, errMap["LinodeDiskImagize"] = c.LinodeDiskImagize(0, 0, nil, nil)
	_, _, errMap["LinodeIPAddPrivate"] = c.LinodeIPAddPrivate(0)

	var ok int

	for fn, err := range errMap {
		assert.Error(t, err, fmt.Sprintf("%s - iteration %d\n", fn, ok))
		if err != nil {
			assert.Equal(t, msg, err.Error(), fmt.Sprintf("%s - iteration %d\n", fn, ok))
		}
		ok++
	}
}
