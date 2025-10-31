# Brainfuck Interpreter Writen by GO

Brainfuck Interpreter by GO with 100 line code

# Build

- install go 1.21 or newer
- install make
- install git

```shell
git clone -b main https://github.com/kanrin/brainfuck-go.git
cd brainfuck-go
make build
```

# How to use

```shell
cat helloworld.bf | ./bf
# also you can mv bf binary into your $PATH
# and you can run with
cat helloworld.bf | bf
# if you want more memory stack(default is 32), can use -s
cat helloworld.bf | bf -s 65535
# we can run it directly using the command line
# command line first
./bf -f helloworld.bf
./bf -f helloworld.bf -s 65535
```

# Command args

```shell
Usage of ./bf:
-f string  input file
-s int     run stack size (default 128)
```

# GoodLuck!
