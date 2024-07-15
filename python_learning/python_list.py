#! bin/user/python3

new_list =[12,34,-23,45,13,33]
#print the item in the index 3
print(new_list[3])
#print the numbers from the index 2 to 4
print(new_list[2:4])
#print the numbers from index 0 to 4
print(new_list[:4])
#print the numbers from index 0 to -4
print(new_list[:-4])
#print the numbers from index 1 to 5 and skip two numbers 
print(new_list[1:5:2])
#print the numbers and skip two numbers 
print(new_list[0::2])
#printthe minument number in the list 
print(min(new_list))
#print the maximant number in the list 
print(max(new_list))
#sorting of the list 
new_list.sort()
print(new_list)
#reversing of the list 
new_list.reverse()
print(new_list)
#addind the number 14 to the ens of list 
new_list.append((14))
print(new_list)
#adding 67 to the index 4 
new_list.insert(4,67)
print(new_list)
#adding new list to new_list 
new_list.extend([37,56,22])
print(new_list)
#editing the index 2 by the value 7
new_list[2]=7
print(new_list)
#remove the number 13 from the list 
new_list.remove(13)
print(new_list)
#print the number that want to remove 
print(new_list.pop())
#remove the last number
print(new_list)
#remove the number of the index 4
new_list.pop(4)
print(new_list)
#clear the list 
new_list.clear()
print(new_list)
