package register

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	
)
//Query input to db
type Query struct{
	Db *sql.DB
}
// Member table
type Member struct {
	Email     string `json:"email"`
	UserID    string `json:"user_id"`
	UserName  string `json:"username"`
	Address   string `json:"address"`
	
}
//Register email
func (q Query) Register(c *gin.Context) {
	sqlStr := "INSERT INTO Member VALUES($1, $2, $3, $4)"
    var member Member
	err := c.BindJSON(&member)
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusBadRequest, err.Error())
		return 
	}
	fmt.Printf("%+v",member)
	_, err = q.Db.Exec(sqlStr, member.Email,member.UserID,member.UserName,member.Address)
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status" :"OK",
	})
}
//ListMember select query
func (q Query) ListMember(c *gin.Context) {
	sqlStr := `SELECT "UserName" FROM member`
	rows, err := q.Db.Query(sqlStr)
	if err != nil {
		fmt.Println("FAIL")
		c.String(http.StatusBadRequest, err.Error())
		return
		}
    defer rows.Close()
	for rows.Next() {
		var username string
		if err := rows.Scan(&username); err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"username": username,
		})
	
    }
}

