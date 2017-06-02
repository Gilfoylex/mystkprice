FROM mydocker

MAINTAINER gilfoyle gilfoylex@outlook.com

COPY nodejs/stocks.js /home/nodejs
COPY golang/messenger /home/golang
COPY golang/messenger.go /home/golang
COPY golang/mystk0.html /home/golang
COPY golang/jquery.min.js /home/golang

WORKDIR /home
