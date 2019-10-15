#!/bin/bash
curl -s https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json | jq --argjson HID "$HERO_ID" '.[] | select (.id == $HID) .connections.relatives' | tr -d '"'

