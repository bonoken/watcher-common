package pagination

import (
	"github.com/jinzhu/gorm"
	"math"
)

// Param
type Param struct {
	DB    *gorm.DB
	Page  int
	Limit int
	//OrderBy      []string
	ShowSQL      bool
	Keyword      string
	FilterColumn []string
	FilterOrder  []string
}

// Paginator
type Paginator struct {
	TotalRecord int         `json:"total_record"`
	TotalPage   int         `json:"total_page"`
	Records     interface{} `json:"records"`
	Offset      int         `json:"offset"`
	Limit       int         `json:"limit"`
	Page        int         `json:"page"`
	PrevPage    int         `json:"prev_page"`
	NextPage    int         `json:"next_page"`
}

// Paging
func Paging(p *Param, result interface{}) *Paginator {
	db := p.DB

	if p.ShowSQL {
		db = db.Debug()
	}
	if p.Page < 1 {
		p.Page = 1
	}

	if len(p.FilterOrder) > 0 {
		for _, o := range p.FilterOrder {
			db = db.Order(o)
		}
	}

	if len(p.Keyword) > 0 && len(p.FilterColumn) > 0 {
		filterColumns := p.FilterColumn
		var line string
		for i := range filterColumns {
			if len(line) == 0 {
				line = "  " + filterColumns[i] + " LIKE " + "'%" + p.Keyword + "%' "
			}
			if len(line) > 0 {
				line = line + "  OR " + filterColumns[i] + " LIKE " + "'%" + p.Keyword + "%' "
			}

		}
		if len(line) > 0 {
			//fmt.Println(line)
			db = db.Where(line)
		}
	}

	done := make(chan bool, 1)
	var paginator Paginator
	var count int
	var offset int

	if p.Limit > 0 {
		go countRecords(db, result, done, &count)

		if p.Page == 1 {
			offset = 0
		} else {
			offset = (p.Page - 1) * p.Limit
		}

		db.Limit(p.Limit).Offset(offset).Find(result)
		<-done

		paginator.TotalRecord = count
		paginator.Records = result
		paginator.Page = p.Page

		paginator.Offset = offset
		paginator.Limit = p.Limit
		paginator.TotalPage = int(math.Ceil(float64(count) / float64(p.Limit)))

		if p.Page > 1 {
			paginator.PrevPage = p.Page - 1
		} else {
			paginator.PrevPage = p.Page
		}

		if p.Page == paginator.TotalPage {
			paginator.NextPage = p.Page
		} else {
			paginator.NextPage = p.Page + 1
		}

	} else {
		go countRecords(db, result, done, &count)

		db.Find(result)
		<-done

		paginator.TotalRecord = count
		paginator.Records = result
		paginator.Page = 1

		paginator.Offset = 0
		paginator.Limit = count
		paginator.TotalPage = 1
	}

	return &paginator
}

func countRecords(db *gorm.DB, anyType interface{}, done chan bool, count *int) {
	db.Model(anyType).Count(count)
	done <- true
}
