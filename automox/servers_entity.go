package automox

import (
	"strconv"
	"strings"
	"time"
)

const automoxTimeFormat = "2006-01-02T15:04:05-0700"

type AutomoxUptime int64

func (a *AutomoxUptime) UnmarshalJSON(data []byte) error {
	// Remove JSON quotes from the string
	s := strings.Replace(string(data), "\"", "", -1)

	// Convert the string to an int64
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	*a = AutomoxUptime(i)
	return nil
}

type AutomoxTime time.Time

func (at *AutomoxTime) UnmarshalJSON(data []byte) error {
	// Remove JSON quotes from the string
	s := strings.Replace(string(data), "\"", "", -1)

	// Some values are null or empty, so we can't parse them
	if s == "null" || s == "" {
		*at = AutomoxTime(time.Time{})
		return nil
	}

	// Its always UTC
	t, err := time.ParseInLocation(automoxTimeFormat, s, time.UTC)
	if err != nil {
		*at = AutomoxTime(time.Time{})
		return err
	}
	*at = AutomoxTime(t)
	return nil
}

type Servers []ServerDetails

// ServerDetails are the details related to a specific server in Automox
type ServerDetails struct {
	ID                            int                 `json:"id"`
	AgentVersion                  string              `json:"agent_version"`
	Commands                      []interface{}       `json:"commands"`
	CompatibilityChecks           CompatibilityChecks `json:"compatibility_checks"`
	Compliant                     bool                `json:"compliant"`
	Connected                     bool                `json:"connected"`
	CreateTime                    AutomoxTime         `json:"create_time"`
	CustomName                    string              `json:"custom_name"`
	Deleted                       bool                `json:"deleted"`
	Detail                        Detail              `json:"detail"`
	DisplayName                   string              `json:"display_name"`
	Exception                     bool                `json:"exception"`
	InstanceID                    string              `json:"instance_id"`
	IPAddrs                       []string            `json:"ip_addrs"`
	IPAddrsPrivate                []string            `json:"ip_addrs_private"`
	IsCompatible                  bool                `json:"is_compatible"`
	IsDelayedByNotification       bool                `json:"is_delayed_by_notification"`
	IsDelayedByUser               bool                `json:"is_delayed_by_user"`
	LastDisconnectTime            AutomoxTime         `json:"last_disconnect_time"`
	LastLoggedInUser              string              `json:"last_logged_in_user"`
	LastProcessTime               AutomoxTime         `json:"last_process_time"`
	LastRefreshTime               AutomoxTime         `json:"last_refresh_time"`
	LastScanFailed                bool                `json:"last_scan_failed"`
	LastUpdateTime                AutomoxTime         `json:"last_update_time"`
	Name                          string              `json:"name"`
	NeedsAttention                bool                `json:"needs_attention"`
	NeedsReboot                   bool                `json:"needs_reboot"`
	NextPatchTime                 AutomoxTime         `json:"next_patch_time"`
	NotificationCount             int                 `json:"notification_count"`
	OrganizationID                int                 `json:"organization_id"`
	OrganizationalUnit            string              `json:"organizational_unit"`
	OsFamily                      string              `json:"os_family"`
	OsName                        string              `json:"os_name"`
	OsVersion                     string              `json:"os_version"`
	OsVersionID                   int                 `json:"os_version_id"`
	PatchDeferralCount            int                 `json:"patch_deferral_count"`
	Patches                       int                 `json:"patches"`
	Pending                       bool                `json:"pending"`
	PendingPatches                int                 `json:"pending_patches"`
	PolicyStatus                  []PolicyStatus      `json:"policy_status"`
	RebootDeferralCount           int                 `json:"reboot_deferral_count"`
	RebootIsDelayedByNotification bool                `json:"reboot_is_delayed_by_notification"`
	RebootIsDelayedByUser         bool                `json:"reboot_is_delayed_by_user"`
	RebootNotificationCount       int                 `json:"reboot_notification_count"`
	RefreshInterval               int                 `json:"refresh_interval"`
	SerialNumber                  string              `json:"serial_number"`
	ServerGroupID                 int                 `json:"server_group_id"`
	ServerPolicies                []ServerPolicies    `json:"server_policies"`
	Status                        Status              `json:"status"`
	Tags                          []interface{}       `json:"tags"`
	Timezone                      string              `json:"timezone"`
	TotalCount                    int                 `json:"total_count"`
	Uptime                        AutomoxUptime       `json:"uptime"`
	UUID                          string              `json:"uuid"`
}

type CompatibilityChecks struct {
	AppStoreDisconnected bool `json:"app_store_disconnected"`
	MissingSecureToken   bool `json:"missing_secure_token"`
	LowDiskspace         bool `json:"low_diskspace"`
}

type UpdateSourceCheck struct {
	Connected string `json:"CONNECTED"`
	Error     string `json:"ERROR"`
}

type LastUserLogon struct {
	User string `json:"USER"`
	Time string `json:"TIME"`
	Src  string `json:"SRC"`
}

type AutoUpdateOptions struct {
	Options string `json:"OPTIONS"`
	Enabled string `json:"ENABLED"`
}

type Nics struct {
	Ips       []string `json:"IPS"`
	Connected bool     `json:"CONNECTED"`
	Vendor    string   `json:"VENDOR"`
	Device    string   `json:"DEVICE"`
	Type      string   `json:"TYPE"`
	Mac       string   `json:"MAC"`
}

type Volume struct {
	Label        string `json:"LABEL"`
	Avail        string `json:"AVAIL"`
	Free         string `json:"FREE"`
	IsSystemDisk string `json:"IS_SYSTEM_DISK"`
	Volume       string `json:"VOLUME"`
	Fstype       string `json:"FSTYPE"`
}

type Detail struct {
	Ips                []string          `json:"IPS"`
	Model              string            `json:"MODEL"`
	Serial             string            `json:"SERIAL"`
	Servicetag         interface{}       `json:"SERVICETAG"`
	Fqdns              []string          `json:"FQDNS"`
	RAM                string            `json:"RAM"`
	UpdateSourceCheck  UpdateSourceCheck `json:"UPDATE_SOURCE_CHECK"`
	WmiIntegrityCheck  interface{}       `json:"WMI_INTEGRITY_CHECK"`
	DistinguishedName  interface{}       `json:"DISTINGUISHED_NAME"`
	Disks              []interface{}     `json:"DISKS"`
	CPU                string            `json:"CPU"`
	SecureTokenAccount string            `json:"SECURE_TOKEN_ACCOUNT"`
	WsusConfig         interface{}       `json:"WSUS_CONFIG"`
	Version            string            `json:"VERSION"`
	LastUserLogon      LastUserLogon     `json:"LAST_USER_LOGON"`
	AutoUpdateOptions  AutoUpdateOptions `json:"AUTO_UPDATE_OPTIONS"`
	Vendor             string            `json:"VENDOR"`
	PsVersion          interface{}       `json:"PS_VERSION"`
	Nics               []Nics            `json:"NICS"`
	Volume             []Volume          `json:"VOLUME"`
}

type PolicyStatus struct {
	ID              int         `json:"id"`
	OrganizationID  int         `json:"organization_id"`
	PolicyID        int         `json:"policy_id"`
	ServerID        int         `json:"server_id"`
	PolicyName      string      `json:"policy_name"`
	PolicyTypeName  string      `json:"policy_type_name"`
	Status          int         `json:"status"`
	Result          string      `json:"result"`
	CreateTime      AutomoxTime `json:"create_time"`
	WillReboot      bool        `json:"will_reboot"`
	PendingCount    int         `json:"pending_count"`
	NextRemediation AutomoxTime `json:"next_remediation"`
}

type ServerPoliciesConfiguration struct {
	AutoPatch                                       bool          `json:"auto_patch"`
	PatchRule                                       string        `json:"patch_rule"`
	AutoReboot                                      bool          `json:"auto_reboot"`
	FilterType                                      string        `json:"filter_type"`
	NotifyUser                                      bool          `json:"notify_user"`
	DeviceFilters                                   []interface{} `json:"device_filters"`
	IncludeOptional                                 bool          `json:"include_optional"`
	NotifyRebootUser                                bool          `json:"notify_reboot_user"`
	MissedPatchWindow                               bool          `json:"missed_patch_window"`
	UseScheduledTimezone                            bool          `json:"use_scheduled_timezone"`
	InstallDeferralEnabled                          bool          `json:"install_deferral_enabled"`
	NotifyDeferredRebootUser                        bool          `json:"notify_deferred_reboot_user"`
	NotifyUserMessageTimeout                        int           `json:"notify_user_message_timeout"`
	CustomNotificationMaxDelays                     int           `json:"custom_notification_max_delays"`
	PendingRebootDeferralEnabled                    bool          `json:"pending_reboot_deferral_enabled"`
	CustomNotificationPatchMessage                  string        `json:"custom_notification_patch_message"`
	NotifyUserAutoDeferralEnabled                   bool          `json:"notify_user_auto_deferral_enabled"`
	CustomNotificationRebootMessage                 string        `json:"custom_notification_reboot_message"`
	CustomNotificationDefermentPeriods              []int         `json:"custom_notification_deferment_periods"`
	CustomNotificationPatchMessageMac               string        `json:"custom_notification_patch_message_mac"`
	CustomNotificationRebootMessageMac              string        `json:"custom_notification_reboot_message_mac"`
	CustomPendingRebootNotificationMessage          string        `json:"custom_pending_reboot_notification_message"`
	NotifyDeferredRebootUserMessageTimeout          int           `json:"notify_deferred_reboot_user_message_timeout"`
	CustomPendingRebootNotificationMaxDelays        int           `json:"custom_pending_reboot_notification_max_delays"`
	CustomPendingRebootNotificationMessageMac       string        `json:"custom_pending_reboot_notification_message_mac"`
	NotifyDeferredRebootUserAutoDeferralEnabled     bool          `json:"notify_deferred_reboot_user_auto_deferral_enabled"`
	CustomPendingRebootNotificationDefermentPeriods []int         `json:"custom_pending_reboot_notification_deferment_periods"`
}

type Configuration struct {
	Filters                                         []interface{} `json:"filters"`
	AutoPatch                                       bool          `json:"auto_patch"`
	PatchRule                                       string        `json:"patch_rule"`
	AutoReboot                                      bool          `json:"auto_reboot"`
	FilterType                                      string        `json:"filter_type"`
	NotifyUser                                      bool          `json:"notify_user"`
	DeviceFilters                                   []interface{} `json:"device_filters"`
	AdvancedFilter                                  []interface{} `json:"advanced_filter"`
	SeverityFilter                                  []string      `json:"severity_filter"`
	IncludeOptional                                 bool          `json:"include_optional"`
	NotifyRebootUser                                bool          `json:"notify_reboot_user"`
	MissedPatchWindow                               bool          `json:"missed_patch_window"`
	UseScheduledTimezone                            bool          `json:"use_scheduled_timezone"`
	InstallDeferralEnabled                          bool          `json:"install_deferral_enabled"`
	NotifyDeferredRebootUser                        bool          `json:"notify_deferred_reboot_user"`
	NotifyUserMessageTimeout                        int           `json:"notify_user_message_timeout"`
	CustomNotificationMaxDelays                     int           `json:"custom_notification_max_delays"`
	PendingRebootDeferralEnabled                    bool          `json:"pending_reboot_deferral_enabled"`
	CustomNotificationPatchMessage                  string        `json:"custom_notification_patch_message"`
	NotifyUserAutoDeferralEnabled                   bool          `json:"notify_user_auto_deferral_enabled"`
	CustomNotificationRebootMessage                 string        `json:"custom_notification_reboot_message"`
	CustomNotificationDefermentPeriods              []int         `json:"custom_notification_deferment_periods"`
	CustomNotificationPatchMessageMac               string        `json:"custom_notification_patch_message_mac"`
	CustomNotificationRebootMessageMac              string        `json:"custom_notification_reboot_message_mac"`
	CustomPendingRebootNotificationMessage          string        `json:"custom_pending_reboot_notification_message"`
	NotifyDeferredRebootUserMessageTimeout          int           `json:"notify_deferred_reboot_user_message_timeout"`
	CustomPendingRebootNotificationMaxDelays        int           `json:"custom_pending_reboot_notification_max_delays"`
	CustomPendingRebootNotificationMessageMac       string        `json:"custom_pending_reboot_notification_message_mac"`
	NotifyDeferredRebootUserAutoDeferralEnabled     bool          `json:"notify_deferred_reboot_user_auto_deferral_enabled"`
	CustomPendingRebootNotificationDefermentPeriods []int         `json:"custom_pending_reboot_notification_deferment_periods"`
}

type ServerPolicies struct {
	Configuration        ServerPoliciesConfiguration `json:"configuration,omitempty"`
	CreateTime           AutomoxTime                 `json:"create_time"`
	ID                   int                         `json:"id"`
	Name                 string                      `json:"name"`
	NextRemediation      AutomoxTime                 `json:"next_remediation"`
	Notes                string                      `json:"notes"`
	OrganizationID       int                         `json:"organization_id"`
	PolicyTypeName       string                      `json:"policy_type_name"`
	Result               string                      `json:"result"`
	ScheduleDays         int                         `json:"schedule_days"`
	ScheduleMonths       int                         `json:"schedule_months"`
	ScheduleTime         string                      `json:"schedule_time"`
	ScheduleWeeksOfMonth int                         `json:"schedule_weeks_of_month"`
	ServerCount          int                         `json:"server_count"`
	ServerGroups         []int                       `json:"server_groups"`
	Status               int                         `json:"status"`
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
	AgentSeverity                 string      `json:"agent_severity"`
	CreateTime                    AutomoxTime `json:"create_time"`
	Cves                          []string    `json:"cves"`
	CveScore                      string      `json:"cve_score"`
	DeferredUntil                 AutomoxTime `json:"deferred_until"`
	DisplayName                   string      `json:"display_name"`
	GroupDeferredUntil            AutomoxTime `json:"group_deferred_until"`
	GroupIgnored                  bool        `json:"group_ignored"`
	ID                            int64       `json:"id"`
	Ignored                       bool        `json:"ignored"`
	Impact                        int64       `json:"impact"`
	Installed                     bool        `json:"installed"`
	IsManaged                     bool        `json:"is_managed"`
	IsUninstallable               bool        `json:"is_uninstallable"`
	Name                          string      `json:"name"`
	OrganizationID                int64       `json:"organization_id"`
	OsName                        string      `json:"os_name"`
	OsVersion                     string      `json:"os_version"`
	OsVersionID                   int         `json:"os_version_id"`
	PackageID                     int64       `json:"package_id"`
	PackageVersionID              int64       `json:"package_version_id"`
	PatchClassificationCategoryID int         `json:"patch_classification_category_id"`
	PatchScope                    string      `json:"patch_scope"`
	Repo                          string      `json:"repo"`
	RequiresReboot                bool        `json:"requires_reboot"`
	SecondaryID                   string      `json:"secondary_id"`
	ServerID                      int64       `json:"server_id"`
	Severity                      string      `json:"severity"`
	SoftwareID                    int64       `json:"software_id"`
	Version                       string      `json:"version"`
}

type CommandQueue []CommandQueueItem

type CommandQueueItem struct {
	CommandTypeName string      `json:"command_type_name"`
	Args            string      `json:"args"`
	ExecTime        AutomoxTime `json:"exec_time"`
}
