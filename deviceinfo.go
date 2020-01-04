package gohome

type DeviceInfo struct {
	BSSID         string   `json:"bssid"`
	BuildVersion  string   `json:"build_version"`
	HasUpdate     bool     `json:"has_update"`
	HotspotBSSID  string   `json:"hotspot_bssid"`
	IP            string   `json:"ip_address"`
	Location      Location `json:"location"`
	MAC           string   `json:"mac_address"`
	Name          string   `json:"name"`
	NoiseLevel    float64  `json:"noise_level"`
	OptIn         OptIn    `json:"opt_in"`
	PublicKey     string   `json:"public_key"`
	SignalLevel   float64  `json:"signal_level"`
	SSID          string   `json:"ssid"`
	Uptime        float64  `json:"uptime"`
	Version       uint16   `json:"version"`
	WPAConfigured bool     `json:"wpa_configured"`
	WPAID         uint16   `json:"wpa_id"`
	WPAState      uint16   `json:"wpa_state"`
}
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
type OptIn struct {
	Crash    bool `json:"crash"`
	OpenCast bool `json:"opencast"`
	Stats    bool `json:"stats"`
}
