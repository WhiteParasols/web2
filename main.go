package main

import (
	"net/http"

	"github.com/WhiteParasols/web2/myapp"
)

func main() {

	http.ListenAndServe(":3000", myapp.NewHandler())
}
