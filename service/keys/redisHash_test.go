package keys

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_getHashKeys(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHashKeys(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getHashKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getHash(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHash(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getHashValue(t *testing.T) {
	type args struct {
		key   string
		field string
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
			if got := getHashValue(tt.args.key, tt.args.field); got != tt.want {
				t.Errorf("getHashValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHINCRBY(t *testing.T) {
	fmt.Println(HINCRBY("myhash","f"))

}
