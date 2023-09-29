#!/bin/zsh
if lsof -i :23462 >/dev/null; then

    pid=$(lsof -i :23462 | grep LISTEN | awk '{print $2}')
    echo "[Warn] already have process on port 23462, ready to kill, pid: $pid"

    # 强制杀死进程
    echo "killing process, pid: $pid"
    kill -9 "$pid"
    echo "process killed, pid: $pid"
else
    echo "no process on port 23462, ready to run"
fi