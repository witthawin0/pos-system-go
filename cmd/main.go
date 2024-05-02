package main

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/witthawin0/pos-system-go/internal/delivery"
	"github.com/witthawin0/pos-system-go/internal/http"
	"github.com/witthawin0/pos-system-go/internal/repository"
	"github.com/witthawin0/pos-system-go/internal/usecase"
)

func main() {
	// Connect to the database
	db, err := sql.Open("mysql", "witthawin:witthawin123@tcp(127.0.0.1:3306)/inventory?parseTime=true")
	if err != nil {
		panic(err)
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	defer db.Close()

	emplyeeRepo := repository.NewEmployeeReposistoryImpl(db)

	emplyeeUc := usecase.NewEmployeeUseCaseImpl(emplyeeRepo)

	employeeHandler := delivery.NewEmployeeHandler(emplyeeUc)

	// Create new instance of Echo Server
	e := http.NewEchoServer()

	e.POST("/api/employee", employeeHandler.AddEmployee)
	e.GET("/api/employee", employeeHandler.ListEmployees)
	e.GET("/api/employee/:id", employeeHandler.GetEmployeeByID)
	e.PUT("/api/employee/:id", employeeHandler.UpdateEmployee)
	e.DELETE("/api/employee/:id", employeeHandler.RemoveEmployee)

	e.Logger.Fatal(e.Start(":3333"))
}
