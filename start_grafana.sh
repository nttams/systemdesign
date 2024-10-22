#!/bin/bash

docker run -d -p 3000:3000 \
    --restart always \
    --name=grafana \
    --volume "/Users/tamngo/personal/data/grafana:/var/lib/grafana" \
    grafana/grafana