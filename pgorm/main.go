// filepath: /home/dio/programacao/goprogram/pgorm/main.go
package main


import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)
import (
	// "fmt"
	// "github.com/diogenesdornelles/pgorm/internal/database"
	// comp "github.com/diogenesdornelles/pgorm/internal/domain/company"
	// user "github.com/diogenesdornelles/pgorm/internal/domain/user"
	// "github.com/diogenesdornelles/pgorm/internal/platform"
	// compRepo "github.com/diogenesdornelles/pgorm/internal/repository/company"
	// userRepo "github.com/diogenesdornelles/pgorm/internal/repository/user"
	// "log"
)

func main() {
	// config := platform.LoadConfig()
	// db, err := platform.ConnectDatabase(config)
	// if err != nil {
	// 	log.Fatal("Failed to connect to database:", err)
	// }

	// if !db.Migrator().HasTable(&comp.Company{}) {
	// 	db.Migrator().CreateTable(&comp.Company{})
	// } else {
	// 	log.Println("Table 'companies' already exists, skipping migration.")
	// }

	// if !db.Migrator().HasTable(&user.User{}) {
	// 	// Create table for `User`
	// 	db.Migrator().CreateTable(&user.User{})
	// } else {
	// 	log.Println("Table 'users' already exists, skipping migration.")
	// }

	// // Injeção no repository
	// userRepo := userRepo.NewRepository(db)
	// companyRepo := compRepo.NewRepository(db)

	// err = userRepo.CreateUser(database.User)
	// if err != nil {
	// 	log.Println("Error creating user:", err)
	// }

	// err = companyRepo.CreateCompany(database.Company)
	// if err != nil {
	// 	log.Println("Error creating company:", err)
	// }

	// fmt.Println("Starting app.")
}
