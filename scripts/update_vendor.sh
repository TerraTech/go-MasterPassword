#!/usr/bin/env bash
set -e

prog=$(basename $0)
dryrun=

if [[ "${prog}" = 'update_vendor-dryrun.sh' ]]; then
	dryrun='--dry-run'
fi

d_src="$1"
d_dst="${2:+/$2}"
shift; shift
subrepos=("$@") #implement at later time when really needed

if [[ ! -d "${d_src}" ]]; then
	echo "[FATAL] Source repo does not exist: ${d_src}"
	exit 1
fi

if [[ ! -d "${d_dst}" ]]; then
	mkdir -p "./vendor${d_dst}"
fi

\rsync -avip ${dryrun} \
	--delete \
	--exclude=.git \
	"${d_src}" \
	"./vendor${d_dst}"
