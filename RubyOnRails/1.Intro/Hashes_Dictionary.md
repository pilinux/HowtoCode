## Hashes/dictionary

### Key - value

**String as key:**

```
2.6.5 :001 > hash1 = {"a" => 1, "b" => 2, "c" => 3}
 => {"a"=>1, "b"=>2, "c"=>3}

2.6.5 :002 > hash2 = {'a' => 1, 'b' => 2, 'c' => 3}
 => {"a"=>1, "b"=>2, "c"=>3}
```

**Symbol as key:**

```
2.6.5 :003 > hash3 = {a: 1, b: 2, c: 3}
 => {:a=>1, :b=>2, :c=>3}
```

**Accessing value of a key:**

```
2.6.5 :004 > hash1["a"]
 => 1
2.6.5 :005 > hash1['b']
 => 2
2.6.5 :006 > hash1[:c]
 => nil
2.6.5 :007 > hash3[:c]
 => 3
```

**Finding all keys:**

```
2.6.5 :008 > hash1.keys
 => ["a", "b", "c"]
2.6.5 :009 > hash2.keys
 => ["a", "b", "c"]
2.6.5 :010 > hash3.keys
 => [:a, :b, :c]
```

**Finding all values:**

```
2.6.5 :011 > hash1.values
 => [1, 2, 3]
2.6.5 :012 > hash2.values
 => [1, 2, 3]
2.6.5 :013 > hash3.values
 => [1, 2, 3]
```

**Add new key - value pair:**

```
2.6.5 :014 > hash3[:d] = 4
 => 4
2.6.5 :016 > hash3[:e] = "Name"
 => "Name"
```

**Update a value:**

```
2.6.5 :018 > hash3[:e] = "Name2"
 => "Name2"
```

**Delete a pair:**

```
2.6.5 :020 > hash3.delete(:e)
 => "Name2"
2.6.5 :021 > hash3
 => {:a=>1, :b=>2, :c=>3, :d=>4}
```

**Iterate through a hash:**

```
2.6.5 :022 > hash1.each { |some_key, some_value| puts "The key is #{some_key} and the value is #{some_value}." }
The key is a and the value is 1.
The key is b and the value is 2.
The key is c and the value is 3.
 => {"a"=>1, "b"=>2, "c"=>3}
```



---