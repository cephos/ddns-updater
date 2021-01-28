package publicip

import (
	"errors"
	"fmt"
)

const (
	Cycle    Provider = "cycle"
	Opendns  Provider = "opendns"
	Ifconfig Provider = "ifconfig"
	Ipinfo   Provider = "ipinfo"
	Ipify    Provider = "ipify"
	Google   Provider = "google"
	Noip     Provider = "noip"
)

var (
	ErrUnknownProvider = errors.New("unknown provider string")
)

func ParseProvider(s string) (provider Provider, err error) {
	provider = Provider(s)
	if _, ok := providersData()[provider]; !ok {
		return Provider(""), fmt.Errorf("%w: %s", ErrUnknownProvider, s)
	}
	return provider, nil
}

var (
	ErrMethodNotFoundForProvider = errors.New("method not found for provider")
)

func getProviderData(provider Provider) (data providerData, err error) {
	data, ok := providersData()[provider]
	if !ok {
		return data, fmt.Errorf("%w: %s", ErrMethodNotFoundForProvider, provider)
	}
	return data, nil
}

const (
	ipv4orv6 ipVersion = iota
	ipv4
	ipv6
)

func getProviders(ipVersion ipVersion) (providers []Provider) {
	for provider, data := range providersData() {
		switch ipVersion {
		case ipv4orv6:
			if data.supportsIPv4OrIpv6() {
				providers = append(providers, provider)
			}
		case ipv4:
			if data.supportsIPv4() {
				providers = append(providers, provider)
			}
		case ipv6:
			if data.supportsIPv6() {
				providers = append(providers, provider)
			}
		}
	}
	return providers
}

func providersData() map[Provider]providerData {
	return map[Provider]providerData{
		"cycle": {},
		"opendns": {
			http: providerHTTPUrls{
				ip: []string{"https://diagnostic.opendns.com/myip"},
			},
		},
		"ifconfig": {
			http: providerHTTPUrls{
				ip: []string{"https://ifconfig.io/ip"},
			},
		},
		"ipinfo": {
			http: providerHTTPUrls{
				ip: []string{"https://ipinfo.io/ip"},
			},
		},
		"ipify": {
			http: providerHTTPUrls{
				ipv4: []string{"https://api.ipify.org"},
				ipv6: []string{"https://api6.ipify.org"},
			},
		},
		"google": {
			http: providerHTTPUrls{
				ip: []string{"https://domains.google.com/checkip"},
			},
		},
		"noip": {
			http: providerHTTPUrls{
				ipv4: []string{"http://ip1.dynupdate.no-ip.com", "http://ip1.dynupdate.no-ip.com:8245"},
				ipv6: []string{"http://ip1.dynupdate6.no-ip.com", "http://ip1.dynupdate6.no-ip.com:8245"},
			},
		},
	}
}
