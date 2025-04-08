package main

import(
    _ "fmt"
    "github.com/dumboiiiy/go_fiber_crm_basic/lead"
    "github.com/gofiber/fiber/v2"
    "github.com/dumboiiiy/go_fiber_crm_basic/database"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite" // Required for database connection
)

func setupRoutes(app *fiber.App) { //setupRoutes function to setup the routes for the app
    app.Get("api/v1/lead", lead.GetLeads) // Function signature now matches fiber.Handler
    app.Get("api/v1/lead/:id", lead.GetLead) // Function signature now matches fiber.Handler
    app.Post("api/v1/lead", lead.NewLead) // Function signature now matches fiber.Handler
    app.Delete("api/v1/lead/:id", lead.DeleteLead) // Function signature now matches fiber.Handler
}

func initDatabase() {
    var err error //err is a variable of type error, used to handle errors
    database.DBcon, err = gorm.Open("sqlite3", "test.db") //gorm.Open function is used to open a connection to the database
    if err != nil { //if there is an error opening the database connection
        panic(err) //panic function is used to stop the execution of the program and print the error message
    }
    println("Database connection opened") //print message to console
    database.DBcon.AutoMigrate(&lead.Lead{}) //    AutoMigrate function is used to automatically migrate the database schema, creating the table if it does not exist
    println("Database migrated") //print message to console
}

func main() {
    app := fiber.New() //app is a pointer to fiber.App struct, instance of fiber framework
    initDatabase() //initDatabase function is called to initialize the database connection
    setupRoutes(app) //setupRoutes function is called to setup the routes for the app
    
    // Improved error handling for Listen
    if err := app.Listen(":3000"); err != nil {
        panic(err)
    }
    
    defer database.DBcon.Close() //defer statement is used to close the database connection when the main function exits 
}
