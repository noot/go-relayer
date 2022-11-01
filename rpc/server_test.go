package rpc

import (
	"context"
	"net"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/athanorlabs/go-relayer/relayer"
)

func TestServer_NewServer(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	s, err := NewServer(&Config{
		Address: "127.0.0.1:0",
		Relayer: &relayer.Relayer{},
		Ctx:     ctx,
	})
	require.NoError(t, err)
	httpURL := s.HttpURL()
	portStr := strings.Replace(httpURL, "http://127.0.0.1:", "", 1)
	port, err := strconv.Atoi(portStr)
	require.GreaterOrEqual(t, port, 1024)
	require.Equal(t, s.listener.Addr().(*net.TCPAddr).Port, port)
	require.NoError(t, s.listener.Close())
}
