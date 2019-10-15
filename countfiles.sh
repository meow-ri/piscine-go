#!/bin/bash

a="$(ls -lR | egrep -c '^-|^d')"
((a++))
echo $a
