package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
var message = make(chan string)

func reverse(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return
	}
	defer ws.Close()

	fmt.Println("conexion establecido")

	//go server.Run("localhost:" + os.Getenv("PORT"))
	for {
		ws.WriteMessage(1, []byte(<-message))
	}
}

func serve() {
	// weebhook
	server := gin.Default()
	server.POST("", func(ctx *gin.Context) {
		res, _ := io.ReadAll(ctx.Request.Body)
		message <- string(res)
	})

	httpSrv := &http.Server{
		Addr:    ":" + os.Getenv("PORT_SERVER"),
		Handler: server,
	}
	defer httpSrv.Close()

	httpSrv.ListenAndServe()
}

func main() {
	go serve()
	http.HandleFunc("/reverse", reverse)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT_WS"), nil))
}
