package user

import "testing"

func TestNewUserName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    UserName
		wantErr bool
	}{
		{
			name: "NoName",
			args: args{
				name: "",
			},
			wantErr: true,
		},
		{
			name: "worbridg",
			args: args{
				name: "worbridg",
			},
			want:    "worbridg",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUserName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUserName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewUserName() = %v, want %v", got, tt.want)
			}
		})
	}
}
