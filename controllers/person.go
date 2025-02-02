
package controllers

import (
    "github.com/gin-gonic/gin"
    "Event-Management-API/database"
    "Event-Management-API/repository"
    "Event-Management-API/structs"
    "net/http"
    "strconv"
)

func GetAllPerson(c *gin.Context) {
    var (
       result gin.H
    )

    person, err := repository.GetAllPerson(database.DbConnection)

    if err != nil {
       result = gin.H{
          "result": err.Error(),
       }
    } else {
       result = gin.H{
          "result": person,
       }
    }

    c.JSON(http.StatusOK, result)
}

func InsertPerson(c *gin.Context) {
    var person structs.Person

    err := c.BindJSON(&person)
    if err != nil {
       panic(err)
    }

    if len(person.Ticket_IDS) == 0 || person.FullName == "" {
      c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
      return
  }

    validStatuses := map[string]bool{"registered": true, "pending": true,"canceled": true}
    if person.Status == "" {
        person.Status = "registered" // Gunakan nilai default jika kosong
    } else if !validStatuses[person.Status] {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status person . Allowed: registered, pending, canceled"})
        return
    }

    err = repository.InsertPerson(database.DbConnection, person)
    if err != nil {
       panic(err)
    }

    c.JSON(http.StatusOK, person)
}

func UpdatePerson(c *gin.Context) {
    var person structs.Person
    id, _ := strconv.Atoi(c.Param("id"))

    err := c.BindJSON(&person)
    if err != nil {
       panic(err)
    }

    validStatuses := map[string]bool{"registered": true, "pending": true,"canceled": true}
    if person.Status == "" {
        person.Status = "registered" 
    } else if !validStatuses[person.Status] {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status person . Allowed: registered, pending, canceled"})
        return
    }
    person.ID = id

    err = repository.UpdatePerson(database.DbConnection, person)
    if err != nil {
       panic(err)
    }

    c.JSON(http.StatusOK, person)
}

func DeletePerson(c *gin.Context) {
    var person structs.Person
    id, _ := strconv.Atoi(c.Param("id"))

    person.ID = id
    err := repository.DeletePerson(database.DbConnection, person)
    if err != nil {
       panic(err)
    }

    c.JSON(http.StatusOK, gin.H{"message": "Person Successfully Deleted"})
}

