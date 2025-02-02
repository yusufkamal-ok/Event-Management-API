
package controllers

import (
    "github.com/gin-gonic/gin"
    "event_api/database"
    "event_api/repository"
    "event_api/structs"
    "net/http"
    "strconv"
)

func GetAllEvent(c *gin.Context) {
    var (
       result gin.H
    )

    person, err := repository.GetAllEvent(database.DbConnection)

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

func InsertEvent(c *gin.Context) {
    var event structs.Event

    err := c.BindJSON(&event)
    if err != nil {
       panic(err)
    }

    err = repository.InsertEvent(database.DbConnection, event)
    if err != nil {
       panic(err)
    }

    c.JSON(http.StatusOK, event)
}

func UpdateEvent(c *gin.Context) {
    var event structs.Event
    id, _ := strconv.Atoi(c.Param("id"))

    err := c.BindJSON(&event)
    if err != nil {
       panic(err)
    }

    event.ID = id

    err = repository.UpdateEvent(database.DbConnection, event)
    if err != nil {
       panic(err)
    }

    c.JSON(http.StatusOK, event)
}

func DeleteEvent(c *gin.Context) {
    var event structs.Event
    id, _ := strconv.Atoi(c.Param("id"))

    event.ID = id
    err := repository.DeleteEvent(database.DbConnection, event)
    if err != nil {
       panic(err)
    }

    c.JSON(http.StatusOK, gin.H{"message": "Event Successfully Deleted"})
}

