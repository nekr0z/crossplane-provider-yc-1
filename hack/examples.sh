#!/usr/bin/env bash
set -aeuo pipefail

function join {
  local f=${1-}
  if shift; then
    printf %s "$f" "${@/#/,}"
  fi
}

all=$(find ${1} -name "*.yaml" -not -path "*/mdb/*" -not -path "*/datatransfer/*" -not -path "*/alb/*" -not -path "*/ydb/*")


join $all

