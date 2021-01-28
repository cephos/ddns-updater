package publicip

import (
	"context"
	"errors"
	"fmt"
	"net"
)

type DNSFetcher interface {
	Fetch(ctx context.Context) (publicIP net.IP, err error)
}

type dnsFetcher struct {
	resolver *net.Resolver
}

func NewDNSFetcher() DNSFetcher {
	dialer := net.Dialer{}
	return &dnsFetcher{
		resolver: &net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network string, address string) (net.Conn, error) {
				return dialer.DialContext(ctx, "udp", net.JoinHostPort("ns1.google.com", "53"))
			},
		},
	}
}

var (
	ErrNoTXTRecordFound  = errors.New("no TXT record found")
	ErrTooManyTXTRecords = errors.New("too many TXT records")
)

func (f *dnsFetcher) Fetch(ctx context.Context) (publicIP net.IP, err error) {
	records, err := f.resolver.LookupTXT(ctx, "o-o.myaddr.l.google.com")
	if err != nil {
		return nil, err
	}

	L := len(records)
	if L == 0 {
		return nil, ErrNoTXTRecordFound
	} else if L > 1 {
		return nil, fmt.Errorf("%w: %d instead of 1", ErrTooManyTXTRecords, L)
	}

	publicIP = net.ParseIP(records[0])
	if publicIP == nil {
		return nil, fmt.Errorf("%w: %s", ErrIPMalformed, records[0])
	}

	return publicIP, nil
}
