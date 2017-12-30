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

This exercise is broken into two parts to help simplify the process of explaining it as well as to make it easier to solve. The second part is harder than the first, so if you get stuck feel free to move on to another problem then come back to part 2 later.

*Note: I didn't break this into multiple exercises like I do for some exercises because both of these combined should only take ~30m to cover in screencasts.*

### Part 1

Create a program that will read in a quiz provided via a CSV file (more details below) and will then give the quiz to a user keeping track of how many questions they get right and how many they get incorrect. Regardless of whether the answer is correct or wrong the next question should be asked immediately afterwards.

The CSV file should default to [problems.csv](problems.csv), but the user should be able to customize the filename via a flag.

The CSV file will be in a format like below, where the first column is a question and the second column in the same row is the answer to that question.

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

You can assume that quizzes will be relatively short (< 100 questions) and will have single word/number answers.

At the end of the quiz the program should output the total number of questions correct and how many questions there were in total. Questions given invalid answers are considered incorrect.

### Part 2

Adapt your program from part 1 to add a timer. The default time limit should be 30 seconds, but should also be customizable via a flag.

Your quiz should stop as soon as the time limit has exceeded. That is, you shouldn't wait for the user to answer one final questions but should ideally stop the quiz entirely even if you are currently waiting on an answer from the end user.

Users should be asked to press enter (or some other key) before the timer starts, and then the questions should be printed out to the screen one at a time until the user provides an answer. Regardless of whether the answer is correct or wrong the next question should be asked.

At the end of the quiz the program should still output the total number of questions correct and how many questions there were in total. Questions given invalid answers or unanswered are considered incorrect.

## Bonus

As a bonus exercises you can also...

1. Add string trimming and cleanup to help ensure that correct answers with extra whitespace, capitalization, etc are not considered incorrect. *Hint: Check out the [strings](https://golang.org/pkg/strings/) package.*
2. Add an option (a new flag) to shuffle the quiz order each time it is run.
