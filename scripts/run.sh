#!/usr/bin/env bash

PROJECT_HOME="$(cd "$(dirname "${0}")/.." && pwd)"
cd "$PROJECT_HOME"

IFS='|' # commands delimiter
read -ra COMMAND <<< $@ # read all commands
for i in ${COMMAND[@]}; do
    echo "Running command: $i"
    # if spinning up agents -> run in background
    if [[ $i == *"start"* ]] || [[ $i == *"extend"* ]]; then
        eval "app/app ${i}" > tmp_out 2>&1 &
        trap 'rm -rf tmp_out' EXIT
        for s in {1..50}
        do
            # wait for agents to be ready, timeout 5s
            if grep -q "ready" tmp_out; then
                break
            fi
            sleep 0.1
        done
        rm tmp_out
    else
        # if not spinning up agents -> run in current shell
        eval "app/app ${i}"
    fi
done
