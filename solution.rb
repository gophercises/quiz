# frozen_string_literal: true

require 'csv'

def ask_question_check_answer(question, answer)
  print question, '? '
  answer_from_user = gets.strip

  return true if answer_from_user == answer

  false
end

def main(file_name = 'problems.csv')
  points = 0

  questions = CSV.read file_name
  questions.each do |question|
    next if question.length != 2

    statement = question[0].strip
    answer = question[1].strip

    points += 1 if ask_question_check_answer statement, answer
  end

  print "\n"
  print 'Your score: ', points, '/', lines.length, "\n"
end

if ARGV.length >= 1
  main ARGV[0]
else
  main
end
