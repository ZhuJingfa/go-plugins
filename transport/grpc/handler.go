package grpc

import (
	"runtime/debug"

	"micro/go-log"
	"micro/go-micro/transport"
	pb "micro/go-plugins/transport/grpc/proto"
)

// microTransport satisfies the pb.TransportServer inteface
type microTransport struct {
	fn func(transport.Socket)
}

func (m *microTransport) Stream(ts pb.Transport_StreamServer) error {
	sock := &grpcTransportSocket{
		stream: ts,
	}

	defer func() {
		if r := recover(); r != nil {
			log.Log(r, string(debug.Stack()))
			sock.Close()
		}
	}()

	// execute socket func
	m.fn(sock)
	return nil
}
