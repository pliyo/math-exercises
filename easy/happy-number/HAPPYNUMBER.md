# Happy number

A happy number is a number defined by the following process:

Starting with any positive integer, replace the number by the sum of the squares of its digits, and repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1. Those numbers for which this process ends in 1 are happy numbers.

First few happy numbers are 1, 7, 10, 13, 19, 23, 28, 31, 32, 44, 49, 68, 70, 79, 82, 86, 91, 94, 97, 100

## Examples:

Input : 23
Output : Yes

# Explanation :

First Iteration:
22 + 32 = 4 + 9 = 13
Second Iteration:
12 + 32 = 1 + 9 = 10
Third Iteration:
12 + 02 = 1 + 0 = 1
Since we reach 1, 23 is Happy.
