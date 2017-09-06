#!/usr/bin/env bash

fatal() {
	echo "[FATAL] $@"
	exit 1
}

_f=

if [[ -z "${CLI}" ]]; then
	fatal '$CLI is not defined'
fi

if [[ -z "${GOPATH}" ]]; then
	fatal '$GOPATH is not defined'
fi

if [[ -z "${GOBIN}" ]]; then
	GOBIN="${GOPATH}/bin"
fi

_f="bin/${CLI}"
if [[ ! -x "${_f}" ]]; then
	fatal "Cannot install, binary does not exist: ${_f}"
fi

\mv -v "${_f}" "${GOBIN}"
\rmdir "bin"
