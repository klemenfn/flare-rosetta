#!/usr/bin/env bash
pid=0

term_handler() {
  if [ $pid -ne 0 ]; then
    kill -SIGTERM "$pid"
    wait "$pid"
  fi
  exit 143;
}

trap 'kill ${!}; term_handler' SIGTERM SIGINT

./rosetta-server -config=./config/$1.json &
pid="$!"

# wait forever
while true
do
  tail -f /dev/null & wait ${!}
done

