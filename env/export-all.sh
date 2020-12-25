#!/usr/bin/env bash

export $(cat $1 | grep -v ^# | xargs)
