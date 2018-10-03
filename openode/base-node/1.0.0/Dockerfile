# base docker for node

FROM ubuntu:16.04

# so that we can source
RUN rm /bin/sh && ln -s /bin/bash /bin/sh

RUN apt update --yes
RUN apt upgrade --yes

RUN apt install curl --yes

ENV NVM_DIR /usr/local/nvm

RUN curl -o- https://raw.githubusercontent.com/creationix/nvm/v0.33.11/install.sh | bash

RUN apt autoclean --yes
