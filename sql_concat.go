package main
import (
        "database/sql"
        //_ "github.com/mattn/go-sqlite3"
        "os"
)
func main(){
        db, err := sql.Open("sqlite3", ":memory:")
        if err != nil {
                panic(err)
        }
        rows, err := db.Query("SELECT * FROM foo WHERE name = " + os.Args[1])
        if err != nil {
                panic(err)
        }
        defer rows.Close()
}
