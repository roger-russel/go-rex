package main

func ExampleMain() {

	main()
	// Output:
	// /^[a-zA-Z0-9_.+-]+@[a-zA-Z-]+\.[a-z.A-Z0-9.-]+$/ matched fake@email.com.br
	// /^[[:alnum:]_.+-]+@[[:alpha:]]+\.[[:alnum:].-]+$/ matched fake@email.com.br

}
