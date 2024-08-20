package server

import (
	hclog "github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

type rpcServer struct {
	log hclog.Logger
	gs  *grpc.Server
}
