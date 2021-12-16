FROM golang:1.17

ENV PORT 8080
WORKDIR /app/cmd/sudokuserver

ADD . /app/

RUN go mod tidy
RUN go mod verify

EXPOSE $PORT

# will compile and run the program on the running os/arch
# (before implementing multi-arch build)
CMD ["go", "run", "sudokuserver.go"]

# alternatively, choose another arch and precompile
# ENV GOOS=$MYOS GOARCH=$MYARCH
# RUN go build -o sudokuserver ./..
# CMD [ ./sudokuserver ]
