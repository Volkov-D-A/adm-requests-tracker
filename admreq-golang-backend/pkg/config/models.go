package config

type config struct {
	Env            string `yaml:"env" env-required:"true"`
	Key            string `yaml:"key" env-required:"true"`
	GrpcUserServer `yaml:"grpc_user_server"`
	GrpcTsrServer  `yaml:"grpc_tsr_server"`
	GrpcGw         `yaml:"grpc_gw"`
	PG             `yaml:"pg" env-required:"true"`
}

type GrpcUserServer struct {
	Address string `yaml:"address" env-default:"localhost:6001"`
}

type GrpcTsrServer struct {
	Address string `yaml:"address" env-default:"localhost:6002"`
}

type GrpcGw struct {
	Address string `yaml:"address" env-default:"localhost:8080"`
}

type PG struct {
	Host     string `yaml:"host" env-required:"true"`
	Port     string `yaml:"port" env-required:"true"`
	User     string `yaml:"user" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
	Database string `yaml:"database" env-required:"true"`
	MP       string `yaml:"mp" env-required:"true"`
}
