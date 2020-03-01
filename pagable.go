package pageable

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"runtime/debug"
	"time"
)

// Response: Base response of query
type Response struct {
	PageNow    uint        //PageNow: current page of query
	PageCount  uint        //PageCount: total page of the query
	RawPerPage uint        //RawPerPage: rpp
	ResultSet  interface{} //ResultSet: result data
	FirstPage  bool        //FirstPage: if the result is the first page
	LastPage   bool        //LastPage: if the result is the last page
	Empty      bool        //Empty: if the result is empty
}

// getLimitOffset (private) get LIMIT and OFFSET keyword in SQL
func getLimitOffset(page, rpp uint) (limit, offset uint) {
	return rpp, page * rpp
}

// recoveryHandler : default type of recovery handler
type recoveryHandler func()

// recovery : handler of panic
var recovery recoveryHandler

// setRecovery Set custom recovery
func setRecovery(handler func()) {
	recovery = handler
}

// defaultRecovery : print base recover info
func defaultRecovery() {
	if err := recover(); err != nil {
		//print panic info
		fmt.Printf("Panic recovered: %s \n\n Time: %s \n\n Stack Trace: \n\n",
			fmt.Sprint(err),
			time.Now().Format("2006-01-02 15:04:05"),
		)
		//stack
		debug.PrintStack()
	}
}

// init: use default recovery
func init() {
	//use default recovery
	recovery = defaultRecovery
}

// PageQuery:  main handler of query
// page: 1 for the first page
// resultPtr : MUST input a Slice or it will be a error
// queryHandler : MUST have DB.Module  or it will be a error
func PageQuery(page uint, rawPerPage uint, queryHandler *gorm.DB, resultPtr interface{}) (*Response, error) {
	//recovery
	defer recovery()
	var count uint
	count = 0
	// use page - 1 in query
	limit, offset := getLimitOffset(page-1, rawPerPage)
	queryHandler.Count(&count)
	queryHandler.Limit(limit).Offset(offset).Find(resultPtr)
	if err := queryHandler.Error; err != nil {
		return nil, err
	}
	PageCount := count / rawPerPage
	if rawPerPage%count != 0 {
		PageCount++
	}
	return &Response{
		PageNow:    page,
		PageCount:  PageCount,
		RawPerPage: rawPerPage,
		ResultSet:  resultPtr,
		FirstPage:  page == 1,
		LastPage:   page == PageCount,
		Empty:      (page > PageCount) || count == 0,
	}, nil
}
