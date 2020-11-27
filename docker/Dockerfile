FROM alpine:latest

RUN apk update && apk add --no-cache libc6-compat
COPY ./recipe /go/src/

CMD ["/go/src/recipe"]