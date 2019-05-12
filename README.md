# VBM - Virtual BrainFucking Machine

[![Go Report Card](https://goreportcard.com/badge/github.com/souvikmaji/vbm)](https://goreportcard.com/report/github.com/souvikmaji/vbm)
[![wercker status](https://app.wercker.com/status/e7cb67d27848f4a08dd86543a1c3b8bf/s/master "wercker status")](https://app.wercker.com/project/byKey/e7cb67d27848f4a08dd86543a1c3b8bf)
![Codecov](https://img.shields.io/codecov/c/github/souvikmaji/vbm.svg?style=popout-square)
[![GoDoc](https://godoc.org/github.com/souvikmaji/vbm?status.svg)](https://godoc.org/github.com/souvikmaji/vbm)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

VBM is a simple yet fast brainfuck interpreter with a fancier name.

For those who does not know Brainfuck is the most famous esoteric programming language(or esolang, a programming language designed to experiment with weird ideas, to be hard to program in, or as a joke, rather than for practical use.)

## Usage:

```shell
# download the machine
go get -v github.com/souvikmaji/vbm/...

# go to the machine
cd $GOPATH/src/github.com/souvikmaji/vbm

# install the machine
go install

# give it a input and receive a output
cd samples
vbm hello_world.b
```

## References
1. All you need to know about BrainFucking - <https://esolangs.org/wiki/Brainfuck>
2. Some brainfuck fluff, includes sample code and other resources- <http://www.hevanet.com/cristofd/brainfuck/>
