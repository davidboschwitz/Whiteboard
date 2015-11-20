EXECUTABLE = main

all: main.go
	go build -o $(EXECUTABLE) main.go

run:
	./$(EXECUTABLE)

clean:
	rm -rf $(EXECUTABLE)
