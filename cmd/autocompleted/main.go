package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/jmbaur/autocomplete/pkg/autocomplete"
	"github.com/jmbaur/autocomplete/pkg/pb"
	"google.golang.org/grpc"
)

const (
	port              = ":50051"
	defaultDictionary = "/usr/share/dict/words"
)

var completer *autocomplete.Completer

type server struct {
	pb.CompleterServer
}

func (s *server) Complete(ctx context.Context, in *pb.WordRequest) (*pb.WordsReply, error) {
	wordToComplete := in.GetFragment()
	log.Printf("completing fragment: \"%v\"", wordToComplete)
	possibleWords, err := completer.Complete(wordToComplete)
	if err != nil {
		log.Fatal(err)
	}
	return &pb.WordsReply{Completions: possibleWords}, nil
}

func main() {
	var err error
	var dictionary string
	switch len(os.Args) {
	case 2:
		dictionary = os.Args[1]
	case 1:
		dictionary = defaultDictionary
	default:
		fmt.Println("Usage: autocompleted <path to dictionary>")
		os.Exit(1)
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
}
