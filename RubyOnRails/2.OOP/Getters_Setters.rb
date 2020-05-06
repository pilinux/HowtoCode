# Attributes, getters, setters

class Student
  attr_accessor :first_name, :last_name, :username, :email, :password

  def initialize(firstname, lastname, username, email, password)
    @first_name = firstname
    @last_name = lastname
    @username = username
    @email = email
    @password = password
  end

  def to_s
    "
	First name: #{@first_name},
	Last name: #{@last_name},
	Username: #{@username},
	email address: #{@email}
	"
  end

end

# Instances
mike = Student.new("Michael", "Corleone", "mc", "mike@example.com", "password1")

john = Student.new("John", "Doe", "jd", "john@example.com", "password2")

# Getters
puts mike
puts john

# Setters
mike.last_name = john.last_name

puts "Michael Corleone is altered"

# Getters
puts mike
