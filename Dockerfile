# build stage
FROM golang:alpine AS build-env

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn

#ENV HTTP_PROXY "http://devnet-proxy.oa.com:8080"
#ENV HTTPS_PROXY "http://devnet-proxy.oa.com:8080"
#ENV NO_PROXY "*.oa.com"


ADD . /src
RUN cd /src && \
	go env -w GOPROXY=https://goproxy.cn && \
	go build -o main main.go

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/main /app/main
#COPY --from=build-env /src/app_conf.yml /app/app_conf.yml
#COPY --from=build-env /src/casbinmodel.conf /app/casbinmodel.conf
#RUN mkdir -p /data/applogs
#VOLUME /data/applogs

#RUN apk add --no-cache tzdata busybox-extras bash bash-completion\
#    && ln -snf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
#    && echo "Asia/Shanghai" > /etc/timezone && \
#    unset HTTP_PROXY &&  unset HTTPS_PROXY  && unset NO_PROXY
#
#ENV TZ Asia/Shanghai

ENV GO_ENV "production"

EXPOSE 10080

ENTRYPOINT ./main
