module github.com/volkov-d-a/adm-requests-tracker

go 1.23.0

toolchain go1.23.5

require (
	github.com/golang-migrate/migrate/v4 v4.18.3
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.27.1
	github.com/ilyakaznacheev/cleanenv v1.5.0
	github.com/jackc/pgx/v5 v5.7.5
	golang.org/x/net v0.43.0
	google.golang.org/genproto/googleapis/api v0.0.0-20250804133106-a7a43d27e69b
	google.golang.org/grpc v1.74.2
	google.golang.org/protobuf v1.36.7
)

require (
	github.com/google/go-github/v39 v39.2.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/lib/pq v1.10.9 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	golang.org/x/crypto v0.41.0 // indirect
	golang.org/x/oauth2 v0.30.0 // indirect
	golang.org/x/sync v0.16.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250804133106-a7a43d27e69b // indirect
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.5.1 // indirect
)

require (
	github.com/BurntSushi/toml v1.5.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	golang.org/x/exp v0.0.0-20250718183923-645b1fa84792
	golang.org/x/sys v0.35.0 // indirect
	golang.org/x/text v0.28.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)

tool (
	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	google.golang.org/grpc/cmd/protoc-gen-go-grpc
	google.golang.org/protobuf/cmd/protoc-gen-go
)
