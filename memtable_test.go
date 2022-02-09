package ldbclone

import (
	"reflect"
	"testing"
)

func TestMemtable_Get(t *testing.T) {
	type fields struct {
		store map[string][]byte
	}
	type args struct {
		key []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Getting value that is present",
			fields: fields{
				store: map[string][]byte{"dog": []byte("3")},
			},
			args: args{
				key: []byte("dog"),
			},
			want:    []byte("3"),
			wantErr: false,
		},
		{
			name: "Trying to access value that is not present",
			fields: fields{
				store: map[string][]byte{"dog": []byte("3")},
			},
			args: args{
				key: []byte("cat"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Memtable{
				store: tt.fields.store,
			}
			m.store["dog"] = []byte("3")

			got, err := m.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Memtable.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Memtable.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemtable_Has(t *testing.T) {
	type fields struct {
		store map[string][]byte
	}
	type args struct {
		key []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Item is present",
			fields: fields{
				store: map[string][]byte{"dog": []byte("3")},
			},
			args: args{
				key: []byte("dog"),
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Item is not present",
			fields: fields{
				store: map[string][]byte{"dog": []byte("3")},
			},
			args: args{
				key: []byte("cat"),
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Memtable{
				store: tt.fields.store,
			}
			got, err := m.Has(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Memtable.Has() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Memtable.Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemtable_Delete(t *testing.T) {
	type fields struct {
		store map[string][]byte
	}
	type args struct {
		key []byte
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantErr  bool
		expected map[string][]byte
	}{
		{
			name: "Deleting item that is present",
			fields: fields{
				store: map[string][]byte{"dog": []byte("3")},
			},
			args: args{
				key: []byte("dog"),
			},
			wantErr:  false,
			expected: map[string][]byte{},
		},
		{
			name: "Deleting item that is not present",
			fields: fields{
				store: map[string][]byte{"dog": []byte("3")},
			},
			args: args{
				key: []byte("cat"),
			},
			wantErr:  true,
			expected: map[string][]byte{"dog": []byte("3")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Memtable{
				store: tt.fields.store,
			}
			if err := m.Delete(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Memtable.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
			if match := reflect.DeepEqual(m.store, tt.expected); !match {
				t.Errorf("Memtable.Delete() test failed \nWanted: %v \nGot: %v", m.store, tt.expected)
			}
		})
	}
}

// Test both that Put doesn't return an error, and that the map has the values expcted after the
// operation
func TestMemtable_Put(t *testing.T) {
	type fields struct {
		keys  []string
		store map[string][]byte
	}
	type args struct {
		key   []byte
		value []byte
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantErr  bool
		expected map[string][]byte
	}{
		{
			name: "Putting in a single value",
			fields: fields{
				store: map[string][]byte{},
			},
			args: args{
				key:   []byte("dog"),
				value: []byte("3"),
			},
			wantErr:  false,
			expected: map[string][]byte{"dog": []byte("3")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Memtable{
				keys:  tt.fields.keys,
				store: tt.fields.store,
			}
			if err := m.Put(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Memtable.Put() error = %v, wantErr %v", err, tt.wantErr)
			}
			if match := reflect.DeepEqual(m.store, tt.expected); !match {
				t.Errorf("Memtable.Delete() test failed \nWanted: %v \nGot: %v", m.store, tt.expected)
			}
		})
	}
}
