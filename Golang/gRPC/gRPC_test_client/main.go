package main

import (
	"google.golang.org/grpc"
	"github.com/gin-gonic/gin"
	"grpc_test"
	"strconv"
	"net/http"
)

func main(){
	conn, err := grpc.Dial("localhost:8882", grpc.WithInsecure())
	if err != nil{
		panic(err)
	}
	client := grpc_test.NewAddServiceClient(conn)
	g := gin.Default()
	g.GET("/add/:a/:b", func(c *gin.Context){
		a, _ := strconv.ParseUint(c.Param("a"), 10, 64)
		b, _ := strconv.ParseUint(c.Param("b"), 10, 64)
		req := &grpc_test.Request{A: int64(a), B: int64(b)}
		if _, err := client.Add(c, req); err == nil{
			c.JSON(http.StatusOK, gin.H{
				"result": "Good",
			})
		}
	})
}