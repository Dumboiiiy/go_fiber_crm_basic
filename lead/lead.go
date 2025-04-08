package lead

import(
    "github.com/jinzhu/gorm"
    "github.com/gofiber/fiber/v2"
    _ "github.com/jinzhu/gorm/dialects/sqlite" // Import for side effects only
    "github.com/dumboiiiy/go_fiber_crm_basic/database"
)

type Lead struct { //lead struct is used to define the structure of the lead object
    gorm.Model //gorm.Model is a struct provided by gorm library, used to define the model for the lead object
    Name string `json:"name"` //Name is a field of type string, used to store the name of the lead
    Company string `json:"company"` //Company is a field of type string, used to store the company name of the lead
    Email string `json:"email"` //Email is a field of type string, used to store the email address of the lead
    Phone string `json:"phone"` //Phone is a field of type string, used to store the phone number of the lead
    // ID int `json:"id"` //ID is a field of type int, used to store the ID of the lead    
}

// Modified function signatures to return error as required by Fiber v2
func GetLeads(c *fiber.Ctx) error { //context is a basic golang concept, used to pass data between functions
    db := database.DBcon //db is a variable of type *gorm.DB, used to connect to the database
    var leads []Lead //leads is a slice of lead struct, used to store the list of leads
    //multiple leads with a particular struct, in short slice is an array of object
    db.Find(&leads) //db.Find is a gorm method, used to find all the records in the database and store them in the leads variable
    return c.JSON(leads) //c.JSON is a fiber method, used to send the response in JSON format to the client
    //c is a variable of type *fiber.Ctx, used to handle the request and response in fiber framework
}

func GetLead(c *fiber.Ctx) error {
    id := c.Params("id")
    db := database.DBcon
    var lead Lead
    db.Find(&lead, id)
    return c.JSON(lead)
}

func NewLead(c *fiber.Ctx) error {
    db := database.DBcon //lead is a variable of type lead struct, used to store the lead object, learn what postman or curl is, what is JSON
    lead := new(Lead)
    if err := c.BodyParser(&lead); err != nil { //c.BodyParser is a fiber method, used to parse the request body and store it in the lead variable
        // Fixed: Use SendString with err.Error() instead of Send with err
        return c.Status(503).SendString(err.Error())
    }
    db.Create(&lead) //db.Create is a gorm method, used to create a new record in the database
    return c.JSON(lead) //c.JSON is a fiber method, used to send the response in JSON format to the client
}

func DeleteLead(c *fiber.Ctx) error {
    id := c.Params("id") //c.Params is a fiber method, used to get the parameters from the request URL
    db := database.DBcon //db is a variable of type *gorm.DB, used to connect to the database
    var lead Lead //lead is a variable of type lead struct, used to store the lead object
    db.First(&lead, id) //db.First is a gorm method, used to find the first record in the database that matches the given ID and store it in the lead variable
    if lead.Name == "" {
        // Fixed: Use SendString instead of Send for string messages
        return c.Status(500).SendString("No Lead found with ID")
    }
    db.Delete(&lead) //db.Delete is a gorm method, used to delete the record from the database
    // Fixed: Use SendString instead of Send for string messages
    return c.SendString("Lead deleted successfully")
}
