package goconvey

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAdd(t *testing.T) {
	Convey("Add two numbers together", t, func() {
		So(Add(1, 2), ShouldEqual, 3)
	})
}

func TestSubtract(t *testing.T) {
	Convey("Subtract two numbers", t, func() {
		So(Subtract(1, 2), ShouldEqual, -1)
	})
}

func TestMultiply(t *testing.T) {
	Convey("Multiply two numbers together", t, func() {
		So(Multiply(3, 2), ShouldEqual, 6)
	})
}

func TestDivision(t *testing.T) {
	Convey("Divide two numbers", t, func() {

		Convey("Division by non-zero number", func() {
			num, err := Division(10, 2)
			So(err, ShouldBeNil)
			So(num, ShouldEqual, 5)
		})

		Convey("divide by 0", func() {
			_, err := Division(10, 0)
			So(err, ShouldNotBeNil)
		})
	})
}

//BDD SAMPLE TEMPLATE

// func TestDivide(t *testing.T) {

// 	Convey("Scenario: Divide numbers", t, func() {

// 		Convey("Given Two Integer numbers", func() {
// 			num1 := 30
// 			num2 := 3

// 			So(num1, ShouldHaveSameTypeAs, 0)
// 			So(num2, ShouldHaveSameTypeAs, num1)
// 			Convey("When  numbers are non-zero", func() {
// 				So(num1, ShouldNotEqual, 0)
// 				So(num2, ShouldNotEqual, 0)
// 				// So("num1", shouldNotBeEqualToZero)
// 				// So(0, shouldNotBeEqualToZero)

// 				Convey("Then the result should be an integer value", func() {
// 					result, err := Division(num1, num2)
// 					So(err, ShouldBeNil)
// 					So(result, ShouldEqual, 10)
// 				})

// 			})

// 			Convey("When numbers can be zero", func() {
// 				num2 = 0
// 				Convey("Then the result should be error", func() {
// 					_, err := Division(num1, num2)
// 					So(err, ShouldNotBeNil)
// 				})

// 			})

// 		})

// 	})

// }

//TEST CUSTOMIZED ASSERTION

// func shouldNotBeEqualToZero(actual interface{}, expected ...interface{}) string {
// 	res, ok := actual.(int)
// 	if !ok {
// 		return "passed value is not an integer"
// 	}
// 	if res != 0 {
// 		return "" // empty string means the assertion passed
// 	}
// 	return "passed value is equal to zer0"
// }

// TEST SETUP AND RESET

// func TestSome(t *testing.T) {
// 	Convey("Top-level", t, func() {
// 		// setup (run before each `Convey` at this scope):
// 		convey.Println("Open")
// 		convey.Println("Initliazee")

// 		Convey("Test a query", func() {
// 			convey.Println("Querry") // TODO: assertions here
// 		})

// 		Convey("Test inserts", func() {
// 			convey.Println("Insert") // TODO: assertions here
// 		})

// 		Reset(func() {
// 			// This reset is run after each `Convey` at the same scope.
// 			convey.Println("Close")
// 		})
// 	})

// }
