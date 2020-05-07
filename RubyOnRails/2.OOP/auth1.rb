# Called by users_list.rb file

module Auth
  require 'bcrypt'
  puts "Auth module activated."

  def Auth.create_hash_digest(password)
    BCrypt::Password.create(password)
  end

  def Auth.verify_hash_digest(password)
    BCrypt::Password.new(password)
  end

  def Auth.create_users_auth(users_list)
    users_list.each do |user_record|
      user_record[:password] = create_hash_digest(user_record[:password])
    end
    users_list
  end

  def Auth.authenticate_user(username, password, users_list)
    users_list.each do |user_record|
      if user_record[:username] == username && verify_hash_digest(user_record[:password]) == password
        return user_record
      end
    end
    "Wrong username and/or password!"
  end
end
