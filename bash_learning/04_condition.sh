#! /bin/bash

echo " welcome to the digital game win or die"

rand=$(( $RANDOM % 2 ))

echo " chose a number between (1/0)"
read num

if [[ rand == num ]]; then
	echo " congratulation you are win "
else
	echo " you are died"
fi
