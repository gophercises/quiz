# Exercise #1: Quiz Game

This exercise is broken into two parts.

#### Part 1

[![topic: csvs](https://img.shields.io/badge/topic-csvs-green.svg?style=flat-square)](https://github.com/search?q=topic%3Acsvs+org%3Agophercises&type=Repositories)
[![topic: flags](https://img.shields.io/badge/topic-flags-green.svg?style=flat-square)](https://github.com/search?q=topic%3Aflags+org%3Agophercises&type=Repositories)
[![topic: opening files](https://img.shields.io/badge/topic-files-green.svg?style=flat-square)](https://github.com/search?q=topic%3Aos%2Dpackage+org%3Agophercises&type=Repositories)
[![topic: strings](https://img.shields.io/badge/topic-strings-green.svg?style=flat-square)](https://github.com/search?q=topic%3Astrings+org%3Agophercises&type=Repositories)

#### Part 2

[![topic: goroutines](https://img.shields.io/badge/topic-goroutines-green.svg?style=flat-square)](https://github.com/search?q=topic%3Agoroutines+org%3Agophercises&type=Repositories)
[![topic: channels](https://img.shields.io/badge/topic-channels-green.svg?style=flat-square)](https://github.com/search?q=topic%3Achannels+org%3Agophercises&type=Repositories)
[![topic: timers](https://img.shields.io/badge/topic-timers-green.svg?style=flat-square)](https://github.com/search?q=topic%3Atime$2Dpackages+org%3Agophercises&type=Repositories)



![video status: unreleased](https://img.shields.io/badge/video%20status-unreleased-red.svg?style=flat-square)
![code status: unreleased](https://img.shields.io/badge/code%20status-unreleased-red.svg?style=flat-square)

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

Users should be asked to press enter (or some other key) before the timer starts, and then the questions should be printed out to the screen one at a time until the user provides an answer. Regardless of whether the answer is correct or wrong the next question should be asked.

At the end of the quiz the program should output the total number of questions correct and how many questions there were in total. Questions given invalid answers or unanswered are considered incorrect.

## Bonus

As a bonus exercises you can also...

1. Add string trimming and cleanup to help ensure that correct answers with extra whitespace, capitalization, etc are not considered incorrect. *Hint: Check out the [strings](https://golang.org/pkg/strings/) package.*
2. Add an option (a new flag) to shuffle the quiz order each time it is run.
