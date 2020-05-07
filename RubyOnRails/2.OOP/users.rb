# Example 2: Classes, Modules, Mixins with bcrypt library

require_relative 'auth2'

class Student
  include Auth
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
michael = Student.new("Michael", "Jackson", "mj", "mj@example.com", "password1")
john = Student.new("John", "Doe", "jd", "jd@example.com", "password2")
arya = Student.new("Arya", "Stark", "as", "as@example.com", "password3")
jon = Student.new("Jon", "Snow", "js", "js@example.com", "password4")
vito = Student.new("Vito", "Corleone", "vc", "vc@example.com", "password5")

puts ""
puts "Users:"
puts michael
puts john
puts arya
puts jon
puts vito

puts ""
puts "Hashing password for #{michael.first_name} #{michael.last_name}..."
hashed_password = michael.create_hash_digest(michael.password)
puts hashed_password

puts ""
puts "Verifying password..."
password_verification = michael.verify_hash_digest(hashed_password)
puts password_verification == michael.password
