#!/usr/bin/env bash

fatal() {
	echo "[FATAL] $@"
	exit 1
}

_f=

if [[ -z "${CLI}" ]]; then
	fatal '$CLI is not defined'
fi

if [[ -f "bin/${CLI}" ]]; then
	\rm -v "bin/${CLI}"
fi

_f="${GOBIN}/${CLI}"
if [[ -n "${GOBIN}" && -f "${_f}" ]]; then
	\rm -v "${_f}"
fi

_f="${GOPATH}/bin/${CLI}"
if [[ -f "${_f}" ]]; then
	\rm -v "${_f}"
fi
