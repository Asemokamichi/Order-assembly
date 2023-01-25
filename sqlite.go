package main

// import (
// 	"database/sql"
// 	"log"
// )

// // Еще в этой задаче надо использовать такую вот схему бд:
// // product [product_id, product_name]
// // order_product [order_id, product_id, quantity]
// // box [box_id, box_name]
// // product_box [product_id, box_id, is_main]
// func main() {
// 	db, err := sql.Open("sqlite3", "store.db")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if err := db.Ping(); err != nil {
// 		log.Fatal(err)
// 	}

// 	query := `
// 		CREATE TABLE product(
// 			product_id INTEGER PRIMARY KEY AUTOINCREMENT,
// 			product_name VARCHAR(32) NOT NULL
// 		);
// 		CREATE TABLE order_product(
// 			order_id INTEGER NOT NULL,
// 			product_id INTEGER NOT NULL,
// 			quantity INTEGER NOT NULL
// 		);
// 		CREATE TABLE box(
// 			box_id INTEGER PRIMARY KEY AUTOINCREMENT,
// 			box_name VARCHAR(32) NOT NULL
// 		);
// 		CREATE TABLE product_box(
// 			product_id INTEGER,
// 			box_id INTEGER NOT NULL,
// 			is_main BLOB NOT NULL
// 		);
// 	`
// 	if _, err := db.Exec(query); err != nil {
// 		log.Fatal(err)
// 	}

// 	product := []Product{
// 		{productName: "Ноутбук"},
// 		{productName: "Телевизор"},
// 		{productName: "Телефон"},
// 		{productName: "Системный блок"},
// 		{productName: "Часы"},
// 		{productName: "Микрофон"},
// 	}

// 	box := []Box{
// 		{boxName: "А"},
// 		{boxName: "Б"},
// 		{boxName: "В"},
// 		{boxName: "Ж"},
// 		{boxName: "З"},
// 	}

// 	productBox := []ProductBox{
// 		{productID: 1, boxID: 1, isMain: 1},
// 		{productID: 2, boxID: 1, isMain: 1},
// 		{productID: 3, boxID: 2, isMain: 1},
// 		{productID: 3, boxID: 3, isMain: -1},
// 		{productID: 3, boxID: 5, isMain: -1},
// 		{productID: 4, boxID: 4, isMain: 1},
// 		{productID: 5, boxID: 4, isMain: 1},
// 		{productID: 5, boxID: 1, isMain: -1},
// 		{productID: 6, boxID: 4, isMain: 1},
// 	}

// 	for _, w := range product {
// 		if _, err := db.Exec(`INSERT INTO product(product_name) VALUES($1)`, w.productName); err != nil {
// 			log.Fatal(err)
// 		}
// 	}

// 	for _, w := range box {
// 		if _, err := db.Exec(`INSERT INTO box(box_name) VALUES($1)`, w.boxName); err != nil {
// 			log.Fatal(err)
// 		}
// 	}

// 	for _, w := range productBox {
// 		if _, err := db.Exec(`INSERT INTO product_box(product_id, box_id, is_main) VALUES($1, $2, $3)`, w.productID, w.boxID, w.isMain); err != nil {
// 			log.Fatal(err)
// 		}
// 	}

// 	orderProduct := []OrderProduct{
// 		{orderID: 10, productID: 1, quantity: 2},
// 		{orderID: 11, productID: 2, quantity: 3},
// 		{orderID: 14, productID: 1, quantity: 3},
// 		{orderID: 10, productID: 3, quantity: 1},
// 		{orderID: 14, productID: 4, quantity: 4},
// 		{orderID: 15, productID: 5, quantity: 1},
// 		{orderID: 10, productID: 6, quantity: 1},
// 	}
// 	for _, w := range orderProduct {
// 		if _, err := db.Exec(`INSERT INTO order_product(order_id, product_id, quantity) VALUES($1, $2, $3)`, w.orderID, w.productID, w.quantity); err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// }
