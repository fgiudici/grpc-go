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
	"github.com/fgiudici/grpc-go/credentials/alts"

	pb "github.com/fgiudici/grpc-go/examples/features/proto/echo"
)

var port = flag.Int("port", 50051, "the port to serve on")

func unaryEcho(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{Message: req.Message}, nil
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Create alts based credential.
	altsTC := alts.NewServerCreds(alts.DefaultServerOptions())

	s := grpc.NewServer(grpc.Creds(altsTC))

	// Register EchoService on the server.
	pb.RegisterEchoService(s, &pb.EchoService{UnaryEcho: unaryEcho})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
