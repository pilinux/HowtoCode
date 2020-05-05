## Numbers

### Integer, float

- Integer: `20`
- Float: `20.0`

```
2.6.5 :001 > 20.to_f
 => 20.0
2.6.5 :002 > 20.0.to_i
 => 20
```


### Commonly used methods

`object.even?`

```
2.6.5 :003 > 20.even?
 => true
```

`object.odd?`

```
2.6.5 :004 > 20.odd?
 => false
```


### Comparisons

```
2.6.5 :005 > 1 == 2
 => false
2.6.5 :006 > 2 == 2
 => true
2.6.5 :007 > 5 < 2
 => false
2.6.5 :008 > 5 > 2
 => true
2.6.5 :009 > 5 >= 2
 => true
2.6.5 :010 > 2 <= 2
 => true
2.6.5 :011 > 5 && 6
 => 6
2.6.5 :012 > 5 || 6
 => 5
2.6.5 :013 > 6 || 5
 => 6
2.6.5 :014 > 6 && 5
 => 5
```


### Random numbers

- Between 0 and less than 20:

```
2.6.5 :015 > rand(20)
 => 9
2.6.5 :016 > rand(20)
 => 8
2.6.5 :017 > rand(20)
 => 14
```


### Convert string object to and from an integer object

```
2.6.5 :018 > "10".to_i
 => 10
2.6.5 :019 > 10.to_s
 => "10"
```


### Printing the same character

```
2.6.5 :020 > 20.times { print "-" }
-------------------- => 20
```



---