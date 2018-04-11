package matematica

//function that adds numbers
func Add(args ...int) (int){
	c := 0
	for _, value := range args{
		c += value
	}
	return c
}

//function that multiply numbers
func Multiply(args ...int) (int){
	c := 1
	for _, value := range args{
		c *= value
	}
	return c
}
