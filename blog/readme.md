## blog application

## run

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