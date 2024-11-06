## purpose

sums two numbers

## pre requisites
You need to have

1. VSCode installed. [Download](https://code.visualstudio.com/download)

2. gRPC and protobuf compiler installed. 
[Instructions](https://grpc.io/docs/languages/go/quickstart/)

## build

### first build 
run `go mod tidy` to download dependencies 

### all builds
Open VSCode terminal and select bash terminal type. (Note: Bash shell is included in Git for Windows. Check internet for other options)
run `./calculator/build.sh` from the project root folder

## run

### server 
Open a terminal and run command `go run ./calculator/server/ .`.
You should see 'Listening on port 0.0.0.0:6666...'

### client
Open a separate terminal from the server terminal and run `go run ./calculator/client/ .`