/*
   example2.go
   Created Fri Aug 19 10:39:24 AKDT 2016
   by Raymond E. Marcil <marcilr@gmail.com>

   Example of using Go MySQL Driver to create database handle
   and execute query against table.

   This example uses package reflect to implement run-time reflection.

   References:
     How to call the Scan variadic function in Golang using reflection?
     I'm looking to call the Rows.Scan() function using reflection.
     However it takes a variable number of pointers, but I'm new to Golang
     and there are not a lot of source examples.  I need to use reflection
     because I plan on filling a slice with the values from a Query call.
     So basically using rows.Columns() to get the length of the row and
     then make() a slice of []interface{} to fill with the data points that
     would normally be filled using the pointers passed to the Scan() function.
     edited Aug 3 at 13:01 - amd
     http://stackoverflow.com/users/6169399/amd
     asked Jul 24 '13 at 22:06 - lucidquiet
     http://stackoverflow.com/users/361836/lucidquiet

     To lucidquiet: you can also assign a interface instead of making a slice
     The following code works good:
         var sql = "select * from table"
         rows, err := db.Query(sql)
         columns, err = rows.Columns()
         colNum := len(columns)

         var values = make([]interface{}, colNum)
         for i, _ := range values {
             var ii interface{}
             values[i] = &ii
         }

         for rows.Next() {
             err := rows.Scan(values...)
             for i, colName := range columns {
                 var raw_value = *(values[i].(*interface{}))
                 var raw_type = reflect.TypeOf(raw_value)

                 fmt.Println(colName,raw_type,raw_value)
             }
         }

         NOTE: This package uses the above code by Tom Liu
	 answered Jun 19 '15 at 17:33 - Tom Liu
	 http://stackoverflow.com/users/3266581/tom-liu
     http://stackoverflow.com/questions/17845619/how-to-call-the-scan-variadic-function-in-golang-using-reflection

     Package reflect
     import "reflect"
     Package reflect implements run-time reflection, allowing a program to
     manipulate objects with arbitrary types.  The typical use is to take a
     value with static type interface{} and extract its dynamic type
     information by calling TypeOf, which returns a Type.

     A call to ValueOf returns a Value representing the run-time data.
     Zero takes a Type and returns a Value representing a zero value for
     that type.

     See "The Laws of Reflection" for an introduction to reflection in
     Go: https://golang.org/doc/articles/laws_of_reflection.html
     https://golang.org/pkg/reflect/

*/
package main


/*
  Use the Go MySQL Driver
  https://github.com/go-sql-driver/mysql
 */
import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "reflect"
//    "strconv"
)


/*
 * Main()
 */
func Main() {

    db, err := sql.Open("mysql",
       "user:password@tcp(127.0.0.1:3306)/hello")
    if err != nil {
        //log.Fatal(err)
        // Just for example purpose. You should use proper error handling instead of panic
        panic(err.Error())
    }
    defer db.Close()

    var sql = "select * from table"
    rows, err := db.Query(sql)
    columns, err := rows.Columns()
    colNum := len(columns)

    var values = make([]interface{}, colNum)
    for i, _ := range values {
       var ii interface{}
       values[i] = &ii
    }

    for rows.Next() {
       err := rows.Scan(values...)
       if err != nil {
           //log.Fatal(err)
	   // Just for example purpose. You should use proper error handling instead of panic
	  panic(err.Error())
       }
       for i, colName := range columns {
          var raw_value = *(values[i].(*interface{}))
          var raw_type = reflect.TypeOf(raw_value)

          fmt.Println(colName,raw_type,raw_value)

          // This prints the same as above.
          value := fmt.Sprint(raw_value)
	  fmt.Println("value: " + value)

          //./example2.go:130: cannot convert value (type string) to type int
          //fmt.Println("value: " + strconv.Itoa(int(value)))

          // cannot convert raw_value (type interface {}) to type string: \
 	  // need type assertion
          //fmt.Println("value: " + string(raw_value))

          //invalid operation: "value: " + raw_value.(int) (mismatched types string and int)
          //fmt.Println("value: " + raw_value.(int))

          // Whole series of errors
          //fmt.Println("value: " + strconv.Itoa(raw_value.(int)))


          // NOTE: Uncertain how to display the uint8 raw value above.
       }  // End for i, colName := range columns {
    } // End for rows.Next() {

} // End Main()
