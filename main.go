package main

import (
    "database/sql"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "Event-Management-API/controllers"
    "Event-Management-API/database"
    "os"

    _ "github.com/lib/pq"
)

var (
    DB  *sql.DB
    err error
)

func main() {

    err = godotenv.Load("config/.env")
    if err != nil {
       panic("Error loading .env file")
    }

    psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
       os.Getenv("DB_HOST"),
       os.Getenv("DB_PORT"),
       os.Getenv("DB_USER"),
       os.Getenv("DB_PASSWORD"),
       os.Getenv("DB_NAME"),
    )

    DB, err = sql.Open("postgres", psqlInfo)
    defer DB.Close()
    err = DB.Ping()
    if err != nil {
       panic(err)
    }

    database.DBMigrate(DB)

    router := gin.Default()
    // event
    router.GET("/event",controllers.GetAllEvent)
    router.POST("/event", controllers.InsertEvent)
    router.PUT("/event/:id", controllers.UpdateEvent)
    router.DELETE("/event/:id", controllers.DeleteEvent)

    // ticket
    router.GET("/ticket", controllers.GetAllTicket)
    router.POST("/ticket", controllers.InsertTicket)
    router.PUT("/ticket/:id", controllers.UpdateTicket)
    router.DELETE("/ticket/:id", controllers.DeleteTicket)


    router.GET("/person", controllers.GetAllPerson)
    router.POST("/person", controllers.InsertPerson)
    router.PUT("/person/:id", controllers.UpdatePerson)
    router.DELETE("/person/:id", controllers.DeletePerson)

    // Auth
	router.POST("/user", controllers.Register)
    router.POST("/login", controllers.Login)

    router.Run(":8000")
}

