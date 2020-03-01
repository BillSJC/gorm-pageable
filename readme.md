# Gorm-Pageable

A page query management of GORM

## Installation

We Recommend to use package manager `vgo` to manage this package

```go
import pageable "github.com/BillSJC/gorm-pageable"
```

## Usage

Just prepare a struct and FOLLOW the code below

```go
package main

import (
    "fmt"
    "github.com/BillSJC/gorm-pageable"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB //your gorm DB connection

// the struct you want to search
type User struct{
    gorm.Model
    Active bool
    UserName string
    Age uint
}

// a function to get ddata
func getResultSet (page int,rowsPerPage int)(*pageable.Response,error){
    //your empty result set
    resultSet := make([]*User,0,30)
    //prepare a handler to query
    handler := DB.
        Module(&User{}).
        Where(&User{Active:true})
    //use PageQuery to get data
    resp,err := pageable.PageQuery(page,rowsPerPage,handler,&resultSet)
    // handle error
    if err != nil {
        panic(err)
    }
    //Here are the response
	fmt.Println(resp.PageNow)    //PageNow: current page of query
	fmt.Println(resp.PageCount)  //PageCount: total page of the query
	fmt.Println(resp.RawCount)   //RawCount: total raw of query
	fmt.Println(resp.RawPerPage) //RawPerPage: rpp
	fmt.Println(resp.ResultSet)  //ResultSet: result data
	fmt.Println(resp.FirstPage)  //FirstPage: if the result is the first page
	fmt.Println(resp.LastPage)   //LastPage: if the result is the last page
	fmt.Println(resp.Empty)  //Empty: if the result is empty
}
```