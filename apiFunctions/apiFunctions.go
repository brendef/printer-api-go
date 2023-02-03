package apifunctions

import (
	// "errors"
	"net/http"

	database "printers-api/databaseController"

	"github.com/gin-gonic/gin"
)

func GetPrinters(context *gin.Context) {

  printers := database.GetPrinters()
  
  context.IndentedJSON(http.StatusOK, printers)
}

func GetPrinterByIp(context *gin.Context) {

  ip := context.Param("ip")
  printer := database.GetPrinterByIp(ip)

  context.IndentedJSON(http.StatusOK, printer)
}

func AddPrinter(context *gin.Context) {
  var newPrinter database.Printer

  context.BindJSON(&newPrinter)

  result := database.AddPrinter(newPrinter)
  context.IndentedJSON(http.StatusCreated, result)
}

func ChangePrinterStatus(context *gin.Context) {

  ip := context.Param("ip")
  printer := database.GetPrinterByIp(ip)

  result := database.ChangePrinterStatus(printer)
  context.IndentedJSON(http.StatusCreated, result)

}

func ChangePrinterName(context *gin.Context) {
  ip := context.Param("ip")
  name, success := context.GetQuery("name")

  if !success {
    context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing Name Param"})
  }

  printer := database.GetPrinterByIp(ip)

  result := database.ChangePrinterName(printer, name)
  context.IndentedJSON(http.StatusCreated, result)
}

func RemovePrinterByIp(context *gin.Context) {

  ip := context.Param("ip")
  printer := database.GetPrinterByIp(ip)

  result := database.RemovePrinter(printer)
  context.IndentedJSON(http.StatusCreated, result)

}
