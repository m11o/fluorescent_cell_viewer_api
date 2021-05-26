FROM golang:latest

ENV WORKSPACE /app
WORKDIR $WORKSPACE
ADD . $WORKSPACE

EXPOSE 8080
CMD ["go", "run", "main.go"]
