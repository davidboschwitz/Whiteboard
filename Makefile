EXECUTABLE = main

all: main.go
	go build -o $(EXECUTABLE) src/main.go

run:
	./$(EXECUTABLE)

clean:
	rm -rf $(EXECUTABLE)
