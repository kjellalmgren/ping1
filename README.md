#Build a container for Raspberry 3B+

    Check external availability from within a linux container to the outside world. Sone dms problem in docker 17.05ce has been reported.
    Container is bild from resin/rpi-raspbian.
    
    # build with tag ping1, -t equals tag
    $ docker build --file Dockerfile.builder -t ping1 .
    # push to hub.docker.com, assumes docker login
    $ push <respository>/<container>:<tag>
    # run container
    $ docker run --name ping1 -t ping1 .
