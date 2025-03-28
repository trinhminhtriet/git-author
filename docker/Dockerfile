FROM golang:1.24-alpine AS builder

RUN mkdir /app
ADD . /app

RUN cd /app && go build -a -o git-author -buildvcs=false -ldflags '-s -w'

FROM alpine/git
RUN mkdir /app && git config --system --add safe.directory /git && git config --system alias.author '!/app/git-author'
COPY --from=builder /app/git-author /app/git-author
