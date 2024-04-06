## Function Signature

```go
func (d Deck) functionName(arg1 string) (int) {
		^			^				^
	Receiver	Arguments		Return-Type
}
```
'd' here is 'this' of C++/JS and 'self' of Python,it refers to current instance. This is called "Receiver".  
It is used as:
```go
myDeck := Deck{}
returnValue = myDeck.functionName(1)
```
And if there is a * then it is a pointer. just like in C/C++
```go
func (d *Deck) functionName(arg1 string) (int) {
	...
}
```
