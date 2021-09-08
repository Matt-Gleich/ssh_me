FROM golang:1.17

# Meta data:
LABEL maintainer="email@mattglei.ch"
LABEL description="👋 Meet me via ssh!"

# Copying over all the files:
COPY . /usr/src/app
WORKDIR /usr/src/app

CMD ["make", "local-test"]
