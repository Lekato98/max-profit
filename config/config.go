package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	rpcPortField   = "RPC_PORT"
	httpPortField  = "PORT"
	secretKeyField = "SECRET_KEY"
)

type Configs struct {
	RpcPort   int64
	HttpPort  int64
	SecretKey string
}

func LoadConfigs() (*Configs, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	rpcPort, err := strconv.ParseInt(os.Getenv(rpcPortField), 10, 64)
	if err != nil {
		return nil, err
	}

	httpPort, err := strconv.ParseInt(os.Getenv(httpPortField), 10, 64)
	if err != nil {
		return nil, err
	}

	return &Configs{
		RpcPort:   rpcPort,
		HttpPort:  httpPort,
		SecretKey: os.Getenv(secretKeyField),
	}, nil
}
