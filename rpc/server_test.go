package rpc

import (
	"context"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

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
	require.NoError(t, err)
	require.GreaterOrEqual(t, port, 1024)
	require.Equal(t, s.listener.Addr().(*net.TCPAddr).Port, port)
	require.NoError(t, s.listener.Close())
}

func TestServer_StartThenContextCanceled(t *testing.T) {
	// Configure a timeout to cancel the context after the server is fully started
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	s, err := NewServer(&Config{
		Address: "127.0.0.1:0",
		Relayer: &relayer.Relayer{},
		Ctx:     ctx,
	})
	require.NoError(t, err)
	err = s.Start()
	require.ErrorIs(t, err, context.DeadlineExceeded)
}

func TestServer_Stop(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	s, err := NewServer(&Config{
		Address: "127.0.0.1:0",
		Relayer: &relayer.Relayer{},
		Ctx:     ctx,
	})
	require.NoError(t, err)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(500 * time.Millisecond) // enough time for the server to fully start
		err := s.Stop()
		require.NoError(t, err)
	}()
	err = s.Start()
	require.ErrorIs(t, err, http.ErrServerClosed)
	wg.Wait()
}
