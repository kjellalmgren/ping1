# start from hypriot/rpi-alpine-scratch (nginx:alpine)
#
# -------------------------------------------------
FROM resin/rpi-raspbian
 MAINTAINER kjell.almgren@tetracon.se
 ADD ping1 /ping1
 CMD ["./ping1"] 
#
# -------------------------------------------------