# Brainfuck Interpreter Writen by GO

# Build
- install go 1.21
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
# you can run with debug mode, at least print all stack
cat helloworld.bf | bf -d
# if you want more memory stack(default is 128), can use -s
cat helloworld.bf | bf -s 65535
# we can run it directly using the command line
# command line first
./bf -f helloworld.bf
```
# Command args
```shell
Usage of ./bf:
-d	debug mode
-f string
input file
-s int
run stack size (default 128)
```

# Enjoy it! GoodLuck!
