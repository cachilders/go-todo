package models

import (
  "database/sql"

  _ "github.com/mattn/go-sqlite3"
)

// A struct is a grouping of data fields, like a schema model, I guess
// The backticks are denoting metadata used in destructuring by c.Bind

type Task struct {
  ID int      `json:"id"`
  Name string `json:"name"`
}

type TaskCollection struct {
  Tasks []Task `json:"items"`
}

func GetTasks(db *sql.DB) TaskCollection {
  sql := "SELECT * FROM tasks"
  
  rows, err := db.Query(sql)

  if err != nil {
    panic(err)
  }

  defer rows.Close()

  result := TaskCollection{}
  for rows.Next() {
    task := Task{}
    err2 := rows.Scan(&task.ID, &task.Name)

    if err2 != nil {
      panic(err2)
    }

    result.Tasks = append(result.Tasks, task)
  }
  return result
}

func PutTask(db *sql.DB, name string) (int64, error) {
  sql := "INSERT INTO tasks(name) VALUES(?)"
  // Create SQL statement
  stmt, err := db.Prepare(sql)

  if err != nil {
    panic(err)
  }

  defer stmt.Close()
  // This is where we assign the ? value to the query
  result, err2 := stmt.Exec(name)
  if err2 != nil {
    panic(err2)
  }
  return result.LastInsertId()
}

func DeleteTask(db *sql.DB, id int) (int64, error) {
  sql := "DELETE FROM tasks WHERE id = ?"
  // OK, the value of these prepared statements is they compile AND they
  // protect against injection attacks
  stmt, err := db.Prepare(sql)

  if err != nil {
    panic(err)
  }

  result, err2 := stmt.Exec(id)

  if err2 != nil {
    panic(err2)
  }

  return result.RowsAffected()
}