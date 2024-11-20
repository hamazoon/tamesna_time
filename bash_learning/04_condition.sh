#! /bin/bash

echo " Welcome to the digital game win or die"

echo "Please select your class:
1 - viking
2 - samuray
3 - othmani"
read class

case $class in
	1)
		type="viking"
		defonce=10
		attack=15
		magic=13
		;;
	2)
		type="samuray"
		defonce=16
		attack=15
		magic=15
		;;
	3)
		type="othmani"
		defonce=16
		attack=17
		magic=18
		;;
esac
echo "You chose the $type class. your defonce is $defonce. your attack is $attack. and you magic is $magic "

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


