package ido

import (
	"os/exec"
	"reflect"
	"testing"
)

func TestNewShell(t *testing.T) {
	tests := []struct {
		name string
		arg  []string
		want *shell
	}{
		{
			name: "foo",
			arg:  []string{"bar", "baz"},
			want: &shell{
				line: "foo bar baz",
				cmd:  exec.Command("foo", "bar", "baz"),
			},
		},
	}

	for _, tt := range tests {
		got := newShell(tt.name, tt.arg...)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got: %#v, tt.want: %#v", got, tt.want)
		}
	}
}
