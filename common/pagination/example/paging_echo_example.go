package example

import (
	"github.com/bonoken/watcher-common/common/pagination"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

var (
	Orm *gorm.DB
)

// echo pagination example
func ListUser() echo.HandlerFunc {
	return func(c echo.Context) error {

		resultMap := map[string]string{}
		type response struct {
			List  []User `json:"list"`
			Count int    `json:"count"`
		}

		var users []User
		var count int

		// paging
		param := pagination.GetPagingSearchFilter(c)
		paginator := pagination.Paging(&pagination.Param{
			DB:           Orm.Where("is_use=1"),
			Page:         param.Page,
			Limit:        param.Limit,
			ShowSQL:      true,
			Keyword:      param.Keyword,
			FilterColumn: param.FilterColumn,
			FilterOrder:  param.FilterOrder,
		}, &users)
		count = paginator.TotalRecord

		if len(users) < 1 {
			// 데이터 없음
			resultMap["msg"] = "no data"
			return c.JSON(http.StatusOK, resultMap)
		}

		return c.JSON(http.StatusOK, response{List: users, Count: count})

	}
}

type User struct {
	UserId     string    `gorm:"primary_key;column:user_id" json:"userId"`
	UserName   string    `gorm:"column:user_name" json:"userName"`
	UserEmail  string    `gorm:"column:user_email" json:"userEmail"`
	Department string    `gorm:"column:department" json:"department"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
	IsActive   int8      `gorm:"column:is_active" json:"isActive"`
	IsUse      int8      `gorm:"column:is_use" json:"isUse"`
}

func (User) TableName() string {
	return "user"
}
