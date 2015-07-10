package linode

// StackScriptCreate maps to the 'stackscript.create' call.
//
// https://www.linode.com/api/stackscript/stackscript.create
func (c *Client) StackScriptCreate(label string, distIDList string, script string,
	description *string, isPublic *bool, revNote *string) (ssID int, err error) {

	args := make(map[string]interface{})
	args["Label"] = label
	args["Description"] = description
	args["DistributionIDList"] = distIDList
	args["rev_note"] = revNote
	args["script"] = script

	if isPublic != nil {
		if *isPublic == true {
			args["isPublic"] = 1
		} else {
			args["isPublic"] = 0
		}
	}

	data, err := c.apiCall("stackscript.create", args)
	if err != nil {
		return 0, err
	}

	err = unmarshalSingle(data, "StackScriptID", &ssID)
	if err != nil {
		return 0, err
	}

	return ssID, nil
}

// StackScriptDelete maps to the 'stackscript.delete' call.
//
// https://www.linode.com/api/stackscript/stackscript.delete
func (c *Client) StackScriptDelete(ssID int) error {
	args := make(map[string]interface{})
	args["StackScriptID"] = ssID

	_, err := c.apiCall("stackscript.delete", args)
	if err != nil {
		return err
	}

	return nil
}

// StackScript is the API response to the 'stackscript.list' and
// 'avail.stackscripts' calls.
type StackScript struct {
	Script        string `mapstructure:"SCRIPT"`
	Description   string `mapstructure:"DESCRIPTION"`
	DistIDList    string `mapstructure:"DISTRIBUTIONIDLIST"`
	Label         string `mapstructure:"LABEL"`
	TotalDeploys  int    `mapstructure:"DEPLOYMENTSTOTAL"`
	LatestRev     int    `mapstructure:"LATESTREV"`
	CreateDT      string `mapstructure:"CREATE_DT"`
	ActiveDeploys int    `mapstructure:"DEPLOYMENTSACTIVE"`
	ID            int    `mapstructure:"STACKSCRIPTID"`
	RevNote       string `mapstructure:"REV_NOTE"`
	RevDT         string `mapstructure:"REV_DT"`
	IsPublic      bool   `mapstructure:"ISPUBLIC"`
	UserID        int    `mapstructure:"USERID"`
}

// StackScriptList maps to the 'stackscript.list' call.
//
// https://www.linode.com/api/stackscript/stackscript.list
func (c *Client) StackScriptList(ssID *int) ([]StackScript, error) {
	args := make(map[string]interface{})
	args["StackScriptID"] = ssID

	data, err := c.apiCall("stackscript.list", args)
	if err != nil {
		return nil, err
	}

	var ss []StackScript
	err = unmarshalMultiMap(data, &ss)
	if err != nil {
		return nil, err
	}

	return ss, nil
}

// StackScriptUpdate maps to the 'stackscript.update' call.
//
// https://www.linode.com/api/stackscript/stackscript.update
func (c *Client) StackScriptUpdate(ssID int, label *string, description *string,
	distIDList *string, isPublic *bool, revNote *string, script *string) error {

	args := make(map[string]interface{})
	args["StackScriptID"] = ssID
	args["Label"] = label
	args["Description"] = description
	args["DistributionIDList"] = distIDList
	args["rev_note"] = revNote
	args["script"] = script

	if isPublic != nil {
		if *isPublic == true {
			args["isPublic"] = 1
		} else {
			args["isPublic"] = 0
		}
	}

	_, err := c.apiCall("stackscript.update", args)
	if err != nil {
		return err
	}

	return nil
}
