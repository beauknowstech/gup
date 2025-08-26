![guptheguppy](images/guptheguppy_small.png?raw=true "guptheguppy")
This is gup the guppy

# gup

gup (go-up) is meant to be a go replacement for `python3 -m http.server` with additional functionalities.

## Features
* Shows what IP(s) and Port its listening on
* Shows what local directory gup is listening from
* Shows the file names of the files in the directory
* File Upload Feature
```
#Upload From Windows:
(New-Object System.Net.WebClient).UploadFile('http://<IP>', '<Full Filepath>')

#Upload From Linux:
curl -F 'file=@<FILENAME>' http://<IP>/
```

## Output:
![gup](images/gup.png?raw=true "output")

## Install

* Run `go install github.com/n000b3r/gup@latest` (ensuring `$GOPATH` is in your `$PATH`)
* Or download the gup binary from the releases page and put it somewhere in your `echo $PATH`. Mine is in /usr/bin/
* Or compile yourself with `go build main.go` after cloning the repo

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

https://github.com/beauknowstech/gup

https://github.com/sc0tfree/updog

https://github.com/MuirlandOracle/up-http-tool

https://github.com/patrickhener/goshs
