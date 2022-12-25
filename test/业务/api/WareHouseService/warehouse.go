package WareHouseService

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"modulename/basic_homework/homework_for_redrock/test/业务/model"
	"modulename/basic_homework/homework_for_redrock/test/业务/service"
	"modulename/basic_homework/homework_for_redrock/test/业务/util"
)

func AddWarehouse(c *gin.Context) {
	warehouseName := c.PostForm("name")
	if warehouseName == "" {
		util.RespParamErr(c)
		return
	}

	w, err := service.SearchWarehouseByWarehouseName(warehouseName)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("search warehouse error:%v", err)
		util.RespInternalErr(c)
		return
	}

	if w.WarehouseName != "" {
		util.NomalErr(c, 300, "仓库已存在")
		return
	}
	err = service.CreateWarehouse(model.Warehouse{
		WarehouseName: warehouseName,
	})
}

func UploadItem() {

}

func DownloadItem() {

}

func ChangeItem() {

}
