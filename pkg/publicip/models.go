package publicip

type Provider string

type ipVersion uint8

type providerData struct {
	http providerHTTPUrls
}

type providerHTTPUrls struct {
	ip   []string
	ipv4 []string
	ipv6 []string
}

func (p *providerData) supportsIPv4OrIpv6() bool {
	return len(p.http.ip) > 0
}

func (p *providerData) supportsIPv4() bool {
	return len(p.http.ipv4) > 0
}

func (p *providerData) supportsIPv6() bool {
	return len(p.http.ipv6) > 0
}
