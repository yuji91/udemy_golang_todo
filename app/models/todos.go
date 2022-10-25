package models

import "log"
import "time"

type Todo struct {
  ID int
  Content string
  UserID int
  CreatedAt time.Time
}

// Userのメソッドとして作成
func (u *User) CreateTodo(content string) (err error) {
  cmd := `INSERT INTO todos(
      content,
      user_id,
      created_at) VALUES (?, ?, ?)`

  _, err = Db.Exec(cmd, content, u.ID, time.Now())

  if err != nil {
    log.Fatalln(err)
  }
  return err
}

func GetTodo(id int) (todo Todo, err error) {
  cmd := `SELECT id, content, user_id, created_at
          FROM todos
          WHERE id = ?`
  todo = Todo{}

  err = Db.QueryRow(cmd, id).Scan(
      &todo.ID,
      &todo.Content,
      &todo.UserID,
      &todo.CreatedAt)

  return todo, err
}

func GetTodos() (todos []Todo, err error) {
  cmd := `SELECT id, content, user_id, created_at
          FROM todos`
  rows, err := Db.Query(cmd)
  if err != nil {
    log.Fatalln(err)
  }
  for rows.Next() {
    var todo Todo
    err = rows.Scan(&todo.ID,
        &todo.Content,
        &todo.UserID,
        &todo.CreatedAt)
    if err != nil {
      log.Fatalln(err)
    }
    todos = append(todos, todo)
  }
  rows.Close()

  return todos, err
}

// Userのメソッドとして作成
func (u *User) GetTodosByUser() (todos []Todo, err error) {
  cmd := `SELECT id, content, user_id, created_at
          FROM todos
          WHERE user_id = ?`
  rows, err := Db.Query(cmd, u.ID)
  if err != nil {
    log.Fatalln(err)
  }
  for rows.Next() {
    var todo Todo
    err = rows.Scan(&todo.ID,
        &todo.Content,
        &todo.UserID,
        &todo.CreatedAt)
    if err != nil {
      log.Fatalln(err)
    }
    todos = append(todos, todo)
  }
  rows.Close()

  return todos, err
}

func (t *Todo) UpdateTodo() error {
  cmd := `UPDATE todos SET  content = ?, user_id = ?
          WHERE id = ?`
  _, err := Db.Exec(cmd, t.Content, t.UserID, t.ID)
  if err != nil {
    log.Fatalln(err)
  }
  return err
}
