package main

import (
	"log"
	"os"
	"time"

	"golang.org/x/net/context"

	"github.com/grafana/grafana-plugin-model/go/datasource"
	hclog "github.com/hashicorp/go-hclog"
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
	f, err := os.OpenFile("/tmp/log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()
	//log.SetOutput(f)
	log.Print("init")
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		Output:     f,
		JSONFormat: true,
	})
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
		Logger:     logger,
	})
	log.Print("inited")
}
