package clipboardprovider

import (
	"testing"
)

func TestPB_Copy(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		pb   *PB
		args args
	}{
		{
			name: "test 1",
			pb:   &PB{},
			args: args{
				data: []byte("this is from golang"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pb := &PB{}
			pb.Copy(tt.args.data)
		})
	}
}

func TestPB_Paste(t *testing.T) {
	tests := []struct {
		name string
		pb   *PB
		want bool
	}{
		// TODO: Add test cases.
        {
            name: "Test Case 1",
            pb: &PB{},
            want: true,
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pb := &PB{}
			if got := pb.Paste(); got != tt.want {
				t.Errorf("PB.Paste() = %v, want %v", got, tt.want)
			}
		})
	}
}
