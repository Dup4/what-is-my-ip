#! /bin/bash

set -e -x

if [ X"${1}" = X"primary" ]; then
    exec /app/what-is-my-ip
else
    exec "$@"
fi
