package main

type DeviceInfo struct {
	BSSID        string   `json:"bssid"`
	BuildVersion string   `json:"build_version"`
	HasUpdate    bool     `json:"has_update"`
	HotspotBSSID string   `json:"hotspot_bssid"`
	IP           string   `json:"ip_address"`
	Location     Location `json:"location"`
	MAC          string   `json:"mac_address"`
	Name         string   `json:"name"`
	NoiseLevel   int16    `json:"noise_level"`
	OptIn        OptIn    `json:"opt_in"`
	PublicKey    string   `json:"public_key"`
	Signallevel  int16    `json:"signal_level"`
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
