package publicip

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_fetcher_IP(t *testing.T) {
	t.Parallel()
	fetcher := NewFetcher()
	ctx := context.Background()

	publicIP, err := fetcher.IP(ctx)
	require.NoError(t, err)
	assert.Equal(t, "", publicIP.String())
}
