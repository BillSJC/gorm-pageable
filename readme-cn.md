# Gorm-Pageable

[![Go Report Card](https://goreportcard.com/badge/github.com/BillSJC/gorm-pageable)](https://goreportcard.com/report/github.com/BillSJC/gorm-pageable)
[![Build Status](https://travis-ci.org/BillSJC/gorm-pageable.svg?branch=master)](https://travis-ci.org/BillSJC/gorm-pageable)
[![codecov](https://codecov.io/gh/BillSJC/gorm-pageable/branch/master/graph/badge.svg)](https://codecov.io/gh/BillSJC/gorm-pageable)

一个快捷的gorm翻页查询器工具

## 使用

强烈推荐使用官方包管理 `vgo` 安装

```go
import pageable "github.com/BillSJC/gorm-pageable"
```

## 使用方法

在需要进行翻页流程的地方按照如下方式接入即可

```go
package main

import (
    "fmt"
    pageable "github.com/BillSJC/gorm-pageable"
    "github.com/jinzhu/gorm"
)

var DB *gorm.DB //your gorm DB connection

// 表结构体
type User struct{
    gorm.Model
    Active bool
    UserName string
    Age uint
}

// 某个需要翻页的函数
func getResultSet (page int,rowsPerPage int)(*pageable.Response,error){
    //你的空的结果数组
    resultSet := make([]*User,0,30)
    //准备一个写好查询条件的gorm.DB，注意要执行过Module()
    handler := DB.
        Module(&User{}).
        Where(&User{Active:true})
    //进行查询
    resp,err := pageable.PageQuery(page,rowsPerPage,handler,&resultSet)
    // 处理报错
    if err != nil {
        panic(err)
    }
    // 获得结果
	fmt.Println(resp.PageNow)    //PageNow: 当前页
	fmt.Println(resp.PageCount)  //PageCount: 总页码数
	fmt.Println(resp.RawCount)   //RawCount: 总行数
	fmt.Println(resp.RawPerPage) //RawPerPage: 每页结果数量
	fmt.Println(resp.ResultSet)  //ResultSet: 返回的数组
	fmt.Println(resp.FirstPage)  //FirstPage: 是否是第一页
	fmt.Println(resp.LastPage)   //LastPage: 是否是最后一页
	fmt.Println(resp.Empty)  //Empty: 该页结果是否为空
}
```