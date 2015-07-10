// +build !integration

package linode

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/alexsacr/linode/_third_party/testify/assert"
	"github.com/alexsacr/linode/_third_party/testify/require"
)

const (
	rootPass   = "foo#23113."
	rootSSHKey = "ssh-rsa foobarbaz"
)

func mockLinodeCreateOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"LinodeID":1139016},"ACTION":"linode.create"}`
	params = map[string]string{
		"DatacenterID": "2",
		"PlanID":       "1",
		"api_action":   "linode.create",
		"api_key":      "foo",
	}
	responses = append(responses, newMockAPIResponse("linode.create", params, output))

	return responses
}

func TestLinodeCreateOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockLinodeCreateOK()))
	defer ts.Close()

	id, err := c.LinodeCreate(2, 1, nil)
	require.NoError(t, err)
	require.Equal(t, 1139016, id)
}

func mockLinodeListOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":[{"ALERT_CPU_ENABLED":1,"ALERT_BWIN_ENABLED":1,"ALERT_BWQUOTA_ENABLED":1,"ALERT_DISKIO_THRESHOLD":10000,"BACKUPWINDOW":0,"WATCHDOG":1,"DISTRIBUTIONVENDOR":"foo","DATACENTERID":2,"STATUS":1,"ALERT_DISKIO_ENABLED":1,"CREATE_DT":"2015-07-02 23:08:52.0","TOTALHD":24576,"ALERT_BWQUOTA_THRESHOLD":80,"TOTALRAM":1024,"ALERT_BWIN_THRESHOLD":10,"LINODEID":1139016,"ALERT_BWOUT_THRESHOLD":10,"ALERT_BWOUT_ENABLED":1,"BACKUPSENABLED":1,"ALERT_CPU_THRESHOLD":90,"PLANID":1,"BACKUPWEEKLYDAY":1,"LABEL":"linode1139016","LPM_DISPLAYGROUP":"bar","TOTALXFER":2000}],"ACTION":"linode.list"}`
	params = map[string]string{
		"LinodeID":   "1139016",
		"api_action": "linode.list",
		"api_key":    "foo",
	}
	responses = append(responses, newMockAPIResponse("linode.list", params, output))

	return responses
}

func TestLinodeListOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockLinodeListOK()))
	defer ts.Close()

	nodes, err := c.LinodeList(Int(1139016))
	require.NoError(t, err)
	require.Len(t, nodes, 1)

	n := nodes[0]

	assert.True(t, n.AlertCPUEnabled, "n.AlertCPUEnabled")
	assert.True(t, n.AlertBWInEnabled, "n.AlertBWINEnabled")
	assert.True(t, n.AlertBWQuotaEnabled, "n.AlertBWQuotaEnabled")
	assert.Equal(t, 10000, n.AlertDiskIOThreshold, "n.AlertDiskIOThreshold")
	assert.Equal(t, 0, n.BackupWindow, "n.BackupWindow")
	assert.Equal(t, true, n.Watchdog, "n.Watchdog")
	assert.Equal(t, "foo", n.DistVendor, "n.DistVendor")
	assert.Equal(t, 2, n.DatacenterID, "n.DatacenterID")
	assert.Equal(t, 1, n.Status, "n.Status")
	assert.True(t, n.AlertDiskIOEnabled, "n.AlertDiskIOEnabled")
	assert.Equal(t, "2015-07-02 23:08:52.0", n.CreateDT, "n.CreateDT")
	assert.Equal(t, 24576, n.TotalHD, "n.TotalHD")
	assert.Equal(t, 80, n.AlertBWQuotaThreshold, "n.AlertBWQUotaThreshold")
	assert.Equal(t, 1024, n.TotalRAM, "n.TotalRAM")
	assert.Equal(t, 10, n.AlertBWInThreshold, "n.AlertBWInThreshold")
	assert.Equal(t, 1139016, n.ID, "n.ID")
	assert.Equal(t, 10, n.AlertBWOutThreshold, "n.AlertBWOutThreshold")
	assert.True(t, n.AlertBWOutEnabled, "n.AlertBWOutEnabled")
	assert.True(t, n.BackupsEnabled, "n.BackupsEnabled")
	assert.Equal(t, 90, n.AlertCPUThreshold, "n.AlertCPUThreshold")
	assert.Equal(t, 1, n.PlanID, "n.PlanID")
	assert.Equal(t, 1, n.BackupWeeklyDay, "n.BackupWeeklyDay")
	assert.Equal(t, "linode1139016", n.Label, "n.Label")
	assert.Equal(t, "bar", n.DisplayGroup, "n.DisplayGroup")
	assert.Equal(t, 2000, n.TotalXfer, "n.TotalXfer")
}

func mockLinodeCloneOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"LinodeID":1140616},"ACTION":"linode.clone"}`
	params = map[string]string{
		"DatacenterID": "2",
		"LinodeID":     "1139016",
		"PlanID":       "1",
		"api_action":   "linode.clone",
		"api_key":      "foo",
	}
	responses = append(responses, newMockAPIResponse("linode.clone", params, output))

	return responses
}

func TestLinodeCloneOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockLinodeCloneOK()))
	defer ts.Close()

	cloneID, err := c.LinodeClone(1139016, 2, 1, nil, nil)
	require.NoError(t, err)
	require.NotEmpty(t, cloneID)
	require.Equal(t, 1140616, cloneID, "cloneID")
}

func mockLinodeDeleteOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"LinodeID":1140616},"ACTION":"linode.delete"}`
	params = map[string]string{
		"LinodeID":   "1140616",
		"api_action": "linode.delete",
		"api_key":    "foo",
		"skipChecks": "true",
	}
	responses = append(responses, newMockAPIResponse("linode.delete", params, output))

	return responses
}

func TestLinodeDeleteOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockLinodeDeleteOK()))
	defer ts.Close()

	err := c.LinodeDelete(1140616, Bool(true))
	require.NoError(t, err)
}

func mockLinodeResizeOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[{"ERRORCODE":0,"ERRORMESSAGE":"ok"}],"DATA":{},"ACTION":"linode.resize"}`
	params = map[string]string{
		"LinodeID":   "1139016",
		"PlanID":     "2",
		"api_action": "linode.resize",
		"api_key":    "foo",
	}
	responses = append(responses, newMockAPIResponse("linode.resize", params, output))

	return responses
}

func TestLinodeResizeOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockLinodeResizeOK()))
	defer ts.Close()

	err := c.LinodeResize(1139016, 2)
	require.NoError(t, err)
}

func mockLinodeUpdateOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"LinodeID":1139016},"ACTION":"linode.update"}`
	params = map[string]string{
		"Alert_bwin_enabled":      "false",
		"Alert_bwin_threshold":    "30",
		"Alert_bwout_enabled":     "false",
		"Alert_bwout_threshold":   "31",
		"Alert_bwquota_enabled":   "false",
		"Alert_bwquota_threshold": "55",
		"Alert_cpu_enabled":       "false",
		"Alert_cpu_threshold":     "50",
		"Alert_diskio_enabled":    "false",
		"Alert_diskio_threshold":  "500",
		"LinodeID":                "1139016",
		"api_action":              "linode.update",
		"api_key":                 "foo",
		"backupWeeklyDay":         "6",
		"backupWindow":            "3",
		"label":                   "unit-test",
		"lpm_displayGroup":        "tests",
		"watchdog":                "false",
	}
	responses = append(responses, newMockAPIResponse("linode.update", params, output))

	return responses
}

func TestLinodeUpdateOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockLinodeUpdateOK()))
	defer ts.Close()

	opts := LinodeOpts{
		Label:                 String("unit-test"),
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
	err := c.LinodeUpdate(1139016, opts)
	require.NoError(t, err)
}

func mockLinodeDiskCreateOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"JobID":25087627,"DiskID":3568984},"ACTION":"linode.disk.create"}`
	params = map[string]string{
		"Label":      "test-swap",
		"LinodeID":   "1139016",
		"Size":       "256",
		"Type":       "swap",
		"api_action": "linode.disk.create",
		"api_key":    "foo",
	}
	responses = append(responses, newMockAPIResponse("linode.disk.create", params, output))

	return responses
}

func TestLinodeDiskCreateOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockLinodeDiskCreateOK()))
	defer ts.Close()

	jobID, diskID, err := c.LinodeDiskCreate(1139016, "test-swap", "swap", 256)
	require.NoError(t, err)
	assert.Equal(t, 3568984, diskID)
	assert.Equal(t, 25087627, jobID)
}

func mockLinodeJobListNotFinished() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":[{"HOST_START_DT":"","HOST_MESSAGE":"","ENTERED_DT":"2015-07-03 23:24:12.0","HOST_FINISH_DT":"","LABEL":"Create Filesystem - test-swap","JOBID":25087627,"HOST_SUCCESS":"","ACTION":"fs.create","LINODEID":1139016,"DURATION":""}],"ACTION":"linode.job.list"}`
	params = map[string]string{
		"JobID":      "25088076",
		"LinodeID":   "1139016",
		"api_action": "linode.job.list",
		"api_key":    "foo",
	}
	responses = append(responses, newMockAPIResponse("linode.job.list", params, output))

	return responses
}

func TestLinodeJobListNotFinished(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockLinodeJobListNotFinished()))
	defer ts.Close()

	jobs, err := c.LinodeJobList(1139016, Int(25088076), nil)
	require.NoError(t, err)
	require.Len(t, jobs, 1)

	j := jobs[0]

	assert.False(t, j.Done())
	assert.False(t, j.Success())

	assert.Equal(t, "2015-07-03 23:24:12.0", j.EnteredDT)
	assert.Equal(t, "fs.create", j.Action)
	assert.Equal(t, "Create Filesystem - test-swap", j.Label)
	assert.Equal(t, "", j.HostStartDT)
	assert.Equal(t, 1139016, j.LinodeID)
	assert.Equal(t, "", j.HostFinishDT)
	assert.Equal(t, 0, j.Duration)
	assert.Equal(t, "", j.HostMessage)
	assert.Equal(t, 25087627, j.ID)
	assert.Equal(t, false, j.HostSuccess)
}

func mockLinodeJobListFinishedOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":[{"HOST_START_DT":"2015-07-03 23:51:51.0","HOST_MESSAGE":"foo","ENTERED_DT":"2015-07-03 23:51:41.0","HOST_FINISH_DT":"2015-07-03 23:51:51.0","LABEL":"Create Filesystem - test-swap","JOBID":25088076,"HOST_SUCCESS":1,"ACTION":"fs.create","LINODEID":1139016,"DURATION":5}],"ACTION":"linode.job.list"}`
	params = map[string]string{
		"JobID":      "25088076",
		"LinodeID":   "1139016",
		"api_action": "linode.job.list",
		"api_key":    "foo",
	}
	responses = append(responses, newMockAPIResponse("linode.job.list", params, output))

	return responses
}

func TestLinodeJobListFinishedOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockLinodeJobListFinishedOK()))
	defer ts.Close()

	jobs, err := c.LinodeJobList(1139016, Int(25088076), Bool(false))
	require.NoError(t, err)
	require.Len(t, jobs, 1)

	j := jobs[0]

	assert.True(t, j.Done())
	assert.True(t, j.Success())

	assert.Equal(t, "2015-07-03 23:51:41.0", j.EnteredDT)
	assert.Equal(t, "fs.create", j.Action)
	assert.Equal(t, "Create Filesystem - test-swap", j.Label)
	assert.Equal(t, "2015-07-03 23:51:51.0", j.HostStartDT)
	assert.Equal(t, 1139016, j.LinodeID)
	assert.Equal(t, "2015-07-03 23:51:51.0", j.HostFinishDT)
	assert.Equal(t, 5, j.Duration)
	assert.Equal(t, "foo", j.HostMessage)
	assert.Equal(t, 25088076, j.ID)
	assert.Equal(t, true, j.HostSuccess)
}

func mockLinodeJobListPendingOnlyFalse() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":[{"HOST_START_DT":"2015-07-03 23:51:51.0","HOST_MESSAGE":"foo","ENTERED_DT":"2015-07-03 23:51:41.0","HOST_FINISH_DT":"2015-07-03 23:51:51.0","LABEL":"Create Filesystem - test-swap","JOBID":25088076,"HOST_SUCCESS":1,"ACTION":"fs.create","LINODEID":1139016,"DURATION":5}],"ACTION":"linode.job.list"}`
	params = map[string]string{
		"JobID":       "25088076",
		"LinodeID":    "1139016",
		"pendingOnly": "0",
		"api_action":  "linode.job.list",
		"api_key":     "foo",
	}
	responses = append(responses, newMockAPIResponse("linode.job.list", params, output))

	return responses
}

func TestLinodeJobListPendingOnlyFalse(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockLinodeJobListPendingOnlyFalse()))
	defer ts.Close()

	jobs, err := c.LinodeJobList(1139016, Int(25088076), Bool(false))
	require.NoError(t, err)
	require.NotEmpty(t, jobs)
}

func mockLinodeDiskCreateFromDistributionOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"JobID":25088891,"DiskID":3569234},"ACTION":"linode.disk.createfromdistribution"}`
	params = map[string]string{
		"DistributionID": "130",
		"Label":          "test-dist",
		"LinodeID":       "1139016",
		"Size":           "600",
		"api_action":     "linode.disk.createfromdistribution",
		"api_key":        "foo",
		"rootPass":       "foo#23113.",
		"rootSSHKey":     "ssh-rsa foobarbaz",
	}
	responses = append(responses, newMockAPIResponse("linode.disk.createfromdistribution", params, output))

	return responses
}

func TestLinodeDiskCreateFromDistributionOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockLinodeDiskCreateFromDistributionOK()))
	defer ts.Close()

	jobID, diskID, err := c.LinodeDiskCreateFromDistribution(1139016, 130, "test-dist", 600,
		rootPass, String(rootSSHKey))
	require.NoError(t, err)
	require.Equal(t, 25088891, jobID)
	require.Equal(t, 3569234, diskID)
}

func mockLinodeConfigCreateOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"ConfigID":1855685},"ACTION":"linode.config.create"}`
	params = map[string]string{
		"Comments":               "foo",
		"DiskList":               "3569234,3569220,",
		"KernelID":               "138",
		"Label":                  "test-conf1",
		"LinodeID":               "1139016",
		"RAMLimit":               "800",
		"RootDeviceNum":          "1",
		"RootDeviceRO":           "true",
		"RunLevel":               "default",
		"api_action":             "linode.config.create",
		"api_key":                "foo",
		"devtmpfs_automount":     "true",
		"helper_depmod":          "true",
		"helper_disableUpdateDB": "true",
		"helper_distro":          "true",
		"helper_network":         "true",
		"helper_xen":             "true",
		"virt_mode":              "paravirt",
	}
	responses = append(responses, newMockAPIResponse("linode.config.create", params, output))

	return responses
}

func TestLinodeConfigCreateOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockLinodeConfigCreateOK()))
	defer ts.Close()

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

	confID, err := c.LinodeConfigCreate(1139016, 138, "test-conf1", "3569234,3569220,", lcco)
	require.NoError(t, err)
	require.Equal(t, 1855685, confID)
}

func mockLinodeConfigListOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":[{"RootDeviceCustom":"bar","Comments":"foo","isRescue":1,"devtmpfs_automount":true,"helper_distro":1,"helper_disableUpdateDB":1,"Label":"test-conf1","helper_network":1,"ConfigID":1855685,"DiskList":"3569234,3569220,,,,,,,","RootDeviceRO":true,"RunLevel":"default","helper_libtls":0,"__validationErrorArray":[],"apiColumnFilterStruct":"","RootDeviceNum":1,"helper_xen":1,"RAMLimit":800,"virt_mode":"paravirt","LinodeID":1139016,"helper_depmod":1,"KernelID":138}],"ACTION":"linode.config.list"}`
	params = map[string]string{
		"ConfigID":   "1855685",
		"LinodeID":   "1139016",
		"api_action": "linode.config.list",
		"api_key":    "foo",
	}
	responses = append(responses, newMockAPIResponse("linode.config.list", params, output))

	return responses
}

func TestLinodeConfigList(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockLinodeConfigListOK()))
	defer ts.Close()

	confs, err := c.LinodeConfigList(1139016, Int(1855685))
	require.NoError(t, err)
	require.Len(t, confs, 1)

	cfg := confs[0]
	assert.Equal(t, "bar", cfg.RootDeviceCustom)
	assert.Equal(t, "foo", cfg.Comments)
	assert.Equal(t, "test-conf1", cfg.Label)
	assert.Equal(t, 1855685, cfg.ID)
	assert.Equal(t, "3569234,3569220,,,,,,,", cfg.DiskList)
	assert.Equal(t, "default", cfg.RunLevel)
	assert.Equal(t, 1, cfg.RootDeviceNum)
	assert.Equal(t, 800, cfg.RAMLimit)
	assert.Equal(t, "paravirt", cfg.VirtMode)
	assert.Equal(t, 1139016, cfg.LinodeID)
	assert.Equal(t, 138, cfg.KernelID)
	assert.True(t, cfg.DevTmpFSAutomount)
	assert.True(t, cfg.HelperDistro)
	assert.True(t, cfg.HelperDisableUpdateDB)
	assert.True(t, cfg.HelperDepmod)
	assert.True(t, cfg.HelperXen)
	assert.True(t, cfg.RootDeviceRO)
	assert.True(t, cfg.HelperNetwork)
	assert.True(t, cfg.IsRescue)
}

func mockLinodeConfigUpdateOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"ConfigID":1855685},"ACTION":"linode.config.update"}`
	params = map[string]string{
		"Comments":               "foo",
		"ConfigID":               "1855685",
		"KernelID":               "138",
		"LinodeID":               "1139016",
		"RAMLimit":               "800",
		"RootDeviceNum":          "1",
		"RootDeviceRO":           "true",
		"RunLevel":               "default",
		"api_action":             "linode.config.update",
		"api_key":                "foo",
		"devtmpfs_automount":     "true",
		"helper_depmod":          "true",
		"helper_disableUpdateDB": "true",
		"helper_distro":          "true",
		"helper_network":         "true",
		"helper_xen":             "true",
		"virt_mode":              "paravirt",
	}
	responses = append(responses, newMockAPIResponse("linode.config.update", params, output))

	return responses
}

func TestLinodeConfigUpdateOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockLinodeConfigUpdateOK()))
	defer ts.Close()

	lcuo := LinodeConfigUpdateOpts{
		LinodeID:              Int(1139016),
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
	err := c.LinodeConfigUpdate(1855685, lcuo)
	require.NoError(t, err)
}

func mockLinodeDiskImagizeOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"JobID":25090408,"ImageID":396235},"ACTION":"linode.disk.imagize"}`
	params = map[string]string{
		"Description": "test-image desc",
		"DiskID":      "3569234",
		"Label":       "test-image label",
		"LinodeID":    "1139016",
		"api_action":  "linode.disk.imagize",
		"api_key":     "foo",
	}
	responses = append(responses, newMockAPIResponse("linode.disk.imagize", params, output))

	return responses
}

func TestLinodeDiskImagizeOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockLinodeDiskImagizeOK()))
	defer ts.Close()

	jobID, imgID, err := c.LinodeDiskImagize(1139016, 3569234, String("test-image desc"),
		String("test-image label"))

	require.NoError(t, err)
	assert.Equal(t, 25090408, jobID)
	assert.Equal(t, 396235, imgID)
}

func mockLinodeDiskCreateFromImageOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"JOBID":25090431,"DISKID":3569532},"ACTION":"linode.disk.createfromimage"}`
	params = map[string]string{
		"ImageID":    "396235",
		"Label":      "test-image",
		"LinodeID":   "1139016",
		"api_action": "linode.disk.createfromimage",
		"api_key":    "foo",
		"rootPass":   "foo#23113.",
		"rootSSHKey": "ssh-rsa foobarbaz",
		"size":       "800",
	}
	responses = append(responses, newMockAPIResponse("linode.disk.createfromimage", params, output))

	return responses
}

func TestLinodeDiskCreateFromImageOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockLinodeDiskCreateFromImageOK()))
	defer ts.Close()

	jobID, imgDiskID, err := c.LinodeDiskCreateFromImage(396235, 1139016, "test-image", Int(800),
		String(rootPass), String(rootSSHKey))
	require.NoError(t, err)
	assert.Equal(t, 25090431, jobID)
	assert.Equal(t, 3569532, imgDiskID)
}

func mockLinodeDiskDeleteOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"JobID":25090443,"DiskID":3569532},"ACTION":"linode.disk.delete"}`
	params = map[string]string{
		"DiskID":     "3569532",
		"LinodeID":   "1139016",
		"api_action": "linode.disk.delete",
		"api_key":    "foo",
	}
	responses = append(responses, newMockAPIResponse("linode.disk.delete", params, output))

	return responses
}

func TestLinodeDiskDeleteOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockLinodeDiskDeleteOK()))
	defer ts.Close()

	jobID, err := c.LinodeDiskDelete(1139016, 3569532)
	require.NoError(t, err)
	assert.Equal(t, 25090443, jobID)
}

func mockLinodeDiskDuplicateOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"JobID":25090699,"DiskID":3569577},"ACTION":"linode.disk.duplicate"}`
	params = map[string]string{
		"DiskID":     "3569220",
		"LinodeID":   "1139016",
		"api_action": "linode.disk.duplicate",
		"api_key":    "foo",
	}
	responses = append(responses, newMockAPIResponse("linode.disk.duplicate", params, output))

	return responses
}

func TestLinodeDiskDuplicateOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockLinodeDiskDuplicateOK()))
	defer ts.Close()

	jobID, dupeDiskID, err := c.LinodeDiskDuplicate(1139016, 3569220)
	require.NoError(t, err)
	assert.Equal(t, 25090699, jobID)
	assert.Equal(t, 3569577, dupeDiskID)
}

func mockLinodeDiskResizeOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"JobID":25090703,"DiskID":3569577},"ACTION":"linode.disk.resize"}`
	params = map[string]string{
		"DiskID":     "3569577",
		"LinodeID":   "1139016",
		"api_action": "linode.disk.resize",
		"api_key":    "foo",
		"size":       "512",
	}
	responses = append(responses, newMockAPIResponse("linode.disk.resize", params, output))

	return responses
}

func TestLinodeDiskResizeOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockLinodeDiskResizeOK()))
	defer ts.Close()

	jobID, err := c.LinodeDiskResize(1139016, 3569577, 512)
	require.NoError(t, err)
	assert.Equal(t, 25090703, jobID)
}

func mockLinodeDiskUpdateOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"DiskID":3569577},"ACTION":"linode.disk.update"}`
	params = map[string]string{
		"DiskID":     "3569577",
		"Label":      "updated-label",
		"LinodeID":   "1139016",
		"api_action": "linode.disk.update",
		"api_key":    "foo",
		"isReadOnly": "true",
	}
	responses = append(responses, newMockAPIResponse("linode.disk.update", params, output))

	return responses
}

func TestLinodeDiskUpdateOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockLinodeDiskUpdateOK()))
	defer ts.Close()

	err := c.LinodeDiskUpdate(1139016, 3569577, String("updated-label"), Bool(true))
	require.NoError(t, err)
}

func mockLinodeDiskListOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":[{"ISREADONLY":1,"LABEL":"test-swap","UPDATE_DT":"2015-07-06 23:30:13.0","STATUS":1,"SIZE":256,"LINODEID":1146420,"CREATE_DT":"2015-07-06 23:29:37.0","TYPE":"swap","DISKID":3582351},{"ISREADONLY":0,"LABEL":"test-dist","UPDATE_DT":"2015-07-06 23:30:32.0","STATUS":1,"SIZE":600,"LINODEID":1146420,"CREATE_DT":"2015-07-06 23:30:14.0","TYPE":"ext4","DISKID":3582352},{"ISREADONLY":0,"LABEL":"updated-label","UPDATE_DT":"2015-07-06 23:32:45.0","STATUS":1,"SIZE":512,"LINODEID":1146420,"CREATE_DT":"2015-07-06 23:32:19.0","TYPE":"swap","DISKID":3582356}],"ACTION":"linode.disk.list"}`
	params = map[string]string{
		"LinodeID":   "1146420",
		"DiskID":     "1",
		"api_action": "linode.disk.list",
		"api_key":    "foo",
	}
	responses = append(responses, newMockAPIResponse("linode.disk.list", params, output))

	return responses
}

func TestLinodeDiskListOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockLinodeDiskListOK()))
	defer ts.Close()

	disks, err := c.LinodeDiskList(1146420, Int(1))
	require.NoError(t, err)
	require.Len(t, disks, 3)

	var foundTestSwap bool
	var foundTestDist bool
	var foundUpdatedLabel bool

	for _, d := range disks {
		if d.ID == 3582351 {
			foundTestSwap = true

			assert.True(t, d.IsReadOnly)
			assert.Equal(t, "test-swap", d.Label)
			assert.Equal(t, "2015-07-06 23:30:13.0", d.UpdateDT)
			assert.Equal(t, 1, d.Status)
			assert.Equal(t, 256, d.Size)
			assert.Equal(t, 1146420, d.LinodeID)
			assert.Equal(t, "2015-07-06 23:29:37.0", d.CreateDT)
			assert.Equal(t, "swap", d.Type)
		}
		if d.ID == 3582352 {
			foundTestDist = true
		}
		if d.ID == 3582356 {
			foundUpdatedLabel = true
		}
	}

	assert.True(t, foundTestSwap)
	assert.True(t, foundTestDist)
	assert.True(t, foundUpdatedLabel)
}

func mockLinodeIPListOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":[{"IPADDRESSID":296963,"RDNS_NAME":"li959-147.members.linode.com","LINODEID":1146420,"ISPUBLIC":1,"IPADDRESS":"45.33.5.147"}],"ACTION":"linode.ip.list"}`
	params = map[string]string{
		"LinodeID":    "1146420",
		"IPAddressID": "1",
		"api_action":  "linode.ip.list",
		"api_key":     "foo",
	}
	responses = append(responses, newMockAPIResponse("linode.ip.list", params, output))

	return responses
}

func TestLinodeIPListOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockLinodeIPListOK()))
	defer ts.Close()

	IPList, err := c.LinodeIPList(Int(1146420), Int(1))
	require.NoError(t, err)
	require.Len(t, IPList, 1)

	IP := IPList[0]
	assert.Equal(t, 296963, IP.ID)
	assert.Equal(t, "li959-147.members.linode.com", IP.RDNSName)
	assert.Equal(t, 1146420, IP.LinodeID)
	assert.True(t, IP.IsPublic)
	assert.Equal(t, "45.33.5.147", IP.Address)
}

func mockLinodeIPAddPrivateOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"IPADDRESSID":374332,"IPADDRESS":"192.168.199.65"},"ACTION":"linode.ip.addprivate"}`
	params = map[string]string{
		"LinodeID":   "1146420",
		"api_action": "linode.ip.addprivate",
		"api_key":    "foo",
	}
	responses = append(responses, newMockAPIResponse("linode.ip.addprivate", params, output))

	return responses
}

func TestLinodeIPAddPrivateOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockLinodeIPAddPrivateOK()))
	defer ts.Close()

	id, addr, err := c.LinodeIPAddPrivate(1146420)
	require.NoError(t, err)
	assert.Equal(t, 374332, id)
	assert.Equal(t, "192.168.199.65", addr)
}

func mockLinodeIPSwapOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":[{"IPADDRESSID":296963,"LINODEID":1146530,"IPADDRESS":"45.33.5.147"},{"IPADDRESSID":296386,"LINODEID":1146420,"IPADDRESS":"45.33.3.72"}],"ACTION":"linode.ip.swap"}`
	params = map[string]string{
		"IPAddressID":     "296963",
		"api_action":      "linode.ip.swap",
		"api_key":         "foo",
		"withIPAddressID": "296386",
		"toLinodeID":      "1",
	}
	responses = append(responses, newMockAPIResponse("linode.ip.swap", params, output))

	return responses
}

func TestLinodeIPSwap(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockLinodeIPSwapOK()))
	defer ts.Close()

	err := c.LinodeIPSwap(296963, Int(296386), Int(1))
	require.NoError(t, err)
}

func mockLinodeDiskCreateFromStackScriptOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"JobID":25166896,"DiskID":3582660},"ACTION":"linode.disk.createfromstackscript"}`
	params = map[string]string{
		"DistributionID": "130",
		"Label":          "test-ss",
		"LinodeID":       "1146420",
		"Size":           "600",
		"StackScriptUDFResponses": "{\"db_password\":\"foo\",\"db_name\":\"bar\",\"db_user\":\"baz\",\"db_user_password\":\"quux\"}",
		"StackScriptID":           "10",
		"api_action":              "linode.disk.createfromstackscript",
		"api_key":                 "foo",
		"rootPass":                "foo#23113.",
		"rootSSHKey":              "ssh-rsa foobarbaz",
	}
	responses = append(responses, newMockAPIResponse("linode.disk.createfromstackscript", params, output))

	return responses
}

func TestLinodeDiskCreateFromStackScriptOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockLinodeDiskCreateFromStackScriptOK()))
	defer ts.Close()

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

	jobID, diskID, err := c.LinodeDiskCreateFromStackScript(1146420, 10, string(ssUDFResp),
		130, "test-ss", 600, rootPass, String(rootSSHKey))
	require.NoError(t, err)
	assert.Equal(t, 25166896, jobID)
	assert.Equal(t, 3582660, diskID)
}

func mockLinodeBootOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"JobID":25167133},"ACTION":"linode.boot"}`
	params = map[string]string{
		"ConfigID":   "1862370",
		"LinodeID":   "1146420",
		"api_action": "linode.boot",
		"api_key":    "foo",
	}
	responses = append(responses, newMockAPIResponse("linode.boot", params, output))

	return responses
}

func TestLinodeBootOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockLinodeBootOK()))
	defer ts.Close()

	jobID, err := c.LinodeBoot(1146420, Int(1862370))
	require.NoError(t, err)
	assert.Equal(t, 25167133, jobID)
}

func mockLinodeRebootOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"JobID":25167140},"ACTION":"linode.reboot"}`
	params = map[string]string{
		"ConfigID":   "1862370",
		"LinodeID":   "1146420",
		"api_action": "linode.reboot",
		"api_key":    "foo",
	}
	responses = append(responses, newMockAPIResponse("linode.reboot", params, output))

	return responses
}

func TestLinodeRebootOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockLinodeRebootOK()))
	defer ts.Close()

	jobID, err := c.LinodeReboot(1146420, Int(1862370))
	require.NoError(t, err)
	assert.Equal(t, jobID, 25167140)
}

func mockLinodeShutdownOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"JobID":25167147},"ACTION":"linode.shutdown"}`
	params = map[string]string{
		"LinodeID":   "1146420",
		"api_action": "linode.shutdown",
		"api_key":    "foo",
	}
	responses = append(responses, newMockAPIResponse("linode.shutdown", params, output))

	return responses
}

func TestLinodeShutdownOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockLinodeShutdownOK()))
	defer ts.Close()

	jobID, err := c.LinodeShutdown(1146420)
	require.NoError(t, err)
	assert.Equal(t, 25167147, jobID)
}

func mockLinodeConfigDeleteOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":{"ConfigID":1862591},"ACTION":"linode.config.delete"}`
	params = map[string]string{
		"ConfigID":   "1862591",
		"LinodeID":   "1146666",
		"api_action": "linode.config.delete",
		"api_key":    "foo",
	}
	responses = append(responses, newMockAPIResponse("linode.config.delete", params, output))

	return responses
}

func TestLinodeConfigDeleteOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockLinodeConfigDeleteOK()))
	defer ts.Close()

	err := c.LinodeConfigDelete(1146666, 1862591)
	require.NoError(t, err)
}

func TestWaitForJobOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockLinodeJobListFinishedOK()))
	defer ts.Close()

	ok, err := c.WaitForJob(1139016, 25088076, 1*time.Nanosecond, 1*time.Second)
	require.NoError(t, err)
	require.True(t, ok)
}

func mockWaitforJobMultiJobs() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":[{"HOST_START_DT":"2015-07-03 23:51:51.0","HOST_MESSAGE":"foo","ENTERED_DT":"2015-07-03 23:51:41.0","HOST_FINISH_DT":"2015-07-03 23:51:51.0","LABEL":"Create Filesystem - test-swap","JOBID":25088076,"HOST_SUCCESS":1,"ACTION":"fs.create","LINODEID":1139016,"DURATION":5},{"HOST_START_DT":"2015-07-03 23:51:51.0","HOST_MESSAGE":"foo","ENTERED_DT":"2015-07-03 23:51:41.0","HOST_FINISH_DT":"2015-07-03 23:51:51.0","LABEL":"Create Filesystem - test-swap","JOBID":25088076,"HOST_SUCCESS":1,"ACTION":"fs.create","LINODEID":1139016,"DURATION":5}],"ACTION":"linode.job.list"}`
	params = map[string]string{
		"JobID":      "25088076",
		"LinodeID":   "1139016",
		"api_action": "linode.job.list",
		"api_key":    "foo",
	}
	responses = append(responses, newMockAPIResponse("linode.job.list", params, output))

	return responses
}

func TestWaitForJobMultiJobs(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockWaitforJobMultiJobs()))
	defer ts.Close()

	ok, err := c.WaitForJob(1139016, 25088076, 1*time.Nanosecond, 1*time.Second)
	require.Error(t, err)
	require.False(t, ok)
}

func TestWaitForJobTimeout(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, []mockAPIResponse{}))
	defer ts.Close()

	ok, err := c.WaitForJob(0, 0, 1*time.Second, 1*time.Nanosecond)
	require.Error(t, err)
	require.False(t, ok)
}

func mockWaitForAllJobsOK() []mockAPIResponse {
	var output string
	var params map[string]string
	var responses []mockAPIResponse

	output = `{"ERRORARRAY":[],"DATA":[],"ACTION":"linode.job.list"}`
	params = map[string]string{
		"LinodeID":    "1139016",
		"pendingOnly": "1",
		"api_action":  "linode.job.list",
		"api_key":     "foo",
	}
	responses = append(responses, newMockAPIResponse("linode.job.list", params, output))

	return responses
}

func TestWaitForAllJobsOK(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, mockWaitForAllJobsOK()))
	defer ts.Close()

	err := c.WaitForAllJobs(1139016, 1*time.Nanosecond, 1*time.Second)
	require.NoError(t, err)
}

func TestWaitForAllJobsTimeout(t *testing.T) {
	c, ts := clientFor(newMockAPIServer(t, []mockAPIResponse{}))
	defer ts.Close()

	err := c.WaitForAllJobs(0, 1*time.Second, 1*time.Nanosecond)
	require.Error(t, err)
}
