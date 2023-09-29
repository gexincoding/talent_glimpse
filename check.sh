#!/bin/zsh
if lsof -i :23462 >/dev/null; then
    pid=$(lsof -i :23462 | grep LISTEN | awk '{print $2}')
    echo "already have process on port 23462, pid: $pid"
else
    echo "no process on port 23462"
fi