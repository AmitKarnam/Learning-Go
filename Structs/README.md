# Struct
In Go, a struct is a composite data type that allows you to group together different fields of various types into a single entity. It provides a way to define your own custom data structures and represent complex data.

Basic Synatax

``` 
    type StructName struct {
        field1 type1 
        field2 type2
}
```

Structs can also be used as fields within other structs, allowing us to create nested or hierarchical data structures. 

Structs in Go are value types, which means that when we assign a struct variable to another variable or pass it as a function argument, a copy of the struct is created. However, we can use pointers to structs to achieve pass-by-reference semantics and modify the original struct.




