package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

type User struct {
	ID   int
	Name string
}

func GetUserById(id int) (string, error) {
	user := new(User) //用new()函数初始化一个结构体对象
	row := DB.QueryRow("SELECT id, name FROM users WHERE id=?", id)
	if err := row.Scan(&user.ID, &user.Name); err != nil {
		if err == sql.ErrNoRows {
			return "", errors.WithMessage(err, fmt.Sprintf("can't match id in users: %d", id))
		}
		return "", errors.Wrap(err, fmt.Sprintf("failed in getUserById %d", id))
	}

	return user.Name, nil
}
