README.md
Created Mon Aug 15 09:56:07 AKDT 2016
Copyright (C) 2016 by Raymond E. Marcil <marcilr@gmail.com>


Package bufio


github: git-dev-lang-go-packages-bufio


Filelist
========
bufio.txt
  Package bufio
  import "bufio"
  Package bufio implements buffered I/O.  It wraps an io.Reader or io.Writer
  object, creating another object (Reader or Writer) that also implements the
  interface but provides buffering and some help for textual I/O.
  https://golang.org/pkg/bufio/

links.txt
  Package bufio related links

read-file-line-by-line.txt 
  How to read a file starting from a specific line number using Scanner?
    
  edited Jan 8 at 9:03 - icza
  http://stackoverflow.com/users/1705598/icza
    
  asked Jan 7 at 11:54 - Amyth
  http://stackoverflow.com/users/1373822/amyth
    
  Instead of using a Scanner, use a bufio.Reader, specifically the ReadBytes
  or ReadString methods.  This way you can read up to each line termination,
  and still receive the full line with line endings.
    
  Has great piece of code for reading a file line by line.
  http://stackoverflow.com/questions/34654514/how-to-read-a-file-starting-from-a-specific-line-number-using-scanner
