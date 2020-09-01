/*
 *
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Binary server is an example server.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/fgiudici/grpc-go"

	ecpb "github.com/fgiudici/grpc-go/examples/features/proto/echo"
	hwpb "github.com/fgiudici/grpc-go/examples/helloworld/helloworld"
)

var port = flag.Int("port", 50051, "the port to serve on")

// sayHello implements helloworld.GreeterServer.SayHello
func sayHello(ctx context.Context, in *hwpb.HelloRequest) (*hwpb.HelloReply, error) {
	return &hwpb.HelloReply{Message: "Hello " + in.Name}, nil
}

// unaryEcho implements echo.Echo.UnaryEcho
func unaryEcho(ctx context.Context, req *ecpb.EchoRequest) (*ecpb.EchoResponse, error) {
	return &ecpb.EchoResponse{Message: req.Message}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Printf("server listening at %v\n", lis.Addr())

	s := grpc.NewServer()

	// Register Greeter on the server.
	hwpb.RegisterGreeterService(s, &hwpb.GreeterService{SayHello: sayHello})

	// Register Echo on the same server.
	ecpb.RegisterEchoService(s, &ecpb.EchoService{UnaryEcho: unaryEcho})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
