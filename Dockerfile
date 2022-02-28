FROM golang:1.17

WORKDIR /app

CMD ["go","run","."]

#build
#docker build -t my-golang-app . 
#launch
#docker run -v "$PWD":/app my-golang-app