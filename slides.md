# 1 - Go Generics
Cover Pic


# 2 - Generics Basics
### Concepts
    * Generic types
    * Generic Parameters
    * Generic Constrains
    * Type approximation


# 3 - Generic types
Generic types allow us to declare a parametised `type` where these parameters 
can be used to define the type of a slice, map or another generic type.
This means that we can now avoid code duplication 
as well as reducing the use of `interface{}` and losing valuable type information.

### Example
We can create a generic pagination type as well as a generic Collection type:
Note that the in Collection[T], the `T` Parameter must be passed through.
``` go
    type Collection[T any] []T

	type PaginatedResults[T any] struct {
		Page int
		Total int
		Items Collection[T]
	}
```

### generic types must be instantiated with `[string]`:
``` go
	result := PaginatedResults[string]{
		Page:  1,
		Total: 10,
		Items: []string{"1", "2"},
	}
```


# 4 - Generic Parameters
``` go
    func Map[T any, P any](items []T, fn func(T) P) []P {
        results := make([]P, 0, len(items))
        for _, elem := range items {
            results = append(results, fn(elem))
        }
        return results
    }
```
Generic functions can make use of type inference:
``` go
    func IntToString(i int) string { return fmt.Sprint(i) }

    func main() {
        ints := []int{1, 2, 3, 4, 5}
        strs := Map(ints, IntToString)
    }  
```

The type of `strs` can be inferred (`[]string`) based on the return type of `IntToString` being a `string`

### Limitations:
Generic Parameters can only be used in functions, you cannot define generic methods, only functions with no receiver.
``` go
func (r Receiver) [T any]NotAllowed(t T) {...
func [T any]Allowed(t T) {...
```


# 5 - Generic Constrains
generic constraints allow you to continue to use go language primitives in your functions.
Such as: `==`, `>`, `<` or any of the arithmetic or bitwise operators:
    * any
	* comparable (works with comparison operator: ==)
	* constraints.Ordered (works with order operators: >, <, >= , etc...)
    * constraints.Integer (works with bitwise/bitshift operators: &, |, ^, <<, >>)
    * constraints.Float (works with arithmetic operators: +, -, *, /, %)
    * constraints.Unsigned
    * constraints.Signed
    * constraints.Complex
    

# 6 - Custom constraints:
A constraint can be defined as any interface with a particalar set of functions: 
``` go
    type Stringer interface {
        String() string
    }
```
or You can define a custom constraint as a union of multiple types or constraints
``` go
    type Decimal interface {
        float32 | float64
    }

    type Integer interface {
        int64 | int32
    }

    type Number interface {
        Decimal | Integer
    }
```


# 7 - Type approximation
`~int` matches type int but also types derived from int eg. `type Integer int`
``` go
    type Float interface {
        ~float32 | ~float64
    }
```
