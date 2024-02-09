package main

import "testing"

func Test_httpArchive_String(t *testing.T) {
	tests := []struct {
		name string
		achv *httpArchive
		want string
	}{
		{
			name: "rules_go v0.45.1",
			achv: rules_go_0_45_1,
			want: "jasjdfjadf",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.achv.String()
			if err != nil {
				t.Errorf("String() returned an error: %v", err)
			}
			if got != tt.want {
				t.Errorf("String() got = %v, want %v", got, tt.want)
			}
		})
	}
}
