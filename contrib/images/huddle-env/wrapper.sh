#!/usr/bin/env sh

BINARY=/huddle/${BINARY:-huddle}
ID=${ID:-0}
LOG=${LOG:-huddle.log}

if ! [ -f "${BINARY}" ]; then
	echo "The binary $(basename "${BINARY}") cannot be found. Please add the binary to the shared folder. Please use the BINARY environment variable if the name of the binary is not 'huddle'"
	exit 1
fi

BINARY_CHECK="$(file "$BINARY" | grep 'ELF 64-bit LSB executable, x86-64')"

if [ -z "${BINARY_CHECK}" ]; then
	echo "Binary needs to be OS linux, ARCH amd64"
	exit 1
fi

export HUDDLEDHOME="/huddle/node${ID}/huddle"

if [ -d "$(dirname "${HUDDLEDHOME}"/"${LOG}")" ]; then
  "${BINARY}" --home "${HUDDLEDHOME}" "$@" | tee "${HUDDLEDHOME}/${LOG}"
else
  "${BINARY}" --home "${HUDDLEDHOME}" "$@"
fi
