// +build !integration

package linode

import (
	"testing"

	"github.com/alexsacr/linode/_third_party/testify/assert"
	"github.com/alexsacr/linode/_third_party/testify/require"
)

func mockStackScriptCreateOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"StackScriptID":12567},"ACTION":"stackscript.create"}`
	params = map[string]string{
		"Description":        `foo`,
		"DistributionIDList": `130,`,
		"Label":              `test`,
		"api_action":         `stackscript.create`,
		"api_key":            `foo`,
		"isPublic":           `0`,
		"rev_note":           `bar`,
		"script":             `#! /bin/bash foo`,
	}
	responses = append(responses, newMockAPIResponse("stackscript.create", params, output))

	return responses
}

func TestStackScriptCreateOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockStackScriptCreateOK()))
	defer ts.Close()

	ssID, err := c.StackScriptCreate("test", "130,", "#! /bin/bash foo", String("foo"),
		Bool(false), String("bar"))
	require.NoError(t, err)
	require.Equal(t, 12567, ssID)
}

func mockStackScriptCreatePublicOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"StackScriptID":12567},"ACTION":"stackscript.create"}`
	params = map[string]string{
		"Description":        `foo`,
		"DistributionIDList": `130,`,
		"Label":              `test`,
		"api_action":         `stackscript.create`,
		"api_key":            `foo`,
		"isPublic":           `1`,
		"rev_note":           `bar`,
		"script":             `#! /bin/bash foo`,
	}
	responses = append(responses, newMockAPIResponse("stackscript.create", params, output))

	return responses
}

func TestStackScriptCreatePublicOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockStackScriptCreatePublicOK()))
	defer ts.Close()

	ssID, err := c.StackScriptCreate("test", "130,", "#! /bin/bash foo", String("foo"),
		Bool(true), String("bar"))
	require.NoError(t, err)
	require.Equal(t, 12567, ssID)
}

func mockStackScriptListOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":[{"REV_NOTE":"bar","SCRIPT":"#! \/bin\/bash foo","DISTRIBUTIONIDLIST":"130,","DESCRIPTION":"foo","REV_DT":"2015-07-08 23:46:50.0","LABEL":"test","DEPLOYMENTSTOTAL":1,"LATESTREV":60999,"STACKSCRIPTID":12567,"ISPUBLIC":1,"DEPLOYMENTSACTIVE":2,"CREATE_DT":"2015-07-08 23:46:50.0","USERID":1337}],"ACTION":"stackscript.list"}`
	params = map[string]string{
		"StackScriptID": `12567`,
		"api_action":    `stackscript.list`,
		"api_key":       `foo`,
	}
	responses = append(responses, newMockAPIResponse("stackscript.list", params, output))

	return responses
}

func TestStackScriptListOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockStackScriptListOK()))
	defer ts.Close()

	ssList, err := c.StackScriptList(Int(12567))
	require.NoError(t, err)
	require.Len(t, ssList, 1)

	ss := ssList[0]
	assert.Equal(t, "bar", ss.RevNote)
	assert.Equal(t, "#! /bin/bash foo", ss.Script)
	assert.Equal(t, "130,", ss.DistIDList)
	assert.Equal(t, "2015-07-08 23:46:50.0", ss.RevDT)
	assert.Equal(t, "test", ss.Label)
	assert.Equal(t, 1, ss.TotalDeploys)
	assert.Equal(t, 60999, ss.LatestRev)
	assert.Equal(t, 12567, ss.ID)
	assert.True(t, ss.IsPublic)
	assert.Equal(t, 2, ss.ActiveDeploys)
	assert.Equal(t, "2015-07-08 23:46:50.0", ss.CreateDT)
	assert.Equal(t, 1337, ss.UserID)
}

func mockStackScriptUpdateOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"StackScriptID":12567},"ACTION":"stackscript.update"}`
	params = map[string]string{
		"Description":        `quux`,
		"DistributionIDList": `129,130`,
		"Label":              `test-2`,
		"StackScriptID":      `12567`,
		"api_action":         `stackscript.update`,
		"api_key":            `foo`,
		"isPublic":           `1`,
		"rev_note":           `baz`,
		"script":             `#! /bin/bash baz`,
	}
	responses = append(responses, newMockAPIResponse("stackscript.update", params, output))

	return responses
}

func TestStackScriptUpdateOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockStackScriptUpdateOK()))
	defer ts.Close()

	err := c.StackScriptUpdate(12567, String("test-2"), String("quux"), String("129,130"),
		Bool(true), String("baz"), String("#! /bin/bash baz"))
	require.NoError(t, err)
}

func mockStackScriptUpdatePrivateOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"StackScriptID":12567},"ACTION":"stackscript.update"}`
	params = map[string]string{
		"Description":        `quux`,
		"DistributionIDList": `129,130`,
		"Label":              `test-2`,
		"StackScriptID":      `12567`,
		"api_action":         `stackscript.update`,
		"api_key":            `foo`,
		"isPublic":           `0`,
		"rev_note":           `baz`,
		"script":             `#! /bin/bash baz`,
	}
	responses = append(responses, newMockAPIResponse("stackscript.update", params, output))

	return responses
}

func TestStackScriptUpdatePrivateOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockStackScriptUpdatePrivateOK()))
	defer ts.Close()

	err := c.StackScriptUpdate(12567, String("test-2"), String("quux"), String("129,130"),
		Bool(false), String("baz"), String("#! /bin/bash baz"))
	require.NoError(t, err)
}

func mockStackScriptDeleteOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"StackScriptID":12567},"ACTION":"stackscript.delete"}`
	params = map[string]string{
		"StackScriptID": `12567`,
		"api_action":    `stackscript.delete`,
		"api_key":       `foo`,
	}
	responses = append(responses, newMockAPIResponse("stackscript.delete", params, output))

	return responses
}

func TestStackScriptDeleteOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockStackScriptDeleteOK()))
	defer ts.Close()

	err := c.StackScriptDelete(12567)
	require.NoError(t, err)
}
