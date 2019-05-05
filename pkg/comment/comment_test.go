Generated TestRender
// Package comment templates response letters for grading permit and subdivision
// applications to the Public Works for the City of Little Rock.
package comment

import "testing"

func TestRender(t *testing.T) {
	type args struct {
		m master
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Render(tt.args.m)
		})
	}
}
