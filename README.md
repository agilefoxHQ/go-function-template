# Go Function

This function was generated using the `agilefoxHQ/go-template-function repository. It is a Google Cloud Function written in Go 1.19 with some bootstrapping code provided for a sample deploy pipeline.

It also has a main package for local testing using Google Functions Framework.

The main function in this template is an HTTP function, please refer to https://cloud.google.com/functions/docs/concepts/events-triggers to adapt your function.


### Changes to make

 - [ ] Change function name
 - [ ] Change the workflow yaml files to reflect the right flag values in the `Deploy to GCP Functions` step.
 - [ ] Update README

## Run local environment:

`export $(grep -v '^#' .env.dev | xargs) && go run cmd/main.go`

## Files:

### `service.go`

This contains all the interfaces and struct this application defines. Since this
application is a "service" in our world, this is service.go. For a cloud
function, a cloud function running for this purpose is function.go.

## Directories:

We follow the project structure from https://github.com/golang-standards/project-layout and the
handler-service-respository pattern described in https://chidiwilliams.com/post/writing-cleaner-go-web-servers/. I
have over time started writing code that kind of looks like what these articles talk about.

### `/cmd`

Main applications for this project. We keep it minimal, only initialising clients from inner packages.
This usually has one small main function that imports and invokes the code from the `/handler` and initialises the
stores and APIs in `/repositories` directories and the service in `/service`.

### `/service`

The way I define it, it's the business logic of this application. This takes a repo that should satisfy the Repo
interfaces defined in `service.go`.

### `/repository`

These are the data stores, external APIs and all data handling parts of the code. These aren't used directly but via
the service and these satisfy some interface that the services would use.

### `/.gcloud`

This directory contains Google Cloud artifacts. Example: the service definition on Cloud
Run or Kubernetes helm charts or GAE application.yaml. This allows the service deployable
to a multitude of services. We could change the Dockerfile, Makefile and Github actions to
load the definitions from this path.
