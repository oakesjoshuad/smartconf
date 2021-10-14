package smartconf

import "encoding/json"

type ZoneConfig struct {
	ZoneName           string            `json:"zonename,omitempty"`
	Autoboot           bool              `json:"autoboot"`
	Brand              string            `json:"brand"`
	LimitPriv          string            `json:"limit_priv"`
	V                  int               `json:"v,omitempty"`
	CreateTimestamp    string            `json:"create_timestamp,omitempty"`
	ImageUUID          string            `json:"image_uuid"`
	CPUShares          int               `json:"cpu_shares,omitempty"`
	MaxLWPS            int               `json:"max_lwps,omitempty"`
	MaxMsgIds          int               `json:"max_msg_ids,omitempty"`
	MaxSemIds          int               `json:"max_sem_ids,omitempty"`
	MaxShmIds          int               `json:"max_shm_ids,omitempty"`
	MaxShmMemory       int               `json:"max_shm_memory,omitempty"`
	ZfsIOPriority      int               `json:"zfs_io_priority,omitempty"`
	MaxPhysicalMemory  int               `json:"max_physical_memory"`
	MaxLockedMemory    int               `json:"max_locked_memory,omitempty"`
	MaxSwap            int               `json:"max_swap,omitempty"`
	BillingId          string            `json:"billing_id,omitempty"`
	OwnerUUID          string            `json:"owner_uuid,omitempty"`
	Tmpfs              int               `json:"tmpfs,omitempty"`
	Hostname           string            `json:"hostname"`
	DNSDomain          string            `json:"dns_domain"`
	Resolvers          []string          `json:"resolvers"`
	Alias              string            `json:"alias"`
	Nics               []Nic             `json:"nics"`
	UUID               string            `json:"uuid,omitempty"`
	ZoneState          string            `json:"zone_state,omitempty"`
	ZonePath           string            `json:"zonepath,omitempty"`
	HVM                bool              `json:"hvm,omitempty"`
	ZoneId             int               `json:"zoneid,omitempty"`
	ZoneDId            int               `json:"zonedid,omitempty"`
	LastModified       string            `json:"last_modified,omitempty"`
	FirewallEnabled    bool              `json:"firewall_enabled"`
	ServerUUID         string            `json:"server_uuid,omitempty"`
	PlatformBuildstamp string            `json:"platform_buildstamp,omitempty"`
	State              string            `json:"state,omitempty"`
	BootTimestamp      string            `json:"boot_timestamp,omitempty"`
	InitRestarts       int               `json:"init_restarts,omitempty"`
	PID                int               `json:"pid,omitempty,omitempty"`
	CustomerMetadata   map[string]string `json:"customer_metadata,omitempty"`
	InternalMetadata   map[string]string `json:"internal_metadata,omitempty"`
	Routes             map[string]string `json:"routes,omitempty"`
	Tags               map[string]string `json:"tags,omitempty"`
	Quota              int               `json:"quota,omitempty"`
	ZfsRootRecsize     int               `json:"zfs_root_recsize,omitempty"`
	ZfsFilesystem      string            `json:"zfs_filesystem,omitempty"`
	Zpool              string            `json:"zpool,omitempty"`
}

type Nic struct {
	InterfaceName string   `json:"interface"`
	MacAddr       string   `json:"mac,omitempty"`
	NicTag        string   `json:"nic_tag"`
	Gateways      []string `json:"gateways"`
	Netmask       string   `json:"netmask"`
	Ips           []string `json:"ips"`
	VLanId        int      `json:"vlan_id,omitempty"`
	Primary       bool     `json:"primary"`
}

func (zconf *ZoneConfig) JSON() ([]byte, error) {
	return json.MarshalIndent(zconf, "", "\t")
}

func NewZoneConfig(image, alias, hostname string, memory int) *ZoneConfig {
	return &ZoneConfig{
	}
}
