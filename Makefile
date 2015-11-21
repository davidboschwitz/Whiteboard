EXECUTABLE = Whiteboard

all: main.go
	go build

run:
	./$(EXECUTABLE) web

clean:
	rm -rf $(EXECUTABLE)
