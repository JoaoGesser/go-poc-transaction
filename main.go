package main
import (
  "fmt"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

type DataTest struct {
  id int `json:"id"`
}


func main() {
  db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()
  //insert, err := db.Query("INSERT INTO dataTest (text) VALUES ('Inseting 1')")
  results, err := db.Query("SELECT id FROM dataTest WHERE id = 1 FOR UPDATE")
  if err != nil {
    panic(err.Error())
  }
  for results.Next() {
    var data DataTest
    err = results.Scan(&data.id)
    if err != nil {
      panic(err.Error())
    }
    fmt.Println("First query: " + fmt.Sprint(data.id))
  }
  //defer insert.Close()
  updateQuery, err := db.Query("UPDATE dataTest SET text = 'Jhon' WHERE id = 1")
  if err != nil {
    panic(err.Error())
  }
  fmt.Println("Executed update!")
  defer updateQuery.Close()
  defer results.Close()
  fmt.Println("Success!")


}

