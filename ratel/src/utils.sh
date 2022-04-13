#!/usr/bin/env bash

kill_python() {
  pkill -f python || true
}

kill_mpc() {
  pkill -f ./malicious-shamir-party.x || true
  pkill -f ./random-shamir.x || true
}

clear_db() {
  rm -rf /opt/hbswap/db/*
}

refill() {
  players=$1
  token_id=0
  for (( server_id = 0; server_id < $players; server_id++ )) do
    python3 -m ratel.src.python.refill server_$server_id $token_id
  done
}

create_ids() {
  init_players=$1
  ids='0'
  for ((server_id = 1; server_id < $init_players; server_id++ )) do
    ids="${ids},${server_id}"
  done
  echo $ids
}