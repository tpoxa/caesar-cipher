package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"tech-challenge/internal/caesar"
	"tech-challenge/pkg/proto/cipher"
	grpcserver "tech-challenge/server/grpc"
)

var listenAddr string

var rootCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Run the gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		// Create new gRPC server instance
		s := grpc.NewServer()
		// enables auto discovery
		reflection.Register(s)
		// create cipher instance with latin alphabet
		cipherImp := caesar.NewCipher([]rune("abcdefghijklmnopqrstuvwxyz"))
		// create new service instance
		srv := grpcserver.NewServer(cipherImp)
		// Register gRPC server
		cipher.RegisterCipherServer(s, srv)
		// Listen on port 8080
		l, err := net.Listen("tcp", listenAddr)
		if err != nil {
			log.Fatal(err)
		}
		// Start gRPC server
		log.Println("Server is running ...")
		log.Printf("Listening %s\n", listenAddr)
		if err := s.Serve(l); err != nil {
			log.Fatal(err)
		}
	},
}

func main() {
	rootCmd.Flags().StringVarP(&listenAddr, "listen", "l", ":8080", "Listen address")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
