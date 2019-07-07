package ido

import (
	"reflect"
	"testing"
)

func TestNewDocker(t *testing.T) {
	tests := []struct {
		want *docker
	}{
		{
			want: &docker{},
		},
	}

	for _, tt := range tests {
		got := newDocker()
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got: %#v, tt.want: %#v", got, tt.want)
		}
	}
}
