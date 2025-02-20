# Description

This is Davids's repository for the solution of the Computer science challenge of Endava's internship 2025-1. In this repo, you can find my implemented solutions for all of the problems with the approach and explanation of the solutions. Additionally, the time and space complexity of each one.

  

Here is a summary of all the complexities in the exercises:

  
| **Code** | **Name**                                                                                                              | Time Complexity | Space Complexity |
| -------- | --------------------------------------------------------------------------------------------------------------------- | --------------- | ---------------- |
| 78       | Subsets                                                                                                               | $O(n*2^n)$      | $O(2^n)$         |
| 13       | [Roman to Integer](https://leetcode.com/problems/roman-to-integer/)                                                   | $O(n)$          | $O(1)$           |
| 151      | [Reverse Words in a String](https://leetcode.com/problems/reverse-words-in-a-string/)                                 | $O(n)$          | $O(n)$           |
| 206      | [Reverse Linked List](https://confluence.endava.com/spaces/ARCHDISC/pages/140422221/Learning+Materials+System+Design) | $O(n)$          | $O(n)$           |
| 92       | [Reverse Linked List II](https://leetcode.com/problems/reverse-linked-list-ii/description/)                           | $O(n)$          | $O(n)$           |
| 214      | Shortest Palindrome                                                                                                   | $O(n^2)$        | $O(1)$           |
| 48       | Rotate Image                                                                                                          | $O(n^2)$        | $O(1)$           |
| 32       | [Longest Valid Parentheses](https://leetcode.com/problems/longest-valid-parentheses/)                                 | $O(n)$          | $O(n)$           |

# 78 Subsets

To find all of the subsets first we need to define in which subset we are going to focus first. In this case we will always focus on the bigger, so we will find all of the subsets with length M, after that it will find the subsets with M-1 length (until we reach 1, as the remaining are trivial enough to find) 

When finding all the subsets of M length we will be using an array with some static points and one moving cursor. In the case of M = 4 and N= 6 we could have something like this: `[1,3,4]` these are the static elements, and we need to find the remaining permutations `[1,3,4,5]` `[1,3,4,6]` We left some elements static, and one with the ability to move. When the permutations with that number ends, we continue adding to the static elements 1 until there are no more possible combinations 

| **Code** | **Name**                                                                                                              | Time Complexity | Space Complexity |
| -------- | --------------------------------------------------------------------------------------------------------------------- | --------------- | ---------------- |
| 78       | Subsets                                                                                                               | $O(n*2^n)$      | $O(2^n)$         |

The reason of this complexity is that the growth of this approach follows the pascal triangle, as for every N there will be exponentially more subsets than in N-1. And as we iterate over every N that gives us this complexity. The space needed to store all the subsets also follows the pascal triangle, for the same reason

# 13 Roman to Integer

  

This is a simple problem. At the start I created a HashMap with the values of the roman string as a key, and it's number equivalent as it's value. With that done, i iterated over the roman number in **reverse**. The reason for this is that i at the right the only possibility is to add. If a character is to the absolute right it's adding, With that done i also kept a variable with the latest large character. And depending on what it is (L,C,M,V) if the following character (the one on the left) was smaller it would subtract instead of adding
  
| **Code** | **Name**                                                                                                              | Time Complexity | Space Complexity |
| -------- | --------------------------------------------------------------------------------------------------------------------- | --------------- | ---------------- |
| 13       | [Roman to Integer](https://leetcode.com/problems/roman-to-integer/)                                                   | $O(n)$          | $O(1)$           |

# 151 Reverse Words in a String

This is a simple problem. Trim the ending and beginning spaces, split the words into an array on spaces, And iterate in reverse over this array, checking if the array is empty don't include it, else add it to the string.  


| **Code** | **Name**                                                                                                              | Time Complexity | Space Complexity |
| -------- | --------------------------------------------------------------------------------------------------------------------- | --------------- | ---------------- |
| 151      | [Reverse Words in a String](https://leetcode.com/problems/reverse-words-in-a-string/)                                 | $O(n)$          | $O(n)$           |


# 206 Reverse Linked List

This problem uses the following approach: Store the elements of the node in an array, and reverse the values of the array.

| **Code** | **Name**                                                                                                              | Time Complexity | Space Complexity |
| -------- | --------------------------------------------------------------------------------------------------------------------- | --------------- | ---------------- |
| 206      | [Reverse Linked List](https://confluence.endava.com/spaces/ARCHDISC/pages/140422221/Learning+Materials+System+Design) | $O(n)$          | $O(n)$           |


# 92 Reverse Linked List II

For this problem, i took the same approach as in the reverse linked list. storing the elements of the node in an array, and reversing it's values. The difference is only doing it to the parts that are between the Left and right parameters. And finally adding the right parameter to the list

| **Code** | **Name**                                                                                                              | Time Complexity | Space Complexity |
| -------- | --------------------------------------------------------------------------------------------------------------------- | --------------- | ---------------- |
| 92       | [Reverse Linked List II](https://leetcode.com/problems/reverse-linked-list-ii/description/)                           | $O(n)$          | $O(n)$           |


# 214 Shortest Palindrome

  A surprisingly easy problem. The main concept comes with understanding two pointers, you create a pointer at the left and right `(l,r)`. With them you can check if the word is a palindrome, if not move the starting point of r until the word is a palindrome or no palindrome is found, either way you know know what to add at the front of the word (the same characters that come after the starting point of r)

| **Code** | **Name**                                                                                                              | Time Complexity | Space Complexity |
| -------- | --------------------------------------------------------------------------------------------------------------------- | --------------- | ---------------- |
| 206      | [Reverse Linked List](https://confluence.endava.com/spaces/ARCHDISC/pages/140422221/Learning+Materials+System+Design) | $O(n^2)$        | $O(1)$           |
# 48 Rotate Image

For this problem i transposed the matrix and then i "mirrored" it to the correct order. This takes advantage of the compiler allowing code like this `x,y=y,x`, as i am NOT creating any extra space, but the compiler will do (as it has to package and unpackage the values to an structure)

| **Code** | **Name**     | Time Complexity | Space Complexity |
| -------- | ------------ | --------------- | ---------------- |
| 48       | Rotate Image | $O(n^2)$        | $O(1)$           |



# 32 Longest Valid Parentheses

For this problem, we find the longest valid parentheses (LVP) from the left, and then from the right. The greater one will be the longest valid parenthesis of the whole problem.

To find the LVP of any side we need a stack (that i had to implement myself as go does not have a direct implementation of stack), the current streak, a Boolean that will track if an streak is occurring, and debt.

A streak is started when we start creating valid parentheses, in other words, when we start closing parenthesis a streak is started. If an opening parentheses appears when a streak is occurring, we add one to the debt, so in other words, the debt tracks how many closing parenthesis must appear in a row.

The reason we divide this problem in two sides, is that an occurring streak will assume any of the following parentheses belong to the same streak (instead of being able to start a new streak). Therefore having two sides would allow for previously left sided streaks not to affect right sided streaks and vice versa 

| **Code** | **Name**                                                                                                              | Time Complexity | Space Complexity |
| -------- | --------------------------------------------------------------------------------------------------------------------- | --------------- | ---------------- |
| 32       | [Longest Valid Parentheses](https://leetcode.com/problems/longest-valid-parentheses/)                                 | $O(n)$          | $O(n)$           |