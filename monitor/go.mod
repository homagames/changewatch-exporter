module github.com/homagames/changewatch-exporter/monitor

go 1.21.5

require (
	github.com/google/go-github/v58 v58.0.0
	github.com/homagames/changewatch-exporter/logger v0.0.0-00010101000000-000000000000
	github.com/homagames/changewatch-exporter/utils v0.0.0-00010101000000-000000000000
	github.com/prometheus/client_golang v1.18.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/matttproud/golang_protobuf_extensions/v2 v2.0.0 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.45.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)

replace github.com/homagames/changewatch-exporter/utils => ../utils

replace github.com/homagames/changewatch-exporter/logger => ../logger
