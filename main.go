package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var server *gin.Engine
var receipts []Receipt

func main() {
	initServer()
}

func initServer() {
	server = gin.Default()

	// setup routes
	server.GET("/receipts/:id/points", GetPointsHandler)
	server.POST("/receipts/process", processReceiptHandler)

	server.Run()
}

func processReceiptHandler(context *gin.Context) {
	var receipt Receipt
	err := context.BindJSON(&receipt)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"description": "The receipt is invalid."})
		return
	}

	receipt.ID = uuid.New().String()
	receipts = append(receipts, receipt)

	context.JSON(http.StatusCreated, gin.H{"id": receipt.ID})
}

func GetPointsHandler(context *gin.Context) {
	id := context.Param("id")
	receipt, err := findReceipt(id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"description": "Receipt not found."})
		return
	}

	points := calculatePoints(receipt)
	context.JSON(http.StatusOK, gin.H{"points": points})
}

func findReceipt(id string) (Receipt, error) {
	for _, v := range receipts {
		if v.ID == id {
			return v, nil
		}
	}

	return Receipt{}, errors.New("Receipt not found")
}

func calculatePoints(r Receipt) int {
	points := 0

	points += r.checkTotal()
	points += r.checkTotalMultiple()
	points += r.checkCharacters()
	points += r.checkItemPairs()
	points += r.checkItemDescriptions()
	points += r.checkPurchaseDate()
	points += r.checkTime()

	return points
}
