package api

import (
	"github.com/gin-gonic/gin"
	"modulename/basic_homework/homework_for_redrock/test/业务/api/UserService"
	"modulename/basic_homework/homework_for_redrock/test/业务/api/WareHouseService"
)

func InitRouter() {
	r := gin.Default()

	u := r.Group("/user")
	{
		u.POST("/register", UserService.Register) //此处不为Register（），若加括号则不表示中间件而表示立即执行
		u.POST("/login", UserService.Login)
		u.PUT("/password")
	}

	w := r.Group("/warehouse")
	{
		w.POST("/addWarehouse", WareHouseService.AddWarehouse)
		//w.POST("/upload", WareHouseService.UploadItem)
		//w.POST("/download", WareHouseService.DownloadItem)
		//w.POST("/download", WareHouseService.ChangeItem)

	}
	r.Run()
}
