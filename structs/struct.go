package structs
import (
    "time"
)

type User struct {

    Email string `json:"email" binding:"required"`
    Password  string `json:"pass_word" binding: "required"`
}

type Event struct {
    ID int `json:"id"`
    Title string `json:"title"`
    Description  string `json:"description_event"`
    Location  string `json:"location_event"`
    Date_Event time.Time `json:"date_event"`
    Created_At      time.Time `json:"created_at"`
}

type Ticket struct {
    ID int `json:"id"`
    Event_ID int `json:"event_id"`
    Price int `json:"price"`
    Status string `json:"status"`
    Quantity int `json:"quantity"`
    Created_At    time.Time `json:"created_at"`
}

type TicketWithEvent struct {
    ID         int       `json:"id"`
    Event_ID   int       `json:"event_id"`
    Price      int       `json:"price"`
    Status     string    `json:"status"`
    Quantity   int       `json:"quantity"`
    Event      struct {
        Title      string    `json:"title"`
        Location string     `json:"location_event"`
        Date_Event time.Time `json:"date_event"`
    } `json:"event"`
    Created_At time.Time `json:"created_at"`
}



type Person struct {
    ID        int    `json:"id"`
    FullName string `json:"full_name"`
    Address_Person  string `json:"address_person"`
    Ticket_IDS []int `json:"ticket_ids`
    Status string `json:"status"`
}

type PersonWithTicket struct {
    ID            int       `json:"id"`
    FullName      string    `json:"full_name"`
    AddressPerson string    `json:"address_person"`
    Status        string    `json:"status"`
    CreatedAt     time.Time `json:"created_at"`
    Tickets       []TicketWithEvent `json:"tickets"`
    Created_At  time.Time   `json:"created_at"`
}