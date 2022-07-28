FROM golang:1.18-alpine AS build

WORKDIR /src/
COPY main.go go.* /src/
COPY ubi/ /src/ubi
RUN CGO_ENABLED=0 go build -o /bin/u2ug

FROM scratch
COPY --from=build /bin/u2ug /bin/u2ug
ENTRYPOINT ["/bin/u2ug"]