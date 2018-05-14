margarine _spreads_ an empty file throughout a file system.

Usage:
```
$ go get github.com/derekargueta/margarine
$ margarine -p 3 -r /Users/dereka -f uselessfile.txt > paths_of_new_files.txt
```

clean.sh is a helper script.

`$ ./clean.sh check path_of_new_files.txt` will list out the paths of all the new files and a "yes/no" as to whether they exist.

You can clean them up with

`$ ./clean.sh clean path_of_new_files.txt` which `rm`'s all the paths in the file.
