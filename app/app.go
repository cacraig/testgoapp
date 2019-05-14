package app

import (
    "net/http"
    "fmt"
    //"html"
    "log"
    "database/sql"
    "encoding/json"
    _ "gopkg.in/goracle.v2"
)

type App struct {
    DB *string
}

type DBDate struct {
    Date string `json:"date"`
}

// respondJSON makes the response with payload as json format
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
    response, err := json.Marshal(payload)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(err.Error()))
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    w.Write([]byte(response))
}

// respondError makes the error response with payload as json format
func respondError(w http.ResponseWriter, code int, message string) {
    respondJSON(w, code, map[string]string{"error": message})
}

func (a *App) Run(p string) {

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

        dbdate := &DBDate{
            Date: thedate }

        res, _ := json.Marshal(dbdate)

        fmt.Println(string(res))
        fmt.Printf("The DB date is: %s\n", thedate)
        respondJSON(w, http.StatusOK, dbdate)
        //fmt.Fprintf(w, "The DB date is: %s\n", thedate)
        //fmt.Fprintf(w, "URI, %q", html.EscapeString(r.URL.Path))
    })
    log.Fatal(http.ListenAndServe(p, nil))
}




