# dir-scan :open_file_folder:
[![Go Report Card](https://goreportcard.com/badge/github.com/eze-kiel/dir-scan)](https://goreportcard.com/report/github.com/eze-kiel/dir-scan)

dir-scan is a multi threads web content scanner. It looks for existing and/or hidden web pages on a specific target.

## Usage
```
-target <URL>
    The target URL

-type <TYPE>
    The type of scan. Values can be : 
        wp (for WordPress)
        common (for standard names)
        joomla (for Joomla websites)
        apache (for Apache web servers)
        linuxfiles (for interesting Linux files. Warning : > 87.000 lines, it is a long scan !)

-v
    Verbose mode (display 404 status codes)
    Default : false

-mt
    Enable multi threading. It will create 2 threads
    Default : false
```

The `mt` flag will separate the work into 2 differents goroutines. It will also display results in an unalphabetical order, due to the list splitting.

## About lists
This tool uses lists which are inside the `lists/` folder.
The lists used here are from [github.com/danielmiessler/SecLists](github.com/danielmiessler/SecLists). You should definitely check his GitHub repo for a lot more lists for different purposes.

You can add your own lists by copying them inside the `lists/` folder. They have to be .txt.