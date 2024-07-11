#!/bin/bash
trap "rm server;kill 0" EXIT

go build -o server
./server -port=8333 &
./server -port=8334 &
./server -port=8335 -api=1 &

# sleep 2
# echo ">>> start test"
# curl -x "" "http://localhost:8300/api?member=scores&key=Tom" -s -w '\n' &
# curl -x "" "http://localhost:8300/api?member=scores&key=Tom" -s -w '\n' &
# curl -x "" "http://localhost:8300/api?member=scores&key=Tom" -s -w '\n' &
# curl -x "" "http://localhost:8300/api?member=scores&key=Tom" -s -w '\n'

# sleep 2
# curl -x "" "http://localhost:8300/api?member=scores&key=Tom" -s -w '\n' &
# curl -x "" "http://localhost:8300/api?member=scores&key=Tom" -s -w '\n' &
# curl -x "" "http://localhost:8300/api?member=scores&key=Tom" -s -w '\n' &
# curl -x "" "http://localhost:8300/api?member=scores&key=Tom" -s -w '\n'


wait