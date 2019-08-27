package pagination

import (
	"encoding/json"
	"github.com/bonoken/watcher-common/common/strcase"
	"github.com/labstack/echo/v4"
	"strconv"
	"strings"
)

type PagingFilter struct {
	FilterColumn, FilterOrder []string
	Keyword                   string
	Page, Limit               int
}

func GetPagingSearchFilter(c echo.Context) (p PagingFilter) {

	if len(c.QueryParam("filterOrder")) > 0 {
		var filterOrders []string

		if err := json.Unmarshal([]byte(c.QueryParam("filterOrder")), &filterOrders); err != nil {
			//fmt.Println("error")
		}

		for i := range filterOrders {
			splitFn := func(c rune) bool {
				return c == ' '
			}
			s := strings.FieldsFunc(filterOrders[i], splitFn)

			if len(s) > 0 {
				column := strcase.ToSnake(s[0])
				order := "ASC"
				if len(s) > 1 {
					str := strings.ToLower(s[1])
					if strings.Contains(str, "d") {
						order = "DESC"
					}
				}
				p.FilterOrder = append(p.FilterOrder, column+" "+order)
				//fmt.Println(p.FilterOrder, column+" "+order)
			}
		}

	}

	if len(c.QueryParam("filterColumn")) > 0 {
		var filterColumns []string
		if err := json.Unmarshal([]byte(c.QueryParam("filterColumn")), &filterColumns); err != nil {
			//fmt.Println("error")
		}

		for i := range filterColumns {
			p.FilterColumn = append(p.FilterColumn, strcase.ToSnake(filterColumns[i]))
			//fmt.Println(strcase.ToSnake(filterColumns[i]))
		}
	}

	p.Keyword = c.QueryParam("keyword")

	if pg, err := strconv.Atoi(isNull(c.QueryParam("page"), "1")); err == nil {
		p.Page = pg
	}

	if l, err := strconv.Atoi(isNull(c.QueryParam("limit"), "0")); err == nil {
		p.Limit = l
	}
	return
}

func isNull(value, defaultValue string) string {
	if len(value) > 0 {
		return value
	}

	return defaultValue
}
