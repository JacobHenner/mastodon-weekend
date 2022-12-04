FROM golang:1.19-bullseye AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY cmd cmd

RUN go mod download

RUN cd cmd && go build -o /mastodon-weekend

FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /mastodon-weekend /mastodon-weekend

ENTRYPOINT [ "/mastodon-weekend" ]