#!/bin/sh
set -e

# Disabling nginx daemon mode
export KONG_NGINX_DAEMON=off

# shellcheck disable=SC2002
cat usr/local/share/kong-template.yml | sed 's/"/\\"/g' | sed 's/^/echo "/g' | sed 's/$/"/g' | sh > usr/local/share/kong.yml

# shellcheck disable=SC2039
if [ "$1" == "kong" ]; then
  PREFIX=${KONG_PREFIX:=/usr/local/kong}

  if [ "$2" == "start" ]; then
    kong prepare -p "$PREFIX" --v

    exec /usr/local/openresty/nginx/sbin/nginx \
      -p "$PREFIX" \
      -c nginx.conf
  fi
fi

exec "$@"
