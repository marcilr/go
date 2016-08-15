/*
   info.go
   Created Mon Aug 15 14:32:11 AKDT 2016
   Copyright (C) 2016 by Raymond E. Marcil <marcilr@gmail.com>

   Package to show use of info() function to display debugging
   output.  info() used to be naming debug() but info() is a 
   bit more manager friendly.
*/
package main


import "fmt"


/*
 * This is how you declare a global variable
 * Flag it enable/disable info logging
 */
var infoflag bool = true

/*
 * info()
 * Display info messags if infoflag is set to true.
 */
func info(msg string) {

//   fmt.Println(msg)

    if infoflag == true {
        fmt.Println(msg)
    }

}   // End info()


/*
 * main()
 * Main body to show info calls.
 */
func main() {

  info("main() START ***")

  info("main() main body")

  info("main() END ***")

}   // End main()


