package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type Item struct {
	ShortDescription string `json:"shortDescription" binding:"required"`
	Price            string `json:"price" binding:"required"`
}

type Receipt struct {
	ID           string `json:"id" binding:"omitnil"`
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Total        string `json:"total"`
	Items        []Item `json:"items"`
}

// One point for every alphanumeric character in the retailer name.
func (r Receipt) checkCharacters() int {
	result := 0

	for _, char := range r.Retailer {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			result += 1
		}
	}

	return result
}

// 50 points if the total is a round dollar amount with no cents.
func (r Receipt) checkTotal() int {
	parts := strings.Split(r.Total, ".")

	if len(parts) == 1 {
		return 50
	}

	if parts[1] == "00" || parts[1] == "0" {
		return 50
	}

	return 0
}

// 25 points if the total is a multiple of 0.25.
func (r Receipt) checkTotalMultiple() int {
	total, err := strconv.ParseFloat(r.Total, 64)

	if err != nil {
		return 0
	}

	if int(total*100)%25 == 0 {
		return 25
	}

	return 0
}

// 5 points for every two items on the receipt.
func (r Receipt) checkItemPairs() int {
	if len(r.Items) > 1 {
		result := (len(r.Items) / 2) * 5
		return result
	}

	return 0
}

// If the trimmed length of the item description is a multiple of 3, multiply
// the price by 0.2 and round up to the nearest integer. The result is the
// number of points earned.
func (r Receipt) checkItemDescriptions() int {
	result := 0

	for _, item := range r.Items {
		trimmedDesc := strings.TrimSpace(item.ShortDescription)

		if len(trimmedDesc)%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)

			fmt.Printf("Original price: %v \n", item.Price)
			fmt.Printf("Parsed price: %v \n", price)

			if err == nil {
				temp := price * 0.2
				fmt.Println("Point calculation: ", temp)
				v := int(math.Ceil(price * 0.2))
				fmt.Printf("%v Points earned with description: %v \n", v, item.ShortDescription)
				result += v
			}
		}
	}

	return result
}

// 6 points if the day in the purchase date is odd
func (r Receipt) checkPurchaseDate() int {
	purchaseDate, _ := time.Parse("2006-01-02", r.PurchaseDate)

	if purchaseDate.Day()%2 != 0 {
		return 6
	}

	return 0
}

// 10 points if the time of purchase is after 2:00pm and before 4:00pm
func (r Receipt) checkTime() int {
	timeLayout := "15:04"
	receiptTime, _ := time.Parse(timeLayout, r.PurchaseTime)
	startTime, _ := time.Parse(timeLayout, "14:00")
	endTime, _ := time.Parse(timeLayout, "16:00")

	if receiptTime.After(startTime) && receiptTime.Before(endTime) {
		return 10
	}

	return 0
}
