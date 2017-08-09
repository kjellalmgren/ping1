# start from hypriot/rpi-alpine-scratch (nginx:alpine)
#
# -------------------------------------------------
#FROM resin/rpi-raspbian
FROM alpine
 MAINTAINER kjell.almgren@tetracon.se
 ADD ping1 /ping1
 ENTRYPOINT ["./ping1"] 
#
# -------------------------------------------------