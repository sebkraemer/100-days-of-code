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
