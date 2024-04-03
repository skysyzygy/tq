package auth

import (
	"os"
	"reflect"
	"testing"

	"bitbucket.org/TN_WebShare/gotess"
	"github.com/99designs/keyring"
)

// Setup the test environment by making a separate keystore for testing
func TestMain(m *testing.M) {
	// setup code
	keys, _ = keyring.Open(keyring.Config{
		ServiceName: "tq_test",
	})
	code := m.Run()
	// teardown code
	os.Exit(code)
}

func TestAuth_AuthString(t *testing.T) {
	tests := []struct {
		name    string
		a       Auth
		want    string
		wantErr bool
	}{
		{name: "generates string",
			a:    Auth{"a", "b", "c", "d", nil},
			want: "a:b:c:d"},
		{name: "complains when there are ':' present in hostname",
			a:       Auth{"a:", "b", "c", "d", nil},
			wantErr: true},
		{name: "complains when there are ':' present in username",
			a:       Auth{"a", "b:", "c", "d", nil},
			wantErr: true},
		{name: "complains when there are ':' present in usergroup",
			a:       Auth{"a", "b", "c:", "d", nil},
			wantErr: true},
		{name: "complains when there are ':' present in location",
			a:       Auth{"a", "b", "c", "d:", nil},
			wantErr: true},
		{name: "doesn't care about password",
			a:    Auth{"a", "b", "c", "d", []byte(":::::")},
			want: "a:b:c:d"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.AuthString()
			if (err != nil) != tt.wantErr {
				t.Errorf("Auth.AuthString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Auth.AuthString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthFromString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    Auth
		wantErr bool
	}{
		{name: "parses string into Auth",
			args: args{"a:b:c:d"},
			want: Auth{"a", "b", "c", "d", nil}},
		{name: "complains when there are too many parts in the string",
			args:    args{"a:b:c:d:e"},
			wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AuthFromString(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuth_SaveAuth(t *testing.T) {
	tests := []struct {
		name    string
		a       Auth
		wantErr bool
	}{
		{name: "saves Auth to keystore",
			a: Auth{"a", "b", "c", "d", []byte("e")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.a.SaveAuth(); (err != nil) != tt.wantErr {
				t.Errorf("Auth.SaveAuth() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAuth_LoadAuth(t *testing.T) {
	a := Auth{"a", "b", "c", "d", nil}
	tests := []struct {
		name    string
		wantErr bool
		want    []byte
	}{
		{name: "loads Auth password from keystore",
			want: []byte("e")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := a.LoadAuth()
			got := a.password
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListAuths(t *testing.T) {
	tests := []struct {
		name    string
		want    []Auth
		wantErr bool
	}{
		{name: "lists all auths in keystore",
			want: []Auth{{"a", "b", "c", "d", nil}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ListAuths()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListAuths() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAuths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuth_DeleteAuth(t *testing.T) {
	tests := []struct {
		name    string
		a       Auth
		wantErr bool
	}{
		{name: "deletes auth from keystore",
			a: Auth{"a", "b", "c", "d", nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.a.DeleteAuth(); (err != nil) != tt.wantErr {
				t.Errorf("Auth.DeleteAuth() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestListAuths_after_delete(t *testing.T) {
	tests := []struct {
		name    string
		want    []Auth
		wantErr bool
	}{
		{name: "lists all auths in keystore",
			want: []Auth{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ListAuths()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListAuths() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAuths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuth_Login(t *testing.T) {
	tests := []struct {
		name    string
		a       Auth
		want    *gotess.Client
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.Login()
			if (err != nil) != tt.wantErr {
				t.Errorf("Auth.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Auth.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}
