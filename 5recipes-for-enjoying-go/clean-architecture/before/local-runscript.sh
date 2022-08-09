#!/bin/sh

export $(cat .env.local | grep -v ^# | xargs)

go run .
