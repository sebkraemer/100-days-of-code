# Sudoku server, solving sudoku riddles


## Docker

Inspired by "Developing Go Apps with Docker" @pluralsight

Build docker image:
`docker build -t sudokuserver:latest -f Dockerfile ../..`

Run it on defualt port:
`docker run -d --rm --name sudoku -p 8080:8080 sudokuserver`

Run it on port `8082` but expose it to the host (e.g.. for the browser) on `8083`:
`docker run  -ti --rm --name sudoku --env PORT=8082 -p 8083:8082 sudokuserver`

(Would have loved to continue with `docker buildx` but that was not available on my system, despite setting the experimental flag, and it was not important enough for me. Maybe you want to check out multi-platform build! https://www.docker.com/blog/multi-platform-docker-builds/)


## Minikube

Some notes about using Minikube


### exposing 0.0.0.0
.. is probably 'doing it wrong' makes the started services available without specifying k8s deployments

```
minikube start --driver=hyperkit --listen-address=0.0.0.0
eval $(minikube docker-env) 
```

A `docker login` might be needed for pulling base images.

### "local access"

With `Wminikube ssh` it's possible to see the services' ports that are not exported to outside minikube.
I used it for easy test of the webserver to create logs, e.g. `curl`ing a URL.

`ifconfig` will show me information about the networks so that I can also access them outside the ssh shell, e.g. dockerd interface. `eth0` gives me the address to connect from outside, e.g. my browser

The same ip is also visible from the docker-env since that is also where the docker daemon is running.

### Misc

- save some power: `minikube pause`
- get back to work! `minikube unpause`
