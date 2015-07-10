package linode

import "encoding/json"

// EstimatedInvoice is the API response to the 'account.estimateinvoice' call.
type EstimatedInvoice struct {
	InvoiceTo string  `json:"INVOICE_TO"`
	Price     float64 `json:"AMOUNT"`
}

// AccountEstimateInvoice maps to the 'account.estimateinvoice' call.
//
// https://www.linode.com/api/account/account.estimateinvoice
func (c *Client) AccountEstimateInvoice(mode string, term *int, planID *int,
	linodeID *int) (EstimatedInvoice, error) {

	args := make(map[string]interface{})
	args["mode"] = mode
	args["PaymentTerm"] = term
	args["PlanID"] = planID
	args["LinodeID"] = linodeID

	data, err := c.apiCall("account.estimateinvoice", args)
	if err != nil {
		return EstimatedInvoice{}, err
	}

	var inv EstimatedInvoice
	err = json.Unmarshal(data, &inv)
	if err != nil {
		return EstimatedInvoice{}, err
	}

	return inv, nil
}

// AccInfo is the API response to the 'account.info' call.
type AccInfo struct {
	ActiveSince      string  `json:"ACTIVE_SINCE"`
	TransferPool     int     `json:"TRANSFER_POOL"`
	TransferUsed     int     `json:"TRANSFER_USED"`
	TransferBillable int     `json:"TRANSFER_BILLABLE"`
	Managed          bool    `json:"MANAGED"`
	Balance          float64 `json:"BALANCE"`
	BillingMethod    string  `json:"BILLING_METHOD"`
}

// AccountInfo maps to the 'account.info' call.
//
// https://www.linode.com/api/account/account.info
func (c *Client) AccountInfo() (AccInfo, error) {
	data, err := c.apiCall("account.info", nil)
	if err != nil {
		return AccInfo{}, err
	}

	var info AccInfo
	err = json.Unmarshal(data, &info)
	if err != nil {
		return AccInfo{}, err
	}

	return info, nil
}

// UserGetAPIKey maps to the 'user.getapikey' call.  It is used to generate an
// API key from your Linode Manager credentials.
//
// Token is required if you have enabled 2-factor auth.
//
// Expires is the number of hours an API key will be valid for, from 0-8760
// hours.  The default (if you pass in nil) is 168 hours.  If you pass 0, the
// key will never expire.
//
// https://www.linode.com/api/account/user.getapikey
func (c *Client) UserGetAPIKey(username string, password string, token *string, expires *int,
	label *string) (apiKey string, err error) {

	args := make(map[string]interface{})
	args["username"] = username
	args["password"] = password
	args["token"] = token
	args["expires"] = expires
	args["label"] = label

	data, err := c.apiCall("user.getAPIKey", args)
	if err != nil {
		return "", err
	}

	err = unmarshalSingle(data, "API_KEY", &apiKey)
	if err != nil {
		return "", err
	}

	return apiKey, nil
}
