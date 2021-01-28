package publicip

import (
	"context"
	"net"
	"net/http"
	"time"
)

type HTTPFetcher interface {
	IP(ctx context.Context) (ip net.IP, err error)
}

type httpFetcher struct {
	client       *http.Client
	ipProvider   Provider
	ipv4Provider Provider
	ipv6Provider Provider
	ipCycler     cycler
	ipv4Cycler   cycler
	ipv6Cycler   cycler
}

type HTTPOptions struct {
	Timeout      time.Duration
	IPProvider   Provider
	IPv4Provider Provider
	IPv6Provider Provider
}

func NewHTTPFetcher(options HTTPOptions) HTTPFetcher {
	return &httpFetcher{
		client: &http.Client{
			Timeout: options.Timeout,
		},
		ipProvider:   options.IPProvider,
		ipv4Provider: options.IPv4Provider,
		ipv6Provider: options.IPv6Provider,
		ipCycler:     newCycler(getProviders(ipv4orv6)),
		ipv4Cycler:   newCycler(getProviders(ipv4)),
		ipv6Cycler:   newCycler(getProviders(ipv6)),
	}
}

func (f *httpFetcher) IP(ctx context.Context) (ip net.IP, err error) {
	provider := f.ipProvider
	if provider == Cycle {
		provider = f.ipCycler.next()
	}

	providerData, err := getProviderData(provider)
	if err != nil {
		return nil, err
	}

	return network.GetPublicIP(ctx, i.client, providerData.http.ipv4, constants.IPv4OrIPv6)
}

func (f *httpFetcher) IPv4(ctx context.Context) (ip net.IP, err error) {
	provider := f.ipv4Provider
	if provider == Cycle {
		provider = f.ipv4Cycler.next()
	}

	providerData, err := getProviderData(provider)
	if err != nil {
		return nil, err
	}

	return network.GetPublicIP(ctx, i.client, providerData.http.ipv4, constants.IPv4)
}

func (f *httpFetcher) IPv6(ctx context.Context) (ip net.IP, err error) {
	provider := f.ipv6Provider
	if provider == Cycle {
		provider = f.ipv6Cycler.next()
	}

	providerData, err := getProviderData(provider)
	if err != nil {
		return nil, err
	}

	return network.GetPublicIP(ctx, i.client, providerData.http.ipv6, constants.IPv6)
}
