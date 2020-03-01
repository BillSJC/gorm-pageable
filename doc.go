// Package pageable is for page query based on GORM package
//
// As a quick start:
// 	func getResultSet (page int,rowsPerPage int)(*pageable.Response,error){
// 	//your empty result set
// 		resultSet := make([]*User,0,30)
// 		//prepare a handler to query
// 		handler := DB.
// 			Module(&User{}).
// 			Where(&User{Active:true})
// 		//use PageQuery to get data
// 		resp,err := pageable.PageQuery(page,rowsPerPage,handler,&resultSet)
// 		// handle error
// 		f err != nil {
// 			panic(err)
// 	}
// And then you can print this value to see the page info
// 		fmt.Println(resp.PageNow)    //PageNow: current page of query
// 		fmt.Println(resp.PageCount)  //PageCount: total page of the query
// 		fmt.Println(resp.RawCount)   //RawCount: total raw of query
// 		fmt.Println(resp.RawPerPage) //RawPerPage: rpp
// 		fmt.Println(resp.ResultSet)  //ResultSet: result data
// 		fmt.Println(resp.FirstPage)  //FirstPage: if the result is the first page
// 		fmt.Println(resp.LastPage)   //LastPage: if the result is the last page
// 		fmt.Println(resp.Empty)  //Empty: if the result is empty
// 	}
// And here a clear JSON object of the Response LIKE Spring Pageable
// 	{
// 		"PageNow": 2,
// 		"PageCount": 1,
// 		"RawCount": 1,
// 		"RawPerPage": 25,
// 		"ResultSet": [],
// 		"FirstPage": false,
// 		"LastPage": false,
// 		"Empty": true
// 	}
package pageable
