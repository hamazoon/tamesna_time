#! /bin/bash

echo "what is your name"

read name
greating=$1

user=$(whoami)
date=$(date)
whereami=$(pwd)

echo "$greating morning $name"
sleep 2
echo "have a nice day $name"
sleep 2
echo "$greating bye $name"
sleep 3

echo "you are currently logging in as $user and you are in the directory $whereami . also now is: $date"

