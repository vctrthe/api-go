# Go API
## Background
This API is developed using Go Language and MySQL is being used as the database. This API will be publicly open and will be act as endpoints for learning Frontend Development using Nuxt, Vue, Angular, React, etc.
<br>
This API is the outcome of my progress of learning via an online class, called <a href="https://buildwithangga.com/kelas/full-stack-golang-vue-nuxtjs-website-crowdfunding"> BWA.</a> This repository will be updated regularly and futher changes may applied.

## Feature
### List of features
- [x] All the endpoints are running
- [x] Clean, modular code
- [x] Separate secret values like db config, JWT Secret Key, and Payment Keys
- [x] MidTrans payment gateway integration
- [x] Packed inside a Docker Container, complete with MySQL and PMA
- [ ] Properly working payment gateway

## Project Directory Map
```
api-go/
|-- auth/
|   |-- service.go
|
|-- campaign/
|   |-- entity.go
|   |-- formatter.go
|   |-- input.go
|   |-- repository.go
|   |-- service.go
|
|-- config/
|   |-- config.go
|   |-- config.yml
|   |-- config.yml.example  --> *rename this to `config.yml` and change the values according to your own.
|
|-- database/
|   |-- db.go
|
|-- docker/
|   |-- Dockerfile
|
|-- handler/
|   |-- campaign.go
|   |-- transaction.go
|   |-- user.go
|
|-- helper/
|   |-- helper.go
|
|-- images/
|   |-- campaign-images/
|   |   |-- .gitkeep    --> *not imported but the directory is a must
|   |
|   |-- notes.txt   --> *complete explanation here
|
|-- middleware/
|   |-- middleware.go
|
|-- payment/
|   |-- entity.go
|   |-- service.go
|
|-- routes/
|   |-- routes.go
|
|-- transaction/
|   |-- entity.go
|   |-- formatter.go
|   |-- input.go
|   |-- repository.go
|   |-- service.go
|
|-- user/
|   |-- entity.go
|   |-- formatter.go
|   |-- input.go
|   |-- repository.go
|   |-- service.go
|
|-- .gitignore
|
|-- docker-compose.yml.example --> remove `.example` and fill your own values
|
|-- go.mod
|
|-- go.sum
|
|-- main.go
|
|-- README.md   <-- You are here
```
## Changelog
This project is the same project as <a href="github.com/vctrthe/go-api"> this,</a> but due to continuous problems and errors, I decided to create a new repository that is working.

## Future Plans
This project has a ton of plans for the future, but for now, the list will be short.
### Future development for this `go-api` project
- [x] Docker Containerization
- [ ] Host this project, with the Docker container of course, in my self-hosted server

**(Stay tuned, this will be updated as the progression of learning process)**
