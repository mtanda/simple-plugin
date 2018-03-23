package main

import (
	"log"
	"time"

	"golang.org/x/net/context"

	"github.com/grafana/grafana-plugin-model/go/datasource"
	plugin "github.com/hashicorp/go-plugin"
)

type Tsdb struct {
	plugin.NetRPCUnsupportedPlugin
}

func (t *Tsdb) Query(ctx context.Context, req *datasource.DatasourceRequest) (*datasource.DatasourceResponse, error) {
	log.Print("Tsdb.Get() from plugin")

	return &datasource.DatasourceResponse{
		Results: []*datasource.QueryResult{
			&datasource.QueryResult{
				RefId: "A",
				Series: []*datasource.TimeSeries{
					&datasource.TimeSeries{
						Name: "serie 1",
						Tags: map[string]string{
							"key1": "value1",
							"key2": "value2",
						},
						Points: []*datasource.Point{&datasource.Point{Timestamp: time.Now().Unix(), Value: 234}},
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
			"backend-datasource": &datasource.DatasourcePluginImpl{Plugin: &Tsdb{}},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
