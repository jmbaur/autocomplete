package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmbaur/autocomplete/pkg/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: autocomplete <word to complete>")
		os.Exit(1)
	}
	fragment := os.Args[1]

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewCompleterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Complete(ctx, &pb.WordRequest{Fragment: fragment})
	if err != nil {
		log.Fatalf("could not complete fragment \"%s\": %v", fragment, err)
	}

	completions := r.GetCompletions()
	for _, v := range completions {
		fmt.Printf("%s\n", v)
	}
}
