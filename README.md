# msgpack

> :warning: DO NOT USE, STILL WIP!

![](https://img.shields.io/github/v/tag/arivum/msgpack?label=latest&color=%234488BB)
![](https://img.shields.io/github/go-mod/go-version/arivum/msgpack?color=%234488BB)
![](https://img.shields.io/github/workflow/status/arivum/msgpack/Go)
![](https://img.shields.io/github/license/arivum/msgpack?color=%234488BB)


At the minute, this module contains use-case-specific msgpack decoding improved for maximum performance.

[See package documentation](https://pkg.go.dev/github.com/arivum/msgpack)

## How to use

```go
package main

import (
	"fmt"
	"strings"

	"github.com/arivum/json2msgpackStreamer"
	"github.com/arivum/msgpack"
)

func main() {

	conv := json2msgpackStreamer.NewJSON2MsgPackStreamer(strings.NewReader(
		`{"a": "testvar", "b": ["b", "c"], "c": null, "d": false, "e": true, "g": -10, "f": -2.1, "h": "this is a longer text that needs a broader length indicator", "i": "sldkfj"}
{"a": 4, "b": 18000000043000000233, "b2":  -9000000043000000233, "c": null, "d": false, "e": true, "g": -10, "f": 2.0, "h": "this is a longer text that needs a broader length indicator", "i": "sldkfj"}
{"a": "testvar", "b": ["b", "c"], "c": null}

{"a":"this is larger map test","b":"b","c":"c","d":"d","e":"e","f":"f","g":"g","h":"h","i":"i","j":"j","k":"k","l":"l","m":"m","n":"n","o":"o","p":"p","q": 1}

{"cmplx": {"a":"this is larger map test","b":"b","c":"c","d":"d","e":"e","f":"f","g":"g","h":"h","i":"i","j":"j","k":"k","l":"l","m":"m","n":"n","o":"o","p":"p","q": 1}}

["this is a larger slice test", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o"]

{"cmplx": ["2nd", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n"]}
`,
	))

	d := msgpack.NewDecoder(conv)
	for entry := range d.Stream() {
		fmt.Printf("%+v\n", entry)
	}
	fmt.Println(d.LastError())
}

```
