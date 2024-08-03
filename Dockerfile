#starts with a base image
FROM golang:1.22.5 as base 

#setting workdir inside the container 
WORKDIR /app

#copying the go.mod 
COPY go.mod ./

#downloading all the dependencies required
RUN go mod download

#copying all the source code to Workdir
COPY . .

#Building the application
RUN go build -o main .

############################################################
#multi-stage builds used to reduce image size 

FROM gcr.io/distroless/base

#copy binary from previous stage i.e base
COPY --from=base /app/main .

#copy the static files from previous stage
COPY --from=base /app/static_files ./static_files

#Expose on which application would run
EXPOSE 8080

#command to run application
CMD [ "./main" ]
