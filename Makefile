EXECUTABLE = main
VIEWS = src/views

all: src/main.go
	cp -r $(VIEWS) .
	go build -o $(EXECUTABLE) src/main.go

run:
	./$(EXECUTABLE)

clean:
	rm -rf $(EXECUTABLE) views/
