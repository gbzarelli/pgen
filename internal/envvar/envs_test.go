package envvar

import (
	"os"
	"strconv"
	"testing"
)

func TestGetIntegerValueEnv(t *testing.T) {
	realValue1 := 20
	realValue2 := 30
	os.Setenv("REAL_ENV_1", strconv.Itoa(realValue1))
	os.Setenv("REAL_ENV_2", strconv.Itoa(realValue2))
	os.Setenv("REAL_ENV_3", "ABC")

	type args struct {
		envName      string
		defaultValue int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"get default value", args{envName: "TEST", defaultValue: 1}, 1},
		{"get default value", args{envName: "TEST-2", defaultValue: 10}, 10},
		{"get real value", args{envName: "REAL_ENV_1", defaultValue: 1}, realValue1},
		{"get real value", args{envName: "REAL_ENV_2", defaultValue: 1}, realValue2},
		{"get default value when the real value is not int", args{envName: "REAL_ENV_3", defaultValue: 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetIntegerEnv(tt.args.envName, tt.args.defaultValue); got != tt.want {
				t.Errorf("GetIntegerEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStringValueEnv(t *testing.T) {
	realValue1 := "ABC"
	realValue2 := "ABC-DEF"
	os.Setenv("REAL_ENV_1", realValue1)
	os.Setenv("REAL_ENV_2", realValue2)

	type args struct {
		envName      string
		defaultValue string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"get default value", args{envName: "TEST", defaultValue: "a"}, "a"},
		{"get default value", args{envName: "TEST-2", defaultValue: "b"}, "b"},
		{"get real value", args{envName: "REAL_ENV_1", defaultValue: "a"}, realValue1},
		{"get real value", args{envName: "REAL_ENV_2", defaultValue: "a"}, realValue2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStringEnv(tt.args.envName, tt.args.defaultValue); got != tt.want {
				t.Errorf("GetIntegerEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}
