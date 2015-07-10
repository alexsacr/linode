// +build !integration

package linode

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/alexsacr/linode/_third_party/testify/assert"
	"github.com/alexsacr/linode/_third_party/testify/require"
)

func mockAvailDatacentersOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":[{"LOCATION":"Dallas, TX, USA","DATACENTERID":2,"ABBR":"dallas"},{"LOCATION":"Fremont, CA, USA","DATACENTERID":3,"ABBR":"fremont"},{"LOCATION":"Atlanta, GA, USA","DATACENTERID":4,"ABBR":"atlanta"},{"LOCATION":"Newark, NJ, USA","DATACENTERID":6,"ABBR":"newark"},{"LOCATION":"London, England, UK","DATACENTERID":7,"ABBR":"london"},{"LOCATION":"Tokyo, JP","DATACENTERID":8,"ABBR":"tokyo"},{"LOCATION":"Singapore, SG","DATACENTERID":9,"ABBR":"singapore"}],"ACTION":"avail.datacenters"}`
	params = map[string]string{
		"api_action": "avail.datacenters",
		"api_key":    "foo",
	}
	responses = append(responses, newMockAPIResponse("avail.datacenters", params, output))

	return responses
}

func TestAvailDatacentersOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockAvailDatacentersOK()))
	defer ts.Close()

	datacenters, err := c.AvailDatacenters()
	require.NoError(t, err)
	require.NotEmpty(t, datacenters)

	expected := []Datacenter{
		{
			Location: "Dallas, TX, USA",
			ID:       2,
			Abbr:     "dallas",
		},
		{
			Location: "Fremont, CA, USA",
			ID:       3,
			Abbr:     "fremont",
		},
		{
			Location: "Atlanta, GA, USA",
			ID:       4,
			Abbr:     "atlanta",
		},
		{
			Location: "Newark, NJ, USA",
			ID:       6,
			Abbr:     "newark",
		},
		{
			Location: "London, England, UK",
			ID:       7,
			Abbr:     "london",
		},
		{
			Location: "Tokyo, JP",
			ID:       8,
			Abbr:     "tokyo",
		},
		{
			Location: "Singapore, SG",
			ID:       9,
			Abbr:     "singapore",
		},
	}

	for i, eDC := range expected {
		for _, aDC := range datacenters {
			if reflect.DeepEqual(aDC, eDC) {
				expected[i] = Datacenter{}
				break
			}
		}
	}

	for _, dc := range expected {
		assert.Equal(t, Datacenter{}, dc, fmt.Sprintf("%+v not returned.", dc))
	}
}

func mockAvailDistributionsOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":[{"REQUIRESPVOPSKERNEL":1,"DISTRIBUTIONID":138,"IS64BIT":1,"LABEL":"Arch Linux 2015.02","MINIMAGESIZE":800,"CREATE_DT":"2015-02-20 14:17:16.0"},{"REQUIRESPVOPSKERNEL":1,"DISTRIBUTIONID":129,"IS64BIT":1,"LABEL":"CentOS 7","MINIMAGESIZE":750,"CREATE_DT":"2014-07-08 10:07:21.0"},{"REQUIRESPVOPSKERNEL":1,"DISTRIBUTIONID":130,"IS64BIT":1,"LABEL":"Debian 7","MINIMAGESIZE":600,"CREATE_DT":"2014-09-24 13:59:32.0"},{"REQUIRESPVOPSKERNEL":1,"DISTRIBUTIONID":140,"IS64BIT":1,"LABEL":"Debian 8.1","MINIMAGESIZE":900,"CREATE_DT":"2015-04-27 16:26:41.0"},{"REQUIRESPVOPSKERNEL":1,"DISTRIBUTIONID":141,"IS64BIT":1,"LABEL":"Fedora 22","MINIMAGESIZE":650,"CREATE_DT":"2015-05-26 13:50:58.0"},{"REQUIRESPVOPSKERNEL":1,"DISTRIBUTIONID":137,"IS64BIT":1,"LABEL":"Gentoo 2014.12","MINIMAGESIZE":2000,"CREATE_DT":"2014-12-24 18:00:09.0"},{"REQUIRESPVOPSKERNEL":1,"DISTRIBUTIONID":135,"IS64BIT":1,"LABEL":"openSUSE 13.2","MINIMAGESIZE":700,"CREATE_DT":"2014-12-17 17:55:42.0"},{"REQUIRESPVOPSKERNEL":1,"DISTRIBUTIONID":117,"IS64BIT":1,"LABEL":"Slackware 14.1","MINIMAGESIZE":875,"CREATE_DT":"2013-11-25 11:11:02.0"},{"REQUIRESPVOPSKERNEL":1,"DISTRIBUTIONID":124,"IS64BIT":1,"LABEL":"Ubuntu 14.04 LTS","MINIMAGESIZE":750,"CREATE_DT":"2014-04-17 15:42:07.0"},{"REQUIRESPVOPSKERNEL":1,"DISTRIBUTIONID":139,"IS64BIT":1,"LABEL":"Ubuntu 15.04","MINIMAGESIZE":1200,"CREATE_DT":"2015-04-23 11:08:05.0"},{"REQUIRESPVOPSKERNEL":1,"DISTRIBUTIONID":132,"IS64BIT":1,"LABEL":"Arch Linux 2014.10","MINIMAGESIZE":600,"CREATE_DT":"2014-10-06 15:32:20.0"},{"REQUIRESPVOPSKERNEL":1,"DISTRIBUTIONID":60,"IS64BIT":1,"LABEL":"CentOS 5.6","MINIMAGESIZE":950,"CREATE_DT":"2009-08-17 00:00:00.0"},{"REQUIRESPVOPSKERNEL":1,"DISTRIBUTIONID":127,"IS64BIT":1,"LABEL":"CentOS 6.5","MINIMAGESIZE":675,"CREATE_DT":"2014-04-28 15:19:34.0"},{"REQUIRESPVOPSKERNEL":1,"DISTRIBUTIONID":78,"IS64BIT":1,"LABEL":"Debian 6","MINIMAGESIZE":550,"CREATE_DT":"2011-02-08 16:54:31.0"},{"REQUIRESPVOPSKERNEL":1,"DISTRIBUTIONID":122,"IS64BIT":1,"LABEL":"Fedora 20","MINIMAGESIZE":1536,"CREATE_DT":"2013-01-27 10:00:00.0"},{"REQUIRESPVOPSKERNEL":1,"DISTRIBUTIONID":134,"IS64BIT":1,"LABEL":"Fedora 21","MINIMAGESIZE":650,"CREATE_DT":"2014-12-10 16:56:28.0"},{"REQUIRESPVOPSKERNEL":1,"DISTRIBUTIONID":118,"IS64BIT":1,"LABEL":"Gentoo 2013-11-26","MINIMAGESIZE":3072,"CREATE_DT":"2013-11-26 15:20:31.0"},{"REQUIRESPVOPSKERNEL":1,"DISTRIBUTIONID":120,"IS64BIT":1,"LABEL":"openSUSE 13.1","MINIMAGESIZE":1024,"CREATE_DT":"2013-12-02 12:53:29.0"},{"REQUIRESPVOPSKERNEL":1,"DISTRIBUTIONID":87,"IS64BIT":1,"LABEL":"Slackware 13.37","MINIMAGESIZE":600,"CREATE_DT":"2011-06-05 15:11:59.0"},{"REQUIRESPVOPSKERNEL":1,"DISTRIBUTIONID":126,"IS64BIT":1,"LABEL":"Ubuntu 12.04 LTS","MINIMAGESIZE":550,"CREATE_DT":"2014-04-28 14:16:59.0"},{"REQUIRESPVOPSKERNEL":1,"DISTRIBUTIONID":133,"IS64BIT":1,"LABEL":"Ubuntu 14.10","MINIMAGESIZE":650,"CREATE_DT":"2014-10-24 15:48:04.0"},{"REQUIRESPVOPSKERNEL":1,"DISTRIBUTIONID":86,"IS64BIT":0,"LABEL":"Slackware 13.37 32bit","MINIMAGESIZE":600,"CREATE_DT":"2011-06-05 15:11:59.0"}],"ACTION":"avail.distributions"}`
	params = map[string]string{
		"api_action": "avail.distributions",
		"api_key":    "foo",
	}
	responses = append(responses, newMockAPIResponse("avail.distributions", params, output))

	return responses
}

func TestAvailDistributionsOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockAvailDistributionsOK()))
	defer ts.Close()

	dists, err := c.AvailDistributions(nil)
	require.NoError(t, err)
	require.NotEmpty(t, dists)

	assert.Len(t, dists, 22)
}

func testDistsNotEmpty(t *testing.T, dists []Distribution) {
	var everRequiresPVOps bool
	var ever64Bit bool

	for _, d := range dists {
		if d.RequiresPVOps {
			everRequiresPVOps = true
		}
		if d.Is64Bit {
			ever64Bit = true
		}
		assert.NotEmpty(t, d.ID, "d.ID")
		assert.NotEmpty(t, d.Label, "d.Label")
		assert.NotEmpty(t, d.MinImageSize, "d.MinImageSize")
		assert.NotEmpty(t, d.CreateDT, "d.CreateDT")
	}

	assert.True(t, everRequiresPVOps, "everRequiresPVOps")
	assert.True(t, ever64Bit, "ever64Bit")
}

func mockAvailDistributionsSingle() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":[{"REQUIRESPVOPSKERNEL":1,"DISTRIBUTIONID":130,"IS64BIT":1,"LABEL":"Debian 7","MINIMAGESIZE":600,"CREATE_DT":"2014-09-24 13:59:32.0"}],"ACTION":"avail.distributions"}`
	params = map[string]string{
		"DistributionID": "130",
		"api_action":     "avail.distributions",
		"api_key":        "foo",
	}
	responses = append(responses, newMockAPIResponse("avail.distributions", params, output))

	return responses
}

func TestAvailDistributionsSingle(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockAvailDistributionsSingle()))
	defer ts.Close()

	dists, err := c.AvailDistributions(Int(130))
	require.NoError(t, err)
	require.Len(t, dists, 1)

	d := dists[0]
	assert.Equal(t, true, d.RequiresPVOps, "d.RequiresPVOps")
	assert.Equal(t, 130, d.ID, "d.ID")
	assert.Equal(t, true, d.Is64Bit, "d.Is64Bit")
	assert.Equal(t, "Debian 7", d.Label, "d.Label")
	assert.Equal(t, 600, d.MinImageSize, "d.MinImageSize")
	assert.Equal(t, "2014-09-24 13:59:32.0", d.CreateDT, "d.CreateDT")
}

func mockAvailDistributionsEmpty() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":[],"ACTION":"avail.distributions"}`
	params = map[string]string{
		"DistributionID": "38201938",
		"api_action":     "avail.distributions",
		"api_key":        "foo",
	}
	responses = append(responses, newMockAPIResponse("avail.distributions", params, output))

	return responses
}

func TestAvailDistributionsEmpty(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockAvailDistributionsEmpty()))
	defer ts.Close()

	dists, err := c.AvailDistributions(Int(38201938))
	require.NoError(t, err)
	require.Empty(t, dists)
}

func mockAvailKernelsOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":[{"LABEL":"Latest 32 bit (4.1.0-x86-linode78)","ISXEN":1,"ISPVOPS":1,"KERNELID":137},{"LABEL":"3.13.7-x86-linode57","ISXEN":1,"ISPVOPS":1,"KERNELID":178},{"LABEL":"3.14.1-x86-linode58","ISXEN":1,"ISPVOPS":1,"KERNELID":180},{"LABEL":"3.14.4-x86-linode59","ISXEN":1,"ISPVOPS":1,"KERNELID":182},{"LABEL":"3.14.5-x86-linode60","ISXEN":1,"ISPVOPS":1,"KERNELID":184},{"LABEL":"3.14.5-x86-linode61","ISXEN":1,"ISPVOPS":1,"KERNELID":186},{"LABEL":"3.15.2-x86-linode62","ISXEN":1,"ISPVOPS":1,"KERNELID":188},{"LABEL":"3.15.3-x86-linode63","ISXEN":1,"ISPVOPS":1,"KERNELID":190},{"LABEL":"3.15.4-x86-linode64","ISXEN":1,"ISPVOPS":1,"KERNELID":192},{"LABEL":"3.16.5-x86-linode65","ISXEN":1,"ISPVOPS":1,"KERNELID":194},{"LABEL":"3.16.7-x86-linode67","ISXEN":1,"ISPVOPS":1,"KERNELID":198},{"LABEL":"3.18.1-x86-linode68","ISXEN":1,"ISPVOPS":1,"KERNELID":200},{"LABEL":"3.18.3-x86-linode69","ISXEN":1,"ISPVOPS":1,"KERNELID":203},{"LABEL":"3.18.5-x86-linode70","ISXEN":1,"ISPVOPS":1,"KERNELID":204},{"LABEL":"3.19.1-x86-linode71","ISXEN":1,"ISPVOPS":1,"KERNELID":206},{"LABEL":"4.0-x86-linode72","ISXEN":1,"ISPVOPS":1,"KERNELID":208},{"LABEL":"4.0.1-x86-linode73","ISXEN":1,"ISPVOPS":1,"KERNELID":211},{"LABEL":"4.0.4-x86-linode75","ISXEN":1,"ISPVOPS":1,"KERNELID":217},{"LABEL":"4.0.5-x86-linode76","ISXEN":1,"ISPVOPS":1,"KERNELID":219},{"LABEL":"4.0.5-x86-linode77","ISXEN":1,"ISPVOPS":1,"KERNELID":221},{"LABEL":"4.1.0-x86-linode78","ISXEN":1,"ISPVOPS":1,"KERNELID":222},{"LABEL":"Latest 64 bit (4.1.0-x86_64-linode59)","ISXEN":1,"ISPVOPS":1,"KERNELID":138},{"LABEL":"3.13.7-x86_64-linode38","ISXEN":1,"ISPVOPS":1,"KERNELID":177},{"LABEL":"3.14.1-x86_64-linode39","ISXEN":1,"ISPVOPS":1,"KERNELID":179},{"LABEL":"3.14.4-x86_64-linode40","ISXEN":1,"ISPVOPS":1,"KERNELID":181},{"LABEL":"3.14.5-x86_64-linode41","ISXEN":1,"ISPVOPS":1,"KERNELID":183},{"LABEL":"3.14.5-x86_64-linode42","ISXEN":1,"ISPVOPS":1,"KERNELID":185},{"LABEL":"3.15.2-x86_64-linode43","ISXEN":1,"ISPVOPS":1,"KERNELID":187},{"LABEL":"3.15.3-x86_64-linode44","ISXEN":1,"ISPVOPS":1,"KERNELID":189},{"LABEL":"3.15.4-x86_64-linode45","ISXEN":1,"ISPVOPS":1,"KERNELID":191},{"LABEL":"3.16.5-x86_64-linode46","ISXEN":1,"ISPVOPS":1,"KERNELID":193},{"LABEL":"3.16.7-x86_64-linode49","ISXEN":1,"ISPVOPS":1,"KERNELID":199},{"LABEL":"3.18.1-x86_64-linode50","ISXEN":1,"ISPVOPS":1,"KERNELID":201},{"LABEL":"3.18.3-x86_64-linode51","ISXEN":1,"ISPVOPS":1,"KERNELID":202},{"LABEL":"3.18.5-x86_64-linode52","ISXEN":1,"ISPVOPS":1,"KERNELID":205},{"LABEL":"3.19.1-x86_64-linode53","ISXEN":1,"ISPVOPS":1,"KERNELID":207},{"LABEL":"4.0-x86_64-linode54","ISXEN":1,"ISPVOPS":1,"KERNELID":209},{"LABEL":"4.0.1-x86_64-linode55","ISXEN":1,"ISPVOPS":1,"KERNELID":212},{"LABEL":"4.0.4-x86_64-linode57","ISXEN":1,"ISPVOPS":1,"KERNELID":218},{"LABEL":"4.0.5-x86_64-linode58","ISXEN":1,"ISPVOPS":1,"KERNELID":220},{"LABEL":"4.1.0-x86_64-linode59 ","ISXEN":1,"ISPVOPS":1,"KERNELID":223},{"LABEL":"pv-grub-x86_32","ISXEN":1,"ISPVOPS":0,"KERNELID":92},{"LABEL":"pv-grub-x86_64","ISXEN":1,"ISPVOPS":0,"KERNELID":95},{"LABEL":"Recovery - Finnix (kernel)","ISXEN":1,"ISPVOPS":0,"KERNELID":61},{"LABEL":"Latest 2.6 (2.6.39.1-linode34)","ISXEN":1,"ISPVOPS":1,"KERNELID":110},{"LABEL":"Latest Legacy (2.6.18.8-linode22)","ISXEN":1,"ISPVOPS":0,"KERNELID":60},{"LABEL":"2.6.18.8-domU-linode7","ISXEN":1,"ISPVOPS":0,"KERNELID":81},{"LABEL":"2.6.18.8-linode10","ISXEN":1,"ISPVOPS":0,"KERNELID":89},{"LABEL":"2.6.18.8-linode16","ISXEN":1,"ISPVOPS":0,"KERNELID":98},{"LABEL":"2.6.18.8-linode19","ISXEN":1,"ISPVOPS":0,"KERNELID":103},{"LABEL":"2.6.18.8-linode22","ISXEN":1,"ISPVOPS":0,"KERNELID":113},{"LABEL":"2.6.24.4-linode8","ISXEN":1,"ISPVOPS":1,"KERNELID":84},{"LABEL":"2.6.25-linode9","ISXEN":1,"ISPVOPS":1,"KERNELID":88},{"LABEL":"2.6.25.10-linode12","ISXEN":1,"ISPVOPS":1,"KERNELID":90},{"LABEL":"2.6.26-linode13","ISXEN":1,"ISPVOPS":1,"KERNELID":91},{"LABEL":"2.6.27.4-linode14","ISXEN":1,"ISPVOPS":1,"KERNELID":93},{"LABEL":"2.6.28-linode15","ISXEN":1,"ISPVOPS":1,"KERNELID":96},{"LABEL":"2.6.28.3-linode17","ISXEN":1,"ISPVOPS":1,"KERNELID":99},{"LABEL":"2.6.29-linode18","ISXEN":1,"ISPVOPS":1,"KERNELID":101},{"LABEL":"2.6.30.5-linode20","ISXEN":1,"ISPVOPS":1,"KERNELID":105},{"LABEL":"2.6.31.5-linode21","ISXEN":1,"ISPVOPS":1,"KERNELID":109},{"LABEL":"2.6.32-linode23","ISXEN":1,"ISPVOPS":1,"KERNELID":115},{"LABEL":"2.6.32.12-linode25","ISXEN":1,"ISPVOPS":1,"KERNELID":119},{"LABEL":"2.6.32.16-linode28","ISXEN":1,"ISPVOPS":1,"KERNELID":123},{"LABEL":"2.6.33-linode24","ISXEN":1,"ISPVOPS":1,"KERNELID":117},{"LABEL":"2.6.34-linode27","ISXEN":1,"ISPVOPS":1,"KERNELID":120},{"LABEL":"2.6.35.7-linode29","ISXEN":1,"ISPVOPS":1,"KERNELID":126},{"LABEL":"2.6.37-linode30","ISXEN":1,"ISPVOPS":1,"KERNELID":127},{"LABEL":"2.6.38-linode31","ISXEN":1,"ISPVOPS":1,"KERNELID":128},{"LABEL":"2.6.38.3-linode32","ISXEN":1,"ISPVOPS":1,"KERNELID":130},{"LABEL":"2.6.39-linode33","ISXEN":1,"ISPVOPS":1,"KERNELID":131},{"LABEL":"2.6.39.1-linode34","ISXEN":1,"ISPVOPS":1,"KERNELID":134},{"LABEL":"3.0.0-linode35","ISXEN":1,"ISPVOPS":1,"KERNELID":135},{"LABEL":"3.0.17-linode41","ISXEN":1,"ISPVOPS":1,"KERNELID":147},{"LABEL":"3.0.18-linode43","ISXEN":1,"ISPVOPS":1,"KERNELID":149},{"LABEL":"3.0.4-linode36","ISXEN":1,"ISPVOPS":1,"KERNELID":139},{"LABEL":"3.0.4-linode37","ISXEN":1,"ISPVOPS":1,"KERNELID":141},{"LABEL":"3.0.4-linode38","ISXEN":1,"ISPVOPS":1,"KERNELID":142},{"LABEL":"3.1.0-linode39","ISXEN":1,"ISPVOPS":1,"KERNELID":143},{"LABEL":"3.1.10-linode42","ISXEN":1,"ISPVOPS":1,"KERNELID":148},{"LABEL":"3.10.3-x86-linode53","ISXEN":1,"ISPVOPS":1,"KERNELID":169},{"LABEL":"3.11.6-x86-linode54","ISXEN":1,"ISPVOPS":1,"KERNELID":171},{"LABEL":"3.12.6-x86-linode55","ISXEN":1,"ISPVOPS":1,"KERNELID":174},{"LABEL":"3.12.9-x86-linode56","ISXEN":1,"ISPVOPS":1,"KERNELID":176},{"LABEL":"3.2.1-linode40","ISXEN":1,"ISPVOPS":1,"KERNELID":145},{"LABEL":"3.4.2-linode44","ISXEN":1,"ISPVOPS":1,"KERNELID":152},{"LABEL":"3.5.2-linode45","ISXEN":1,"ISPVOPS":1,"KERNELID":153},{"LABEL":"3.5.3-linode46","ISXEN":1,"ISPVOPS":1,"KERNELID":156},{"LABEL":"3.6.5-linode47","ISXEN":1,"ISPVOPS":1,"KERNELID":157},{"LABEL":"3.7.10-linode49","ISXEN":1,"ISPVOPS":1,"KERNELID":161},{"LABEL":"3.7.5-linode48","ISXEN":1,"ISPVOPS":1,"KERNELID":159},{"LABEL":"3.8.4-linode50","ISXEN":1,"ISPVOPS":1,"KERNELID":163},{"LABEL":"3.9.2-x86-linode51","ISXEN":1,"ISPVOPS":1,"KERNELID":166},{"LABEL":"3.9.3-x86-linode52","ISXEN":1,"ISPVOPS":1,"KERNELID":167},{"LABEL":"4.0.2-x86-linode74","ISXEN":1,"ISPVOPS":1,"KERNELID":214},{"LABEL":"Latest 2.6 (2.6.39.1-x86_64-linode19)","ISXEN":1,"ISPVOPS":1,"KERNELID":111},{"LABEL":"Latest Legacy (2.6.18.8-x86_64-linode10)","ISXEN":1,"ISPVOPS":0,"KERNELID":107},{"LABEL":"2.6.16.38-x86_64-linode2","ISXEN":1,"ISPVOPS":0,"KERNELID":85},{"LABEL":"2.6.18.8-x86_64-linode1","ISXEN":1,"ISPVOPS":0,"KERNELID":86},{"LABEL":"2.6.18.8-x86_64-linode10","ISXEN":1,"ISPVOPS":0,"KERNELID":114},{"LABEL":"2.6.18.8-x86_64-linode7","ISXEN":1,"ISPVOPS":0,"KERNELID":104},{"LABEL":"2.6.27.4-x86_64-linode3","ISXEN":1,"ISPVOPS":1,"KERNELID":94},{"LABEL":"2.6.28-x86_64-linode4","ISXEN":1,"ISPVOPS":1,"KERNELID":97},{"LABEL":"2.6.28.3-x86_64-linode5","ISXEN":1,"ISPVOPS":1,"KERNELID":100},{"LABEL":"2.6.29-x86_64-linode6","ISXEN":1,"ISPVOPS":1,"KERNELID":102},{"LABEL":"2.6.30.5-x86_64-linode8","ISXEN":1,"ISPVOPS":1,"KERNELID":106},{"LABEL":"2.6.31.5-x86_64-linode9","ISXEN":1,"ISPVOPS":1,"KERNELID":112},{"LABEL":"2.6.32-x86_64-linode11","ISXEN":1,"ISPVOPS":1,"KERNELID":116},{"LABEL":"2.6.32.12-x86_64-linode12","ISXEN":1,"ISPVOPS":1,"KERNELID":118},{"LABEL":"2.6.32.12-x86_64-linode15","ISXEN":1,"ISPVOPS":1,"KERNELID":124},{"LABEL":"2.6.34-x86_64-linode13","ISXEN":1,"ISPVOPS":1,"KERNELID":121},{"LABEL":"2.6.34-x86_64-linode14","ISXEN":1,"ISPVOPS":1,"KERNELID":122},{"LABEL":"2.6.35.4-x86_64-linode16","ISXEN":1,"ISPVOPS":1,"KERNELID":125},{"LABEL":"2.6.38-x86_64-linode17","ISXEN":1,"ISPVOPS":1,"KERNELID":129},{"LABEL":"2.6.39-x86_64-linode18","ISXEN":1,"ISPVOPS":1,"KERNELID":132},{"LABEL":"2.6.39.1-x86_64-linode19","ISXEN":1,"ISPVOPS":1,"KERNELID":133},{"LABEL":"3.0.0-x86_64-linode20","ISXEN":1,"ISPVOPS":1,"KERNELID":136},{"LABEL":"3.0.18-x86_64-linode24 ","ISXEN":1,"ISPVOPS":1,"KERNELID":150},{"LABEL":"3.0.4-x86_64-linode21","ISXEN":1,"ISPVOPS":1,"KERNELID":140},{"LABEL":"3.1.0-x86_64-linode22","ISXEN":1,"ISPVOPS":1,"KERNELID":144},{"LABEL":"3.10.3-x86_64-linode34","ISXEN":1,"ISPVOPS":1,"KERNELID":170},{"LABEL":"3.11.6-x86_64-linode35","ISXEN":1,"ISPVOPS":1,"KERNELID":172},{"LABEL":"3.12.6-x86_64-linode36","ISXEN":1,"ISPVOPS":1,"KERNELID":173},{"LABEL":"3.12.9-x86_64-linode37","ISXEN":1,"ISPVOPS":1,"KERNELID":175},{"LABEL":"3.2.1-x86_64-linode23","ISXEN":1,"ISPVOPS":1,"KERNELID":146},{"LABEL":"3.4.2-x86_64-linode25","ISXEN":1,"ISPVOPS":1,"KERNELID":151},{"LABEL":"3.5.2-x86_64-linode26","ISXEN":1,"ISPVOPS":1,"KERNELID":154},{"LABEL":"3.5.3-x86_64-linode27","ISXEN":1,"ISPVOPS":1,"KERNELID":155},{"LABEL":"3.6.5-x86_64-linode28","ISXEN":1,"ISPVOPS":1,"KERNELID":158},{"LABEL":"3.7.10-x86_64-linode30","ISXEN":1,"ISPVOPS":1,"KERNELID":162},{"LABEL":"3.7.5-x86_64-linode29","ISXEN":1,"ISPVOPS":1,"KERNELID":160},{"LABEL":"3.8.4-x86_64-linode31","ISXEN":1,"ISPVOPS":1,"KERNELID":164},{"LABEL":"3.9.2-x86_64-linode32","ISXEN":1,"ISPVOPS":1,"KERNELID":165},{"LABEL":"3.9.3-x86_64-linode33","ISXEN":1,"ISPVOPS":1,"KERNELID":168},{"LABEL":"4.0.2-x86_64-linode56","ISXEN":1,"ISPVOPS":1,"KERNELID":215}],"ACTION":"avail.kernels"}`
	params = map[string]string{
		"api_action": "avail.kernels",
		"api_key":    "foo",
	}
	responses = append(responses, newMockAPIResponse("avail.kernels", params, output))

	return responses
}

func TestAvailKernelsOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockAvailKernelsOK()))
	defer ts.Close()

	kerns, err := c.AvailKernels(nil, nil)
	require.NoError(t, err)
	require.Len(t, kerns, 135)

	var sample Kernel
	var everXen bool
	var everPVOps bool

	for _, k := range kerns {
		if k.IsXen {
			everXen = true
		}
		if k.IsPVOps {
			everPVOps = true
		}
		assert.NotEmpty(t, k.Label, "k.Label")
		assert.NotEmpty(t, k.ID, "k.ID")

		if k.ID == 138 {
			sample = k
		}
	}
	assert.True(t, everXen, "everXen")
	assert.True(t, everPVOps, "everPVOps")

	assert.Equal(t, "Latest 64 bit (4.1.0-x86_64-linode59)", sample.Label, "sample.Label")
	assert.True(t, sample.IsXen, "sample.IsXen")
	assert.True(t, sample.IsPVOps, "sample.IsPVOps")
	assert.Equal(t, 138, sample.ID, "sample.ID")
}

func mockAvailKernelsSingle() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":[{"LABEL":"Latest 64 bit (4.1.0-x86_64-linode59)","ISXEN":1,"ISPVOPS":1,"KERNELID":138}],"ACTION":"avail.kernels"}`
	params = map[string]string{
		"KernelID":   "138",
		"api_action": "avail.kernels",
		"api_key":    "foo",
		"isXen":      "true",
	}
	responses = append(responses, newMockAPIResponse("avail.kernels", params, output))

	return responses
}

func TestAvailKernelsSingle(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockAvailKernelsSingle()))
	defer ts.Close()

	kerns, err := c.AvailKernels(Int(138), Bool(true))
	require.NoError(t, err)
	require.Len(t, kerns, 1)

	k := kerns[0]

	assert.Equal(t, "Latest 64 bit (4.1.0-x86_64-linode59)", k.Label, "k.Label")
	assert.True(t, k.IsXen, "k.IsXen")
	assert.True(t, k.IsPVOps, "k.IsPVOps")
	assert.Equal(t, 138, k.ID, "k.ID")
}

func mockAvailKernelsEmpty() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":[],"ACTION":"avail.kernels"}`
	params = map[string]string{
		"KernelID":   "4389048234",
		"api_action": "avail.kernels",
		"api_key":    "foo",
		"isXen":      "false",
	}
	responses = append(responses, newMockAPIResponse("avail.kernels", params, output))

	return responses
}

func TestAvailKernelsEmpty(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockAvailKernelsEmpty()))
	defer ts.Close()

	kerns, err := c.AvailKernels(Int(4389048234), Bool(false))
	require.NoError(t, err)
	require.Len(t, kerns, 0)
}

func mockAvailLinodePlansOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":[{"CORES":1,"PRICE":10.00,"RAM":1024,"XFER":2000,"PLANID":1,"LABEL":"Linode 1024","AVAIL":{"3":500,"2":500,"7":500,"6":500,"4":500,"9":500,"8":500},"DISK":24,"HOURLY":0.0150},{"CORES":2,"PRICE":20.00,"RAM":2048,"XFER":3000,"PLANID":2,"LABEL":"Linode 2048","AVAIL":{"3":500,"2":500,"7":500,"6":500,"4":500,"9":500,"8":500},"DISK":48,"HOURLY":0.0300},{"CORES":4,"PRICE":40.00,"RAM":4096,"XFER":4000,"PLANID":4,"LABEL":"Linode 4096","AVAIL":{"3":500,"2":500,"7":500,"6":500,"4":500,"9":500,"8":500},"DISK":96,"HOURLY":0.0600},{"CORES":6,"PRICE":80.00,"RAM":8192,"XFER":8000,"PLANID":6,"LABEL":"Linode 8192","AVAIL":{"3":500,"2":500,"7":500,"6":500,"4":500,"9":500,"8":500},"DISK":192,"HOURLY":0.1200},{"CORES":8,"PRICE":160.00,"RAM":16384,"XFER":16000,"PLANID":7,"LABEL":"Linode 16384","AVAIL":{"3":500,"2":500,"7":500,"6":500,"4":500,"9":500,"8":500},"DISK":384,"HOURLY":0.2400},{"CORES":12,"PRICE":320.00,"RAM":32768,"XFER":20000,"PLANID":8,"LABEL":"Linode 32768","AVAIL":{"3":500,"2":500,"7":500,"6":500,"4":500,"9":500,"8":500},"DISK":768,"HOURLY":0.4800},{"CORES":16,"PRICE":480.00,"RAM":49152,"XFER":20000,"PLANID":9,"LABEL":"Linode 49152","AVAIL":{"3":500,"2":500,"7":500,"6":500,"4":500,"9":500,"8":500},"DISK":1152,"HOURLY":0.7200},{"CORES":20,"PRICE":640.00,"RAM":65536,"XFER":20000,"PLANID":10,"LABEL":"Linode 65536","AVAIL":{"3":500,"2":500,"7":500,"6":500,"4":500,"9":500,"8":500},"DISK":1536,"HOURLY":0.9600},{"CORES":20,"PRICE":960.00,"RAM":98304,"XFER":20000,"PLANID":12,"LABEL":"Linode 98304","AVAIL":{"3":500,"2":500,"7":500,"6":500,"4":500,"9":500,"8":500},"DISK":1920,"HOURLY":1.4400}],"ACTION":"avail.linodeplans"}`
	params = map[string]string{
		"api_action": "avail.linodeplans",
		"api_key":    "foo",
	}
	responses = append(responses, newMockAPIResponse("avail.linodeplans", params, output))

	return responses
}

func TestAvailLinodePlansOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockAvailLinodePlansOK()))
	defer ts.Close()

	plans, err := c.AvailLinodePlans(nil)
	require.NoError(t, err)
	require.Len(t, plans, 9)

	testPlanNotEmpty(t, plans)

	p := plans[0]

	assert.Equal(t, 1, p.Cores)
	assert.Equal(t, 10.00, p.Price)
	assert.Equal(t, 1024, p.RAM)
	assert.Equal(t, 2000, p.Xfer)
	assert.Equal(t, 1, p.ID)
	assert.Equal(t, "Linode 1024", p.Label)
	assert.Equal(t, 24, p.Disk)
	assert.Equal(t, 0.015, p.Hourly)
}

func testPlanNotEmpty(t *testing.T, plans []LinodePlan) {
	for _, p := range plans {
		assert.NotEmpty(t, p.Cores, "p.Cores")
		assert.NotEmpty(t, p.Price, "p.Price")
		assert.NotEmpty(t, p.RAM, "p.RAM")
		assert.NotEmpty(t, p.Xfer, "p.Xfer")
		assert.NotEmpty(t, p.ID, "p.ID")
		assert.NotEmpty(t, p.Label, "p.Label")
		assert.NotEmpty(t, p.Disk, "p.Disk")
		assert.NotEmpty(t, p.Hourly, "p.Hourly")
	}
}

func mockAvailLinodePlansSingle() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":[{"CORES":1,"PRICE":10.00,"RAM":1024,"XFER":2000,"PLANID":1,"LABEL":"Linode 1024","AVAIL":{"3":500,"2":500,"7":500,"6":500,"4":500,"9":500,"8":500},"DISK":24,"HOURLY":0.0150}],"ACTION":"avail.linodeplans"}`
	params = map[string]string{
		"PlanID":     "1",
		"api_action": "avail.linodeplans",
		"api_key":    "foo",
	}
	responses = append(responses, newMockAPIResponse("avail.linodeplans", params, output))

	return responses
}

func TestAvailLinodePlansSingle(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockAvailLinodePlansSingle()))
	defer ts.Close()

	plans, err := c.AvailLinodePlans(Int(1))
	require.NoError(t, err)
	require.Len(t, plans, 1)

	p := plans[0]

	assert.Equal(t, 1, p.Cores)
	assert.Equal(t, 10.00, p.Price)
	assert.Equal(t, 1024, p.RAM)
	assert.Equal(t, 2000, p.Xfer)
	assert.Equal(t, 1, p.ID)
	assert.Equal(t, "Linode 1024", p.Label)
	assert.Equal(t, 24, p.Disk)
	assert.Equal(t, 0.015, p.Hourly)
}

func mockAvailLinodePlansEmpty() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":[],"ACTION":"avail.linodeplans"}`
	params = map[string]string{
		"PlanID":     "3498230",
		"api_action": "avail.linodeplans",
		"api_key":    "foo",
	}
	responses = append(responses, newMockAPIResponse("avail.linodeplans", params, output))

	return responses
}

func TestAvailLinodePlansEmpty(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockAvailLinodePlansEmpty()))
	defer ts.Close()

	plans, err := c.AvailLinodePlans(Int(3498230))
	require.NoError(t, err)
	require.Len(t, plans, 0)
}

func mockAvailStackScriptsOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":[{"REV_NOTE":"Professional Services","SCRIPT":"#!\/bin\/bash\n\nhostname $(hostname).members.linode.com && wget -N http:\/\/httpupdate.cpanel.net\/latest && sh latest","DISTRIBUTIONIDLIST":127,"DESCRIPTION":"cPanel takes an hour to install. Track the installation progress in the Lish console.\r\n\r\n** DO NOT REBOOT THE SYSTEM FOR AT LEAST ONE HOUR **","REV_DT":"2015-01-07 15:48:57.0","LABEL":"cPanel","DEPLOYMENTSTOTAL":207,"LATESTREV":53081,"STACKSCRIPTID":11078,"ISPUBLIC":1,"DEPLOYMENTSACTIVE":125,"CREATE_DT":"2015-01-07 15:48:57.0","USERID":307510}],"ACTION":"avail.stackscripts"}`
	params = map[string]string{
		"DistributionID":     `1`,
		"DistributionVendor": `bar`,
		"api_action":         `avail.stackscripts`,
		"api_key":            `foo`,
		"keywords":           `baz`,
	}
	responses = append(responses, newMockAPIResponse("avail.stackscripts", params, output))

	return responses
}

func TestAvailStackScriptsOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockAvailStackScriptsOK()))
	defer ts.Close()

	scripts, err := c.AvailStackScripts(Int(1), String("bar"), String("baz"))
	require.NoError(t, err)
	require.Len(t, scripts, 1)

	s := scripts[0]
	assert.Equal(t, "Professional Services", s.RevNote)
	assert.Contains(t, s.Script, "#!/bin/bash")
	assert.Equal(t, "127", s.DistIDList)
	assert.Contains(t, s.Description, "cPanel takes an hour to install.")
	assert.Equal(t, "2015-01-07 15:48:57.0", s.RevDT)
	assert.Equal(t, "cPanel", s.Label)
	assert.Equal(t, 207, s.TotalDeploys)
	assert.Equal(t, 53081, s.LatestRev)
	assert.Equal(t, 11078, s.ID)
	assert.True(t, s.IsPublic)
	assert.Equal(t, 125, s.ActiveDeploys)
	assert.Equal(t, "2015-01-07 15:48:57.0", s.CreateDT)
	assert.Equal(t, 307510, s.UserID)
}

func mockTestEchoOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"FOO":"bar"},"ACTION":"test.echo"}`
	params = map[string]string{
		"api_action": "test.echo",
		"api_key":    "foo",
		"foo":        "bar",
	}
	responses = append(responses, newMockAPIResponse("test.echo", params, output))

	return responses
}

func TestEchoOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockTestEchoOK()))
	defer ts.Close()

	err := c.TestEcho()
	require.NoError(t, err)
}

func mockTestEchoWrongVal() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"FOO":"unexpected"},"ACTION":"test.echo"}`
	params = map[string]string{
		"api_action": "test.echo",
		"api_key":    "foo",
		"foo":        "bar",
	}
	responses = append(responses, newMockAPIResponse("test.echo", params, output))

	return responses
}

func TestEchoWrongVal(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockTestEchoWrongVal()))
	defer ts.Close()

	err := c.TestEcho()
	require.Error(t, err)
}
