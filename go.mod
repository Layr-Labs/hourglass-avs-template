module github.com/Layr-Labs/hourglass-avs-template

go 1.23.6

toolchain go1.24.2

require (
	github.com/Layr-Labs/hourglass-monorepo/ponos v0.0.0-20250505155921-6f6074f3e11d
	go.uber.org/zap v1.27.0
)

require (
	github.com/BurntSushi/toml v1.5.0 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.36.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240903143218-8af14fe29dc1 // indirect
	google.golang.org/grpc v1.69.0-dev // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)

replace github.com/Layr-Labs/hourglass-monorepo => ./temp_external/hourglass-monorepo

replace github.com/Layr-Labs/hourglass-monorepo/contracts => ./temp_external/hourglass-monorepo/contracts
