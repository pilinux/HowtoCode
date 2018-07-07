fn main() {
	println!("Hello, world!");	//Hello, world!
	//println! macro is used to print a line
	
	println!("{}, {}!", "Hello", "world");	//Hello, world!
	
	println!("{0}, {1}!", "Hello", "world");	//Hello, world!
	
	println!("{1}, {0}!", "world", "Hello");	//Hello, world!
	
	println!("{greeting}, {name}!", greeting="Hello", name="world");	//Hello, world!
	
	println!("{:?}", [1,2,3]);	//[1, 2, 3]
	
	println!("{:#?}", [1,2,3]);
	/*
	[
    	1,
    	2,
    	3
	]
	*/

	let x = format!("{}, {}!", "Hello", "world");	//format! macro is used to store the formatted STRING
	println!("{}", x);	//Hello, world!
}
