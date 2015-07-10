package linode

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// LinodeBoot maps to the 'linode.boot' call.
//
// https://www.linode.com/api/linode/linode.boot
func (c *Client) LinodeBoot(linodeID int, configID *int) (jobID int, err error) {
	args := make(map[string]interface{})
	args["LinodeID"] = linodeID
	args["ConfigID"] = configID

	data, err := c.apiCall("linode.boot", args)
	if err != nil {
		return 0, err
	}

	err = unmarshalSingle(data, "JobID", &jobID)
	if err != nil {
		return 0, err
	}

	return jobID, nil
}

// LinodeClone maps to the 'linode.clone' call.
//
// https://www.linode.com/api/linode/linode.clone
func (c *Client) LinodeClone(linodeID int, datacenterID int, planID int, term *int,
	hypervisor *string) (cloneLinodeID int, err error) {

	args := make(map[string]interface{})
	args["LinodeID"] = linodeID
	args["DatacenterID"] = datacenterID
	args["PlanID"] = planID
	args["PaymentTerm"] = term
	args["hypervisor"] = hypervisor

	data, err := c.apiCall("linode.clone", args)
	if err != nil {
		return 0, err
	}

	err = unmarshalSingle(data, "LinodeID", &cloneLinodeID)
	if err != nil {
		return 0, err
	}

	return cloneLinodeID, nil
}

// LinodeCreate maps to the 'linode.create' call.
//
// https://www.linode.com/api/linode/linode.create
func (c *Client) LinodeCreate(datacenterID int, planID int,
	term *int) (linodeID int, err error) {

	args := make(map[string]interface{})
	args["DatacenterID"] = datacenterID
	args["PlanID"] = planID
	args["PaymentTerm"] = term

	data, err := c.apiCall("linode.create", args)
	if err != nil {
		return 0, err
	}

	err = unmarshalSingle(data, "LinodeID", &linodeID)
	if err != nil {
		return 0, err
	}

	return linodeID, nil
}

// LinodeDelete maps to the 'linode.delete' call.
//
// https://www.linode.com/api/linode/linode.delete
func (c *Client) LinodeDelete(linodeID int, skipChecks *bool) error {
	args := make(map[string]interface{})
	args["LinodeID"] = linodeID
	args["skipChecks"] = skipChecks

	_, err := c.apiCall("linode.delete", args)
	if err != nil {
		return err
	}

	return nil
}

// Linode is the API response to the 'linode.list' call.
type Linode struct {
	TotalXfer             int    `mapstructure:"TOTALXFER"`
	BackupsEnabled        bool   `mapstructure:"BACKUPSENABLED"`
	Watchdog              bool   `mapstructure:"WATCHDOG"`
	DisplayGroup          string `mapstructure:"LPM_DISPLAYGROUP"`
	Status                int    `mapstructure:"STATUS"`
	TotalRAM              int    `mapstructure:"TOTALRAM"`
	BackupWindow          int    `mapstructure:"BACKUPWINDOW"`
	Label                 string `mapstructure:"LABEL"`
	BackupWeeklyDay       int    `mapstructure:"BACKUPWEEKLYDAY"`
	DatacenterID          int    `mapstructure:"DATACENTERID"`
	TotalHD               int    `mapstructure:"TOTALHD"`
	ID                    int    `mapstructure:"LINODEID"`
	CreateDT              string `mapstructure:"CREATE_DT"`
	PlanID                int    `mapstructure:"PLANID"`
	DistVendor            string `mapstructure:"DISTRIBUTIONVENDOR"`
	AlertBWQuotaEnabled   bool   `mapstructure:"ALERT_BWQUOTA_ENABLED"`
	AlertBWQuotaThreshold int    `mapstructure:"ALERT_BWQUOTA_THRESHOLD"`
	AlertDiskIOEnabled    bool   `mapstructure:"ALERT_DISKIO_ENABLED"`
	AlertDiskIOThreshold  int    `mapstructure:"ALERT_DISKIO_THRESHOLD"`
	AlertCPUEnabled       bool   `mapstructure:"ALERT_CPU_ENABLED"`
	AlertCPUThreshold     int    `mapstructure:"ALERT_CPU_THRESHOLD"`
	AlertBWInEnabled      bool   `mapstructure:"ALERT_BWIN_ENABLED"`
	AlertBWInThreshold    int    `mapstructure:"ALERT_BWIN_THRESHOLD"`
	AlertBWOutEnabled     bool   `mapstructure:"ALERT_BWOUT_ENABLED"`
	AlertBWOutThreshold   int    `mapstructure:"ALERT_BWOUT_THRESHOLD"`
}

// LinodeList maps to the 'linode.list' call.
//
// https://www.linode.com/api/linode/linode.list
func (c *Client) LinodeList(linodeID *int) ([]Linode, error) {
	args := make(map[string]interface{})
	args["LinodeID"] = linodeID

	data, err := c.apiCall("linode.list", args)
	if err != nil {
		return nil, err
	}

	var out []Linode
	err = unmarshalMultiMap(data, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// LinodeReboot maps to the 'linode.reboot' call.
//
// https://www.linode.com/api/linode/linode.reboot
func (c *Client) LinodeReboot(linodeID int, configID *int) (jobID int, err error) {
	args := make(map[string]interface{})
	args["LinodeID"] = linodeID
	args["ConfigID"] = configID

	data, err := c.apiCall("linode.reboot", args)
	if err != nil {
		return 0, err
	}

	err = unmarshalSingle(data, "JobID", &jobID)
	if err != nil {
		return 0, err
	}

	return jobID, nil
}

// LinodeResize maps to the 'linode.resize' call.
//
// https://www.linode.com/api/linode/linode.resize
func (c *Client) LinodeResize(linodeID int, planID int) error {
	args := make(map[string]interface{})
	args["LinodeID"] = linodeID
	args["PlanID"] = planID

	_, err := c.apiCall("linode.resize", args)
	if err != nil {
		return err
	}
	return nil
}

// LinodeShutdown maps to the 'linode.shutdown' call.
//
// https://www.linode.com/api/linode/linode.shutdown
func (c *Client) LinodeShutdown(linodeID int) (jobID int, err error) {
	args := make(map[string]interface{})
	args["LinodeID"] = linodeID

	data, err := c.apiCall("linode.shutdown", args)
	if err != nil {
		return 0, err
	}

	err = unmarshalSingle(data, "JobID", &jobID)
	if err != nil {
		return 0, err
	}

	return jobID, nil
}

// LinodeOpts contains the optional arguments to LinodeUpdate().
type LinodeOpts struct {
	Label                 *string `args:"label"`
	DisplayGroup          *string `args:"lpm_displayGroup"`
	AlertCPUEnabled       *bool   `args:"Alert_cpu_enabled"`
	AlertCPUThreshold     *int    `args:"Alert_cpu_threshold"`
	AlertDiskIOEnabled    *bool   `args:"Alert_diskio_enabled"`
	AlertDiskIOThreshold  *int    `args:"Alert_diskio_threshold"`
	AlertBWInEnabled      *bool   `args:"Alert_bwin_enabled"`
	AlertBWInThreshold    *int    `args:"Alert_bwin_threshold"`
	AlertBWOutEnabled     *bool   `args:"Alert_bwout_enabled"`
	AlertBWOutThreshold   *int    `args:"Alert_bwout_threshold"`
	AlertBWQuotaEnabled   *bool   `args:"Alert_bwquota_enabled"`
	AlertBWQuotaThreshold *int    `args:"Alert_bwquota_threshold"`
	BackupWindow          *int    `args:"backupWindow"`
	BackupWeeklyDay       *int    `args:"backupWeeklyDay"`
	Watchdog              *bool   `args:"watchdog"`
}

// LinodeUpdate maps to the 'linode.update' call.
//
// https://www.linode.com/api/linode/linode.update
func (c *Client) LinodeUpdate(linodeID int, conf LinodeOpts) error {
	args, err := c.argMarshal(conf)
	if err != nil {
		return err
	}
	args["LinodeID"] = linodeID

	_, err = c.apiCall("linode.update", args)
	if err != nil {
		return err
	}
	return nil
}

// LinodeConfigCreateOpts contains the optional arguments to
// LinodeConfigCreate().
type LinodeConfigCreateOpts struct {
	Comments              *string `args:"Comments"`
	RAMLimit              *int    `args:"RAMLimit"`
	VirtMode              *string `args:"virt_mode"`
	RunLevel              *string `args:"RunLevel"`
	RootDeviceNum         *int    `args:"RootDeviceNum"`
	RootDeviceCustom      *string `args:"RootDeviceCustom"`
	RootDeviceRO          *bool   `args:"RootDeviceRO"`
	HelperDisableUpdateDB *bool   `args:"helper_disableUpdateDB"`
	HelperDistro          *bool   `args:"helper_distro"`
	HelperXen             *bool   `args:"helper_xen"`
	HelperDepmod          *bool   `args:"helper_depmod"`
	HelperNetwork         *bool   `args:"helper_network"`
	DevTmpFSAutomount     *bool   `args:"devtmpfs_automount"`
}

// LinodeConfigCreate maps to the 'linode.config.create' call.
//
// https://www.linode.com/api/linode/linode.config.create
func (c *Client) LinodeConfigCreate(linodeID int, kernelID int, label string,
	diskList string, conf LinodeConfigCreateOpts) (configID int, err error) {

	args, err := c.argMarshal(conf)
	if err != nil {
		return 0, err
	}
	args["LinodeID"] = linodeID
	args["KernelID"] = kernelID
	args["Label"] = label
	args["DiskList"] = diskList

	data, err := c.apiCall("linode.config.create", args)
	if err != nil {
		return 0, err
	}

	err = unmarshalSingle(data, "ConfigID", &configID)
	if err != nil {
		return 0, err
	}

	return configID, nil
}

// LinodeConfigDelete maps to the 'linode.config.delete' call.
//
// https://www.linode.com/api/linode/linode.config.delete
func (c *Client) LinodeConfigDelete(linodeID int, configID int) error {
	args := make(map[string]interface{})
	args["LinodeID"] = linodeID
	args["ConfigID"] = configID

	_, err := c.apiCall("linode.config.delete", args)
	if err != nil {
		return err
	}

	return nil
}

// LinodeConfig is the API response to the 'linode.config.list' call.
type LinodeConfig struct {
	RootDeviceCustom      string `mapstructure:"RootDeviceCustom"`
	Comments              string `mapstructure:"Comments"`
	IsRescue              bool   `mapstructure:"isRescue"`
	DevTmpFSAutomount     bool   `mapstructure:"devtmpfs_automount"`
	HelperDistro          bool   `mapstructure:"helper_distro"`
	HelperDisableUpdateDB bool   `mapstructure:"helper_disableUpdateDB"`
	Label                 string `mapstructure:"label"`
	HelperNetwork         bool   `mapstructure:"helper_network"`
	ID                    int    `mapstructure:"ConfigID"`
	DiskList              string `mapstructure:"DiskList"`
	RootDeviceRO          bool   `mapstructure:"RootDeviceRO"`
	RunLevel              string `mapstructure:"RunLevel"`
	RootDeviceNum         int    `mapstructure:"RootDeviceNum"`
	HelperXen             bool   `mapstructure:"helper_xen"`
	RAMLimit              int    `mapstructure:"RAMLimit"`
	VirtMode              string `mapstructure:"virt_mode"`
	LinodeID              int    `mapstructure:"LinodeID"`
	HelperDepmod          bool   `mapstructure:"helper_depmod"`
	KernelID              int    `mapstructure:"KernelID"`
}

// LinodeConfigList maps to the 'linode.config.list' call.
//
// https://www.linode.com/api/linode/linode.config.list
func (c *Client) LinodeConfigList(linodeID int, configID *int) ([]LinodeConfig, error) {
	args := make(map[string]interface{})
	args["LinodeID"] = linodeID
	args["ConfigID"] = configID

	data, err := c.apiCall("linode.config.list", args)
	if err != nil {
		return nil, err
	}

	var ret []LinodeConfig
	err = unmarshalMultiMap(data, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// LinodeConfigUpdateOpts contains the optional arguments to
// LinodeConfigUpdate().
type LinodeConfigUpdateOpts struct {
	LinodeID              *int    `args:"LinodeID"`
	KernelID              *int    `args:"KernelID"`
	Comments              *string `args:"Comments"`
	RAMLimit              *int    `args:"RAMLimit"`
	VirtMode              *string `args:"virt_mode"`
	RunLevel              *string `args:"RunLevel"`
	RootDeviceNum         *int    `args:"RootDeviceNum"`
	RootDeviceCustom      *string `args:"RootDeviceCustom"`
	RootDeviceRO          *bool   `args:"RootDeviceRO"`
	HelperDisableUpdateDB *bool   `args:"helper_disableUpdateDB"`
	HelperDistro          *bool   `args:"helper_distro"`
	HelperXen             *bool   `args:"helper_xen"`
	HelperDepmod          *bool   `args:"helper_depmod"`
	HelperNetwork         *bool   `args:"helper_network"`
	DevTmpFSAutomount     *bool   `args:"devtmpfs_automount"`
}

// LinodeConfigUpdate maps to the 'linode.config.update' call.
//
// https://www.linode.com/api/linode/linode.config.update
func (c *Client) LinodeConfigUpdate(configID int, conf LinodeConfigUpdateOpts) error {
	args, err := c.argMarshal(conf)
	if err != nil {
		return err
	}
	args["ConfigID"] = configID

	_, err = c.apiCall("linode.config.update", args)
	if err != nil {
		return err
	}

	return nil
}

// LinodeDiskCreate maps to the 'linode.disk.create' call.
//
// https://www.linode.com/api/linode/linode.disk.create
func (c *Client) LinodeDiskCreate(linodeID int, label string, dType string,
	size int) (jobID int, diskID int, err error) {

	args := make(map[string]interface{})
	args["LinodeID"] = linodeID
	args["Label"] = label
	args["Type"] = dType
	args["Size"] = size

	data, err := c.apiCall("linode.disk.create", args)
	if err != nil {
		return 0, 0, err
	}

	out := struct {
		JobID  int `json:"JobID"`
		DiskID int `json:"DiskID"`
	}{}

	err = json.Unmarshal(data, &out)
	if err != nil {
		return 0, 0, err
	}

	return out.JobID, out.DiskID, err
}

// LinodeDiskCreateFromDistribution maps to the 'linode.disk.createfromdistribution'
// call.
//
// https://www.linode.com/api/linode/linode.disk.createfromdistribution
func (c *Client) LinodeDiskCreateFromDistribution(linodeID int, distID int, label string,
	size int, rootPass string, rootSSHKey *string) (jobID int, diskID int, err error) {

	args := make(map[string]interface{})
	args["LinodeID"] = linodeID
	args["DistributionID"] = distID
	args["Label"] = label
	args["Size"] = size
	args["rootPass"] = rootPass
	args["rootSSHKey"] = rootSSHKey

	data, err := c.apiCall("linode.disk.createfromdistribution", args)
	if err != nil {
		return 0, 0, err
	}

	out := struct {
		JobID  int `json:"JobID"`
		DiskID int `json:"DiskID"`
	}{}

	err = json.Unmarshal(data, &out)
	if err != nil {
		return 0, 0, err
	}

	return out.JobID, out.DiskID, err
}

// LinodeDiskCreateFromImage maps to the 'linode.disk.createfromimage' call.
//
// https://www.linode.com/api/linode/linode.disk.createfromimage
func (c *Client) LinodeDiskCreateFromImage(imageID int, linodeID int, label string, size *int,
	rootPass *string, rootSSHKey *string) (jobID int, diskID int, err error) {

	args := make(map[string]interface{})
	args["ImageID"] = imageID
	args["LinodeID"] = linodeID
	args["Label"] = label
	args["size"] = size
	args["rootPass"] = rootPass
	args["rootSSHKey"] = rootSSHKey

	data, err := c.apiCall("linode.disk.createfromimage", args)
	if err != nil {
		return 0, 0, err
	}

	out := struct {
		JobID  int `json:"JobID"`
		DiskID int `json:"DiskID"`
	}{}

	err = json.Unmarshal(data, &out)
	if err != nil {
		return 0, 0, err
	}

	return out.JobID, out.DiskID, err
}

// LinodeDiskCreateFromStackScript maps to the 'linode.disk.createfromstackscript' call.
//
// https://www.linode.com/api/linode/linode.disk.createfromstackscript
func (c *Client) LinodeDiskCreateFromStackScript(linodeID int, ssID int, ssUDFResp string,
	distID int, label string, size int, rootPass string,
	rootSSHKey *string) (jobID int, diskID int, err error) {

	args := make(map[string]interface{})
	args["LinodeID"] = linodeID
	args["StackScriptID"] = ssID
	args["StackScriptUDFResponses"] = ssUDFResp
	args["DistributionID"] = distID
	args["Label"] = label
	args["Size"] = size
	args["rootPass"] = rootPass
	args["rootSSHKey"] = rootSSHKey

	data, err := c.apiCall("linode.disk.createfromstackscript", args)
	if err != nil {
		return 0, 0, err
	}

	out := struct {
		JobID  int `json:"JobID"`
		DiskID int `json:"DiskID"`
	}{}

	err = json.Unmarshal(data, &out)
	if err != nil {
		return 0, 0, err
	}

	return out.JobID, out.DiskID, err
}

// LinodeDiskDelete maps to the 'linode.disk.delete' call.
//
// https://www.linode.com/api/linode/linode.disk.delete
func (c *Client) LinodeDiskDelete(linodeID int, diskID int) (jobID int, err error) {
	args := make(map[string]interface{})
	args["LinodeID"] = linodeID
	args["DiskID"] = diskID

	data, err := c.apiCall("linode.disk.delete", args)
	if err != nil {
		return 0, err
	}

	err = unmarshalSingle(data, "JobID", &jobID)
	if err != nil {
		return 0, err
	}

	return jobID, nil
}

// LinodeDiskDuplicate maps to the 'linode.disk.duplicate' call.
//
// https://www.linode.com/api/linode/linode.disk.duplicate
func (c *Client) LinodeDiskDuplicate(linodeID int, diskID int) (jobID int, nDiskID int,
	err error) {

	args := make(map[string]interface{})
	args["LinodeID"] = linodeID
	args["DiskID"] = diskID

	data, err := c.apiCall("linode.disk.duplicate", args)
	if err != nil {
		return 0, 0, err
	}

	out := struct {
		JobID  int `json:"JobID"`
		DiskID int `json:"DiskID"`
	}{}

	err = json.Unmarshal(data, &out)
	if err != nil {
		return 0, 0, err
	}

	return out.JobID, out.DiskID, nil
}

// LinodeDiskImagize maps to the 'linode.disk.imagize' call.
//
// https://www.linode.com/api/linode/linode.disk.imagize
func (c *Client) LinodeDiskImagize(linodeID int, diskID int, description *string,
	label *string) (jobID int, imageID int, err error) {

	args := make(map[string]interface{})
	args["LinodeID"] = linodeID
	args["DiskID"] = diskID
	args["Description"] = description
	args["Label"] = label

	data, err := c.apiCall("linode.disk.imagize", args)
	if err != nil {
		return 0, 0, err
	}

	out := struct {
		JobID   int `json:"JobID"`
		ImageID int `json:"ImageID"`
	}{}

	err = json.Unmarshal(data, &out)
	if err != nil {
		return 0, 0, err
	}

	return out.JobID, out.ImageID, nil
}

// LinodeDisk is the API response to the 'linode.disk.list' call.
type LinodeDisk struct {
	UpdateDT   string `mapstructure:"UPDATE_DT"`
	ID         int    `mapstructure:"DISKID"`
	Label      string `mapstructure:"LABEL"`
	Type       string `mapstructure:"TYPE"`
	LinodeID   int    `mapstructure:"LINODEID"`
	IsReadOnly bool   `mapstructure:"ISREADONLY"`
	Status     int    `mapstructure:"STATUS"`
	CreateDT   string `mapstructure:"CREATE_DT"`
	Size       int    `mapstructure:"SIZE"`
}

// LinodeDiskList maps to the 'linode.disk.list' call.
//
// https://www.linode.com/api/linode/linode.disk.list
func (c *Client) LinodeDiskList(linodeID int, diskID *int) ([]LinodeDisk, error) {
	args := make(map[string]interface{})
	args["LinodeID"] = linodeID
	args["DiskID"] = diskID

	data, err := c.apiCall("linode.disk.list", args)
	if err != nil {
		return nil, err
	}

	var ret []LinodeDisk
	err = unmarshalMultiMap(data, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// LinodeDiskResize maps to the 'linode.disk.resize' call.
//
// https://www.linode.com/api/linode/linode.disk.resize
func (c *Client) LinodeDiskResize(linodeID int, diskID int, size int) (jobID int, err error) {
	args := make(map[string]interface{})
	args["LinodeID"] = linodeID
	args["DiskID"] = diskID
	args["size"] = size

	data, err := c.apiCall("linode.disk.resize", args)
	if err != nil {
		return 0, err
	}

	err = unmarshalSingle(data, "JobID", &jobID)
	if err != nil {
		return 0, err
	}

	return jobID, nil
}

// LinodeDiskUpdate maps to the 'linode.disk.update' call.
//
// https://www.linode.com/api/linode/linode.disk.update
func (c *Client) LinodeDiskUpdate(linodeID int, diskID int, label *string, readOnly *bool) error {
	args := make(map[string]interface{})
	args["LinodeID"] = linodeID
	args["DiskID"] = diskID
	args["Label"] = label
	args["isReadOnly"] = readOnly

	_, err := c.apiCall("linode.disk.update", args)
	if err != nil {
		return err
	}
	return nil
}

// LinodeIPAddPrivate maps to the 'linode.ip.addprivate' call.
//
// https://www.linode.com/api/linode/linode.ip.addprivate
func (c *Client) LinodeIPAddPrivate(linodeID int) (ipID int, ipAddr string, err error) {
	args := make(map[string]interface{})
	args["LinodeID"] = linodeID

	data, err := c.apiCall("linode.ip.addprivate", args)
	if err != nil {
		return 0, "", err
	}

	out := struct {
		IPAddrID int    `json:"IPAddressID"`
		IPAddr   string `json:"IPAddress"`
	}{}

	err = json.Unmarshal(data, &out)
	if err != nil {
		return 0, "", err
	}

	return out.IPAddrID, out.IPAddr, nil
}

// LinodeIP is the API response to the 'linode.ip.list' call.
type LinodeIP struct {
	LinodeID int    `mapstructure:"LINODEID"`
	IsPublic bool   `mapstructure:"ISPUBLIC"`
	Address  string `mapstructure:"IPADDRESS"`
	RDNSName string `mapstructure:"RDNS_NAME"`
	ID       int    `mapstructure:"IPADDRESSID"`
}

// LinodeIPList maps to the 'linode.ip.list' call.
//
// https://www.linode.com/api/linode/linode.ip.list
func (c *Client) LinodeIPList(linodeID *int, ipID *int) ([]LinodeIP, error) {
	args := make(map[string]interface{})
	args["LinodeID"] = linodeID
	args["IPAddressID"] = ipID

	data, err := c.apiCall("linode.ip.list", args)
	if err != nil {
		return nil, err
	}

	var out []LinodeIP
	err = unmarshalMultiMap(data, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// LinodeIPSwap maps to the 'linode.ip.swap' call.
//
// https://www.linode.com/api/linode/linode.ip.swap
func (c *Client) LinodeIPSwap(ipID int, withIPID *int, toLinodeID *int) error {
	args := make(map[string]interface{})
	args["IPAddressID"] = ipID
	args["withIPAddressID"] = withIPID
	args["toLinodeID"] = toLinodeID

	_, err := c.apiCall("linode.ip.swap", args)
	if err != nil {
		return err
	}
	return nil
}

// LinodeJob is the API response to the 'linode.job.list' call.
type LinodeJob struct {
	EnteredDT    string `mapstructure:"ENTERED_DT"`
	Action       string `mapstructure:"ACTION"`
	Label        string `mapstructure:"LABEL"`
	HostStartDT  string `mapstructure:"HOST_START_DT"`
	LinodeID     int    `mapstructure:"LINODEID"`
	HostFinishDT string `mapstructure:"HOST_FINISH_DT"`
	Duration     int    `mapstructure:"DURATION"`
	HostMessage  string `mapstructure:"HOST_MESSAGE"`
	ID           int    `mapstructure:"JOBID"`
	HostSuccess  bool   `mapstructure:"HOST_SUCCESS"`
}

// Done returns true if the job has finished.
func (j LinodeJob) Done() bool {
	if j.HostFinishDT != "" {
		return true
	}
	return false
}

// Success returns true if the job completely successfully.  If the job is not
// done, or was unsuccessful, Success returns false.
func (j LinodeJob) Success() bool {
	if j.HostSuccess {
		return true
	}
	return false
}

// WaitForJob will wait for the specified jobID to complete, checking every
// checkInterval up to the timeout.
//
// Ok indicates whether the job was successful or not.
//
// Error will only be non-nil if the timeout is reached, there is an API error,
// or the passed job doesn't exist.
func (c *Client) WaitForJob(linodeID int, jobID int, checkInterval time.Duration,
	timeout time.Duration) (ok bool, err error) {

	ticker := time.NewTicker(checkInterval)
	defer ticker.Stop()
	deadline := time.NewTimer(timeout)
	defer deadline.Stop()

	for {
		select {
		case <-ticker.C:
			jobs, err := c.LinodeJobList(linodeID, Int(jobID), nil)
			if err != nil {
				return false, err
			}
			if len(jobs) != 1 {
				return false, fmt.Errorf("job id %d not found", jobID)
			}
			j := jobs[0]
			if j.Done() {
				return j.Success(), nil
			}
		case <-deadline.C:
			return false, fmt.Errorf("timed out waiting for job ID %d", jobID)
		}
	}
}

// WaitForAllJobs waits for all jobs on the passed Linode to be completed.
//
// Error will be non-nil if there is an API error, or the timeout is reached.
func (c *Client) WaitForAllJobs(linodeID int, checkInterval time.Duration,
	timeout time.Duration) error {

	ticker := time.NewTicker(checkInterval)
	defer ticker.Stop()
	deadline := time.NewTimer(timeout)
	defer deadline.Stop()

	for {
		select {
		case <-ticker.C:
			jobs, err := c.LinodeJobList(linodeID, nil, Bool(true))
			if err != nil {
				return err
			}
			if len(jobs) == 0 {
				return nil
			}
		case <-deadline.C:
			return errors.New("timed out waiting for all jobs to complete")
		}
	}
}

// LinodeJobList maps to the 'linode.job.list' call.
//
// https://www.linode.com/api/linode/linode.job.list
func (c *Client) LinodeJobList(linodeID int, jobID *int, pendingOnly *bool) ([]LinodeJob, error) {
	args := make(map[string]interface{})
	args["LinodeID"] = linodeID
	args["JobID"] = jobID

	// special handle weird bool arg for this method
	if pendingOnly != nil {
		if *pendingOnly == true {
			args["pendingOnly"] = 1
		} else {
			args["pendingOnly"] = 0
		}
	}

	data, err := c.apiCall("linode.job.list", args)
	if err != nil {
		return nil, err
	}

	var out []LinodeJob
	err = unmarshalMultiMap(data, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
