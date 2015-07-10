package linode

// ImageDelete maps to the 'image.delete' call.
//
// https://www.linode.com/api/image/image.delete
func (c *Client) ImageDelete(imgID int) error {
	args := make(map[string]interface{})
	args["ImageID"] = imgID

	_, err := c.apiCall("image.delete", args)
	if err != nil {
		return err
	}

	return nil
}

// Image is the API response to the 'image.list' call.
type Image struct {
	CreateDT    string `mapstructure:"CREATE_DT"`
	Creator     string `mapstructure:"CREATOR"`
	Description string `mapstructure:"DESCRIPTION"`
	FSType      string `mapstructure:"FS_TYPE"`
	ID          int    `mapstructure:"IMAGEID"`
	IsPublic    bool   `mapstructure:"ISPUBLIC"`
	Label       string `mapstructure:"LABEL"`
	LastUsedDT  string `mapstructure:"LAST_USED_DT"`
	MinSize     int    `mapstructure:"MINSIZE"`
	Status      string `mapstructure:"STATUS"`
	Type        string `mapstructure:"TYPE"`
}

// ImageList maps to the 'image.list' call.
//
// https://www.linode.com/api/image/image.list
func (c *Client) ImageList(imgID *int, pendingOnly *bool) ([]Image, error) {
	args := make(map[string]interface{})
	args["ImageID"] = imgID

	// special handle weird bool arg for this method
	if pendingOnly != nil {
		if *pendingOnly == true {
			args["pendingOnly"] = 1
		} else {
			args["pendingOnly"] = 0
		}
	}

	data, err := c.apiCall("image.list", args)
	if err != nil {
		return nil, err
	}

	var imgs []Image
	err = unmarshalMultiMap(data, &imgs)
	if err != nil {
		return nil, err
	}

	return imgs, nil
}

// ImageUpdate maps to the 'image.update' call.
//
// https://www.linode.com/api/image/image.update
func (c *Client) ImageUpdate(imgID int, label *string, description *string) error {
	args := make(map[string]interface{})
	args["ImageID"] = imgID
	args["label"] = label
	args["description"] = description

	_, err := c.apiCall("image.update", args)
	if err != nil {
		return err
	}

	return nil
}
