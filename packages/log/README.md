README.md
Created Mon Aug 15 14:42:18 AKDT 2016
Copyright (C) 2016 by Raymond E. Marcil <marcilr@gmail.com>


Package log


github: git-dev-lang-go-packages-log


Filelist
========
links.txt
  Package log related links

log.txt
  Package log implements a simple logging package.  It defines a type, Logger,
  with methods for formatting output.  It also has a predefined 'standard'
  Logger accessible through helper functions Print[f|ln], Fatal[f|ln], and
  Panic[f|ln], which are easier to use than creating a Logger manually.  That
  logger writes to standard error and prints the date and time of each logged
  message.  The Fatal functions call os.Exit(1) after writing the log message.
  The Panic functions call panic after writing the log message.
  https://golang.org/pkg/log/


Links
=====
Package log
import "log"
The Go Programming Language
https://golang.org/pkg/log/
