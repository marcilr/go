README.md
Created Mon Aug 15 15:05:17 AKDT 2016
Copyright (C) 2016 by Raymond E. Marcil <marcilr@gmail.com>


Package os
import 'os'

github: git-dev-lang-go-packages-os


Filelist
========
links.txt
  Package os related links

os.txt
  Package os provides a platform-independent interface to operating system
  functionality.  The design is Unix-like, although the error handling is
  Go-like; failing calls return values of type error rather than error numbers.
  Often, more information is available within the error.  For example, if a call
  that takes a file name fails, such as Open or Stat, the error will include the
  failing file name when printed and will be of type *PathError, which may be
  unpacked for more information.
  https://golang.org/pkg/os/
