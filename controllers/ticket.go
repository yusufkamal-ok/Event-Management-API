
package controllers

import (
    "github.com/gin-gonic/gin"
    "Event-Management-API/database"
    "Event-Management-API/repository"
    "Event-Management-API/structs"
    "net/http"
    "strconv"
)

func GetAllTicket(c *gin.Context) {
    var (
       result gin.H
    )

    event, err := repository.GetAllTicket(database.DbConnection)

    if err != nil {
       result = gin.H{
          "result": err.Error(),
       }
    } else {
       result = gin.H{
          "result": event,
       }
    }

    c.JSON(http.StatusOK, result)
}

func InsertTicket(c *gin.Context) {
    var ticket structs.Ticket

    err := c.BindJSON(&ticket)
    if err != nil {
       panic(err)
    }

	// Validasi nilai yang harus ada
    if ticket.Event_ID == 0 || ticket.Price <= 0 || ticket.Quantity <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
        return
    }

	validStatuses := map[string]bool{"VIP": true, "Regular": true}
    if ticket.Status == "" {
        ticket.Status = "Regular" 
    } else if !validStatuses[ticket.Status] {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket . Allowed: VIP, Regular"})
        return
    }

    err = repository.InsertTicket(database.DbConnection, ticket)
    if err != nil {
       panic(err)
    }

    c.JSON(http.StatusOK, ticket)
}

func UpdateTicket(c *gin.Context) {
    var ticket structs.Ticket
    id, _ := strconv.Atoi(c.Param("id"))

    err := c.BindJSON(&ticket)
    if err != nil {
       panic(err)
    }

	  // Validasi ENUM status hanya boleh 'VIP' atau 'Regular'
	validStatuses := map[string]bool{"VIP": true, "Regular": true}
	  if _, valid := validStatuses[ticket.Status]; !valid {
		  c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket status. Allowed: VIP, Regular"})
		  return
	}

    ticket.ID = id

    err = repository.UpdateTicket(database.DbConnection, ticket)
    if err != nil {
       panic(err)
    }

    c.JSON(http.StatusOK, ticket)
}

func DeleteTicket(c *gin.Context) {
    var ticket structs.Ticket
    id, _ := strconv.Atoi(c.Param("id"))

    ticket.ID = id
    err := repository.DeleteTicket(database.DbConnection, ticket)
    if err != nil {
       panic(err)
    }

    c.JSON(http.StatusOK, gin.H{"message": "Ticket Successfully Deleted"})
}

