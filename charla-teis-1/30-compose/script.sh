#!/bin/bash
while true
do
    curl -s server:8090/hello
    sleep 1
done
