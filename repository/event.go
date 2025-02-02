
package repository

import (
    "database/sql"
    "event_api/structs"
)

func GetAllEvent(db *sql.DB) (result []structs.Event, err error) {
    sql := "SELECT * FROM event_ac"

    rows, err := db.Query(sql)
    if err != nil {
       return
    }

    defer rows.Close()
    for rows.Next() {
       var event structs.Event

       err = rows.Scan(&event.ID, &event.Title, &event.Description, &event.Location, &event.Date_Event, &event.Created_At)
       if err != nil {
          return
       }

       result = append(result, event)
    }

    return
}

func InsertEvent(db *sql.DB, event structs.Event) (err error) {
    sql := "INSERT INTO event_ac(title, description_event, location_event, date_event) VALUES ($1, $2, $3, $4)"

    errs := db.QueryRow(sql, event.Title, event.Description, event.Location, event.Date_Event)

    return errs.Err()
}

func UpdateEvent(db *sql.DB, event structs.Event) (err error) {
    sql := "UPDATE event_ac SET title = $1, description_event = $2, location_event = $3, date_event = $4 WHERE id = $5"

    errs := db.QueryRow(sql, event.Title, event.Description, event.Location, event.Date_Event, event.ID)

    return errs.Err()
}

func DeleteEvent(db *sql.DB, event structs.Event) (err error) {
    sql := "DELETE FROM event_ac WHERE id = $1"

    errs := db.QueryRow(sql, event.ID)
    return errs.Err()
}

