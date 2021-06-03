[![Build Status](https://www.travis-ci.com/erictmc/url_shortener.svg?branch=master)](https://www.travis-ci.com/erictmc/url_shortener)

## Description 

This is a basic proof of concept for a Url Shortener App. It has
the following components:
 - A Golang-based API, backed by a PostGres DB
   - The entry point for the API code is the `main` function in [./api/api.go](https://github.com/erictmc/url_shortener/blob/master/api/api.go)
 - A React-based web-app, which is served by the API.
   - Entry point: [./api/web/src/App.tsx](https://github.com/erictmc/url_shortener/blob/master/api/web/src/App.tsx)

## Setup 

Currently, there are two configurations for running 
the app:
 - **Local Dev Environment** -  if want to edit the code
 - **Test Environment** - if you want to test the code locally.

Both require Docker (along with docker-compose) to be available
on one's machine.

### Test Environment
This is run with the following command, while in the root of the repo directory:
`docker-compose -f docker-compose.test.yml up`.

This builds the API code and places it on a Docker image. As part of this,
it also builds the React web app, and places the optimized bundle in a
static directory on the image.  If docker-compose is successful,
then you should be able to visit `http://localhost:8080`,
and see a webapp. 

**IMPORTANT:** If you are running in this environment, it will not have
hot-reloads available. Any changes you make to the code-- either to the server or 
client-side code -- WILL NOT BE VISIBLE to the docker container. To 
have hot-reloads enabled, see the *Local Development* environment below.

Note: The Testing environment is also used in the CI pipeline.

### Local Development 
This prioritizes using hot reloads for the api and web code. In addition to Docker,
you will need both `node` and `yarn` installed on your machine.

In the root path of the repo, run the following command: `docker-compose -f docker-compose.local-dev.yml up`
This will start the api in a `local_development` mode, with hot-reloads available for server code.

Run the following in a new shell:
```
cd ./api/web
yarn install
yarn start
```

The above commands will start the webapp on `localhost:3000.` Note it will still communicate with the server running on `8080` via a proxy.
This will allow hot-reloads for the React web app.

**IMPORTANT** If you go to localhost:8080, you will not see
your most recent changes for the web app. Instead, go to localhost:3000.

### Integration Testing and CI/CD

For testing, the "Test Environment" is used. The `.travis.yml` file specifies testing configuration. Currently,
integration tests have been prioritized over stand-alone unit tests.

There are some node-based integration tests in the `./testing` directory. To run these,
you will need to run `yarn install.` The tests can be run with `yarn test.` You can 
run these tests from the repo root directory via `./run-tests.sh`

**IMPORTANT:** In order for the tests to successfully run, you will need to have the Docker-compose environment
running. Otherwise, they will fail. 

These tests will also be run the CI pipeline.

### Items to Improve Scalability, Maintainability and Security 

TODO:

- [ ] Consistent URL validation between client and server-side code.
- [ ] Frontend: Improve Responsiveness.
- [ ] Frontend: Incorporate modern CSS lib, such as Tailwind or SCSS.
- [ ] Cookie-based auth to track users.
- [ ] Incorporate Swagger to reduce redundant client-side api code.
- [ ] Include Cypress-based tests to improve true e2e coverage.   
- [ ] Look at caching mechanism (Redis/Memcache) for validating users.
- [ ] Look at more scalable/performant mechanism for reading short-urls (such Redis or Cassandra/DynamoDB)