package app

import (
	"net/http"
	"fmt"
	"html"
	"log"
	"database/sql"
	_ "gopkg.in/goracle.v2"
)

type App struct {
	DB *string
}

func (a *App) Run(p string) {
	fmt.Sprintf("Hello, %s", p)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		db, err := sql.Open("goracle", "sys/Oradoc_db1@ORCLCDB as sysdba")
		// sys/Oradoc_db1@127.0.0.1:1521/ORCLPDB1
	    if err != nil {
	        fmt.Println(err)
	        return
	    }
	    defer db.Close()
		rows,err := db.Query("select sysdate from dual")
	    if err != nil {
	        fmt.Println("Error running query")
	        fmt.Println(err)
	        return
	    }
	    defer rows.Close()
	 
	    var thedate string
	    for rows.Next() {
	 
	        rows.Scan(&thedate)
	    }
	    fmt.Printf("The DB date is: %s\n", thedate)
        fmt.Fprintf(w, "The DB date is: %s\n", thedate)
        fmt.Fprintf(w, "URI, %q", html.EscapeString(r.URL.Path))
    })
    log.Fatal(http.ListenAndServe(p, nil))
}


