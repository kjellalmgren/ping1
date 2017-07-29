#Build a container for Raspberry 3B+

Check external availability from within a linux container to the outside world. Sone dms problem in docker 17.05ce has been reported.
Container is bild from resin/rpi-raspbian.

##Build with go (1.8.3) 


    # build for raspberry PI 3 (64bit)
    $ GOOS=linux GOARCH=arm64 go build -v

##build container

    # build with tag ping1, -t equals tag
    $ docker build --file Dockerfile.builder -t ping1 .
    # push to hub.docker.com, assumes docker login
    $ push <respository>/<container>:<tag>
    # run container
    $ docker run --name ping1 -t ping1 .

##Dockerfile.builder used
    # -------------------------------------------------
    FROM resin/rpi-raspbian
    MAINTAINER kjell.almgren@tetracon.se
    ADD ping1 /ping1
    CMD ["./ping1"] 
    #
    # -------------------------------------------------
