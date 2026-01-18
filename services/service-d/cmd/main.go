package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	commonpb "grpc-vs-rest-poc/proto"
	grpcserver "grpc-vs-rest-poc/services/service-d/internal/grpc"
	"grpc-vs-rest-poc/services/service-d/internal/rest"
	"grpc-vs-rest-poc/services/service-d/internal/service"
	"log"
	"net"
	"net/http"
)

func main() {
	processor := &service.Processor{ServiceName: "service-d"}

	grpcSer := grpc.NewServer()

	commonpb.RegisterProcessorServer(grpcSer, grpcserver.New(processor))
	reflection.Register(grpcSer)

	go func() {
		lis, err := net.Listen("tcp", ":9084")
		if err != nil {
			log.Fatal(err)
		}

		log.Println("gRPC listening on :9084")
		log.Fatal(grpcSer.Serve(lis))
	}()

	handler := rest.New(processor)
	mux := http.NewServeMux()
	mux.HandleFunc("/process", handler.Process)

	log.Println("REST listening on :8084")
	log.Fatal(http.ListenAndServe(":8084", mux))
}
