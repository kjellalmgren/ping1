# Try to minimize each container

Check external availability from within a linux container to the outside world. Sone dms problem in docker 17.05ce has been reported.
Container is bild from scratch.

##Build with go (1.8.3) 


    # build for raspberry PI 3 (64bit)
    $ GOOS=linux GOARCH=arm64 go build -v

##build container

    $ GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-w -s' -a -installsuffix cgo -o ping1
    # build with tag ping1, -t equals tag
    $ docker build --file Dockerfile.builder -t tetracon/ping1:0.1 .
    # build from scratch, -t equals tag
    $ docker build --file Dockerfile.scratch -t tetracon/ping1:0.2 .
    # push to hub.docker.com, assumes docker login
    $ push <respository>/<container>:<tag>
    # run container version 1
    $ docker run --name ping1 -t tetracon/ping1:0.1
    # run container version 2
    $ docker run --name ping1 -t tetracon/ping1:0.2
    # run shell to look into cantainer
    $ docker run -it -t tetracon/ping1:0.2 sh

##Dockerfile.builder used
    # -------------------------------------------------
    FROM resin/rpi-raspbian
    MAINTAINER kjell.almgren@tetracon.se
    ADD ping1 /ping1
    CMD ["./ping1"] 
    #
    # -------------------------------------------------
