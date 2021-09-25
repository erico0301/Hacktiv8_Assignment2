package helper

import (
	"encoding/json"
	"fmt"

	"io/ioutil"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"Assignment2/config"
	"Assignment2/structs"
)

// func QrySelectbyId () (structs.Order, error){
// 	db, err := config.Connect()
// 		if err != nil {
// 			fmt.Println(err.Error())
// 			return structs.Order{}, err
// 		}
// 		defer db.Close()

// 	var order structs.Order

// }

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Connect DB
		db, err := config.Connect()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer db.Close()

		// Exec Query read last order_id
		rows, err := db.Query("select order_id from orders order by order_id desc limit 1")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer rows.Close()
		//Baca Query
		var result string
		for rows.Next() {
			var err = rows.Scan(&result)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
		if err = rows.Err(); err != nil {
			fmt.Println(err.Error())
			return
		}
		// id = last orderId
		orderid, _ := strconv.Atoi(result)
		fmt.Println("Last order_id : ", orderid)

		// Exec Query Read last item_id
		rows, err = db.Query("select item_id from items order by item_id desc limit 1")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer rows.Close()
		//Baca Query
		var result1 string
		for rows.Next() {
			var err = rows.Scan(&result1)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
		if err = rows.Err(); err != nil {
			fmt.Println(err.Error())
			return
		}
		itemid, _ := strconv.Atoi(result1)
		fmt.Println("Last item_id : ", itemid)

		//Baca Req Body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		var orders structs.CrtOrder
		err = json.Unmarshal(body, &orders)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		//fmt.Println(len(orders.Items))
		itemLen := len(orders.Items)

		_, err = db.Exec("insert into orders value (?, ?, ?)", strconv.Itoa(orderid+1), orders.CustomerName, orders.OrderedAt)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		addItemId := itemid + 1
		fmt.Println("Add item_id : ", addItemId)
		for i := 0; i < itemLen; i++ {
			_, err = db.Exec("insert into items value(?, ?, ?, ?, ?)", strconv.Itoa(addItemId), orders.Items[i].ItemCode, orders.Items[i].Description, orders.Items[i].Quantity, strconv.Itoa(orderid+1))
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			addItemId += 1
		}
		fmt.Println("Create Order Succesfull...")

	}
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := config.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	rows, err := db.Query("select * from orders")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()
	var result []structs.Order
	for rows.Next() {
		var each = structs.Order{}
		var err = rows.Scan(&each.OrderId, &each.CustomerName, &each.OrderedAt)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		rowItems, err := db.Query("select * from items where order_id = ?", each.OrderId)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer rowItems.Close()
		var resultItems []structs.Items
		for rowItems.Next() {
			var temp = structs.Items{}
			var err = rowItems.Scan(&temp.ItemId, &temp.ItemCode, &temp.Description, &temp.Quantity, &temp.OrderId)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			resultItems = append(resultItems, temp)
		}
		each.Items = resultItems
		result = append(result, each)
	}

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}
	res, _ := json.Marshal(result)
	w.Write(res)
	return

}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {

	// Database Connection
	w.Header().Set("Content-Type", "application/json")
	db, err := config.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	//Read HTTP Parameter
	params := mux.Vars(r)
	fmt.Println(params["id"])

	//Delete on Table Items
	_, err = db.Exec("delete from items where order_id = ?", params["id"])
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = db.Exec("delete from orders where order_id = ?", params["id"])
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Delete Success...")
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	// Connect DB
	w.Header().Set("Content-Type", "application/json")
	db, err := config.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	//req Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var order structs.Order
	err = json.Unmarshal(body, &order)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(order)

	_, err = db.Exec("update orders set customer_name = ?, ordered_at = ? where order_id = ?", order.CustomerName, order.OrderedAt, order.OrderId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i := 0; i < len(order.Items); i++ {
		_, err = db.Exec("update items set item_code = ?, description = ?, quantity = ? where item_id = ?", order.Items[i].ItemCode, order.Items[i].Description, order.Items[i].Quantity, order.Items[i].ItemId)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	fmt.Println("Update Success...")

}
