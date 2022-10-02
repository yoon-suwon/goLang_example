package main
 
import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "html/template"
    "net/http"
    "log"
)
 
const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "yoon"
    dbname   = "postgres"
)

var templates = make(map[string]*template.Template)

var db *sql.DB


func main() {    
    var err error
    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
    // open database
    db, err = sql.Open("postgres", psqlconn)
    CheckError(err)
 
    defer db.Close()
     err = db.Ping()
    CheckError(err)
 
   fmt.Println("Connected!") 

    errr := http.ListenAndServe("localhost:8080", nil)
     if errr != nil {
      panic(errr)
    }

    
}

func loadTemplate(name string) *template.Template {
    t, err := template.ParseFiles(
        "templates/"+name+".html",
        "./templates/_header.html",
    )
    if err != nil {
        log.Fatal("template ParseFiles: ", err)
    }
    return t
}

func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}
