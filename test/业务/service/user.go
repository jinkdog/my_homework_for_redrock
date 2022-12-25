package service

import (
	"modulename/basic_homework/homework_for_redrock/test/业务/dao"
	"modulename/basic_homework/homework_for_redrock/test/业务/model"
)

func SearchUserByUserName(name string) (u model.User, err error) {
	u, err = dao.SearchUserByUserName(name)
	return
}

func CreateUser(u model.User) error {
	err := dao.InsertUser(u)
	return err
}

func SearchWarehouseByWarehouseName(name string) (w model.Warehouse, err error) {
	w, err = dao.SearchWarehouseByWarehouseName(name)
	return
}

func CreateWarehouse(w model.Warehouse) error {
	err := dao.InsertWarehouse(w)
	return err
}
