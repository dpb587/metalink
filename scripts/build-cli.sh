#!/bin/bash

set -euo pipefail

outdir="${1}"
version_number="${EXPORTS_VERSION_NUMBER:-0.0.0}"

function build() {
  pkgpath="${1}"
  cliname="${2}"
  target="${3}"

  target_os="$( cut -d- -f1 <<< "${target}" )"
  target_arch="$( cut -d- -f2 <<< "${target}" )"

  cmd="${cliname}-${version_number}-${target_os}-${target_arch}"

  if [ "${target_os}" == "windows" ]
  then
    cmd="${cmd}.exe"
  fi

  echo "build/cmd: os=${target_os} arch=${target_arch} cmd=${cmd}" >&2

  GOOS="${target_os}" GOARCH="${target_arch}" go build \
    -o "${outdir}/${cmd}" \
    "${pkgpath}"
}

for target in darwin-amd64 darwin-arm64 linux-amd64 windows-amd64
do
  build ./cli meta4 "${target}"
  build ./repository/cli meta4-repo "${target}"
done
