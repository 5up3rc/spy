# Watcher

Watcher is a simple, general purpose, cross platform, file watcher written in Go (Golang). Watcher takes a `directory` and a `program` and runs `program` whenever a file in `directory` changes.

### Install

```
go get -v github.com/jpillora/watcher
```

*Currently, `watcher` only works on **Unix** systems as it uses process groups to ensure all sub-processes have exited between restarts. A pull request to add Windows support would be greatly appreciated!*

### Usage

```
$ watcher --help

	Usage: watcher [options] program ...args

	program (along with it's args) is initially
	run and then it is restarted with every file
	change. program will always be run from the
	current working directory.

	Options:

	--inc INCLUDE - Describes a path to files to
	watch. Use ** to describe any number of
	directories. Use * to describe any file name.
	For example, you could watch all Go source
	files with "**/*.go" or all	JavaScript source
	files in './lib/' with "lib/**/*.js".

	--exc EXCLUDE - Describes a path to files not
	to watch. Inverse of INCLUDE.

	--dir DIR - Watches for changes to all files in
	DIR (defaults to the current directory). After
	each change, program will be restarted.

	--delay DELAY - Restarts are debounced by DELAY
	(defaults to '0.5s').

	-v - Enable verbose logging

	Read more:
	https://github.com/jpillora/watcher

```

### Examples

Go - Auto restart server

```
$ cd $GOPATH/github.com/user/repo
$ watcher go run main.go
watcher 2015/03/02 01:40:40.248290 Watching '.../github.com/user/repo'
2015/03/02 01:40:40 listening on 8080...
watcher 2015/03/02 01:40:43.996842 Restarting...
2015/03/02 01:40:44 listening on 8080...
```

Go - Auto re-run tests

```
$ watcher go test
```

Node

```
$ watcher node server.js
```

Bash

```
$ watcher bash program.sh
```

### Todo

* Port Unix code to Windows

#### MIT License

Copyright © 2015 &lt;dev@jpillora.com&gt;

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
'Software'), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
