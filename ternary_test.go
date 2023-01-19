package ternary

import (
	"reflect"
	"testing"
)

type T any

func TestTernary(t *testing.T) {
	type args struct {
		condition bool
		a         T
		b         T
	}
	tests := []struct {
		name string
		args args
		want T
	}{
		{
			name: "String True",
			args: args{
				condition: true,
				a:         "foo",
				b:         "bar",
			},
			want: "foo",
		},

		{
			name: "String False",
			args: args{
				condition: false,
				a:         "foo",
				b:         "bar",
			},
			want: "bar",
		},

		{
			name: "Int True",
			args: args{
				condition: true,
				a:         1,
				b:         2,
			},
			want: 1,
		},

		{
			name: "Int False",
			args: args{
				condition: false,
				a:         1,
				b:         2,
			},
			want: 2,
		},

		{
			name: "Float True",
			args: args{
				condition: true,
				a:         1.1,
				b:         2.2,
			},
			want: 1.1,
		},

		{
			name: "Float False",
			args: args{
				condition: false,
				a:         1.1,
				b:         2.2,
			},
			want: 2.2,
		},

		{
			name: "Bool True",
			args: args{
				condition: true,
				a:         true,
				b:         false,
			},
			want: true,
		},

		{
			name: "Bool False",
			args: args{
				condition: false,
				a:         true,
				b:         false,
			},
			want: false,
		},

		{
			name: "Struct True",
			args: args{
				condition: true,
				a:         struct{ foo string }{foo: "foo"},
				b:         struct{ foo string }{foo: "bar"},
			},
			want: struct{ foo string }{foo: "foo"},
		},

		{
			name: "Struct False",
			args: args{
				condition: false,
				a:         struct{ foo string }{foo: "foo"},
				b:         struct{ foo string }{foo: "bar"},
			},
			want: struct{ foo string }{foo: "bar"},
		},

		{
			name: "Slice True",
			args: args{
				condition: true,
				a:         []string{"foo"},
				b:         []string{"bar"},
			},
			want: []string{"foo"},
		},

		{
			name: "Slice False",
			args: args{
				condition: false,
				a:         []string{"foo"},
				b:         []string{"bar"},
			},
			want: []string{"bar"},
		},

		{
			name: "Map True",
			args: args{
				condition: true,
				a:         map[string]string{"foo": "foo"},
				b:         map[string]string{"foo": "bar"},
			},
			want: map[string]string{"foo": "foo"},
		},

		{
			name: "Map False",
			args: args{
				condition: false,
				a:         map[string]string{"foo": "foo"},
				b:         map[string]string{"foo": "bar"},
			},
			want: map[string]string{"foo": "bar"},
		},

		{
			name: "Interface True",
			args: args{
				condition: true,
				a:         interface{}("foo"),
				b:         interface{}("bar"),
			},
			want: interface{}("foo"),
		},

		{
			name: "Interface False",
			args: args{
				condition: false,
				a:         interface{}("foo"),
				b:         interface{}("bar"),
			},
			want: interface{}("bar"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := If(tt.args.condition, tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ternary() = %v, want %v", got, tt.want)
			}
		})
	}
}
