# blog application

# build and run code

## purpose

blog

## pre requisites
You need to have:

1. VSCode installed. [Download](https://code.visualstudio.com/download)

2. gRPC and protobuf compiler installed. 
[Instructions](https://grpc.io/docs/languages/go/quickstart/)

## build

### first build 
run `go mod tidy` to download dependencies 

### all builds
Open VSCode terminal and select bash terminal type. (Note: Bash shell is included in Git for Windows. Check internet for other options)
run `./blog/build.sh` from the project root folder

## run

### server 
Open a terminal and run command `go run ./blog/server/ .`.
You should see 'Listening on port 0.0.0.0:6666...'

### client
Open a separate terminal from the server terminal and run `go run ./blog/client/ .`



## run Container

### Start container
In terminal 
1. navigate to blog folder
2. run `docker-compose up`
3. Open `http://localhost:8081/` in browser. You should see Mongo Express UI

### Stop container
3. Open another terminal and navigate to blog folder
4. run `docker-compose down`

## .env file
This is in the repo but ignored by git by using command `git update-index --assume-unchanged ./blog/.env` 
We want to be able to see the format in github but not the values


## MongoDb

We are using the driver form this [repo](https://github.com/mongodb/mongo-go-driver).
Run command `go get go.mongodb.org/mongo-driver/v2/mongo` in terminal to add dependency


