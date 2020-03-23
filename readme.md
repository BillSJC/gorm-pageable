# Gorm-Pageable

[![Go Report Card](https://goreportcard.com/badge/github.com/BillSJC/gorm-pageable)](https://goreportcard.com/report/github.com/BillSJC/gorm-pageable)
[![Build Status](https://travis-ci.org/BillSJC/gorm-pageable.svg?branch=master)](https://travis-ci.org/BillSJC/gorm-pageable)
![Go](https://github.com/BillSJC/gorm-pageable/workflows/Go/badge.svg)
[![GoDoc](https://godoc.org/github.com/BillSJC/gorm-pageable?status.svg)](https://godoc.org/github.com/BillSJC/gorm-pageable)
[![codecov](https://codecov.io/gh/BillSJC/gorm-pageable/branch/master/graph/badge.svg)](https://codecov.io/gh/BillSJC/gorm-pageable)

A page query management of GORM 

[->中文文档](readme-cn.md)

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
    pageable "github.com/BillSJC/gorm-pageable"
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
    // if need, you can turn to next/last page
    resp,err = resp.GetNextPage() //to next page
    resp,err = resp.GetLastPage() //to last page
    resp,err = resp.GetFirstPage() //to first page
    resp,err = resp.GetEndPage() //to end page
    // you can use both resp.ResultSet or the resultSet you input to access the result
    // handle error
    if err != nil {
        //print the err or do sth else
        fmt.Println("something happened...")
        fmt.Println(err)
        return nil,err
    }
    //Here are the response
	fmt.Println(resp.PageNow)    //PageNow: current page of query
	fmt.Println(resp.PageCount)  //PageCount: total page of the query
	fmt.Println(resp.RawCount)   //RawCount: total raw of query
	fmt.Println(resp.RawPerPage) //RawPerPage: rpp
	fmt.Println(resp.ResultSet)  //ResultSet: result data
	fmt.Println(resultSet)          //the same as resp.ResultSet and have the raw type
	fmt.Println(resp.FirstPage)  //FirstPage: if the result is the first page
	fmt.Println(resp.LastPage)   //LastPage: if the result is the last page
	fmt.Println(resp.Empty)  //Empty: if the result is empty
	fmt.Println(resp.StartRow)  //StartRow: the first row of the result set, 0 when result set is empty
	fmt.Println(resp.EndRow)  //EndRow: the last row of the result set, 0 when result set is empty
}
```

`Notice:` when access to the result, you can use both `resp.ResultSet` in Response or the param `resultSet` you input into the function, both of then have same pointer and same data, but the type of `resp.ResultSet` is `interface{}` and you mat need to convert to the raw type if you need to do any operation of the result set

## Guidance

### Use `0` as the first page

the default first page is `1`. However,if u want to use `0` as the first page, just follow this step:

```go
    pageable.Use0AsFirstPage()
```

### Set Default Result Per Page(rpp)

Sometimes you just want to use same `rpp` in every query, then u just need do this:

```go
    pageable.SetDefaultRPP(25) //set 25 rows per page in every query
```

And next time, you can use rpp=0 to use Default rpp

```go
    pageable.PageQuery(page:1, rpp:0, queryHandler: ..., resultPtr: ...)
```

### Use custom recovery

The default recovery will only print stack trace. If you want to use your custom Recovery handler, just follow the step:

```go
package main
import (
    "fmt"
    pageable "github.com/BillSJC/gorm-pageable"
)

//your recovery
func myRecovery(){
    if err := recover ; err != nil {
        fmt.Println("something happend")
        fmt.Println(err)
        //then you can do some logs...
    } 
}

func init(){
    //setup your recovery
    pageable.SetRecovery(myRecovery)
}
```