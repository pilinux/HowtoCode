# Example 1: Classes, Modules, Mixins with bcrypt library

require_relative 'auth1'

users = [
          { username: "michael", password: "password1" },
          { username: "john", password: "password2" },
          { username: "arya", password: "password3" },
          { username: "jon", password: "password4" },
          { username: "vito", password: "password5" }
        ]

puts ""
puts "Users:"
puts users

puts ""
puts "Hashing..."
hashed_users = Auth.create_users_auth(users)
puts hashed_users

puts ""
puts "Verifying michael with pass:<password2>"
verify_hashed_users = Auth.authenticate_user("michael", "password2", users)
puts verify_hashed_users

puts ""
puts "Verifying john with pass:<password2>"
verify_hashed_users = Auth.authenticate_user("john", "password2", users)
puts verify_hashed_users
