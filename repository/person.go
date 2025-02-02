
package repository

import (
    "database/sql"
    "event_api/structs"
    "github.com/lib/pq"
    "fmt"
)

func GetAllPerson(db *sql.DB) (result []structs.PersonWithTicket, err error) {
    sql := `SELECT person.id, person.full_name, person.address_person, person.status, person.created_at, ticket.id, ticket.price, ticket.status, ticket.quantity, event_ac.title, event_ac.location_event, event_ac.date_event
            FROM person
            JOIN ticket ON ticket.id = ANY(person.ticket_ids)
            JOIN event_ac ON ticket.event_id = event_ac.id
            ORDER BY person.id`

            rows, err := db.Query(sql)
            if err != nil {
                return nil, err
            }
            defer rows.Close()
        
            var currentPerson structs.PersonWithTicket
            for rows.Next() {
                var person structs.PersonWithTicket
                var ticket structs.TicketWithEvent
                fmt.Println("ini jalan")
        
                // Scan data 
                err = rows.Scan(
                    &person.ID, 
                    &person.FullName, 
                    &person.AddressPerson, 
                    &person.Status, 
                    &person.CreatedAt,
                    &ticket.ID, 
                    &ticket.Price, 
                    &ticket.Status, 
                    &ticket.Quantity, 
                    &ticket.Event.Title, 
                    &ticket.Event.Location, 
                    &ticket.Event.Date_Event,
                )
                if err != nil {
                    return nil, err
                }

            // Menambahkan tiket ke dalam array tickets
            if len(currentPerson.Tickets) == 0 || currentPerson.ID != person.ID {
                if currentPerson.ID != 0 {
                    result = append(result, currentPerson) 
                }
                currentPerson = person 
                currentPerson.Tickets = []structs.TicketWithEvent{ticket} 
            } else {
                currentPerson.Tickets = append(currentPerson.Tickets, ticket)
            }
    }

    if len(result) == 0 {
        fmt.Println("No data in result slice.")
    }

    if currentPerson.ID != 0 {
        result = append(result, currentPerson)
    }

        
    return
}

func InsertPerson(db *sql.DB, person structs.Person) (err error) {
    sql := "INSERT INTO person(full_name, address_person, ticket_ids, status) VALUES ($1, $2, $3, $4)"

    errs := db.QueryRow(sql, person.FullName, person.Address_Person, pq.Array(person.Ticket_IDS), person.Status)

    return errs.Err()
}



func UpdatePerson(db *sql.DB, person structs.Person) (err error) {
    sql := "UPDATE person SET full_name = $1, address_person = $2, ticket_ids = $3, status = $4 WHERE id = $5"

    errs := db.QueryRow(sql, person.FullName, person.Address_Person, pq.Array(person.Ticket_IDS), person.Status, person.ID)

    return errs.Err()
}

func DeletePerson(db *sql.DB, person structs.Person) (err error) {
    sql := "DELETE FROM person WHERE id = $1"

    errs := db.QueryRow(sql, person.ID)
    return errs.Err()
}

