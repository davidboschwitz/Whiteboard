EXECUTABLE = build/main

all: src/main.go
	mkdir -p build
	go build -o $(EXECUTABLE) src/main.go

run:
	./$(EXECUTABLE)

clean:
	rm -rf build
