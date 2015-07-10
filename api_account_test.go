// +build !integration

package linode

import (
	"testing"

	"github.com/alexsacr/linode/_third_party/testify/assert"
	"github.com/alexsacr/linode/_third_party/testify/require"
)

func mockAccountEstimateInvoiceOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"INVOICE_TO":"2015-07-31 23:59:59","AMOUNT":8.07},"ACTION":"account.estimateinvoice"}`
	params = map[string]string{
		"PaymentTerm": "1",
		"PlanID":      "1",
		"LinodeID":    "1",
		"api_action":  "account.estimateinvoice",
		"api_key":     "foo",
		"mode":        "linode_new",
	}
	responses = append(responses, newMockAPIResponse("account.estimateinvoice", params, output))

	return responses
}

func TestAccountEstimateInvoiceOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockAccountEstimateInvoiceOK()))
	defer ts.Close()

	inv, err := c.AccountEstimateInvoice("linode_new", Int(1), Int(1), Int(1))
	require.NoError(t, err)
	assert.Equal(t, "2015-07-31 23:59:59", inv.InvoiceTo)
	assert.Equal(t, 8.07, inv.Price)
}

func mockAccountInfo() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"TRANSFER_USED":1,"BALANCE":-4.8200,"TRANSFER_BILLABLE":2,"BILLING_METHOD":"metered","TRANSFER_POOL":104,"ACTIVE_SINCE":"2015-06-25 23:52:06.0","MANAGED":true},"ACTION":"account.info"}`
	params = map[string]string{
		"api_action": "account.info",
		"api_key":    "foo",
	}
	responses = append(responses, newMockAPIResponse("account.info", params, output))

	return responses
}

func TestMockAccountInfoOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockAccountInfo()))
	defer ts.Close()

	i, err := c.AccountInfo()
	require.NoError(t, err)
	assert.Equal(t, 1, i.TransferUsed)
	assert.Equal(t, -4.8200, i.Balance)
	assert.Equal(t, 2, i.TransferBillable)
	assert.Equal(t, "metered", i.BillingMethod)
	assert.Equal(t, 104, i.TransferPool)
	assert.Equal(t, "2015-06-25 23:52:06.0", i.ActiveSince)
	assert.True(t, i.Managed)
}

func mockUserGetAPIKeyOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"USERNAME":"","API_KEY":"foo"},"ACTION":"user.getAPIKey"}`
	params = map[string]string{
		"api_action": `user.getAPIKey`,
		"api_key":    `foo`,
		"expires":    `1`,
		"token":      `baz`,
		"label":      `test`,
		"password":   `bar`,
		"username":   `foo`,
	}
	responses = append(responses, newMockAPIResponse("user.getAPIKey", params, output))

	return responses
}

func TestUserGetAPIKeyOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockUserGetAPIKeyOK()))
	defer ts.Close()

	key, err := c.UserGetAPIKey("foo", "bar", String("baz"), Int(1), String("test"))
	require.NoError(t, err)
	require.Equal(t, "foo", key)
}
