package models

type Config struct {
	Env        string `yaml:"env" env-required:"true"`
	Key        string `yaml:"key" env-required:"true"`
	GrpcServer `yaml:"grpc_server"`
	GrpcGw     `yaml:"grpc_gw"`
}

type GrpcServer struct {
	Address string `yaml:"address" env-default:"localhost:5000"`
}

type GrpcGw struct {
	Address string `yaml:"address" env-default:"localhost:8080"`
}
