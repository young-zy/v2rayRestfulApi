package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	logCommand "v2rayRestfulApi/log/command"
	statsCommand "v2rayRestfulApi/stats/command"
)

func runRestServer(endpoint string, listener net.Listener) error {
	mux := runtime.NewServeMux()
	dialOptions := []grpc.DialOption{grpc.WithInsecure()}
	ctx := context.Background()
	err := statsCommand.RegisterStatsServiceHandlerFromEndpoint(ctx, mux, endpoint, dialOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = logCommand.RegisterLoggerServiceHandlerFromEndpoint(ctx, mux, endpoint, dialOptions)
	if err != nil {
		log.Fatal(err)
	}
	return http.Serve(listener, mux)
}

func main() {
	port := flag.Int("port", 0, "restful server port")
	endpoint := flag.String("endpoint", "", "endpoint of grpc server， example: 127.0.0.1:10086")
	flag.Parse()
	log.Printf("Listening on: %d", *port)
	log.Printf("transforming data to: %s", *endpoint)
	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	err = runRestServer(*endpoint, listener)
	if err != nil {
		log.Fatal(err)
	}
}
