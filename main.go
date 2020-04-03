package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/matkinhig/developer-community/server"
)

// const port = ":50051"

func main() {
	fmt.Println("start golang...")
	g := gin.Default()
	server.Run(g)
}

// func premain() {
// 	// 1. Listen/Open a TPC connect at port
// 	lis, _ := net.Listen("tcp", port)
// 	// 2. Tao server tu GRP
// 	grpcServer := grpc.NewServer()
// 	// 3. Map service to server
// 	pb.RegisterNoteServiceServer(grpcServer, &noteService{})
// 	// 4. Binding port
// 	grpcServer.Serve(lis)
// }
