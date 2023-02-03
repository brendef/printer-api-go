package main

import (
	"github.com/gin-gonic/gin"

  api "printers-api/apiFunctions"
)

func main() {

  router := gin.Default()

  router.GET("/printers", api.GetPrinters)
  router.GET("/printer/:ip", api.GetPrinterByIp)
  router.POST("/add_printer", api.AddPrinter)
  router.GET("/change_status/:ip", api.ChangePrinterStatus)
  router.GET("/update_name/:ip", api.ChangePrinterName)
  router.DELETE("/remove_printer/:ip", api.RemovePrinterByIp)

  router.Run("localhost:8080")
}
