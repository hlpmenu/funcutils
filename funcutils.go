package funcutils

// DoubleFuncs is a struct that holds two functions

type ABFuncs struct {
	A Function
	B Function
}

type FuncMap map[string]Function

type FunctionSlice []Function

func (f FunctionSlice) RunAll() {

}
