# syntax=docker/dockerfile:experimental

FROM alpine:latest

RUN apk update && apk upgrade

RUN apk add openssh

# owner can read, can write and can execute.
# (G)roup can't read, can't write and can't execute. 
# (O)thers can't read, can't write and can't execute.
# === chmod a+rwx,g-rwx,o-rwx
RUN mkdir -m 700 $HOME/.ssh

RUN touch -m $HOME/.ssh/known_hosts
RUN ssh-keyscan github.com >> $HOME/.ssh/known_hosts

RUN --mount=type=ssh \
    ssh -vT git@github.com || true