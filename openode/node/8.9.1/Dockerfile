# node specific docker
FROM openode/base-node:latest


ENV NVM_DIR /usr/local/nvm
ENV NODE_VERSION 8.9.1

RUN source /usr/local/nvm/nvm.sh       && nvm install 8.9.1       && nvm alias default 8.9.1       && nvm use default

ENV NODE_PATH /usr/local/nvm/v8.9.1/lib/node_modules
ENV PATH /usr/local/nvm/versions/node/v8.9.1/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games

RUN npm install -g pm2

RUN node -v
RUN npm -v
