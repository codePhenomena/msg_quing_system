package utility

import (
	
	"html/template"
	"os"
	"log"
	"encoding/json"
	"net/http"
	"github.com/jmoiron/sqlx"
	
)

// Template Pool
var View *template.Template


// DB Connections
var Db *sqlx.DB



type AjaxRequest struct {
	Token   string
	Userid  string
	Payload map[string]string
	Email   string
}

type AjaxResponse struct {
	Status  string
	Message string
	Payload interface{}
}



func RedirectTo(w http.ResponseWriter, r *http.Request, path string) {
	http.Redirect(w, r, os.Getenv("APP_URL")+"/"+path, http.StatusFound)
}

/* if isCurlAPI w.Write json otherwise ExcuteTemplate() */
func RenderTemplate(w http.ResponseWriter, r *http.Request, template string, data interface{}) {

	jsonresponse, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	} else {
		w.Write([]byte(jsonresponse))
	}
}





