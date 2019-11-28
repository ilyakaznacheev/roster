# Roster Service 
 
The simple service for a roster management. 
 
[![Go Report Card](https://goreportcard.com/badge/github.com/ilyakaznacheev/roster)](https://goreportcard.com/report/github.com/ilyakaznacheev/roster)  
[![GoDoc](https://godoc.org/github.com/ilyakaznacheev/roster?status.svg)](https://godoc.org/github.com/ilyakaznacheev/roster) 
[![Coverage Status](https://codecov.io/github/ilyakaznacheev/roster/coverage.svg?branch=master)](https://codecov.io/gh/ilyakaznacheev/roster) 
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](/LICENSE) 
 
## Contents 
 
- [About](#about) 
- [Tools and Technologies](#tools-and-technologies) 
- [Tests](#tests) 
- [Production readiness](#production-readiness) 
- [Dependencies](#dependencies) 
- [Installation](#installation) 
- [Run](#run) 
- [Docker](#docker) 
- [Cloud](#cloud) 
- [API](#api) 
 
## About 
 
This tiny service was built to manage rosters with very limited abilities. The API is described in [API](#api) section. 
 
Tools, libraries and approaches were chosen with a focus on production readiness.  
 
### Tools and Technologies 
 
There are some technical decisions made to bring more flexibility and scalability to the service. It may look overcomplicated, but there is extensibility and reliability behind the simplicity. 
 
There are no frameworks due to simple infrastructure.  
 
- MongoDB - chosen because it's easy to shard and has enough search engine power to fulfill service needs; 
- Swagger - perfectly describes REST APIs and can be used to generate server and client API-code in a single command; 
- JWT auth - used to secure changes via API. Used because it's a common choice for API auth. 
- Docker - used to satisfy requirements and build and run the service in isolated environment. Multistage build is made to make resulted image as small as possible. Can be used on any cloud environment or as a part of orchestration systems like k8s; 
- Heroku - chosen as a simplest hosting with CI/CD features because it is easy to setup and monitor. 
 
### Tests 
 
All core logic (request handlers) is covered with unit-tests. No need to check infrastructure and generated core, because it doesn't change so much, but really hard to test. DB and web layers are isolated. 
 
### Production readiness 
 
The service is ready to run as a standalone microservice in any environment, e.g. K8s. 
 
It is built with a closer look to 12-factor app principle, and can be easily integrated into 12-f infrastructure with small changes or additions. 
 
## Dependencies 
 
The tool has no runtime dependencies, but some development tools which you may need to contribute: 
 
### go-swagger 
 
go-swagger is a tool for automate code generation based on [`swagger.yml`](/swagger.yml) schema. 
 
Install on MacOS 
 
```bash 
brew install go-swagger 
``` 
 
Install on Linux 
```bash 
echo "deb https://dl.bintray.com/go-swagger/goswagger-debian ubuntu main" | sudo tee -a /etc/apt/sources.list 
``` 
 
[Other installation options](https://github.com/go-swagger/go-swagger#installing). 
 
## mockery 
 
mockery is a tool for automate code generation for testing mocks. 
 
Install 
 
```bash 
go get github.com/vektra/mockery/.../ 
``` 
 
[More information](https://github.com/vektra/mockery). 
 
## Installation 
 
Installation as a Go package 
 
```bash 
go get -v -u https://github.com/ilyakaznacheev/roster 
``` 
 
Installation as a git repo 
 
```bash 
git clone https://github.com/ilyakaznacheev/roster.git 
``` 
 
## Run 
 
There are several options. You can run the service as a standalone binary, but you should satisfy infrastructure requirements in this case. For more information run 
 
```bash 
go run cmd/roster/main.go -h 
``` 
 
> Note! There is no possibility to create a roster via API, so you need to initialize the database by your own. 
 
But there are some more useful ways to run it preconfigured for you. 
 
### Docker 
 
To start the service in a docker-compose, you need `docker` and `docker-compose` to be installed on your machine. 
 
Then you can run 
 
```bash 
make docker/up 
``` 
 
it will start to serve on `localhost:8080`. 
 
To shut it down run 
 
```bash 
make docker/down 
``` 
 
### Cloud 
 
The project is preconfigured to run on Heroku. You can clone the repo and connect your own Heroku app - it should start Dockerfile. For more information read my [article](https://dev.to/ilyakaznacheev/setup-build-automate-deploy-a-dockerized-app-to-heroku-fast-167). 
 
## API 
 
API is tiny but powerful. For detailed information check [`swagger.yml`](/swagger.yml). For nice picture paste it into [Swagger Editor](https://editor.swagger.io/). 
 
Shortly you can do the following: 
 
- [POST] `/api/login` - get JWT token (you need it to make changes) 
- [GET] `/api/rosters` - get a list of rosters 
- [GET] `/api/rosters/{id}` - get certain roster 
- [GET] `/api/rosters/{id}/active` - get a roster with active players only 
- [GET] `/api/rosters/{id}/benched` - get a roster with benched players only 
- [POST] `/api/rosters/{id}/add_player` - add a new player 
- [POST] `/api/rosters/{id}/rearrange` - rearrange players 
 