// +build !integration

package linode

import (
	"testing"

	"github.com/alexsacr/linode/_third_party/testify/assert"
	"github.com/alexsacr/linode/_third_party/testify/require"
)

func mockImageListOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":[{"LAST_USED_DT":"baz","MINSIZE":600,"DESCRIPTION":"foo","LABEL":"bar","CREATOR":"quux","STATUS":"available","ISPUBLIC":1,"CREATE_DT":"2015-07-07 23:55:59.0","TYPE":"manual","FS_TYPE":"ext4","IMAGEID":402716}],"ACTION":"image.list"}`
	params = map[string]string{
		"ImageID":     `402716`,
		"api_action":  `image.list`,
		"api_key":     `foo`,
		"pendingOnly": `0`,
	}
	responses = append(responses, newMockAPIResponse("image.list", params, output))

	return responses
}

func TestImageListOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockImageListOK()))
	defer ts.Close()

	imgs, err := c.ImageList(Int(402716), Bool(false))
	require.NoError(t, err)
	require.Len(t, imgs, 1)

	i := imgs[0]
	assert.Equal(t, "baz", i.LastUsedDT)
	assert.Equal(t, 600, i.MinSize)
	assert.Equal(t, "foo", i.Description)
	assert.Equal(t, "bar", i.Label)
	assert.Equal(t, "quux", i.Creator)
	assert.Equal(t, "available", i.Status)
	assert.True(t, i.IsPublic)
	assert.Equal(t, "2015-07-07 23:55:59.0", i.CreateDT)
	assert.Equal(t, "manual", i.Type)
	assert.Equal(t, "ext4", i.FSType)
	assert.Equal(t, 402716, i.ID)
}

func mockImageListPendingOnlyOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":[],"ACTION":"image.list"}`
	params = map[string]string{
		"ImageID":     `402716`,
		"api_action":  `image.list`,
		"api_key":     `foo`,
		"pendingOnly": `1`,
	}
	responses = append(responses, newMockAPIResponse("image.list", params, output))

	return responses
}

func TestImageListPendingOnlyOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockImageListPendingOnlyOK()))
	defer ts.Close()

	imgs, err := c.ImageList(Int(402716), Bool(true))
	require.NoError(t, err)
	require.Len(t, imgs, 0)
}

func mockImageUpdateOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"last_used_dt":"","minSize":600,"description":"quux","label":"baz","creator":"foo","status":"available","isPublic":0,"create_dt":"2015-07-07 23:55:59.0","type":"manual","fs_type":"ext4","imageID":402716},"ACTION":"image.update"}`
	params = map[string]string{
		"ImageID":     `402716`,
		"api_action":  `image.update`,
		"api_key":     `foo`,
		"description": `quux`,
		"label":       `baz`,
	}
	responses = append(responses, newMockAPIResponse("image.update", params, output))

	return responses
}

func TestLinodeImageUpdateOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockImageUpdateOK()))
	defer ts.Close()

	err := c.ImageUpdate(402716, String("baz"), String("quux"))
	require.NoError(t, err)
}

func mockImageDeleteOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"last_used_dt":"","minSize":600,"description":"quux","label":"baz","creator":"foo","status":"deleted","isPublic":0,"create_dt":"2015-07-07 23:55:59.0","type":"manual","fs_type":"ext4","imageID":402716},"ACTION":"image.delete"}`
	params = map[string]string{
		"ImageID":    `402716`,
		"api_action": `image.delete`,
		"api_key":    `foo`,
	}
	responses = append(responses, newMockAPIResponse("image.delete", params, output))

	return responses
}

func TestLinodeImageDeleteOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockImageDeleteOK()))
	defer ts.Close()

	err := c.ImageDelete(402716)
	require.NoError(t, err)
}
