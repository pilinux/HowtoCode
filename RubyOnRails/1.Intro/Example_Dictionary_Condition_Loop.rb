iso_country_code = {
  "Argentina" => "032",
  "Australia" => "036",
  "Austria" => "040",
  "Belgium" => "056",
  "Canada" => "124",
  "China" => "156",
  "Denmark" => "208",
  "Estonia" => "233",
  "Germany" => "276",
  "Greece" => "300",
  "Italy" => "380",
  "Norway" => "578",
  "Sweden" => "752",
  "Switzerland" => "756",
  "United Kingdom" => "826",
  "United States of America" => "840"
}

def get_country_names(countries)
  countries.keys
end

def get_country_code(country, key)
  country[key]
end

loop do
  puts "Find an ISO code based on a country's name? (Y/N)"
  answer = gets.chomp.downcase
  break if answer != "y"

  puts "Countries available in the database:"
  puts get_country_names(iso_country_code)
  puts ""
  puts "Enter a country's name:"
  prompt = gets.chomp.capitalize

  if iso_country_code.include?(prompt)
    puts "The ISO code for #{prompt} is #{get_country_code(iso_country_code, prompt)}"
  else
    puts "Country is not available in the dictionary!"
  end
end
