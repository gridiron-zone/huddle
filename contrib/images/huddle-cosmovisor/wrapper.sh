#!/usr/bin/env sh

BINARY=cosmovisor
ID=${ID:-0}
LOG=${LOG:-huddle.log}

export HUDDLEDHOME="/huddle/node${ID}/huddle"

# Set environment variables
export DAEMON_NAME=huddle
export DAEMON_HOME="$HUDDLEDHOME"
export DAEMON_ALLOW_DOWNLOAD_BINARIES=true
export DAEMON_RESTART_AFTER_UPGRADE=true

# Setup Cosmovisor
COSMOVISOR_GENESIS="$HUDDLEDHOME/cosmovisor/genesis/bin"
if [ ! -d "$COSMOVISOR_GENESIS" ]; then
  mkdir -p $COSMOVISOR_GENESIS
  cp $(which huddle) "$COSMOVISOR_GENESIS/huddle"
fi

# Run the command
if [ -d "$(dirname "${HUDDLEDHOME}"/"${LOG}")" ]; then
  "${BINARY}" "$@" --home "${HUDDLEDHOME}" | tee "${HUDDLEDHOME}/${LOG}"
else
  "${BINARY}" "$@" --home "${HUDDLEDHOME}"
fi