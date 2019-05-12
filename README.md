# VBM - Virtual BrainFucking Machine

VBM is a simple brainfuck interpreter with a fancier name.

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
1. All you need to about BrainFucking - <https://esolangs.org/wiki/Brainfuck>
2. Some brainfuck fluff, includes sample code and other resources- <http://www.hevanet.com/cristofd/brainfuck/>
