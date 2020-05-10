# dir-scan :open_file_folder:
[![Go Report Card](https://goreportcard.com/badge/github.com/eze-kiel/dir-scan)](https://goreportcard.com/report/github.com/eze-kiel/dir-scan)

dir-scan is a multi threads web content scanner. It launches a dictionnary-based attack to find existing and/or hidden web pages on a specific target.

## Usage
```
-t <URL>
    The target's URL

-d <PATH-TO-DICT>
   Dictionnary's path

-v
    Verbose mode (display 404 status codes)
    Default : false

-mt
    Enable multi threading. It will create 2 threads
    Default : false

-w
    Waiting time between requests to avoid flood, in milliseconds
```

The `mt` flag will separate the work into 2 differents goroutines. It will also display results in an unalphabetical order, due to the list splitting.

## About lists
If you need some lists, you should try [github.com/danielmiessler/SecLists](github.com/danielmiessler/SecLists).

## Notes
If this program doesn't fit your need, you should try [dirb](https://tools.kali.org/web-applications/dirb).