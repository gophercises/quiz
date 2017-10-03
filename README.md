# Exercise #1: Quiz Game

**Topics**: [![topic](https://img.shields.io/badge/-CSVs-green.svg?style=flat-square)]() [![topic](https://img.shields.io/badge/-goroutines-green.svg?style=flat-square)]() [![topic](https://img.shields.io/badge/-channels-green.svg?style=flat-square)]() [![topic](https://img.shields.io/badge/-flags-green.svg?style=flat-square)]()

## Exercise details

Given a CSV like below, where the first column is the question and the second column is the answer:

```
5+5,10
7+3,10
1+1,2
8+3,11
1+2,3
8+6,14
3+1,4
1+4,5
5+1,6
2+3,5
3+3,6
2+4,6
5+2,7
```

Create a program that will accept the CSV filepath and a time limit (in seconds) as flags and will then run the the quiz reading each problem in order and stopping the quiz as soon as the time limit has been exceeded.

## Bonus

As a bonus exercise you can also add an option to shuffle the quiz order each time it is run with another flag.
