package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmbaur/autocomplete/pkg/pb"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	for {
		fmt.Printf("autosh> ")
		var fragment string
		_, err := fmt.Scanln(&fragment)
		if err != nil {
			fmt.Println("Could not complete fragment")
		}
		if fragment == "quit" || fragment == "q" {
			os.Exit(0)
		}
		c := pb.NewCompleterClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.Complete(ctx, &pb.WordRequest{Fragment: fragment})
		if err != nil {
			fmt.Printf("could not complete fragment \"%s\": %v\n", fragment, err)
		}
		completions := r.GetCompletions()
		for _, v := range completions {
			fmt.Printf("%s\n", v)
		}
	}
}
