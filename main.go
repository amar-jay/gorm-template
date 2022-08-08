package main

import (
  "fmt"

  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
)

// Model 
type Todo struct {
  gorm.Model
  Name  string
  Price uint
  Done bool
}

func main() {
  db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }
  fmt.Println("Connected to database successfully")

  // Migrate the schema
  db.AutoMigrate(&Todo{})
  fmt.Println("Database Migrated successfully")

  // Create
  db.Create(&Todo{Name: "A walk with the dog", Price: 100, Done: true})
  fmt.Println("Todo created successfully")

  // Read
  var todo Todo
  db.First(&todo, 1) // find product with integer primary key
  db.First(&todo, "done = ?", false) // find product with code D42

  fmt.Println("Todo read successfully" + todo.Name)
  // Update - update product's price to 200
  db.Model(&todo).Update("Done", false)
  // Update - update multiple fields
  db.Model(&todo).Updates(Todo{Price: 200, Done: true}) // non-zero fields
  // db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

  fmt.Println("Product updated successfully")
  // Delete - delete product
  // db.Delete(&todo, 1)
  // fmt.Println("Product deleted successfully")
}

