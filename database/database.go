package database

import(
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite" // Import for side effects only
)


var(
	DBcon *gorm.DB //DBcon is a pointer to gorm.DB struct, instance of gorm framework
	//DBcon is a pointer to gorm.DB struct, instance of gorm framework
)