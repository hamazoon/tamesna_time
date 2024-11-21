#! /bin/bash

echo "Welcome to the tower hotel"
sleep 2
echo "Going up"
sleep 2

for  x in {1..17};
do
	if [[ $x == 13 ]]; then
		continue
	fi
	echo "Floor number $x"
	sleep 2
done

