package main

import (
	"fmt"
	"github.com/XmahopAbuse/hydra-billing-api/internal/user"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	fmt.Println("Create router...")
	router := httprouter.New()

	handler := user.NewHandler()
	handler.Register(router)

	http.ListenAndServe(":8080", router)
}
