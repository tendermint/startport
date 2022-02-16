package network

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/x/launch/types"
)

func TestPeerAddress(t *testing.T) {
	tests := []struct {
		name string
		peer types.Peer
		want string
		err  error
	}{
		{
			name: "simple peer connection",
			peer: types.NewPeerConn("simple-conn", "200.100.50.20"),
			want: "simple-conn@200.100.50.20",
		},
		{
			name: "http tunnel peer",
			peer: types.NewPeerTunnel("httpTunnel", "tunnel", "200.100.50.20"),
			want: "httpTunnel@200.100.50.20",
		},
		{
			name: "invalid peer",
			peer: types.Peer{Id: "invalid-peer", Connection: nil},
			err:  errors.New("invalid peer connection type: <nil>"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PeerAddress(tt.peer)
			if tt.err != nil {
				require.Error(t, err)
				require.Equal(t, tt.err.Error(), err.Error())
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
