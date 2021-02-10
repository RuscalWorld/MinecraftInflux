FROM golang:latest
WORKDIR /home/container
ADD minecraftinflux .
CMD [ "./minecraftinflux" ]
