## IRB, Built-in Class, Methods

Interactive Ruby Shell (irb) console: `irb`

Exit irb console: `exit`


### What class an object belongs to

`variable_name.class` :

```
10.class
 => Integer

"10".class
 => String
```


### Methods available to an object

`objectName.methods` :

```
Integer.methods

String.methods
```


### Some methods commonly used on objects

`objectName.nil?` :

```
Integer.nil?
 => false

nil.nil?
 => true
```

`objectName.empty?` :

```
"".empty?
 => true

" ".empty?
 => false

"Hello".empty?
 => false
```

`objectName.Length` :

```
"Hello".length
 => 5
```

`objectName.reverse` :

```
"Hello".reverse
 => "olleH"
```



---