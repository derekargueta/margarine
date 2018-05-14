#!/usr/bin/env bash

# simple script that accepts a file of new-line-separated paths to rm.

function cleanup() {
    # iterate through the file and delete each path
    cat $2 | while read line; do
        rm "$line"
    done
}

function check() {
    cat $2 | while read; do
        echo $REPLY
        if [ ! -f "$REPLY" ]; then
            echo "no";
        else
            echo "yes";
        fi
    done
}


if [ $# -eq 2 ]; then
    if [ "$1" = "clean" ]; then
        cleanup $@
        exit $?
    elif [ "$1" = "check" ]; then
        check $@
        exit $?
    fi
fi

echo "Unknown command, usage:"
echo "./clean.sh clean FILE"
echo "./clean.sh check FILE"
echo ""
echo "FILE - file to read list of paths from"

exit 1

