## String Basics

### String concatenation

```
first_name = "Name1"
last_name = "Name2"
full_name = first_name + " " + last_name

p full_name
 => "Name1 Name2"
```


### String interpolation

```
p "My first name is #{first_name} and last name is #{last_name}"
 => "My first name is Name1 and last name is Name2"

p 'My first name is #{first_name} and last name is #{last_name}'
 => "My first name is \#{first_name} and last name is \#{last_name}"

[Difference between "" and '']
```


### Get input from the command line

```
2.6.5 :001 > variableName1 = gets.chomp
hello
 => "hello"
2.6.5 :002 > p variableName1
"hello"
 => "hello"

2.6.5 :003 > variableName2 = gets
hello
 => "hello\n"
2.6.5 :004 > p variableName2
"hello\n"
 => "hello\n"
```



---