package main
import "testing"

type Rectest struct{
	recs Rectangle
	result float64
}

type CircleTest struct{
	cirs Circle
	result float64
}

var recTest = []Rectest{{Rectangle{2,8},{Rectangle{4,16}}
var circleTest = []CircleTest{{Circle{4}}

func PerimeterTest(t *testing.T){
	for _,cases:= range recTest{
		x:=cases.recs.perimeter()
		if x!=cases.result{
			t.Error("Expected result was",cases.result, "got",x)
		}
	}

	for _,cases:= range circleTest{
		x:=cases.cirs.Perimeter()
		if x!=cases.result{
			t.Error("Expected result was",cases.result, "got",x)
		}
	}
}
