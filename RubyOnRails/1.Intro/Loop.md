## Loop

### Loop through an array

```
2.6.5 :001 > i = [1, 2, 3, 4, 5, 6, 7, 8]
 => [1, 2, 3, 4, 5, 6, 7, 8]
2.6.5 :002 > i.each { |j| puts j }
1
2
3
4
5
6
7
8
 => [1, 2, 3, 4, 5, 6, 7, 8]
```


### Array of strings

**Create:**

```
2.6.5 :001 > %w(my name is nobody)
 => ["my", "name", "is", "nobody"]
```

**Assign to a variable:**

```
2.6.5 :002 > z = _
 => ["my", "name", "is", "nobody"]
```

**Loop through z:**

```
2.6.5 :003 > z.each do |element|
2.6.5 :004 >     print element + " "
2.6.5 :005?>   end
my name is nobody  => ["my", "name", "is", "nobody"]


2.6.5 :006 > z.each { |element| print element + " " }
my name is nobody  => ["my", "name", "is", "nobody"]

2.6.5 :007 > z.each { |element| print element.capitalize + " " }
My Name Is Nobody  => ["my", "name", "is", "nobody"]
```


### Array of numbers

```
2.6.5 :012 > x = (1..10).to_a
 => [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
2.6.5 :013 > x.select { |num| num.even? }
 => [2, 4, 6, 8, 10]
```


### Join elements of an array

```
2.6.5 :014 > x.join
 => "12345678910"
2.6.5 :015 > x
 => [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
2.6.5 :016 > x.join(" ")
 => "1 2 3 4 5 6 7 8 9 10"
2.6.5 :017 > x.join("-")
 => "1-2-3-4-5-6-7-8-9-10"
```



---