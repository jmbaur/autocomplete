package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jmbaur/autocomplete/pkg/pb"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	cliCmd = &cobra.Command{
		Use:   "get",
		Short: "CLI short description",
		Long:  "CLI long description",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				log.Fatal("did not provide one word to complete")
			}
			fragment := args[0]

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
		},
	}
)
