GOOS=windows GOARCH=amd64 go build -o ./bin/flora-amd64.exe main.go
GOOS=windows GOARCH=386 go build -o ./bin/flora-386.exe main.go
GOOS=linux GOARCH=amd64 go build -o ./bin/flora-amd64-linux main.go
GOOS=darwin GOARCH=amd64 go build -o ./bin/flora-amd64-darwin main.go