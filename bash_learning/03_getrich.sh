#! /bin/bash

echo "what is your name"
read name 
echo "how old are you"
read age
sleep 2
echo "hello $name ! you are $age years old."

echo "calculating"
sleep 1
echo "*......"
sleep 1
echo "**....."
sleep 1 
echo "***...."
sleep 1
echo "****..."
sleep 1
echo "*****.."
sleep 1
echo "******."
sleep 1
echo "*******"

getrich=$((( $RANDOM % 17) + $age ))
echo "$name you will be rich when you are  $getrich years old "
