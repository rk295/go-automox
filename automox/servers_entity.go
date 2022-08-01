package automox

import (
	"strconv"
	"time"
)

type AutomoxUptime int64

func (a *AutomoxUptime) UnmarshalJSON(data []byte) error {
	v := string(data)
	// fmt.Printf("attempting to parse %s\n", v)
	i, err := strconv.ParseInt(v, 10, 64)
	if err == nil {
		return err
	}
	// fmt.Printf("parsed %d\n", i)
	*a = AutomoxUptime(i)
	return nil
}

type Servers []ServerDetails

// Server holds the details of a specific Automox server
// type Server struct {
// 	Details ServerDetails `json:"details"`
// }

// ServerDetails are the details related to a specific server in Automox
type ServerDetails struct {
	ID              int64         `json:"id"`
	OsVersionID     int64         `json:"os_version_id"`
	ServerGroupID   int64         `json:"server_group_id"`
	OrganizationID  int64         `json:"organization_id"`
	UUID            string        `json:"uuid"`
	Name            string        `json:"name"`
	InstanceID      string        `json:"instance_id"`
	RefreshInterval int64         `json:"refresh_interval"`
	LastUpdateTime  string        `json:"last_update_time"`
	LastRefreshTime string        `json:"last_refresh_time"`
	Uptime          AutomoxUptime `json:"uptime"`
	NeedsReboot     bool          `json:"needs_reboot"`
	Timezone        string        `json:"timezone"`
	Tags            []string      `json:"tags"`
	Deleted         bool          `json:"deleted"`
	// TODO: Fails unmarshalling:
	// 2022/07/31 07:10:12 parsing time "\"2022-07-21T10:10:06+0000\"" as "\"2006-01-02T15:04:05Z07:00\"": cannot parse "+0000\"" as "Z07:00"
	// CreateTime                    time.Time           `json:"create_time"`
	CreateTime                    string              `json:"create_time"`
	OsVersion                     string              `json:"os_version"`
	OsName                        string              `json:"os_name"`
	OsFamily                      string              `json:"os_family"`
	IPAddrs                       []string            `json:"ip_addrs"`
	IPAddrsPrivate                []string            `json:"ip_addrs_private"`
	Patches                       int64               `json:"patches"`
	Details                       Details             `json:"details"`
	AgentVersion                  string              `json:"agent_version"`
	CustomName                    string              `json:"custom_name"`
	Exception                     bool                `json:"exception"`
	TotalCount                    int64               `json:"total_count"`
	IsCompatible                  bool                `json:"is_compatible"`
	CompatibilityChecks           CompatibilityChecks `json:"compatibility_checks"`
	PolicyStatus                  []PolicyStatus      `json:"policy_status"`
	LastScanFailed                bool                `json:"last_scan_failed"`
	Pending                       bool                `json:"pending"`
	Compliant                     bool                `json:"compliant"`
	DisplayName                   string              `json:"display_name"`
	Commands                      []Commands          `json:"commands"`
	PendingPatches                int64               `json:"pending_patches"`
	Connected                     bool                `json:"connected"`
	LastProcessTime               string              `json:"last_process_time"`
	NextPatchTime                 string              `json:"next_patch_time"`
	NotificationCount             int64               `json:"notification_count"`
	RebootNotificationCount       int64               `json:"reboot_notification_count"`
	PatchDeferralCount            int64               `json:"patch_deferral_count"`
	IsDelayedByNotification       bool                `json:"is_delayed_by_notification"`
	RebootIsDelayedByNotification bool                `json:"reboot_is_delayed_by_notification"`
	IsDelayedByUser               bool                `json:"is_delayed_by_user"`
	RebootIsDelayedByUser         bool                `json:"reboot_is_delayed_by_user"`
	// TODO: Fails unmarshalling, see above
	// LastDisconnectTime            time.Time           `json:"last_disconnect_time"`
	LastDisconnectTime string `json:"last_disconnect_time"`
	NeedsAttention     bool   `json:"needs_attention"`
	SerialNumber       string `json:"serial_number"`
	Status             Status `json:"status"`
	LastLoggedInUser   string `json:"last_logged_in_user"`
}
type Disks struct {
	Size string `json:"SIZE"`
	Type string `json:"TYPE"`
}
type Nics struct {
	Connected bool     `json:"CONNECTED"`
	Device    string   `json:"DEVICE"`
	Ips       []string `json:"IPS"`
	Mac       string   `json:"MAC"`
	Type      string   `json:"TYPE"`
	Vendor    string   `json:"VENDOR"`
}
type Details struct {
	CPU        string  `json:"CPU"`
	Disks      []Disks `json:"DISKS"`
	Model      string  `json:"MODEL"`
	Nics       []Nics  `json:"NICS"`
	RAM        string  `json:"RAM"`
	Serial     string  `json:"SERIAL"`
	Servicetag string  `json:"SERVICETAG"`
	Vendor     string  `json:"VENDOR"`
	Version    string  `json:"VERSION"`
}
type CompatibilityChecks struct {
	LowDiskspace                    bool `json:"low_diskspace"`
	MissingSecureToken              bool `json:"missing_secure_token"`
	AppStoreDisconnected            bool `json:"app_store_disconnected"`
	MissingPowershell               bool `json:"missing_powershell"`
	MissingWmiIntegrityCheck        bool `json:"missing_wmi_integrity_check"`
	WsusDisconnected                bool `json:"wsus_disconnected"`
	WindowsUpdateServerDisconnected bool `json:"windows_update_server_disconnected"`
}
type PolicyStatus struct {
	ID             int64  `json:"id"`
	OrganizationID int64  `json:"organization_id"`
	PolicyID       int64  `json:"policy_id"`
	ServerID       int64  `json:"server_id"`
	PolicyName     string `json:"policy_name"`
	PolicyTypeName string `json:"policy_type_name"`
	Status         int64  `json:"status"`
	Result         string `json:"result"`
	CreateTime     string `json:"create_time"`
}
type Commands struct {
	CommandTypeName string    `json:"command_type_name"`
	Args            string    `json:"args"`
	ExecTime        time.Time `json:"exec_time"`
}
type PolicyStatuses struct {
	ID        int  `json:"id"`
	Compliant bool `json:"compliant"`
}
type Status struct {
	DeviceStatus   string           `json:"device_status"`
	AgentStatus    string           `json:"agent_status"`
	PolicyStatus   string           `json:"policy_status"`
	PolicyStatuses []PolicyStatuses `json:"policy_statuses"`
}

type Packages []PackageDetails

type PackageDetails struct {
	AgentSeverity string `json:"agent_severity"`
	CreateTime    string `json:"create_time"`
	// CreateTime                    time.Time `json:"create_time"`
	Cves          []string `json:"cves"`
	CveScore      string   `json:"cve_score"`
	DeferredUntil string   `json:"deferred_until"`
	// DeferredUntil                 time.Time `json:"deferred_until"`
	DisplayName        string `json:"display_name"`
	GroupDeferredUntil string `json:"group_deferred_until"`
	// GroupDeferredUntil            time.Time `json:"group_deferred_until"`
	GroupIgnored                  bool   `json:"group_ignored"`
	ID                            int64  `json:"id"`
	Ignored                       bool   `json:"ignored"`
	Impact                        int64  `json:"impact"`
	Installed                     bool   `json:"installed"`
	IsManaged                     bool   `json:"is_managed"`
	IsUninstallable               bool   `json:"is_uninstallable"`
	Name                          string `json:"name"`
	OrganizationID                int64  `json:"organization_id"`
	OsName                        string `json:"os_name"`
	OsVersion                     string `json:"os_version"`
	OsVersionID                   int    `json:"os_version_id"`
	PackageID                     int64  `json:"package_id"`
	PackageVersionID              int64  `json:"package_version_id"`
	PatchClassificationCategoryID int    `json:"patch_classification_category_id"`
	PatchScope                    string `json:"patch_scope"`
	Repo                          string `json:"repo"`
	RequiresReboot                bool   `json:"requires_reboot"`
	SecondaryID                   string `json:"secondary_id"`
	ServerID                      int64  `json:"server_id"`
	Severity                      string `json:"severity"`
	SoftwareID                    int64  `json:"software_id"`
	Version                       string `json:"version"`
}

type CommandQueue []CommandQueueItem

type CommandQueueItem struct {
	CommandTypeName string `json:"command_type_name"`
	Args            string `json:"args"`
	ExecTime        string `json:"exec_time"`
	// ExecTime        time.Time `json:"exec_time"`
}
