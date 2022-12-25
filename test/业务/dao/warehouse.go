package dao

import "modulename/basic_homework/homework_for_redrock/test/业务/model"

func InsertWarehouse(w model.Warehouse) (err error) {
	_, err = DB.Exec("insert into messageboard.user(name, password) values (?,?)", w.WarehouseName, w.Item)
	return
} //可在选择架构中把homework架构更改为message架构

func SearchWarehouseByWarehouseName(name string) (w model.Warehouse, err error) {
	row := DB.QueryRow("select id,name,password from messageboard.user  where name=?")
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&w.WarehouseName)
	return
}
