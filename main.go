package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := openDB("store.db")
	if err != nil {
		log.Fatalf("The database cannot be opened due to the following error: %v", err)
		return
	}

	args := strings.Split(os.Args[1:][0], ",")
	fmt.Printf("=+=+=+=\nСтраница сборки заказов %v", os.Args[1:])

	orders := make([]int, 0, len(args))
	for i, w := range args {
		num, err := strconv.Atoi(w)
		if err != nil {
			log.Printf("Order %v entered incorrectly\n", i+1)
			continue
		}
		orders = append(orders, num)
	}

	box, err := db.receivingOrders(orders)
	if err != nil {
		log.Println()
	}

	for key, value := range box {
		fmt.Printf("\n\n===Стеллаж %v", key)
		for _, w := range value {
			fmt.Printf("\n\n%v (id=%v)", w.product.productName, w.product.productID)
			fmt.Printf("\nзаказ %v, %v шт", w.orderProduct.orderID, w.orderProduct.quantity)
			if len(w.additionalBox) != 0 {
				fmt.Printf("\nдоп стеллаж: %v", w.additionalBox)
			}
		}
	}
	fmt.Println()
}
