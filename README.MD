# Beego auth with React

Implementation of an auth flow using Beego (go) and React. 

## Set up 

For db connection change the string to what you want in `models/storage.go`  
Once you have the desired db, you can use the command below to create the user table.
```bash
go build main.go  
./main orm syncdb -db=default -force=false  
```

Make sure you have Beego installed 
```bash
go get -u github.com/beego/bee
```

Once you have the command you can run `bee run` inside the root folder. This will run on `port: 8080`

For the frontend go into the frontend folder and run the first time

```bash
yarn install
```  
once installed run

```bash
yarn start
```  
the react project will run `port: 3000`.  

## API endpoints

## User
### POST
createUser: `/user` 
```json
{
  "username": "example",
  "password": "password",
}
```

## Auth
### POST
login: `/auth/login`  
```json
{
    "username": "example",
    "password": "password",
}
{
    "token": "abc...."
}
```
validate: `/auth/validate`  
```json
{
    "token": "abc....",
}
200 OK
```
refresh: `/auth/refresh`  
```json
{
    "token": "abc....",
}
{
    "token": "zyx...."
}
200 OK
```


## Future Improvments

1. Dockerise for easy deployment or use beego to statically serve the frontend.  
2. Form validate and flesh out of frontend.
3. Configuration improvments for beego.
