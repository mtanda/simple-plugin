package main

import (
	"github.com/golang/protobuf/ptypes"
	"golang.org/x/net/context"

	"log"

	proto "github.com/grafana/grafana/pkg/tsdb/models"
	shared "github.com/grafana/grafana/pkg/tsdb/models/proxy"
	plugin "github.com/hashicorp/go-plugin"
)

type Tsdb struct {
	plugin.NetRPCUnsupportedPlugin
}

func (t *Tsdb) Query(ctx context.Context, req *proto.TsdbQuery) (*proto.Response, error) {
	log.Print("Tsdb.Get() from plugin")

	return &proto.Response{
		Message: "from plugins! meta meta",
		Results: []*proto.QueryResult{
			&proto.QueryResult{
				Series: []*proto.TimeSeries{
					&proto.TimeSeries{
						Name: "serie 1",
						Tags: map[string]string{
							"key1": "value1",
							"key2": "value2",
						},
						Points: []*proto.Point{&proto.Point{Timestamp: ptypes.TimestampNow(), Value: 234}},
					},
				},
			},
		},
	}, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "GRAFANA_BACKEND_DATASOURCE",
			MagicCookieValue: "55d2200a-6492-493a-9353-73b728d468aa",
		},
		Plugins: map[string]plugin.Plugin{
			"backend-datasource": &shared.TsdbPluginImpl{Plugin: &Tsdb{}},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
