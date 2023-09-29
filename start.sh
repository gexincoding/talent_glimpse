#!/bin/zsh

RUN_NAME="talent.glimpse"

rm -rf output

mkdir -p output output/conf

find conf -type f | xargs -I{} cp {} ./output/conf/

go mod tidy

go build -o ./output/${RUN_NAME} main.go

cd ./output

chmod +x ${RUN_NAME}

# 检查端口23462是否有进程
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

# 后台运行可执行文件abc，并将日志输出到aaa.log文件
nohup ./${RUN_NAME} > ${RUN_NAME}.log 2>&1 &

echo "[Success] Running Success! "