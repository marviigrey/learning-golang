The API server fetches the Product details from within the go code because it was 
created inside the go file using the struct type.

Here is a breakdown of the code inside the main.go file.

1. we created a struct type called Product with the values Id(int),Name(string),
Quantity(int), Price(float64).

2. Then we declared an array of type Product(from our struct which we created) called 
Products. The array Products is what we will use to define our values of our struct.
inside our main function we redeclared the Products array and assign values to it.

3. created our Homepage function that simply displays a message.

4. We created our returnAllProducts function that we use to encode our Products array
into a JSON format, which means when we call the function it displays our products in 
a JSON format. 

5. we created a handledRequests function to call all other functions inside our file.
And we also specified the port of our page.

