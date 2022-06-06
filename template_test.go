package main

import "testing"

func TestGenDependencyTemplate(t *testing.T) {
	type args struct {
		platFormType PlatFormType
		gAVEntity    GAVEntity
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenDependencyTemplate(tt.args.platFormType, tt.args.gAVEntity); got != tt.want {
				t.Errorf("GenDependencyTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}
