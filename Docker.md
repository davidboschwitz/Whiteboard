To run in Docker, we first need to build it:
```
go get github.com/hunterpraska/Whiteboard
cd $GOPATH/src/github.com/hunterpraska/Whiteboard
docker build -t whiteboard
```

Then just ``` docker run -p 8080:8080 --name whiteboard whiteboard```


### TODO
Add Docker build to one of the Docker repositories.
