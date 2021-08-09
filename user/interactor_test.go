package user

import (
	"errors"
	"reflect"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestNewInteractor(t *testing.T) {
	type args struct {
		repo Repository
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				repo: NewMockRepository(gomock.NewController(t)),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInteractor(tt.args.repo); (got == nil) != tt.wantErr {
				t.Errorf("NewInteractor() = %v, wantErr %v", got, tt.wantErr)
			}
		})
	}
}

func TestInteractor_GetUser(t *testing.T) {
	type args struct {
		r *Request
	}
	tests := []struct {
		name          string
		newRepository func() *MockRepository
		args          args
		want          *Response
		wantErr       bool
	}{
		{
			name: "OK",
			newRepository: func() *MockRepository {
				repo := NewMockRepository(gomock.NewController(t))
				repo.EXPECT().
					GetByID(UserID(1)).
					Return(&User{1, "worbridg"}, nil)
				return repo
			},
			args: args{
				r: &Request{
					UserID:   1,
					UserName: "",
				},
			},
			want: &Response{
				Name: "worbridg",
			},
		},
		{
			name: "RepositoryError",
			newRepository: func() *MockRepository {
				repo := NewMockRepository(gomock.NewController(t))
				repo.EXPECT().
					GetByID(UserID(1)).
					Return(nil, errors.New("repository error"))
				return repo
			},
			args: args{
				r: &Request{
					UserID:   1,
					UserName: "",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				repo: tt.newRepository(),
			}
			got, err := i.GetUser(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Interactor.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interactor.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteractor_CreateUser(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		r *Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				repo: tt.fields.repo,
			}
			if err := i.CreateUser(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("Interactor.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
