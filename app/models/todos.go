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
