# Microproxy
Testing echo framework

# Objective
Have an / with all static content and proxy pass to /api endpoint

## Running
`docker run -ti -e PROXY=localhost:8081 -e ADDRESS=0.0.0.0:8080 -p 8080:8080 microproxy`
