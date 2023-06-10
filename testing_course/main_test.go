package main
 
import (
    "testing"
)
 
// test function
func TestReturnGeeks(t *testing.T) {
	type testcase struct {
		name string
		amount int
		returnvalue int
	}

	testcases :=[] testcase {
		{name: "should return 15", amount: 3, returnvalue: 15},
		{name: "should return 20", amount: 4, returnvalue: 20},
		{name: "should return 25", amount: 5, returnvalue: 20},
	}

	for _,tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			expectedvalue := ReturnGeeks(tc.amount)
			if  tc.returnvalue != expectedvalue{
				t.Errorf("Expected value(%v) is not same as actual return value (%v)", expectedvalue,tc.returnvalue)
		}
	})
			

	}
    
// Errorf = Fail + Log
// Fatalf = FailNow(stops execution at that point only)+ Log
}