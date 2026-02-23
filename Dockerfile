FROM golang:1.26-alpine AS build

WORKDIR /src

COPY . .

RUN apk add build-base \
    && go build -ldflags "-linkmode external -extldflags -static" -o ipsv -a .

FROM scratch

COPY --from=build /src/ipsv /ipsv

EXPOSE 80

CMD ["/ipsv"]
