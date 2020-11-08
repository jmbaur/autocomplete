package cmd

import (
	"context"
	"log"
	"net"

	"github.com/jmbaur/autocomplete/pkg/autocomplete"
	"github.com/jmbaur/autocomplete/pkg/pb"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

type server struct {
	pb.CompleterServer
}

var completer *autocomplete.Completer

var (
	serverCmd = &cobra.Command{
		Use:   "serve",
		Short: "Server short description",
		Long:  "Server long description",
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			var dictionary string
			switch len(args) {
			case 1:
				dictionary = args[0]
			case 0:
				dictionary = defaultDictionary
			default:
				log.Fatal("did not provide path to a dictionary file")
			}
			completer, err = autocomplete.NewCompleter(dictionary)
			if err != nil {
				log.Fatal(err)
			}

			lis, err := net.Listen("tcp", port)
			if err != nil {
				log.Fatalf("failed to listen: %v", err)
			}
			s := grpc.NewServer()
			pb.RegisterCompleterServer(s, &server{})
			if err := s.Serve(lis); err != nil {
				log.Fatalf("failed to serve: %v", err)
			}
		},
	}
)

func (s *server) Complete(ctx context.Context, in *pb.WordRequest) (*pb.WordsReply, error) {
	wordToComplete := in.GetFragment()
	log.Printf("completing fragment: \"%v\"", wordToComplete)
	possibleWords, err := completer.Complete(wordToComplete)
	if err != nil {
		log.Fatal(err)
	}
	return &pb.WordsReply{Completions: possibleWords}, nil
}
