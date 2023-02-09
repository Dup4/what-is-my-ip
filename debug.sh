#! /bin/bash

CUR_DIR="$(dirname "$(realpath "${BASH_SOURCE[0]}")")"

make
"${CUR_DIR}"/output/what-is-my-ip
