<<<<<<< HEAD
# go-weaver
Estudos arquitetura modular go
=======
# Esutos com service-weaver

Repo de referência utilizado Elton Minetto [microservices-serviceweaver](https://github.com/eminetto/microservices-serviceweaver)

## Build

```
make build
```

## Run single

```
make run-single
```

## Run multi

```
make run-multi
```

## Get status

```
make status
```

## Show dashboard

```
make dashboard
```

## Deploy to cloud

You must follow [this documentation](https://serviceweaver.dev/docs.html#gke) to run this option.

```
make gke-run-multi
```



## Using the services

### Auth

```
curl -X "POST" "http://localhost:12345/auth" \
     -H 'Accept: application/json' \
     -H 'Content-Type: application/json' \
     -d $'{
  "email": "eminetto@gmail.com",
  "password": "1234567"
}'

```

The result should be a token, like:

```
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImVtaW5ldHRvQGdtYWlsLmNvbSIsImV4cCI6MTY3Nzc2NzA5NSwiaWF0IjoxNjc3NzYzNDY1LCJuYmYiOjE2Nzc3NjM0NjV9.XXNnS35c0D1H2kdJzKIs4sJrNlICCbWgwe1cZNu3ZbQ"
}
```

### Feedback

You need to use the token generated by the ```Auth``` service:

```
curl -X "POST" "http://localhost:12345/feedback" \
     -H 'Accept: application/json' \
     -H 'Content-Type: application/json' \
	 -H 'Authorization:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImVtaW5ldHRvQGdtYWlsLmNvbSIsImV4cCI6MTY3Nzc2NzA5NSwiaWF0IjoxNjc3NzYzNDY1LCJuYmYiOjE2Nzc3NjM0NjV9.XXNnS35c0D1H2kdJzKIs4sJrNlICCbWgwe1cZNu3ZbQ' \
     -d $'{
  "title": "Feedback test",
  "body": "Feedback body"
}'
```

### Vote

You need to use the token generated by the ```Auth``` service:

```
curl -X "POST" "http://localhost:12345/vote" \
     -H 'Accept: application/json' \
     -H 'Content-Type: application/json' \
	 -H 'Authorization:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImVtaW5ldHRvQGdtYWlsLmNvbSIsImV4cCI6MTY3Nzc2NzA5NSwiaWF0IjoxNjc3NzYzNDY1LCJuYmYiOjE2Nzc3NjM0NjV9.XXNnS35c0D1H2kdJzKIs4sJrNlICCbWgwe1cZNu3ZbQ' \
     -d $'{
  "talk_name": "Go e Microserviços",
  "score": "10"
}'
```
>>>>>>> 471dee6 (First commit estudo arquitetura modular go)
