#!/usr/bin/env bash
set -e


d_vsrc='futurequest.net'
d_tmp=$(\mktemp -d)

trap "\rmdir ${d_tmp}" EXIT

echo "Move: vendor/${d_vsrc} => ${d_tmp}/${d_vsrc}"
\mv "vendor/${d_vsrc}" "${d_tmp}"
glide up
echo "Move: ${d_tmp}/${d_vsrc} => vendor/${d_vsrc}"
\mv "${d_tmp}/${d_vsrc}" "vendor/${d_vsrc}"
