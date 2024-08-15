package ipinfo

import "fmt"

type IPInfo struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
}

func (ip *IPInfo) Info() string {
	return fmt.Sprintf("IP: %s\nHostname: %s\nCity: %s\nRegion: %s\nCountry: %s\nLocation: %s\nOrganization: %s\nPostal Code: %s\nTimezone: %s",
		ip.IP, ip.Hostname, ip.City, ip.Region, ip.Country, ip.Loc, ip.Org, ip.Postal, ip.Timezone)
}
