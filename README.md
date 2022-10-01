# gup

gup (go-up) is meant to be a go replacement for `python3 -m http.server` with features that I find useful when practicing on HackTheBox, TryHackMe, OSCP etc. 

## Features
* Shows what IP(s) and Port its listening on
* Shows what local directory gup is listening from
* Shows the file names of the files in the directory

## Flags

-p to change the port. Default 80.

-d to change the directory. Default is pwd.

-r show files recursively in addition to just the files in pwd. Default false. 

## Examples

**Serve from your current directory**

`gup`

**Serve from your current directory and show files recursively**

`gup -r`

**Serve from another directory**

`gup -d /tmp`

**Serve from port 1337**

`gup -p 1337`

  
## Credits/Inspiration

https://github.com/sc0tfree/updog

https://github.com/MuirlandOracle/up-http-tool

https://github.com/patrickhener/goshs


