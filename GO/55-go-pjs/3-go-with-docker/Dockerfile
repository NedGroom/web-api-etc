
	# Original code from video
FROM golang:1.12.0-alpine3.9
# make a directory in the image
RUN mkdir /app
# take all local files into my image app folder
ADD . /app
# workdir is a command to tell docker that this is our main directory to run commands from
WORKDIR /app
RUN go build -o main .
# go into directory and run the executable
CMD ["/app/main"]


# you will need a few more commands to initialise mod files etc
# tried to incorporate this https://docs.docker.com/language/golang/build-images/

# tried this but didnt work
#		FROM golang:1.12.0-alpine3.9
#		WORKDIR /app
#		COPY go.mod ./
#		RUN go mod download
#		COPY *.go ./
#		# RUN go mod init 3-go-with-docker
#		RUN go build -o /3-go-with-docker
#		CMD ["./app/main"]



