# SampleApp Boilerplate

This sample project GQL boilerplate
# Prerequisite
To contribute to this project you need to prepare this requirement in your environment.

* Docker
* Golang 1.16 or latest version
* gqlgen
* Set GOPRIVATE with this command `go env -w GOPRIVATE=github.com/kecci`

# How to run
First, you need to initialize all requirements with a run this command.
```
make init
```
This command will help you to set up everything you need before development. After that, you can run this command to start the services:

```
make services-up
```

For helping your quick development, this project equipped with a hot-reload using [air](https://github.com/cosmtrek/air). So, you do not need to restart the service manually.

To shut down the services, run this command:
```
make services-down
```

If you working with a database and need to summon a table or add/delete/update table and seeding data in your DB, you need to run this command to clean all cache in the docker file:
```
make clear-postgres-docker-data
```
This command will clean all directory cache in docker related with database.

If you want to generate a mocking file, run this command
```
make mocks
```

If you facing issue like this `fatal: could not read Username for "https://github.com": terminal prompts disabled`, run this command on your terminal
```
env GIT_TERMINAL_PROMPT=1
```
This command will allow you to input username and password during download the package from github.com

To install all package after fetch/pull the repository, use this command to download the latest packages in `go.mod`
```
make download
```

If you facing issue like this `open ./.dev/.gql.air.toml: no such file or directory`, check your dockerfile:
```
ENV GROUP={your-group}
ENV SERVICE={your-service}

WORKDIR {your-git}
```

# How to develop
This project build by reference [clean architecture golang uncle bob](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html). 

## Core Concept
```
┌───────────────┐   ┌───┐   ┌───┐
|   resolver    |   |   |   |   |
└───────────────┘   |   |   |   |
┌───────────────┐   | M |   |   |
|   service     |   | O |   | P |
└───────────────┘   | D |   | K |
┌───────────────┐   | E |   | G |
|     repo      |   | L |   |   |
└───────────────┘   |   |   |   |
┌───────────────┐   |   |   |   |
|  Data Source  |   |   |   |   |
└───────────────┘   └───┘   └───┘
```
- *Resolver* use as the handler
- *Service* use as use-case
- *Repo* use as a repository
- *Data Source* is a Database or 3rd-party (API/GRPC/etc.) or resources(Redis, ES, etc.)
- *Model* is a structure business logic
- *PKG* is a package library or utilities

## Structure Project
```
internal
├── config
|   └── ...
├── model
│   └── ...
├── repo
│   └── ...
├── resolver
|   └── ...
└── service
    └── ...
```
## How to create query GraphQL
TBD

## Working with Database in a local environment
TBD


# Contribution
Here the points you need to know before contributing:

* Unit Test is essential. Whatever you code, you should always test your code and create the testing unit. Nothing is impossible.
* Clean code is everything. This project is already implemented with linter, but you need to make sure your code is clean.
* Documentation. Don't forget to documenting your task. You can use anything documentation like a [Confluence](https://kecci.atlassian.net/wiki/spaces/EN/overview). Don't think your code is the best documentation. That's a stupid mindset.
* Follow the standard. We do not care about your style code. In this project, you should follow the standard and style what we made.
* Idea? Yes, we need your idea to improve the quality of this project and cycle development for more efficiently.
* Finally, let's code as a team, not as an individual.