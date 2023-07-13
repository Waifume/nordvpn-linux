package daemon

import (
	"context"
	"log"
	"sort"

	"github.com/NordSecurity/nordvpn-linux/config"
	"github.com/NordSecurity/nordvpn-linux/daemon/pb"
	"github.com/NordSecurity/nordvpn-linux/internal"
)

// Groups provides endpoint and autocompletion.
func (r *RPC) Groups(ctx context.Context, in *pb.GroupsRequest) (*pb.Payload, error) {
	var cfg config.Config
	if err := r.cm.Load(&cfg); err != nil {
		log.Println(internal.ErrorPrefix, err)
	}

	var groupNames []string
	for group := range r.dm.GetAppData().GroupNames[in.GetObfuscate()][cfg.AutoConnectData.Protocol].Iter() {
		groupNames = append(groupNames, group.(string))
	}

	sort.Strings(groupNames)
	return &pb.Payload{
		Type: internal.CodeSuccess,
		Data: groupNames,
	}, nil
}
