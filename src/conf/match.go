/*
   match.go
   Created Thu Oct 20 10:49:12 AKDT 2016
   Copyright (C) 2016 by Raymond E. Marcil <marcilr@gmail.com>

   Package to read name/value pairs from file.
*/
package main

import (
	"fmt"
	"log"
	"regexp"
)

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

} // End info()

func main() {
	// What's next after foo, bar?
	// Wednesday, September 01, 2004
	// http://nerddawg.blogspot.com/2004/09/whats-next-after-foo-bar.html
	valuepair := "foo='bar=baz,quux=quuux'"

	info("main() info")
	log.Println("main() log.Println")

	/*
		    =========================================================
			  Compile regular expression to match
			  comments.  Comments are zero or more spaces followed by #
			  =========================================================
	*/
	commentRegexp, err := regexp.Compile(`^\s*#`)

	if err != nil {
		log.Println("main() Error compiling commentRegexp")
		log.Fatal(err)
	}

	// Determine if line, in valuepair variable, is a comment
	isComment := commentRegexp.MatchString(valuepair)

	if isComment == true {
		log.Println("main() valuepair is comment")
	}

	/*
			==============================================================
			Compile regular expression to match lines with only whitespace
		  ==============================================================
	*/
	whitespaceLineRegexp, err := regexp.Compile(`^\s*$`)

	if err != nil {
		log.Println("main() Error compiling whitespaceLineRegexp")
		log.Fatal(err)
	}

	// Determine if line, in valuepair variable, is a whitespace
	isWhitespaceLine := whitespaceLineRegexp.MatchString(valuepair)

	if isWhitespaceLine == true {
		log.Println("main() valuepair is comment")
	}

	/*
		  =========================================
			Compile regular expression to match value
		  =========================================
	*/
	valueRegexp, err := regexp.Compile(`=.*$`)

	if err != nil {
		log.Println("main() Error compiling valueRegexp")
		log.Fatal(err)
	}

	// Determine if value present
	isValue := valueRegexp.MatchString(valuepair)

	if isValue == true {
		log.Println("main() valuepair has value: ")

		// This print "true"
		fmt.Println(valueRegexp.MatchString(valuepair))

		// This prints "='bar=baz,quux=quuux'"
		value := valueRegexp.FindString(valuepair)
		fmt.Println("main() value: " + value)

		// NOTE: string() is used to not print ASCII value of 39
		// How to index characters in a Golang string?
		// http://stackoverflow.com/questions/15018545/how-to-index-characters-in-a-golang-string
		fmt.Println(string(value[1]))
		fmt.Println(string([]rune(value)[1]))

		// Use -1 since 0 based indexing
		last := value[len(value)-1]
		fmt.Println(string(last))

		fmt.Println("last: " + string(value[1]))

		// Check if leading/trailing single quotes and remove
		// First check if first/last character of string are the same
		if string(value[1]) == string(value[len(value)-1]) {
			fmt.Println("Foo")
		}

	} else {
		log.Println("main() valuepair has *no* value")
	}

} // End main()
