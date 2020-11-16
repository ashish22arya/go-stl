package queue

import (
	"reflect"
	"testing"
)

func TestNewQueue(t *testing.T) {
	tests := []struct {
		name string
		want *queue
	}{
		{name: "Test case 1", want: &queue{nil, nil, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewQueue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queue_Size(t *testing.T) {
	type fields struct {
		front  *node
		back   *node
		length int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{name: "Test case 1", fields: fields{nil, nil, 0}, want: 0},
		{name: "Test case 2", fields: fields{nil, nil, 2}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &queue{
				front:  tt.fields.front,
				back:   tt.fields.back,
				length: tt.fields.length,
			}
			if got := q.Size(); got != tt.want {
				t.Errorf("queue.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queue_IsEmpty(t *testing.T) {
	type fields struct {
		front  *node
		back   *node
		length int
	}
	testNode := &node{1, nil}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{name: "Test case 1", fields: fields{nil, nil, 0}, want: true},
		{name: "Test case 2", fields: fields{testNode, testNode, 1}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &queue{
				front:  tt.fields.front,
				back:   tt.fields.back,
				length: tt.fields.length,
			}
			if got := q.IsEmpty(); got != tt.want {
				t.Errorf("queue.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queue_Front(t *testing.T) {
	type fields struct {
		front  *node
		back   *node
		length int
	}
	testNode := &node{1, nil}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{name: "Test case 1", fields: fields{nil, nil, 0}, want: nil},
		{name: "Test case 2", fields: fields{testNode, testNode, 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &queue{
				front:  tt.fields.front,
				back:   tt.fields.back,
				length: tt.fields.length,
			}
			if got := q.Front(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("queue.Front() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queue_Back(t *testing.T) {
	type fields struct {
		front  *node
		back   *node
		length int
	}
	testNode := &node{1, nil}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{name: "Test case 1", fields: fields{nil, nil, 0}, want: nil},
		{name: "Test case 2", fields: fields{testNode, testNode, 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &queue{
				front:  tt.fields.front,
				back:   tt.fields.back,
				length: tt.fields.length,
			}
			if got := q.Back(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("queue.Back() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queue_Push(t *testing.T) {
	type fields struct {
		front  *node
		back   *node
		length int
	}
	type args struct {
		d interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "Test case 1", fields: fields{nil, nil, 0}, args: args{d: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &queue{
				front:  tt.fields.front,
				back:   tt.fields.back,
				length: tt.fields.length,
			}
			q.Push(tt.args.d)
		})
	}
}

func Test_queue_Pop(t *testing.T) {
	type fields struct {
		front  *node
		back   *node
		length int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{name: "Test case 1", fields: fields{nil, nil, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &queue{
				front:  tt.fields.front,
				back:   tt.fields.back,
				length: tt.fields.length,
			}
			q.Pop()
		})
	}
}
