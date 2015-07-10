// +build integration

package linode

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/alexsacr/linode/_third_party/testify/assert"
	"github.com/alexsacr/linode/_third_party/testify/require"
)

var apiKey string

func init() {
	apiKey = os.Getenv("LINODE_API_KEY")
	if apiKey == "" {
		fmt.Println("$LINODE_API_KEY not set.")
		os.Exit(1)
	}
}

const (
	rootPass          = "foo#23113."
	rootSSHKey        = "ssh-rsa foobarbaz"
	linodeLampStackSS = 10

	sslCert = `
-----BEGIN CERTIFICATE-----
MIIC9zCCAd+gAwIBAgIJAPekgcQOyxXDMA0GCSqGSIb3DQEBBQUAMBIxEDAOBgNV
BAMMB2Zvby5jb20wHhcNMTUwNzA3MjIwNDQ3WhcNMjUwNzA0MjIwNDQ3WjASMRAw
DgYDVQQDDAdmb28uY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA
5Ps5CEXI6QP/lFnk0apONjUyVD7HzRLp29LlFoSSz3SCH8jmsQ2krMcMGivy5qUK
8jhB4uw/Kmk5mwo5vjUURjlgKY967Xz5AGbs2MM6qWhOa4QJ+GCx6f+CW30R9lDB
j2lJEsLNKI4WMQ94eZwNFcfUumpetlPAYgHSG6udJGvdnODaYqI2hQ7R/sfEyO/f
zk6VWqGTBxFSi8+jlNokjO60BGUXpr2db1bg9mHKTosSHIotonLuQIFuLl8HhR4r
zM2aJI0AossvzFAmNJaIFdoyrWYu7NhgEMRpsJLI4SbTPOttlncxg2Ps4SkskPt3
yQhgv1fo5zujEIrOVbPWPQIDAQABo1AwTjAdBgNVHQ4EFgQU11HeU75IbfAjN3tv
vzXgEzHsgz4wHwYDVR0jBBgwFoAU11HeU75IbfAjN3tvvzXgEzHsgz4wDAYDVR0T
BAUwAwEB/zANBgkqhkiG9w0BAQUFAAOCAQEAd7xrZFgaVgZEdUiRuspYb5i7oTsj
j7LCypePzkYHk2t8iqgssHAP/647h6FE37qgI/yQVoSeG5Cj+uPz2wT/VYbDSPLI
Hl4wJlmRHfKFapUaB7H2YlCaIbB5gje8B6vTxKdKxgvm2ZHeS9wZkzAPXQVeptUk
squ0KXrgLbHRZGrN/zpK885EUdRhvdLklWD7dg3qfIxivAmtOFZkXj4KeJxtuQKK
QUSVF0tueKx2kPsYEgEUvxTJdTG4kkMw7N4YZApWQ0Rk2ycG2fLjHBkfDxvl6kDW
RfNfI8M4C2Gnly1miJYpeuD7eQYM2/Mi1wQKh7+B/VV45FMkchIa4KZ4wg==
-----END CERTIFICATE-----
`
	sslKey = `
-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEA5Ps5CEXI6QP/lFnk0apONjUyVD7HzRLp29LlFoSSz3SCH8jm
sQ2krMcMGivy5qUK8jhB4uw/Kmk5mwo5vjUURjlgKY967Xz5AGbs2MM6qWhOa4QJ
+GCx6f+CW30R9lDBj2lJEsLNKI4WMQ94eZwNFcfUumpetlPAYgHSG6udJGvdnODa
YqI2hQ7R/sfEyO/fzk6VWqGTBxFSi8+jlNokjO60BGUXpr2db1bg9mHKTosSHIot
onLuQIFuLl8HhR4rzM2aJI0AossvzFAmNJaIFdoyrWYu7NhgEMRpsJLI4SbTPOtt
lncxg2Ps4SkskPt3yQhgv1fo5zujEIrOVbPWPQIDAQABAoIBAQDTOoyZ/QYhLfcO
uv5CC4CqsNgErwVRQClLB3kdFJ75kNiYyZNFsPhACj43xmMEMSuv1gWxd42tafQh
+YTa2cwiG7uBm0Ii4i4xGMFUFISA0h/FPsCTx19YJMPz8aQoPMbRrKYkEF+BEdGW
4FaamEHZ3cV3UbidKRVHU7amn+EOPlQ2UUSW7OtRACQVDA3IFVa+pENUPaZGNiyS
YVEshAYJvyfcZ/8uAy/XiHM9mS+1hBpEIP/x95LeIQ+6gCCfAGg1N1nwZkSrl4vB
XHFLQcXdV31Pv4GN4NC11Ac1tmrWnks6SJP4YxwxwhQ1H5GCP7DWIoKJK27JbpCZ
D6hecWUBAoGBAPMtz89RgHWJ5xXc3Oz059YiOCa+suTjDt8+kQwq09yL0NwT1F83
3hFRHvGvaxbtQYV2PYLkMXCNH7gJcL9QzuWd/d6JO1gYTMflBRekolLHb/D2q6IT
faX+keaxdtxtjaJYwyTpm7q6L+QwchVo9cw9HwMECspkR7vClOE+4I91AoGBAPEN
ySV0lcTrY2cVApP9wNXtSTQjIHeViuytOUqxyRcy27YDzrK3bOL22n5GidXFAHI5
tqW3jOmA49Z7CzuSg52gZi3fAIm9bNxmFvp5po9aJbbtc6R4uvlkDiQ9AqZf+r7l
fwwii8S/YvXtKur5eLu/l42VKmWHVX3IYxFyLlqpAoGADLhEum9k7MD92WLoG2zk
A4terIglC1vlF0BKjfxfgJW/owNWhHSDpRw9Jh8q1tQqLTT2Grac72oyUooL37X2
XIRbFxzOgdkjiwErtwThiLrt3AjLCXtDgz9BmnOF6BnC4s8JxhCCGM5MWv89uGj5
EmnQCXRYrCV6qxOOfgmv7VECgYEArOuLwN+6VJHbM+ZBfDJLM+tNWSZdswtGpnO5
JpkUvRyRuidPDqtAZCxbHryxQfVZVZeLK6PZZOQ+DO0laK241slqoztW4nhNcGmX
0ESWND2h0nDSRUkKL78T2fEeRoWRlYGCOw1JIHF+pxZkdD1T27McH8fCqySDMoEe
eDRlhkkCgYEAtTfiwPfNjkVsnCtD2pmsOiBKTiBcv6OwC1gbIHl55RyZY/n1mE32
B0WBG+MNCJJfksblfLe8kfgddOR97+lpHpvqf8aM9Bo3S7W/fNBNSBzXJTvyR39M
rxDXhMjWqHBzcbswkgHskwb2OQQz7eBh+9PGk7P6hd1O+ih/W73rKOo=
-----END RSA PRIVATE KEY-----
`
)

func TestAvailDatacentersIntegration(t *testing.T) {
	c := NewClient(apiKey)

	datacenters, err := c.AvailDatacenters()
	require.NoError(t, err)
	require.NotEmpty(t, datacenters)

	// Test subset
	var foundDallas bool
	var foundFremont bool

	for _, dc := range datacenters {
		switch dc.ID {
		case 2:
			assert.Equal(t, "Dallas, TX, USA", dc.Location, "dc.Location")
			assert.Equal(t, "dallas", dc.Abbr, "dc.Abbr")
			foundDallas = true
		case 3:
			assert.Equal(t, "Fremont, CA, USA", dc.Location, "dc.Location")
			assert.Equal(t, "fremont", dc.Abbr, "dc.Abbr")
			foundFremont = true
		}
	}

	assert.True(t, foundDallas, "Dallas not returned.")
	assert.True(t, foundFremont, "Fremont not returned.")
}

func TestAvailDistributionsIntegration(t *testing.T) {
	c := NewClient(apiKey)

	dists, err := c.AvailDistributions(nil)
	require.NoError(t, err)
	require.NotEmpty(t, dists)

	testDistsNotEmpty(t, dists)

	dists, err = c.AvailDistributions(Int(130))
	assert.NoError(t, err)
	assert.Len(t, dists, 1)

	d := dists[0]
	assert.Equal(t, true, d.RequiresPVOps, "d.RequiresPVOps")
	assert.Equal(t, 130, d.ID, "d.ID")
	assert.Equal(t, true, d.Is64Bit, "d.Is64Bit")
	assert.Equal(t, "Debian 7", d.Label, "d.Label")
	assert.Equal(t, 600, d.MinImageSize, "d.MinImageSize")
	assert.Equal(t, "2014-09-24 13:59:32.0", d.CreateDT, "d.CreateDT")

	dists, err = c.AvailDistributions(Int(38201938))
	assert.NoError(t, err)
	assert.Empty(t, dists)
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

func TestAvailKernelsIntegration(t *testing.T) {
	c := NewClient(apiKey)

	kerns, err := c.AvailKernels(nil, nil)
	require.NoError(t, err)
	require.NotEmpty(t, kerns)

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
	}
	assert.True(t, everXen, "everXen")
	assert.True(t, everPVOps, "everPVOps")

	kerns, err = c.AvailKernels(Int(138), Bool(true))
	require.NoError(t, err)
	require.Len(t, kerns, 1)

	k := kerns[0]

	assert.True(t, k.IsXen)
	assert.True(t, k.IsPVOps)
	assert.NotEmpty(t, k.ID)
	assert.NotEmpty(t, k.Label)

	kerns, err = c.AvailKernels(Int(4389048234), Bool(false))
	require.NoError(t, err)
	require.Empty(t, kerns)
}

func TestAvailLinodePlansIntegration(t *testing.T) {
	c := NewClient(apiKey)

	plans, err := c.AvailLinodePlans(nil)
	require.NoError(t, err)
	require.NotEmpty(t, plans)

	testPlanNotEmpty(t, plans)

	// Spot check
	plans, err = c.AvailLinodePlans(Int(1))
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

	plans, err = c.AvailLinodePlans(Int(3498230))
	require.NoError(t, err)
	require.Empty(t, plans)
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

func TestAvailStackScriptsIntegration(t *testing.T) {
	c := NewClient(apiKey)

	sscripts, err := c.AvailStackScripts(nil, nil, nil)
	require.NoError(t, err)
	require.NotEmpty(t, sscripts)

	testSSNotEmpty(t, sscripts)

	sscripts, err = c.AvailStackScripts(Int(3133731337), String("bar"), String("baz"))
	require.NoError(t, err)
	require.Empty(t, sscripts)
}

func testSSNotEmpty(t *testing.T, sscripts []StackScript) {
	var everTotalDeploys bool
	var everActiveDeploys bool
	var everPublic bool
	var everDescription bool
	var everRevNote bool

	for _, ss := range sscripts {
		if ss.TotalDeploys != 0 {
			everTotalDeploys = true
		}
		if ss.ActiveDeploys != 0 {
			everActiveDeploys = true
		}
		if ss.IsPublic == true {
			everPublic = true
		}
		if ss.Description != "" {
			everDescription = true
		}
		if ss.RevNote != "" {
			everRevNote = true
		}

		assert.NotEmpty(t, ss.Script, "ss.Script")
		assert.NotEmpty(t, ss.Label, "ss.Label")
		assert.NotEmpty(t, ss.CreateDT, "ss.CreateDT")
		assert.NotEmpty(t, ss.ID, "ss.ID")
		assert.NotEmpty(t, ss.RevDT, "ss.RevDT")
		assert.NotEmpty(t, ss.UserID, "ss.UserID")
	}

	if len(sscripts) > 1 {
		assert.True(t, everTotalDeploys, "everTotalDeploys")
		assert.True(t, everActiveDeploys, "everActiveDeploys")
		assert.True(t, everPublic, "everPublic")
		assert.True(t, everDescription, "everDescription")
		assert.True(t, everRevNote, "everRevNote")
	}
}

func TestTestEchoIntegration(t *testing.T) {
	c := NewClient(apiKey)

	err := c.TestEcho()
	require.NoError(t, err)
}

func TestLinodeIntegration(t *testing.T) {
	c := NewClient(apiKey)

	t.Log("c.LinodeCreate...")
	id, err := c.LinodeCreate(2, 1, nil)
	require.NoError(t, err)
	require.NotEmpty(t, id)

	t.Log("c.WaitForAllJobs...")
	err = c.WaitForAllJobs(id, 2*time.Second, 60*time.Second)
	require.NoError(t, err)

	t.Log("c.LinodeList...")
	nodes, err := c.LinodeList(Int(id))
	require.NoError(t, err)
	require.Len(t, nodes, 1)

	t.Log("c.LinodeResize...")
	err = c.LinodeResize(id, 2)
	require.NoError(t, err)

	t.Log("c.LinodeUpdate...")
	linodeOpts := LinodeOpts{
		Label:                 String("integration-test"),
		DisplayGroup:          String("tests"),
		AlertCPUEnabled:       Bool(false),
		AlertCPUThreshold:     Int(50),
		AlertDiskIOEnabled:    Bool(false),
		AlertDiskIOThreshold:  Int(500),
		AlertBWInEnabled:      Bool(false),
		AlertBWInThreshold:    Int(30),
		AlertBWOutEnabled:     Bool(false),
		AlertBWOutThreshold:   Int(31),
		AlertBWQuotaEnabled:   Bool(false),
		AlertBWQuotaThreshold: Int(55),
		BackupWindow:          Int(3),
		BackupWeeklyDay:       Int(6),
		Watchdog:              Bool(false),
	}
	err = c.LinodeUpdate(id, linodeOpts)
	require.NoError(t, err)

	t.Log("c.LinodeDiskCreate...")
	jobID, swapDiskID, err := c.LinodeDiskCreate(id, "test-swap", "swap", 256)
	require.NoError(t, err)
	require.NotEmpty(t, swapDiskID, "swapDiskID")
	require.NotEmpty(t, jobID, "jobID")

	t.Log("c.WaitForJob...")
	ok, err := c.WaitForJob(id, jobID, 2*time.Second, 60*time.Second)
	require.NoError(t, err)
	require.True(t, ok, "creating swap failed")

	t.Log("c.LinodeDiskCreateFromDistribution...")
	jobID, distDiskID, err := c.LinodeDiskCreateFromDistribution(id, 130, "test-dist", 600,
		rootPass, String(rootSSHKey))
	require.NoError(t, err)
	require.NotEmpty(t, jobID)
	require.NotEmpty(t, distDiskID)

	t.Log("c.WaitForJob...")
	ok, err = c.WaitForJob(id, jobID, 2*time.Second, 80*time.Second)
	require.NoError(t, err)
	require.True(t, ok)

	t.Log("c.LinodeConfigCreate...")
	lcco := LinodeConfigCreateOpts{
		Comments:              String("foo"),
		RAMLimit:              Int(800),
		VirtMode:              String("paravirt"),
		RunLevel:              String("default"),
		RootDeviceNum:         Int(1),
		RootDeviceRO:          Bool(true),
		HelperDisableUpdateDB: Bool(true),
		HelperDistro:          Bool(true),
		HelperXen:             Bool(true),
		HelperDepmod:          Bool(true),
		HelperNetwork:         Bool(true),
		DevTmpFSAutomount:     Bool(true),
	}
	diskList := strconv.Itoa(distDiskID) + "," + strconv.Itoa(swapDiskID) + ","
	confID, err := c.LinodeConfigCreate(id, 138, "test-conf1", diskList, lcco)
	require.NoError(t, err)
	require.NotEmpty(t, confID)

	t.Log("c.LinodeConfigList...")
	confs, err := c.LinodeConfigList(id, Int(confID))
	require.NoError(t, err)
	require.Len(t, confs, 1)

	cfg := confs[0]
	assert.Equal(t, "", cfg.RootDeviceCustom)
	assert.Equal(t, "foo", cfg.Comments)
	assert.Equal(t, "test-conf1", cfg.Label)
	assert.Equal(t, confID, cfg.ID)
	assert.NotEmpty(t, cfg.DiskList)
	assert.Equal(t, "default", cfg.RunLevel)
	assert.Equal(t, 1, cfg.RootDeviceNum)
	assert.Equal(t, 800, cfg.RAMLimit)
	assert.Equal(t, "paravirt", cfg.VirtMode)
	assert.Equal(t, id, cfg.LinodeID)
	assert.Equal(t, 138, cfg.KernelID)
	assert.True(t, cfg.DevTmpFSAutomount)
	assert.True(t, cfg.HelperDistro)
	assert.True(t, cfg.HelperDisableUpdateDB)
	assert.True(t, cfg.HelperDepmod)
	assert.True(t, cfg.HelperXen)
	assert.True(t, cfg.RootDeviceRO)
	assert.True(t, cfg.HelperNetwork)
	assert.False(t, cfg.IsRescue)

	t.Log("c.LinodeConfigUpdate...")
	lcuo := LinodeConfigUpdateOpts{
		LinodeID:              Int(id),
		KernelID:              Int(138),
		Comments:              String("foo"),
		RAMLimit:              Int(800),
		VirtMode:              String("paravirt"),
		RunLevel:              String("default"),
		RootDeviceNum:         Int(1),
		RootDeviceRO:          Bool(true),
		HelperDisableUpdateDB: Bool(true),
		HelperDistro:          Bool(true),
		HelperXen:             Bool(true),
		HelperDepmod:          Bool(true),
		HelperNetwork:         Bool(true),
		DevTmpFSAutomount:     Bool(true),
	}
	err = c.LinodeConfigUpdate(confID, lcuo)
	require.NoError(t, err)

	t.Log("c.LinodeDiskImagize...")
	jobID, imgID, err := c.LinodeDiskImagize(id, distDiskID, String("test-image desc"),
		String("test-image label"))
	require.NoError(t, err)
	require.NotEmpty(t, jobID)
	require.NotEmpty(t, imgID)

	t.Log("c.WaitForJob...")
	ok, err = c.WaitForJob(id, jobID, 2*time.Second, 120*time.Second)
	require.NoError(t, err)
	require.True(t, ok)

	t.Log("c.LinodeDiskCreateFromImage...")
	jobID, imgDiskID, err := c.LinodeDiskCreateFromImage(imgID, id, "test-image", Int(800),
		String(rootPass), String(rootSSHKey))
	require.NoError(t, err)
	require.NotEmpty(t, jobID)
	require.NotEmpty(t, imgDiskID)

	t.Log("c.WaitForJob...")
	ok, err = c.WaitForJob(id, jobID, 2*time.Second, 120*time.Second)
	require.NoError(t, err)
	require.True(t, ok)

	t.Log("deleting image...")
	err = c.ImageDelete(imgID)
	assert.NoError(t, err)

	t.Log("c.LinodeDiskDelete...")
	jobID, err = c.LinodeDiskDelete(id, imgDiskID)
	require.NoError(t, err)
	require.NotEmpty(t, jobID)

	t.Log("c.WaitForJob...")
	ok, err = c.WaitForJob(id, jobID, 2*time.Second, 60*time.Second)
	require.NoError(t, err)
	require.True(t, ok)

	t.Log("c.LinodeDiskDuplicate...")
	jobID, dupeDiskID, err := c.LinodeDiskDuplicate(id, swapDiskID)
	require.NoError(t, err)
	require.NotEmpty(t, dupeDiskID)
	require.NotEmpty(t, jobID)

	t.Log("c.WaitForJob...")
	ok, err = c.WaitForJob(id, jobID, 2*time.Second, 60*time.Second)
	require.NoError(t, err)
	require.True(t, ok)

	t.Log("c.LinodeDiskResize...")
	jobID, err = c.LinodeDiskResize(id, dupeDiskID, 512)
	require.NoError(t, err)
	require.NotEmpty(t, jobID)

	t.Log("c.WaitForJob...")
	ok, err = c.WaitForJob(id, jobID, 1*time.Second, 60*time.Second)
	require.NoError(t, err)
	require.True(t, ok)

	t.Log("c.LinodeDiskUpdate...")
	err = c.LinodeDiskUpdate(id, dupeDiskID, String("updated-label"), Bool(true))
	require.NoError(t, err)

	curDisks, err := c.LinodeDiskList(id, nil)
	require.NoError(t, err)
	require.Len(t, curDisks, 3)

	t.Log("c.LinodeIPList...")
	ipList, err := c.LinodeIPList(Int(id), nil)
	require.NoError(t, err)
	require.Len(t, ipList, 1)
	pubID := ipList[0].ID

	t.Log("c.LinodeIPAddPrivate...")
	_, _, err = c.LinodeIPAddPrivate(id)
	require.NoError(t, err)

	t.Log("c.LinodeIPList...")
	ipList, err = c.LinodeIPList(Int(id), nil)
	require.NoError(t, err)
	require.Len(t, ipList, 2)

	t.Log("c.LinodeClone...")
	cloneID, err := c.LinodeClone(id, 2, 1, nil, nil)
	require.NoError(t, err)
	require.NotEmpty(t, cloneID)

	t.Log("c.WaitForAllJobs...")
	err = c.WaitForAllJobs(cloneID, 2*time.Second, 60*time.Second)
	require.NoError(t, err)

	t.Log("c.LinodeIPList...")
	ipList, err = c.LinodeIPList(Int(cloneID), nil)
	require.NoError(t, err)
	require.Len(t, ipList, 1)
	clonePubID := ipList[0].ID

	t.Log("c.LinodeIPSwap...")
	err = c.LinodeIPSwap(pubID, Int(clonePubID), nil)
	require.NoError(t, err)

	t.Log("c.WaitForAllJobs...")
	err = c.WaitForAllJobs(cloneID, 2*time.Second, 60*time.Second)
	require.NoError(t, err)

	t.Log("c.LinodeDelete...")
	err = c.LinodeDelete(cloneID, Bool(true))
	require.NoError(t, err)

	ssOpts := struct {
		DBPassword    string `json:"db_password"`
		DBName        string `json:"db_name"`
		MySQLUsername string `json:"db_user"`
		MySQLUserpass string `json:"db_user_password"`
	}{
		"foo",
		"bar",
		"baz",
		"quux",
	}
	ssUDFResp, err := json.Marshal(ssOpts)
	require.NoError(t, err)

	t.Log("c.DiskCreateFromStackScript")
	jobID, diskID, err := c.LinodeDiskCreateFromStackScript(id, linodeLampStackSS, string(ssUDFResp),
		130, "test-ss", 600, rootPass, String(rootSSHKey))
	require.NoError(t, err)
	require.NotEmpty(t, jobID)
	require.NotEmpty(t, diskID)

	t.Log("c.WaitForJob...")
	ok, err = c.WaitForJob(id, jobID, 2*time.Second, 120*time.Second)
	require.NoError(t, err)
	require.True(t, ok)

	t.Log("c.LinodeBoot...")
	jobID, err = c.LinodeBoot(id, Int(confID))
	require.NoError(t, err)
	require.NotEmpty(t, jobID)

	t.Log("c.WaitForJob...")
	ok, err = c.WaitForJob(id, jobID, 2*time.Second, 120*time.Second)
	require.NoError(t, err)
	require.True(t, ok)

	t.Log("c.LinodeReboot...")
	jobID, err = c.LinodeReboot(id, Int(confID))
	require.NoError(t, err)
	require.NotEmpty(t, jobID)

	t.Log("c.WaitForJob...")
	ok, err = c.WaitForJob(id, jobID, 2*time.Second, 120*time.Second)
	require.NoError(t, err)
	require.True(t, ok)

	t.Log("c.LinodeShutdown...")
	jobID, err = c.LinodeShutdown(id)
	require.NoError(t, err)
	require.NotEmpty(t, jobID)

	t.Log("c.WaitForJob...")
	ok, err = c.WaitForJob(id, jobID, 2*time.Second, 120*time.Second)
	require.NoError(t, err)
	require.True(t, ok)

	err = c.LinodeConfigDelete(id, confID)
	require.NoError(t, err)

	t.Log("c.LinodeDelete...")
	err = c.LinodeDelete(id, Bool(true))

}

func TestAccountEstimateInvoiceIntegration(t *testing.T) {
	c := NewClient(apiKey)

	inv, err := c.AccountEstimateInvoice("linode_new", Int(1), Int(1), nil)
	require.NoError(t, err)
	require.NotEmpty(t, inv.InvoiceTo)
	require.NotEmpty(t, inv.Price)
}

func TestAccountInfoIntegration(t *testing.T) {
	c := NewClient(apiKey)

	info, err := c.AccountInfo()
	require.NoError(t, err)
	require.NotEmpty(t, info.ActiveSince)
	require.NotEmpty(t, info.TransferPool)
}

func TestDomainIntegration(t *testing.T) {
	c := NewClient(apiKey)

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

	t.Log("domain.create...")
	dID, err := c.DomainCreate("foo.com", "master", dco)
	require.NoError(t, err)
	require.NotEmpty(t, dID)

	t.Log("domain.list...")
	domains, err := c.DomainList(Int(dID))
	require.NoError(t, err)
	require.Len(t, domains, 1)
	d := domains[0]

	assert.Equal(t, dID, d.ID)
	assert.Equal(t, "foo", d.Description)
	assert.Equal(t, 300, d.ExpireSec)
	assert.Equal(t, 300, d.RetrySec)
	assert.Equal(t, 1, d.Status)
	assert.Equal(t, "test", d.DisplayGroup)
	assert.Equal(t, "", d.MasterIPs)
	assert.Equal(t, 300, d.RefreshSec)
	assert.Equal(t, "foo@foo.com", d.SOAEmail)
	assert.Equal(t, 300, d.TTLSec)
	assert.Equal(t, "foo.com", d.Domain)
	assert.Equal(t, "master", d.Type)
	assert.Equal(t, "", d.AXFRIPs)

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

	t.Log("domain.update...")
	err = c.DomainUpdate(dID, duo)
	require.NoError(t, err)

	t.Log("domain.list...")
	domains, err = c.DomainList(Int(dID))
	require.NoError(t, err)
	require.Len(t, domains, 1)
	d = domains[0]

	assert.Equal(t, dID, d.ID)
	assert.Equal(t, 3600, d.ExpireSec)
	assert.Equal(t, 3600, d.RetrySec)
	assert.Equal(t, 2, d.Status)
	assert.Equal(t, "still-testing", d.DisplayGroup)
	assert.Equal(t, "", d.MasterIPs)
	assert.Equal(t, 3600, d.RefreshSec)
	assert.Equal(t, "baz@baz.com", d.SOAEmail)
	assert.Equal(t, 3600, d.TTLSec)
	assert.Equal(t, "baz.com", d.Domain)
	assert.Equal(t, "master", d.Type)
	assert.Equal(t, "", d.AXFRIPs)

	drco := DomainResourceCreateOpts{
		Name:     String("_foo._tcp"),
		Target:   String("bar.baz.com"),
		Priority: Int(5),
		Weight:   Int(10),
		Port:     Int(15),
		Protocol: String("bar"),
		TTLSec:   Int(20),
	}

	t.Log("domain.resource.create...")
	rID, err := c.DomainResourceCreate(dID, "srv", drco)
	require.NoError(t, err)
	require.NotEmpty(t, rID)

	t.Log("domain.resource.list...")
	resources, err := c.DomainResourceList(dID, Int(rID))
	require.NoError(t, err)
	require.Len(t, resources, 1)
	r := resources[0]

	assert.Equal(t, dID, r.DomainID)
	assert.Equal(t, 15, r.Port)
	assert.Equal(t, rID, r.ID)
	assert.Equal(t, "_foo._tcp", r.Name)
	assert.Equal(t, 10, r.Weight)
	assert.Equal(t, 300, r.TTLSec)
	assert.Equal(t, "bar.baz.com", r.Target)
	assert.Equal(t, 5, r.Priority)
	assert.Equal(t, "tcp", r.Protocol)
	assert.Equal(t, "srv", r.Type)

	druo := DomainResourceUpdateOpts{
		DomainID: Int(dID),
		Name:     String("_qux._udp"),
		Target:   String("qux.baz.com"),
		Priority: Int(20),
		Weight:   Int(25),
		Port:     Int(30),
		Protocol: String("udp"),
		TTLSec:   Int(301),
	}

	t.Log("domain.resource.update...")
	err = c.DomainResourceUpdate(rID, druo)
	require.NoError(t, err)

	t.Log("domain.resource.list...")
	resources, err = c.DomainResourceList(dID, Int(rID))
	require.NoError(t, err)
	require.Len(t, resources, 1)
	r = resources[0]

	assert.Equal(t, dID, r.DomainID)
	assert.Equal(t, 30, r.Port)
	assert.Equal(t, rID, r.ID)
	assert.Equal(t, "_qux._udp", r.Name)
	assert.Equal(t, 25, r.Weight)
	assert.Equal(t, 3600, r.TTLSec)
	assert.Equal(t, "qux.baz.com", r.Target)
	assert.Equal(t, 20, r.Priority)
	assert.Equal(t, "udp", r.Protocol)

	t.Log("domain.resource.delete...")
	err = c.DomainResourceDelete(dID, rID)
	require.NoError(t, err)

	t.Log("domain.delete...")
	err = c.DomainDelete(dID)
	require.NoError(t, err)
}

func TestImageIntegration(t *testing.T) {
	c := NewClient(apiKey)

	t.Log("creating test linode...")
	linodeID, err := c.LinodeCreate(2, 1, nil)
	require.NoError(t, err)
	require.NotEmpty(t, linodeID)

	err = c.WaitForAllJobs(linodeID, 3*time.Second, 60*time.Second)
	require.NoError(t, err)

	t.Log("creating test disk...")
	jobID, dDiskID, err := c.LinodeDiskCreateFromDistribution(linodeID, 130, "test-image", 600,
		rootPass, nil)
	require.NoError(t, err)
	require.NotEmpty(t, dDiskID)
	require.NotEmpty(t, jobID)

	ok, err := c.WaitForJob(linodeID, jobID, 3*time.Second, 60*time.Second)
	require.NoError(t, err)
	require.True(t, ok)

	t.Log("converting test disk to image...")
	jobID, imageID, err := c.LinodeDiskImagize(linodeID, dDiskID, String("foo"), String("bar"))
	require.NoError(t, err)
	require.NotEmpty(t, imageID)
	require.NotEmpty(t, jobID)

	ok, err = c.WaitForJob(linodeID, jobID, 3*time.Second, 120*time.Second)
	require.NoError(t, err)
	require.True(t, ok)

	t.Log("image.list...")
	images, err := c.ImageList(&imageID, Bool(false))
	require.NoError(t, err)
	require.Len(t, images, 1)

	i := images[0]
	assert.Empty(t, i.LastUsedDT)
	assert.Equal(t, 600, i.MinSize)
	assert.Equal(t, "foo", i.Description)
	assert.Equal(t, "bar", i.Label)
	assert.NotEmpty(t, i.Creator)
	assert.Equal(t, "available", i.Status)
	assert.False(t, i.IsPublic)
	assert.NotEmpty(t, i.CreateDT)
	assert.Equal(t, "manual", i.Type)
	assert.Equal(t, "ext4", i.FSType)
	assert.Equal(t, imageID, i.ID)

	t.Log("image.update...")
	err = c.ImageUpdate(imageID, String("baz"), String("quux"))
	require.NoError(t, err)

	t.Log("image.list...")
	images, err = c.ImageList(&imageID, Bool(false))
	require.NoError(t, err)
	require.Len(t, images, 1)

	i = images[0]
	assert.Equal(t, "baz", i.Label)
	assert.Equal(t, "quux", i.Description)

	t.Log("image.delete...")
	err = c.ImageDelete(imageID)
	require.NoError(t, err)

	t.Log("cleaning up linode...")
	err = c.LinodeDelete(linodeID, Bool(true))
	require.NoError(t, err)
}

func TestNodeBalancerIntegration(t *testing.T) {
	c := NewClient(apiKey)

	t.Log("creating test linode...")
	linodeID, err := c.LinodeCreate(2, 1, nil)
	require.NoError(t, err)
	require.NotEmpty(t, linodeID)

	err = c.WaitForAllJobs(linodeID, 3*time.Second, 60*time.Second)
	require.NoError(t, err)

	t.Log("assigning private IP...")
	_, linodeIP, err := c.LinodeIPAddPrivate(linodeID)
	require.NoError(t, err)
	require.NotEmpty(t, linodeIP)

	t.Log("nodebalancer.create...")
	nbID, err := c.NodeBalancerCreate(2, String("testing"), Int(5))
	require.NoError(t, err)
	require.NotEmpty(t, nbID)

	t.Log("nodebalancer.list...")
	nbList, err := c.NodeBalancerList(Int(nbID))
	require.NoError(t, err)
	require.Len(t, nbList, 1)

	n := nbList[0]
	assert.NotEmpty(t, n.Hostname)
	assert.Equal(t, "testing", n.Label)
	assert.Equal(t, 5, n.Throttle)
	assert.Equal(t, 2, n.DatacenterID)
	assert.NotEmpty(t, n.IPv4Addr)
	assert.NotEmpty(t, n.IPv6Addr)
	assert.Equal(t, nbID, n.ID)

	t.Log("nodebalancer.update...")
	err = c.NodeBalancerUpdate(nbID, String("testing-2"), Int(10))
	require.NoError(t, err)

	t.Log("nodebalancer.list...")
	nbList, err = c.NodeBalancerList(Int(nbID))
	require.NoError(t, err)
	require.Len(t, nbList, 1)

	n = nbList[0]
	assert.Equal(t, "testing-2", n.Label)
	assert.Equal(t, 10, n.Throttle)
	assert.Equal(t, nbID, n.ID)

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
	confID, err := c.NodeBalancerConfigCreate(nbID, nbcco)
	require.NoError(t, err)
	require.NotEmpty(t, confID)

	t.Log("nodebalancer.config.list...")
	nbconfs, err := c.NodeBalancerConfigList(nbID, Int(confID))
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
	assert.Equal(t, confID, nbc.ID)
	assert.Equal(t, 15, nbc.CheckAttempts)
	assert.Equal(t, nbID, nbc.NodeBalancerID)
	assert.Equal(t, 80, nbc.Port)
	assert.Equal(t, "http", nbc.Check)
	assert.False(t, nbc.CheckPassive)
	assert.Equal(t, "http", nbc.Protocol)
	assert.Equal(t, 29, nbc.CheckTimeout)

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
		SSLCert:       String(sslCert),
		SSLKey:        String(sslKey),
	}

	t.Log("nodebalancer.config.update...")
	err = c.NodeBalancerConfigUpdate(confID, nbcuo)
	require.NoError(t, err)

	t.Log("nodebalancer.config.list...")
	nbconfs, err = c.NodeBalancerConfigList(nbID, Int(confID))
	require.NoError(t, err)
	require.Len(t, nbconfs, 1)

	nbc = nbconfs[0]
	assert.Equal(t, "table", nbc.Stickiness)
	assert.Equal(t, "foo.com", nbc.SSLCommonName)
	assert.Equal(t, "/bar", nbc.CheckPath)
	assert.Equal(t, "quux", nbc.CheckBody)
	assert.Equal(t, 24, nbc.CheckInterval)
	assert.Equal(t, "0B:40:09:0C:4E:DA:5B:FB:2A:31:69:C9:4D:80:AE:CE:76:8F:DA:60", nbc.SSLFingerprint)
	assert.Equal(t, "leastconn", nbc.Algorithm)
	assert.Equal(t, confID, nbc.ID)
	assert.Equal(t, 20, nbc.CheckAttempts)
	assert.Equal(t, nbID, nbc.NodeBalancerID)
	assert.Equal(t, 90, nbc.Port)
	assert.Equal(t, "http_body", nbc.Check)
	assert.True(t, nbc.CheckPassive)
	assert.Equal(t, "https", nbc.Protocol)
	assert.Equal(t, 23, nbc.CheckTimeout)

	t.Log("nodebalancer.node.create...")
	nbNodeID, err := c.NodeBalancerNodeCreate(confID, "test", linodeIP+":90", Int(50),
		String("accept"))
	require.NoError(t, err)
	require.NotEmpty(t, nbNodeID)

	t.Log("nodebalancer.node.list...")
	nbNodes, err := c.NodeBalancerNodeList(confID, Int(nbNodeID))
	require.NoError(t, err)
	require.Len(t, nbNodes, 1)

	node := nbNodes[0]
	assert.Equal(t, 50, node.Weight)
	assert.Equal(t, linodeIP+":90", node.Address)
	assert.Equal(t, "test", node.Label)
	assert.Equal(t, nbNodeID, node.ID)
	assert.Equal(t, "accept", node.Mode)
	assert.Equal(t, confID, node.ConfigID)
	assert.NotEmpty(t, node.Status)
	assert.Equal(t, nbID, node.NodeBalancerID)

	t.Log("nodebalancer.node.update...")
	err = c.NodeBalancerNodeUpdate(nbNodeID, String("test-2"), String(linodeIP+":80"),
		Int(60), String("reject"))
	require.NoError(t, err)

	t.Log("nodebalancer.node.list...")
	nbNodes, err = c.NodeBalancerNodeList(confID, Int(nbNodeID))
	require.NoError(t, err)
	require.Len(t, nbNodes, 1)

	node = nbNodes[0]
	assert.Equal(t, 60, node.Weight)
	assert.Equal(t, linodeIP+":80", node.Address)
	assert.Equal(t, "test-2", node.Label)
	assert.Equal(t, nbNodeID, node.ID)
	assert.Equal(t, "reject", node.Mode)
	assert.Equal(t, confID, node.ConfigID)
	assert.NotEmpty(t, node.Status)
	assert.Equal(t, nbID, node.NodeBalancerID)

	t.Log("nodebalancer.node.delete...")
	err = c.NodeBalancerNodeDelete(nbNodeID)
	require.NoError(t, err)

	t.Log("nodebalancer.config.delete...")
	err = c.NodeBalancerConfigDelete(nbID, confID)
	require.NoError(t, err)

	t.Log("nodebalancer.delete...")
	err = c.NodeBalancerDelete(nbID)
	require.NoError(t, err)

	t.Log("deleting test linode...")
	err = c.LinodeDelete(linodeID, Bool(true))
	require.NoError(t, err)
}

func TestStackScriptIntegration(t *testing.T) {
	c := NewClient(apiKey)

	ssID, err := c.StackScriptCreate("test", "130,", "#! /bin/bash foo", String("foo"),
		Bool(false), String("bar"))
	require.NoError(t, err)
	require.NotEmpty(t, ssID)

	ssList, err := c.StackScriptList(Int(ssID))
	require.NoError(t, err)
	require.Len(t, ssList, 1)

	ss := ssList[0]
	assert.Equal(t, "bar", ss.RevNote)
	assert.Equal(t, "#! /bin/bash foo", ss.Script)
	assert.Equal(t, "130,", ss.DistIDList)
	assert.Equal(t, "foo", ss.Description)
	assert.NotEmpty(t, ss.RevDT)
	assert.Equal(t, "test", ss.Label)
	assert.Equal(t, 0, ss.TotalDeploys)
	assert.NotEmpty(t, ss.LatestRev)
	assert.Equal(t, ssID, ss.ID)
	assert.False(t, ss.IsPublic)
	assert.Equal(t, 0, ss.ActiveDeploys)
	assert.NotEmpty(t, ss.CreateDT)
	assert.NotEmpty(t, ss.UserID)

	err = c.StackScriptUpdate(ssID, String("test-2"), String("quux"), String("129,130"), Bool(true),
		String("baz"), String("#! /bin/bash baz"))
	require.NoError(t, err)

	ssList, err = c.StackScriptList(Int(ssID))
	require.NoError(t, err)
	require.Len(t, ssList, 1)

	ss = ssList[0]
	assert.Equal(t, "baz", ss.RevNote)
	assert.Equal(t, "#! /bin/bash baz", ss.Script)
	assert.Equal(t, "129,130", ss.DistIDList)
	assert.Equal(t, "quux", ss.Description)
	assert.Equal(t, "test-2", ss.Label)
	assert.Equal(t, ssID, ss.ID)
	assert.True(t, ss.IsPublic)

	err = c.StackScriptDelete(ssID)
	require.NoError(t, err)
}
