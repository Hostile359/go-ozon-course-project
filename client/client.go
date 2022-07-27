package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	pb "gitlab.ozon.dev/Hostile359/homework-1/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conns, err := grpc.Dial(":8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conns.Close()

	client := pb.NewAdminClient(conns)

	ctx := context.Background()

	cmdHandler := NewCommandHandler(client, ctx)
	fmt.Println(cmdHandler.HandleCommand("help"))
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter command:")
		sc.Scan()
		cmd := sc.Text()
		response := cmdHandler.HandleCommand(cmd)
		fmt.Printf("response: [%v]\n", response)
	}
}
