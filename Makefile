BUILD_TARGET = bf

build: main.go
	go build --ldflags="-s -w" -o $(BUILD_TARGET) $^

clean: $(BUILD_TARGET)
	RM $^

