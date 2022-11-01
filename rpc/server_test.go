package rpc

import (
	"context"
	"net"
	"testing"
)

func TestServer_HttpURL(t *testing.T) {
	type fields struct {
		ctx        context.Context
		listener   net.Listener
		httpServer *http.Server
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				ctx:        tt.fields.ctx,
				listener:   tt.fields.listener,
				httpServer: tt.fields.httpServer,
			}
			if got := s.HttpURL(); got != tt.want {
				t.Errorf("HttpURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
