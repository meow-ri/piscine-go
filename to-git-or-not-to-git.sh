#! /bin/bash

curl -s https://api.github.com/users/meow-ri | jq '.id'
