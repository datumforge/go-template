FROM golang:1.22.6 AS builder

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -o /go/bin/template -a -ldflags '-linkmode external -extldflags "-static"' .

FROM gcr.io/distroless/static:nonroot

# `nonroot` coming from distroless
USER 65532:65532

# Copy the binary that goreleaser built
COPY --from=builder /go/bin/template /bin/template

# Run the web service on container startup.
ENTRYPOINT [ "/bin/template" ]
CMD ["serve"]
