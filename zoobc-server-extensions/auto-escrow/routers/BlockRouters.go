package routers

var (
	items   []Item
	ItemObj Item
)

// func CreateItem(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	err := json.NewDecoder(r.Body).Decode(&ItemObj)

// 	if err != nil {
// 		fmt.Println("Error = ", err)
// 		return
// 	}

// 	fmt.Println("itemObj Raw =", ItemObj.Company)

// 	db, err := sql.Open("sqlite3", "./db/sparepartshop.db")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer db.Close()
// 	feed := model.NewFeed(db)

// 	item, err := feed.CreateItem(ItemObj.Nama, ItemObj.Code, ItemObj.Jenis, ItemObj.Company, ItemObj.BuyPrice, ItemObj.SellPrice1, ItemObj.SellPrice2)
// 	if err != nil {
// 		fmt.Println("CreateItem = ", err)
// 	}
// 	json.NewEncoder(w).Encode(&item)
// }
