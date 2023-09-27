#!/usr/bin/env bash

VERSION=v0.2.0

if [[ ! -f ~/.demo-magic.sh ]]; then
    curl -fsSL https://raw.githubusercontent.com/paxtonhare/demo-magic/master/demo-magic.sh > ~/.demo-magic.sh
fi

rm -rf ./demo
mkdir -p ./demo

DIR="$(pwd)"

# subshell execution
(
cd demo

source ~/.demo-magic.sh

TYPE_SPEED=30
PROMPT_TIMEOUT=1
DEMO_PROMPT="${GREEN}\$${COLOR_RESET} "

pe "ls -l"

p "go run github.com/doomshrine/gocosi/cmd/bootstrap@${VERSION} -module example.com/your/cosi-osp -dir cosi-osp"
go run "${DIR}/cmd/bootstrap" -module example.com/your/cosi-osp -dir cosi-osp

pe "tree -a cosi-osp"
)
