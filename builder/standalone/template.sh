#!/bin/sh
set -e

SCRIPT_DIR=$(cd "$(dirname "$0")" && pwd)
cd "$SCRIPT_DIR"

PYTHON="{{ .PYTHON }}"
chmod +x "./python/$PYTHON"

if [ -d "./__pip_install__/" ]; then
	"./python/$PYTHON" -m ensurepip

	for whl in "./__pip_install__/"*.whl; do
		if [ -f "$whl" ]; then
			"./python/$PYTHON" -m pip install --no-cache-dir --no-index --find-links="./__pip_install__/" "$whl"
		fi
	done

	rm -rf "./__pip_install__/"
fi

"./python/$PYTHON" {{ .RUN }} "$@"
