FROM golang:1.17 AS builder

# just for testing
ARG defaultPort=8080
RUN echo "Setting default env.PORT to ${defaultPort}"

ENV PORT=${defaultPort}
WORKDIR /app/cmd/sudokuserver

ADD . /app/

RUN go mod tidy
RUN go mod verify

EXPOSE $PORT

# will compile and run the program on the running os/arch
# (before implementing multi-arch build)
# CMD ["go", "run", "sudokuserver.go"]
# alternatively, choose another arch and precompile
# ENV GOOS=$MYOS GOARCH=$MYARCH
# RUN go build -o sudokuserver ./..
# CMD [ ./sudokuserver ]

# now implementing multistage builds
# RUN go build sudokuserver.go
# use static linking
RUN CGO_ENABLED=0 go build sudokuserver.go

# second build stage
#FROM alpine:latest
FROM scratch

WORKDIR /app/sudokuserver

COPY --from=builder /app/cmd/sudokuserver /app/sudokuserver

CMD ["./sudokuserver"]
