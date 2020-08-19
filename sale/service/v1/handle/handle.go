package handle

import (
	"database/sql"
	
	"github.com/chayupon/sale/service/v1/register"
	"github.com/gin-gonic/gin"
)
// Router app
func Router(db *sql.DB) {
	regist :=register.Query{
		Db : db,
	}
	r := gin.Default()
	r.POST("/register/email", regist.Register)
	r.GET("/member", regist.ListMember)
	r.Run()
}
