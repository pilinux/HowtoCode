####################
# Conditions
# ==========
#
#
# if condition
#   ... ...
# end
#
#
# if condition
#   ... ...
# else
#   ... ...
# end
#
#
# if condition
#   ... ...
# elsif condition
#   ... ...
# else
#   ... ...
# end
####################


def multiply(first_number, second_number)
  return first_number.to_f * second_number.to_f
end

def divide(first_number, second_number)
  return first_number.to_f / second_number.to_f
end

def subtract(first_number, second_number)
  return first_number.to_f - second_number.to_f
end

def mod(first_number, second_number)
  return first_number.to_f % second_number.to_f
end


puts "Select: 1) multiply 2) divide 3) subtract 4) find remainder"
prompt = gets.chomp
puts "First number:"
first_number = gets.chomp
puts "Second number:"
second_number = gets.chomp

if prompt == '1'
  puts "Multiplication: #{first_number} x #{second_number} ="
  puts multiply(first_number, second_number)
elsif prompt == '2'
  puts "Division: #{first_number} / #{second_number} ="
  puts divide(first_number, second_number)
elsif prompt == '3'
  puts "Subtraction: #{first_number} - #{second_number} ="
  puts subtract(first_number, second_number)
elsif prompt == '4'
  puts "Remainder: #{first_number} % #{second_number} ="
  puts mod(first_number, second_number)
else
  puts "Invalid choice!"
end
