#! /bin/bash

echo "Welcome to pushup program"
x=1

while [[ $x -le 15 ]] 
do 
	read -p "Pushup $x. Press enter to Contiue"
	(( x ++ ))
done
echo "Good job you complited your pushups"
