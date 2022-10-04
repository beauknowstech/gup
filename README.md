# gup

gup (go-up) is meant to be a go replacement for `python3 -m http.server` with features that I find useful when practicing on HackTheBox, TryHackMe, OSCP etc. 

## Features
* Shows what IP(s) and Port its listening on
* Shows what local directory gup is listening from
* Shows the file names of the files in the directory

## Output:
![gup](images/gup.png?raw=true "output")


## Install
Either compile yourself with `go build main.go` or download the gup binary from the releases page and put it somewhere in your `echo $PATH`. Mine is in /usr/bin/

## Flags

-p to change the port. Default is 80.

-d to change the directory. Default is pwd.

-r show files recursively in addition to just the files in pwd. Default is false. 

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


## Disclaimer

I'm not a developer. I only tested this on Kali 2022.2. If you want to make changes then go for it. 

### To-Do
~~Does anyone know how to make a logger for something like this?~~
