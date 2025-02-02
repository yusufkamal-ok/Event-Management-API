
package repository

import (
    "database/sql"
    "event_api/structs"
    "fmt"
)

func AddUser(db *sql.DB, user structs.User) (err error) {
    sql := `INSERT INTO account (email, pass_word) VALUES ($1, $2)`

    _,err = db.Exec(sql, user.Email, user.Password)
    
    if err != nil {
        fmt.Println("Error executing query:", err) 
    }

    return err
};

func LogUser(db *sql.DB, email string) (structs.User, error) {
	var user structs.User
	sql := "SELECT email, pass_word FROM account WHERE email = $1"

	err := db.QueryRow(sql, email).Scan(&user.Email, &user.Password)

	if err != nil {
		return user, err
	}

	return user, nil
}

