# Gorm-Pageable

[![Go Report Card](https://goreportcard.com/badge/github.com/BillSJC/gorm-pageable)](https://goreportcard.com/report/github.com/BillSJC/gorm-pageable)
[![Build Status](https://travis-ci.org/BillSJC/gorm-pageable.svg?branch=master)](https://travis-ci.org/BillSJC/gorm-pageable)
![Go](https://github.com/BillSJC/gorm-pageable/workflows/Go/badge.svg)
[![GoDoc](https://godoc.org/github.com/BillSJC/gorm-pageable?status.svg)](https://godoc.org/github.com/BillSJC/gorm-pageable)
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

    // 如果需要的话你也可以直接翻页
    resp,err = resp.GetNextPage() //下一页
    resp,err = resp.GetLastPage() //上一页
    resp,err = resp.GetFirstPage() //回首页
    resp,err = resp.GetEndPage() //去最后一页
    // 处理报错
    if err != nil {
        panic(err)
    }
    // 获得结果
    // 注意：所有参数都应当是只读的，如果修改值可能导致关联的上下页操作出现异常
	fmt.Println(resp.PageNow)    //PageNow: 当前页
	fmt.Println(resp.PageCount)  //PageCount: 总页码数
	fmt.Println(resp.RawCount)   //RawCount: 总行数
	fmt.Println(resp.RawPerPage) //RawPerPage: 每页结果数量
	fmt.Println(resp.ResultSet)  //ResultSet: 返回的数组
	fmt.Println(resp.FirstPage)  //FirstPage: 是否是第一页
	fmt.Println(resp.LastPage)   //LastPage: 是否是最后一页
	fmt.Println(resp.Empty)  //Empty: 该页结果是否为空
	fmt.Println(resp.StartRow)  //StartRow: 本页开始行
	fmt.Println(resp.EndRow)  //EndRow: 本页结束行
}
```

## 导航

### 使用0作为首页

当前默认首页是1，如果有需要你也可以使用0作为首页，但请在执行任何操作之前执行，否则可能导致查询混乱:

```go
    pageable.Use0AsFirstPage()
```

### 设置默认的每页结果数

有些时候你全局都使用相同的每页结果数，那么只需要设置`SetDefaultRPP`然后每次查询传入0或负数即可

另外若查询中的每页结果数不合法(0或负数)，则会自动采用默认值

```go
    pageable.SetDefaultRPP(25) //设置每页25行
```

下一次你就可以用`rpp=0`来查询了

```go
    pageable.PageQuery(page:1, rpp:0, queryHandler: ..., resultPtr: ...)
```

### 自定义Recovery

默认的recover组件会打印堆栈，如果你需要更多的`Recovery`可以直接在这里注入，会直接替换默认的Recovery:

```go
package main
import (
    "fmt"
    pageable "github.com/BillSJC/gorm-pageable"
)

// 你的自定义Recovery函数
func myRecovery(){
    if err := recover ; err != nil {
        fmt.Println("something happend")
        fmt.Println(err)
        //然后可以加入自定义内容
    } 
}

func init(){
    //自定义你的Recovery
    pageable.SetRecovery(myRecovery)
}
```