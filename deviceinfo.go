package gohome

// DeviceInfo represents the Home device's info
type DeviceInfo struct {
	BSSID             string     `json:"bssid"`
	BuildVersion      string     `json:"build_version"`
	CastBuildVersion  string     `json:"cast_build_version"`
	Connected         bool       `json:"connected"`
	EthernetConnected bool       `json:"ethernet_connected"`
	HasUpdate         bool       `json:"has_update"`
	HotspotBSSID      string     `json:"hotspot_bssid"`
	IP                string     `json:"ip_address"`
	Locale            string     `json:"locale"`
	Location          Location   `json:"location"`
	MAC               string     `json:"mac_address"`
	Name              string     `json:"name"`
	NoiseLevel        float64    `json:"noise_level"`
	OptIn             OptIn      `json:"opt_in"`
	PublicKey         string     `json:"public_key"`
	ReleaseTrack      string     `json:"release_track"`
	SetupState        uint16     `json:"setup_state"`
	SetupStats        SetupStats `json:"setup_stats"`
	SignalLevel       float64    `json:"signal_level"`
	SSID              string     `json:"ssid"`
	Uptime            float64    `json:"uptime"`
	Version           uint16     `json:"version"`
	WPAConfigured     bool       `json:"wpa_configured"`
	WPAID             uint16     `json:"wpa_id"`
	WPAState          uint16     `json:"wpa_state"`
}

// Location is part of DeviceInfo and represents a location
type Location struct {
	CountryCode string  `json:"country_code"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}

// OptIn is part of DeviceInfo and represents some opt-in values
type OptIn struct {
	Crash    bool `json:"crash"`
	OpenCast bool `json:"opencast"`
	Stats    bool `json:"stats"`
}

// SetupStats is part of DeviceInfo and represents setup stats
type SetupStats struct {
	HistoricallySucceeded    bool   `json:"historically_succeeded"`
	NumCheckConnectivity     uint16 `json:"num_check_connectivity"`
	NumConnectWiFi           uint16 `json:"num_connect_wifi"`
	NumConnectedWifiNotSaved uint16 `json:"num_connected_wifi_not_saved"`
	NumInitialEurekaInfo     uint16 `json:"num_initial_eureka_info"`
	NumObtainIP              uint16 `json:"num_obtain_ip"`
}
