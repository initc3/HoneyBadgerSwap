#!/bin/bash

if [ -z $TMUX ]; then
    echo "tmux is not active, will start new session"
    TMUX_CMD="new-session"
else
    echo "tmux is active, will launch into new window"
    TMUX_CMD="new-window"
fi

tmux $TMUX_CMD "docker-compose logs -f eth.chain; sh" \; \
    splitw -v -p 60 "docker-compose logs -f mpcnode0; sh" \; \
    splitw -v -p 60 "docker-compose logs -f mpcnode1; sh" \; \
    splitw -v -p 60 "docker-compose logs -f mpcnode2; sh" \; \
    splitw -v -p 60 "docker-compose logs -f mpcnode3; sh" \; \
    selectp -t 0 \; \
    splitw -h -p 60 "docker-compose logs -f contract.deploy; sh" \; \
    splitw -h -p 60 "docker-compose logs -f contract.deposit; sh"
