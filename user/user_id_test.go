package user

import "testing"

func TestNewUserID(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    UserID
		wantErr bool
	}{
		{
			name: "zero",
			args: args{
				id: 0,
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "negative",
			args: args{
				id: -1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUserID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserID_String(t *testing.T) {
	tests := []struct {
		name string
		id   UserID
		want string
	}{
		{
			name: "zero",
			id:   0,
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.id.String(); got != tt.want {
				t.Errorf("UserID.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
