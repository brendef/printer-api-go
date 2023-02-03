package databaseController

import (
	"database/sql"
	"fmt"
	"log"
	"os"
  "github.com/gin-gonic/gin"


	"github.com/joho/godotenv"
  _"github.com/go-sql-driver/mysql"
)

// type databaseDetails struct {
// 	url      string
// 	user     string
// 	password string
// 	name     string
// }

type Printer struct {
	Name       string `json:"name"`
	Ip_address string `json:"ip_address"`
	Status     int    `json:"status"`
}

func openConnection() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseName := os.Getenv("AWS_DATABASE_NAME")
	databaseUser := os.Getenv("AWS_DATABASE_USER")
	databasePassword := "W0ch4_etppe#3"
	databaseUrl := os.Getenv("AWS_DATABASE_URL")
	databasePort := os.Getenv("AWS_DATABASE_PORT")

	fmt.Println(databaseName)
	fmt.Println(databaseUser)
	fmt.Println(databasePassword)
	fmt.Println(databaseUrl)
	fmt.Println(databasePort)

  fmt.Println(sql.Drivers())

	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", databaseUser, databasePassword, databaseUrl, databasePort, databaseName)

  fmt.Println(connectionString)

	db, error := sql.Open("mysql", connectionString)

	if error != nil {
		log.Fatal(error)
	}

	// defer db.Close()

	return db
}

func GetPrinters() []Printer {

	database := openConnection()

	storedPrinters, error := database.Query("SELECT * FROM Printers")

	if error != nil {
		log.Fatal(error)
	}

	printers := []Printer{}

	for storedPrinters.Next() {
		var printer Printer
		error = storedPrinters.Scan(&printer.Name, &printer.Ip_address, &printer.Status)
		if error != nil {
			panic(error.Error())
		}
		printers = append(printers, printer)
	}

  defer storedPrinters.Close()
  defer database.Close()

	return printers
}

func GetPrinterByIp(ip string) Printer {
  database := openConnection()
	
  var printer Printer
  storedPrinter := database.QueryRow("SELECT * from Printers WHERE ip_address = ?", ip)
  error := storedPrinter.Scan(&printer.Name, &printer.Ip_address, &printer.Status)

  if error != nil {
   log.Fatal(error)
  }

  defer database.Close()

  return printer

}

func AddPrinter(printer Printer) gin.H {
  database := openConnection()

  result, error := database.Prepare("INSERT INTO Printers VALUES (?, ?, ?)")

  if error != nil {
    log.Fatal(error)
  }

  result.Exec(printer.Name, printer.Ip_address, printer.Status)

  defer database.Close()

  return gin.H{"message": "Printer added succesfully"}

}

func ChangePrinterStatus(printer Printer) gin.H {
  database := openConnection()

  result, error := database.Prepare("UPDATE Printers SET status=? WHERE ip_address = ?")

  if error != nil {
    log.Fatal(error)
  }

  if printer.Status == 1 {
    printer.Status = 0
  } else {
    printer.Status = 1
  }

  result.Exec(printer.Status, printer.Ip_address)

  defer database.Close()

  return gin.H{"message": printer.Status}
}

func ChangePrinterName(printer Printer, name string) gin.H {

  database := openConnection()

  result, error := database.Prepare("UPDATE Printers SET name=? WHERE ip_address = ?")

  if error != nil {
    log.Fatal(error)
  }

  result.Exec(name, printer.Ip_address)

  defer database.Close()

  return gin.H{"message": "Printers name changed succesfully"}
}

func RemovePrinter(printer Printer) gin.H {
  database := openConnection()

  deletePrinter, error := database.Prepare("DELETE FROM Printers WHERE ip_address=?")

  if error != nil {
    log.Fatal(error)
  }

  deletePrinter.Exec(printer.Ip_address)

  return gin.H{"message":"Printer removed succesfully"}
  
}
