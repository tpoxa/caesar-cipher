package main

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"tech-challenge/client/grpc"
)

var (
	grpcAddr string
	shifts   int
)

var rootCmd = &cobra.Command{
	Use:   "client",
	Short: "Run gRPC client",
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&grpcAddr, "connect", "c", "127.0.0.1:8080", "GRPC server address")
	rootCmd.PersistentFlags().IntVarP(&shifts, "shifts", "s", 1, "Number of shifts")
	var decrypt = &cobra.Command{
		Use:     "decrypt",
		Short:   "decrypt message",
		Long:    `Decrypt`,
		Example: "client decrypt -s 2 Ecguct",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				log.Fatal("please provide a message to decrypt")
			}
			client := grpc.NewClient(grpcAddr)
			result, err := client.Decrypt(context.TODO(), args[0], shifts)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Result is: %s\n", result)
		},
	}
	var encrypt = &cobra.Command{
		Use:     "encrypt",
		Short:   "encrypt message",
		Long:    `Encrypt`,
		Example: "client encrypt -s 2 Caesar",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				log.Fatal("please provide a message to encrypt")
			}
			client := grpc.NewClient(grpcAddr)
			result, err := client.Encrypt(context.TODO(), args[0], shifts)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Result is: %s\n", result)
		},
	}
	rootCmd.AddCommand(encrypt, decrypt)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
