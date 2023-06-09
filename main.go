package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/olzzhas/anthonygg-microservices/client"
	"log"
)

func main() {
	client := client.New("http://localhost:3000")

	price, err := client.FetchPrice(context.Background(), "GsG")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", price)

	return

	// https://www.youtube.com/watch?v=367qYRy39zw&list=PL0xRBLFXXsP5cru52B5GAQmIrTTAL8A66&index=5
	// 1:02:27

	listenAddr := flag.String("listenaddr", ":3000", "listen address the service is running")
	flag.Parse()

	svc := NewLoggingService(NewMetricService(&priceFetcher{}))

	server := NewJSONAPIServer(*listenAddr, svc)

	server.Run()
}
