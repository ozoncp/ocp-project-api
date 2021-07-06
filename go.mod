module github.com/ozoncp/ocp-project-api

go 1.16

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/HdrHistogram/hdrhistogram-go v1.1.0 // indirect
	github.com/Masterminds/squirrel v1.5.0
	github.com/Shopify/sarama v1.29.0
	github.com/gofrs/uuid v4.0.0+incompatible
	github.com/golang/mock v1.5.0
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/jmoiron/sqlx v1.3.4
	github.com/lib/pq v1.2.0
	github.com/onsi/ginkgo v1.16.2
	github.com/onsi/gomega v1.13.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/ozoncp/ocp-project-api/pkg/ocp-project-api v0.0.0-00010101000000-000000000000
	github.com/ozoncp/ocp-project-api/pkg/ocp-repo-api v0.0.0-00010101000000-000000000000
	github.com/prometheus/client_golang v1.11.0
	github.com/rs/zerolog v1.22.0
	github.com/spf13/afero v1.3.4
	github.com/stretchr/testify v1.7.0
	github.com/uber/jaeger-client-go v2.29.1+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible
	github.com/yandex/pandora v0.3.2
	go.uber.org/atomic v1.4.0 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/grpc v1.38.0
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/ozoncp/ocp-project-api/pkg/ocp-project-api => ./pkg/ocp-project-api

replace github.com/ozoncp/ocp-project-api/pkg/ocp-repo-api => ./pkg/ocp-repo-api
