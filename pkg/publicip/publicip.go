package publicip

import (
	"context"
	"net"
)

type Fetcher interface {
	IP(ctx context.Context) (ip net.IP, err error)
	IPv4(ctx context.Context) (ipv4 net.IP, err error)
	IPv6(ctx context.Context) (ipv6 net.IP, err error)
}

type fetcher struct {
}

func NewFetcher() Fetcher {
	return &fetcher{}
}

func (f *fetcher) IP(ctx context.Context) (ip net.IP, err error) {
	return nil, nil
}

func (f *fetcher) IPv4(ctx context.Context) (ipv4 net.IP, err error) {
	return nil, nil
}

func (f *fetcher) IPv6(ctx context.Context) (ipv6 net.IP, err error) {
	return nil, nil
}
