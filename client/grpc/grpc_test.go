package grpc

import (
	"net"
	"strconv"
	"strings"
	"testing"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"micro/go-micro/client"
	"micro/go-micro/registry"
	"micro/go-micro/registry/mock"
	"micro/go-micro/selector"
)

// server is used to implement helloworld.GreeterServer.
type greeterServer struct{}

// SayHello implements helloworld.GreeterServer
func (g *greeterServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func TestGRPCClient(t *testing.T) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}
	defer l.Close()

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &greeterServer{})

	go s.Serve(l)
	defer s.Stop()

	parts := strings.Split(l.Addr().String(), ":")
	port, _ := strconv.Atoi(parts[len(parts)-1])
	addr := strings.Join(parts[:len(parts)-1], ":")

	// create mock registry
	r := mock.NewRegistry()

	// register service
	r.Register(&registry.Service{
		Name:    "test",
		Version: "test",
		Nodes: []*registry.Node{
			&registry.Node{
				Id:      "test-1",
				Address: addr,
				Port:    port,
			},
		},
	})

	// create selector
	se := selector.NewSelector(
		selector.Registry(r),
	)

	// create client
	c := NewClient(
		client.Registry(r),
		client.Selector(se),
	)

	testMethods := []string{
		"/helloworld.Greeter/SayHello",
		"Greeter.SayHello",
	}

	for _, method := range testMethods {
		req := c.NewRequest("test", method, &pb.HelloRequest{
			Name: "John",
		})

		rsp := pb.HelloReply{}

		err = c.Call(context.TODO(), req, &rsp)
		if err != nil {
			t.Fatal(err)
		}

		if rsp.Message != "Hello John" {
			t.Fatalf("Got unexpected response %v", rsp.Message)
		}
	}
}
