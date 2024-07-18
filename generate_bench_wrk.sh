#!/bin/bash

# 1. Authentication (Same as before)
export TOKEN=$(curl -s -X "POST" "http://localhost:12345/auth" \
    -H 'Accept: application/json' \
    -H 'Content-Type: application/json' \
    -d $'{
        "email": "eminetto@gmail.com",
        "password": "1234567"
    }' | jq -r .token)

# 2. Load Test with wrk (Simplified)
wrk -t12 -c400 -d30s \
    -H 'Accept: application/json' \
    -H 'Content-Type: application/json' \
    -H "Authorization: $TOKEN" \
    -s feedback_data.lua \
    http://localhost:12345/feedback

wrk -t12 -c400 -d30s \
    -H 'Accept: application/json' \
    -H 'Content-Type: application/json' \
    -H "Authorization: $TOKEN" \
    -s vote_data.lua \
    http://localhost:12345/vote
