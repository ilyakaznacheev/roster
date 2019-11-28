// Package roster
/*
This tiny service was built to manage rosters with very limited abilities. Check swagger.yml for more information.

Tools, libraries and approaches were chosen with a focus on production readiness.

Tools and Technologies

There are some technical decisions made to bring more flexibility and scalability to the service. It may look overcomplicated, but there is extensibility and reliability behind the simplicity.

There are no frameworks due to simple infrastructure.

- MongoDB - chosen because it's easy to shard and has enough search engine power to fulfill service needs;
- Swagger - perfectly describes REST APIs and can be used to generate server and client API-code in a single command;
- JWT auth - used to secure changes via API. Used because it's a common choice for API auth.
- Docker - used to satisfy requirements and build and run the service in isolated environment. Multistage build is made to make resulted image as small as possible. Can be used on any cloud environment or as a part of orchestration systems like k8s;
- Heroku - chosen as a simplest hosting with CI/CD features because it is easy to setup and monitor.

Tests

All core logic (request handlers) is covered with unit-tests. No need to check infrastructure and generated core, because it doesn't change so much, but really hard to test. DB and web layers are isolated.

Production readiness

The service is ready to run as a standalone microservice in any environment, e.g. k8s.

It is built with a closer look to 12-factor app principle, and can be easily integrated into 12-f infrastructure with small changes or additions.


To run the service as a library you need to run the Run function. It handles termination, so no need to set extra context.

	var cfg config.Application
	// fill the config
	err := roster.Run(cfg)
	if err != nil {
		os.Exit(1)
	}

To run from command-line just call

	go run cmd/roster/main.go
*/
package roster
