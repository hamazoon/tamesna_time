#! /bin/bash

echo " Welcome to the digital game win or die"

rand=$(( $RANDOM % 2 ))

echo "Chose a number between (1/0)"
read num

if [[ $rand == $num ]]; then
	echo "Great you are done the level 1 "
else
	echo "You are died" 
	exit 1
fi

sleep 2
echo "Welcome to the level 2 chose a number between 0 and 4 (0-4)"
read num1
rand1=$(( $RANDOM % 5))

if [[ $rand1 == $num1 ]]; then 
	echo "Good job you are done the level 2"
else
	echo "You are died"
	exit 2
fi

sleep 2
echo "Welcome to the last level chose a number between 0 and 9 (0-9)"
read num2
rand2=$(( $RANDOM % 10))

if [[ $rand2 == $num2 ]]; then 
	echo "Congratulation you are win"
else
	echo "You are died"
fi


