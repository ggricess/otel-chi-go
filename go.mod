module github.com/ggricess/otel-chi-go

go 1.21

require (
	github.com/ggricess/otel-chi-go/internal v0.0.0-00010101000000-000000000000
	github.com/go-chi/chi/v5 v5.0.10
	github.com/signalfx/splunk-otel-go/instrumentation/internal v1.10.0
	go.opentelemetry.io/otel v1.20.0
	go.opentelemetry.io/otel/trace v1.20.0
)

require (
	github.com/go-logr/logr v1.3.0 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/otel/metric v1.20.0 // indirect
)

replace github.com/ggricess/otel-chi-go/internal => ./internal
