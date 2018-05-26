Things to remember about go 

![alt text](https://raw.githubusercontent.com/FelipeRando/go-workspace/master/pointers-in-a-nutshell.png)

A star in a front of a receiver type means that the function are expecting a pointer to this type (it is not a operator in this moment)

The star become a operator when it is inside the function (now the star indicates that we are modifying the value associated to the pointer)

![alt text](https://raw.githubusercontent.com/FelipeRando/go-workspace/master/pointer-in-a-function-explained.png)



You can pass a variable to a pointer receiver in both ways
![alt text](https://raw.githubusercontent.com/FelipeRando/go-workspace/master/ways-to-pass-a-pointer-to-a-function.png)



However with slices you dont need to pass a pointer of a sclice into a function to get a value modified

![alt text](https://raw.githubusercontent.com/FelipeRando/go-workspace/master/value-and-reference-types.png)


![alt text](https://raw.githubusercontent.com/FelipeRando/go-workspace/master/map-vs-struct.png)

![alt text](https://raw.githubusercontent.com/FelipeRando/go-workspace/master/interface.png)

![alt text](https://raw.githubusercontent.com/FelipeRando/go-workspace/master/response-struct.png)

![alt text](https://raw.githubusercontent.com/FelipeRando/go-workspace/master/reader-interface.png)


Here's a picture on how the Reader interface works
![alt text](https://raw.githubusercontent.com/FelipeRando/go-workspace/master/how_Reader_interface_works.png)