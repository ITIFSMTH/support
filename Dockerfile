FROM golang as back-build-stage
WORKDIR /app/back
COPY ./back/ .
RUN go mod tidy
RUN go build -o ./back

FROM ubuntu:latest
ENV TZ=Europe/Moscow
ENV NVM_DIR /root/.nvm
ENV NODE_VERSION 18.2.0

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN apt-get update && apt --yes upgrade && apt --yes install sudo curl git nano nginx postgresql

USER postgres
RUN /etc/init.d/postgresql start &&\
    psql --command "CREATE USER docker WITH SUPERUSER PASSWORD 'docker';" &&\
    createdb -O docker support
RUN ls /etc/postgresql/
RUN echo "host all  all    0.0.0.0/0  md5" >> /etc/postgresql/14/main/pg_hba.conf
RUN echo "listen_addresses='*'" >> /etc/postgresql/14/main/postgresql.conf
VOLUME  ["/etc/postgresql", "/var/log/postgresql", "/var/lib/postgresql"]

USER root
RUN mkdir $HOME/.nvm && \
    curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.37.2/install.sh | bash && \
    chmod +x $HOME/.nvm/nvm.sh && \
    . $HOME/.nvm/nvm.sh && \
    nvm install --latest-npm "$NODE_VERSION" && \
    nvm alias default "$NODE_VERSION" && \
    nvm use default && \
    DEFAULT_NODE_VERSION=$(nvm version default) && \
    ln -sf /root/.nvm/versions/node/$DEFAULT_NODE_VERSION/bin/node /usr/bin/nodejs && \
    ln -sf /root/.nvm/versions/node/$DEFAULT_NODE_VERSION/bin/node /usr/bin/node && \
    ln -sf /root/.nvm/versions/node/$DEFAULT_NODE_VERSION/bin/npm /usr/bin/npm

WORKDIR /app/front
COPY ./front/package*.json ./
RUN npm install
COPY ./front/ .
RUN npm run build

WORKDIR /app/bot
COPY ./bot/package*.json ./
RUN npm install --force
COPY ./bot/ .
RUN npm install pm2 -g

COPY nginx.conf /etc/nginx/nginx.conf
COPY --from=back-build-stage /app/back/back /app/back/back
COPY /back/env /app/back/env/

COPY ./start.sh /
CMD bash /start.sh
