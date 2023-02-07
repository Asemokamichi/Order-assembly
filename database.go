package main

import (
	"database/sql"
	"log"
	"strings"
)

type Store struct {
	DB *sql.DB
}

func openDB(dbName string) (*Store, error) {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Store{DB: db}, nil
}

func (db *Store) receivingOrders(orders []int) (map[string][]Order, error) {
	queryProduct := `
		SELECT product.product_id, product.product_name, order_product.quantity
		FROM order_product INNER JOIN product ON order_product.product_id=product.product_id
		WHERE order_product.order_id=$1;
	`

	queryBox := `
		SELECT box.box_name
		FROM box, product_box
		WHERE product_box.product_id= $1 AND product_box.box_id=box.box_id AND  product_box.is_main=1
	`

	queryadditionalBox := `
		SELECT box.box_name
		FROM box, product_box
		WHERE product_box.product_id= $1 AND product_box.box_id=box.box_id AND product_box.is_main=-1
	`

	box := map[string][]Order{}

	for _, w := range orders {
		rows, err := db.DB.Query(queryProduct, w)
		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close()

		for rows.Next() {
			var order Order
			order.orderID = w

			if err := rows.Scan(
				&order.Product.productID,
				&order.productName,
				&order.quantity,
			); err != nil {
				log.Fatal(err)
			}

			var orderBox string

			if err := db.DB.QueryRow(queryBox, order.Product.productID).Scan(&orderBox); err != nil {
				log.Fatal(err)
			}

			additionalBox, err := db.DB.Query(queryadditionalBox, order.Product.productID)
			if err != nil {
				if err == sql.ErrNoRows {
					box[orderBox] = append(box[orderBox], order)
					continue
				}
				log.Fatal(err)
			}

			defer additionalBox.Close()

			var str []string
			for additionalBox.Next() {
				var s string
				if err := additionalBox.Scan(&s); err != nil {
					log.Fatal(err)
				}
				str = append(str, s)
			}

			order.additionalBox = strings.Join(str, ",")

			box[orderBox] = append(box[orderBox], order)
		}

	}

	return box, nil
}
