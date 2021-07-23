# base image for our application
# this wil pull docker image from docker hub
FROM golang

# copy root dir to /app/bookapi in docker image
COPY . /app/bookapi

# setting work directory, so all the subsequent commands can be run on this dir
WORKDIR /app/bookapi

# build bookapi code and set it's output binary in dir as /out/booapi
RUN go build -o ./out/bookapi .

# command to execute binaries
CMD ["./out/bookapi"]