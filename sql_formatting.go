package main
import (
        "database/sql"
        "fmt"
        "os"
        //_ "github.com/mattn/go-sqlite3"
)
func main(){
        db, err := sql.Open("sqlite3", ":memory:")
        if err != nil {
                panic(err)
        }
        q := fmt.Sprintf("SELECT * FROM foo where name = '%s'", os.Args[1])
        rows, err := db.Query(q)
        if err != nil {
                panic(err)
        }
        defer rows.Close()
}
