package main

import (
	
	
	"log"
	"net/http"
	"os"
	
	"msg_quing_system/router"
	"msg_quing_system/utility"
	

	_ "github.com/go-sql-driver/mysql"
	
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func init() {
	
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Println(".env file wasn't found, looking at env variables")
	}
	motd()
	// Read Config
	utility.Db, err = sqlx.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@/"+os.Getenv("DB_NAME"))
	if err != nil {
		log.Println("Wowza !, We didn't find the DB or you forgot to setup the env variables")
		panic(err)
	}
	
	
}

func main() {
	http.HandleFunc("/", handler)
	port:=os.Getenv("WEB_PORT")
	log.Fatal(http.ListenAndServe(":"+port, nil))
	if os.Getenv("WEB_PORT")==""{
		port="8080"
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	router.Routes(w, r)
}

func motd() {
	logo := `
______ _____  ___   _   __
S_E_R_V_E_R__I_S__R_E_A_D_Y
                          
----------------------------
Application should now be accessible on port ` + os.Getenv("WEB_PORT") + `

`
	log.Println(logo)
}
