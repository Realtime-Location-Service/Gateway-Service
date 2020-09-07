FROM ubuntu:18.04

RUN apt-get update \
    && apt-get install -y apt-transport-https curl lsb-core apt-utils \
    && echo "deb https://kong.bintray.com/kong-deb `lsb_release -sc` main" | tee -a /etc/apt/sources.list \
    && curl -o bintray.key https://bintray.com/user/downloadSubjectPublicKey?username=bintray \
    && apt-key add bintray.key

RUN apt-get update
RUN apt-get install -y kong

RUN apt-get install -y luarocks libssl-dev
RUN luarocks install lua-resty-template

COPY plugins /usr/local/share/lua/5.1/kong/plugins
COPY kong-template.yml /usr/local/share/kong-template.yml
COPY docker-entrypoint.sh /docker-entrypoint.sh

ENTRYPOINT ["/docker-entrypoint.sh"]

STOPSIGNAL SIGQUIT

CMD ["kong", "start"]
