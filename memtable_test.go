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
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Deleting item that is present",
			fields: fields{
				store: map[string][]byte{"dog": []byte("3")},
			},
			args: args{
				key: []byte("dog"),
			},
			wantErr: false,
		},
		{
			name: "Deleting item that is not present",
			fields: fields{
				store: map[string][]byte{"dog": []byte("3")},
			},
			args: args{
				key: []byte("cat"),
			},
			wantErr: true,
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
		})
	}
}
