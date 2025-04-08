FROM golang:1.23.6-bullseye AS air
WORKDIR /app
RUN go install github.com/air-verse/air@latest
ENV PATH="/go/bin:${PATH}"
CMD ["air"]
