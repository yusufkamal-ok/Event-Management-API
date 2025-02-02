
package repository

import (
    "database/sql"
    "event_api/structs"
)

func GetAllTicket(db *sql.DB) (result []structs.TicketWithEvent, err error) {
    sql := `SELECT ticket.id, ticket.event_id, ticket.price, ticket.status, ticket.quantity, ticket.created_at, event_ac.title, event_ac.location_event, event_ac.date_event
			FROM ticket 
			JOIN event_ac  
			ON ticket.event_id = event_ac.id`

    rows, err := db.Query(sql)
    if err != nil {
       return
    }

    defer rows.Close()
    for rows.Next() {
       var ticket structs.TicketWithEvent

       err = rows.Scan(&ticket.ID, &ticket.Event_ID, &ticket.Price, &ticket.Status, &ticket.Quantity, &ticket.Created_At,
		&ticket.Event.Title, &ticket.Event.Location, &ticket.Event.Date_Event)
       if err != nil {
          return
       }

       result = append(result, ticket)
    }

    return
}

func InsertTicket(db *sql.DB, ticket structs.Ticket) (err error) {
    sql := "INSERT INTO ticket(event_id, price, status, quantity) VALUES ($1, $2, $3, $4)"

    errs := db.QueryRow(sql, ticket.Event_ID, ticket.Price, ticket.Status, ticket.Quantity)

    return errs.Err()
}

func UpdateTicket(db *sql.DB, ticket structs.Ticket) (err error) {
    sql := "UPDATE ticket SET event_id = $1, price = $2, status = $3, quantity = $4 WHERE id = $5"

    errs := db.QueryRow(sql, ticket.Event_ID, ticket.Price, ticket.Status, ticket.Quantity, ticket.ID)

    return errs.Err()
}

func DeleteTicket(db *sql.DB, ticket structs.Ticket) (err error) {
    sql := "DELETE FROM ticket WHERE id = $1"

    errs := db.QueryRow(sql, ticket.ID)
    return errs.Err()
}

