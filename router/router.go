package router

import (
	"net/http"
	"msg_quing_system/controller"
	
	"strings"
)
func Routes(w http.ResponseWriter, r *http.Request) {

	route := strings.Trim(r.URL.Path, "/")
	switch route {
	case "product":
		 if r.Method == "POST" {
			controller.AddProduct(w, r)
		} 
	}
}
