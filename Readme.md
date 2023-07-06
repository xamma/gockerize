# Gockerize - Build Dockerfiles for you Go projects
This project is a very simple CLI tool for creating Dockerfiles for Go projects.  
It always uses the same template and may need to be adjusted according to your needs.  
I primarly made this for learning how to write CLI tools with the **Cobra library in Go**.  

## How to use
Download the binary and execute it. Make sure to have the right permissions as well as the right operating system.  

## Where it works
I tested it on this common project layout, and it will work on similar layouts.  
```
├── Dockerfile
├── LICENSE.md
├── Readme.md
├── config
│   └── config.go
├── database
│   ├── db.go
│   └── pokemon_db.go
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── gokedex.postman_collection.json
├── handlers
│   └── pokemon_handlers.go
├── main.go
└── models
    ├── appconfig.go
    └── pokemon.go
```