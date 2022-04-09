# FROM golang:1.17-alpine

FROM golang:1.17.0-bullseye

# enable go module
ENV GO111MODULE=on
WORKDIR /app/server

# set timezone
# RUN apk update && apk add tzdata
# ENV TZ=America/New_York
# RUN apk add --update nodejs npm

# RUN apk add gcc gtk+3.0-dev webkit2gtk-dev musl-dev
# RUN apk add --no-cache bash


RUN apt-get update
RUN apt-get -y install build-essential
RUN apt-get -y install libgtk-3-dev
# RUN echo "deb http://cz.archive.ubuntu.com/ubuntu bionic main universe" >> /etc/apt/sources.list
RUN apt-get update
# RUN apt-get -y install libwebkitgtk-1.0-0
RUN apt-get -y install libwebkit2gtk-4.0-dev
RUN apt-get -y install curl
RUN curl -sL https://deb.nodesource.com/setup_17.x | bash
RUN apt-get -y install nodejs
RUN apt-get -y install libcanberra-gtk-module
RUN apt-get -y install libcanberra-gtk3-module

#RUN     apt-get install -y x11vnc xvfb
#RUN     mkdir ~/.vnc
#RUN     x11vnc -storepasswd 1234 ~/.vnc/passwd

RUN apt-get update && apt-get install -y \
      x11-apps \
      && rm -rf /usr/share/doc/* && \
      rm -rf /usr/share/info/* && \
      rm -rf /tmp/* && \
      rm -rf /var/tmp/*

#RUN useradd -ms /bin/bash user
#USER user

#ENV DISPLAY :0

COPY . .

EXPOSE 18200 18606

RUN go install github.com/wailsapp/wails/v2/cmd/wails@latest
RUN wails build -platform linux/amd64

#RUN go run go-server.go
CMD ["./build/bin/databridge"]
