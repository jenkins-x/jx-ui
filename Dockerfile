FROM alpine:3.16.2
COPY ./build/linux/jx-ui /usr/bin/jx-ui
COPY ./web/build web/build
ENTRYPOINT ["jx-ui"]
