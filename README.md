#Build a container for Raspberry 3B+

    # build with tag ping1, -t equals tag
    $ docker build --file Dockerfile.builder -t ping1 .
    # push to hub.docker.com, assume docker login
    $ push <respository>/<container>:<tag>
    # run container
    $ docker run --name ping1 -t ping1 .
