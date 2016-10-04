/*
   example1.go
   Created Thu Aug 18 10:21:03 AKDT 2016
   by Raymond E. Marcil <marcilr@gmail.com>

   Example of using Go MySQL Driver to create database handle
   and execute query against table.

   References (initial source for code):
     RawBytes
     Examples
     go-sql-driver/mysql
     The code below is from this page - rock on.
     https://github.com/go-sql-driver/mysql/wiki/Examples
  
   NOTE: RawBytes is a byte slice that holds a reference to memory
         owned by the database itself.  After a Scan into a RawBytes,
         the slice is only valid until the next call to Next, Scan,
         or Close.
         https://golang.org/pkg/database/sql/#RawBytes 
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
)


/*
   main()
   1. Create database handle with:
      sql.Open()
      defer db.Close()
   2. Test database availability with db.Ping()
   3. Execute the query with db.Query()
   4. Get column name with rows.Columns()
   5. Make a slice for the values:
      values := make([]sql.RawBytes, len(columns))
   6. Fetch rows with rows.Next()
   7. Get RawBytes from data with rows.Scan(scanArgs...)
 */
func Main() {

    fmt.Println("Example 1")

    // Create database handle
    db, err := sql.Open("mysql",
        "user:password@tcp(127.0.0.1:3306)/hello")
    if err != nil {
        //log.Fatal(err)
	// Just for example purpose. You should use proper error handling instead of panic
	panic(err.Error())
    }
    defer db.Close()

    /*
       Test database is available and accessible
       --go-database-sql.org/accessing.html
     */
    err = db.Ping()
    if err != nil {
        fmt.Println("Main() db.Ping() Failed: err == nil")
    }

    // Execute the query
    rows, err := db.Query("SELECT * FROM table")

    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    // Get column names
    columns, err := rows.Columns()
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    // Make a slice for the values
    values := make([]sql.RawBytes, len(columns))

    /*
       rows.Scan wants '[]interface{}' as an argument, so we must copy the
       references into such a slice
       See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
     */
    scanArgs := make([]interface{}, len(values))
    for i := range values {
        scanArgs[i] = &values[i]
    }

    // Fetch rows
    for rows.Next() {
        // get RawBytes from data
	err = rows.Scan(scanArgs...)
	if err != nil {
	    panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Now do something with the data.
	// Here we just print each column as a string.
	var value string
	for i, col := range values {
	    // Here we can check if the value is nil (NULL value)
	    if col == nil {
	        value = "NULL"
	    } else {
	        value = string(col)
	    }
	    fmt.Println(columns[i], ": ", value)
	} // End for i, col := range values
    } // End rows.Next()
    if err = rows.Err(); err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

} // End Main()
