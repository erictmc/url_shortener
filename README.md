[![Build Status](https://www.travis-ci.com/erictmc/url_shortener.svg?branch=master)](https://www.travis-ci.com/erictmc/url_shortener)

##Description 

This is a basic proof of concept for a Url Shortener App. It has
the following components:
 - A Golang-based API, backed by a PostGres DB
 - A React-based web-app, which is served by the API.

## Setup 

Currently, there are two configurations for running 
the app, **Local Dev Environment** and **Test Environment** 

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
hot-reloads available. Any changes, you make -- either to the server or 
client-side code -- WILL NOT BE VISIBLE to the docker container. To 
have hot-reloads enabled, see the *Local Development* environment below.

Note: this environment is also used in the CI pipeline.

### Local Development 
This prioritizes using hot reloads for the api and web code. In addition to Docker,
you will need both `node` and `yarn` installed on your machine.

*Running the Web App*: 

The following commands will start the webapp on `localhost:3000.`
Note it will still communicate with the server running on `8080` via a proxy.

**IMPORTANT** If you go to localhost:8080, you will not see
your most recent changes. Instead, go to localhost:3000.

Run the following in a new shell:
```
cd ./api/web
yarn install
yarn start
```

In your previous shell, which should still be in the path `url_shortener`,
running the following command: `docker-compose -f docker-compose.local-dev.yml up`

This will start the api in a `local_development` mode, with hot-reloads available.

### Testing and CI/CD

For testing, the "Test Environment" is used. The testing configuration is specified
in `.travis.yml`. 

There are some node-based tests in the `./testing` directory. To run these,
you will need to run `yarn install.` These will also be run the CI pipeline.