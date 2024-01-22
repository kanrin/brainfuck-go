# Brainfuck Interpreter Writen by GO

# Build
- install go 1.21
- install make
```shell
clone -b main https://github.com/kanrin/brainfuck-go.git
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
./bf -f helloworld.bf
```

# Enjoy it! GoodLuck!
