#! /bin/bash

until [[ $order == "coffee" ]]
do
	echo "What would like to drink"
	read order
done
	echo "Exellent choise. Here is your coffee"
