[![Build Status](https://cloud.drone.io/api/badges/elahe-dastan/HTTP_monitoring/status.svg)](https://cloud.drone.io/elahe-dastan/HTTP_monitoring)

#HTTP endpoints monitoring service
It's service with Golang programming language to monitor HTTP endpoints so that in some configurable periods<br/>
(e.g., the 30s, 1m, 5m) this service sends HTTP requests to the endpoint and logs the response status code<br/>

##Endpoints
To use this service a user should first register at register endpoint then he should login to get a token at login<br/>
endpoint then he can use the token to add url at url endpoint 

##Database
I have used postgres database for this project 

##Example of use
```sh
$ curl -X POST -d '{"Email": "elahe.dstn@gmail.com", "Password": "XXXX"}' 
-H 'Content-Type: application/json' 127.0.0.1:8080/register
```
```sh
$ curl -X POST -d '{"Email": "elahe.dstn@gmail.com", "Password": "XXXX"}' 
-H 'Content-Type: application/json' 127.0.0.1:8080/login
```
this should return a token

```sh
$ curl -X POST -d '{"Token": "token", "URL": "https://www.google.com"}' 
-H 'Content-Type: application/json' 127.0.0.1:8080/url
```
