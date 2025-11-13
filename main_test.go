package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct{
		name  		string 
		a,b   		int 
		expected 	int 
	}{
		{"all positive",2,3,5},
		{"have zero",0,5,5},
		{"all negative",-1,-3,-4},
		{"all positive",4,2,6},
	}

	for _,tt := range tests {
		t.Run(tt.name,func(t *testing.T){
			result := Add(tt.a,tt.b)
			if result != tt.expected {
				t.Errorf("Expected %d got %d",tt.expected,result)
			}
		})
	}
}