#!/bin/sh
set -e

{
    echo "USER_UID=$(id -u)";
    echo "USER_GID=$(id -g)";
    echo "USER=$(id -un)";
} > $(dirname "$0")/../../docker/.env
