package main

import (
	"strings"
	"testing"
)

func TestParseCSV(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    [][]string
		wantErr bool
	}{
		{
			name:  "простой csv",
			input: "name,age\nAlice,30\nBob,25\n",
			want: [][]string{
				{"name", "age"},
				{"Alice", "30"},
				{"Bob", "25"},
			},
			wantErr: false,
		},
		{
			name:  "значение в кавычках с запятой",
			input: "name,city\nBob,\"SPb, Russia\"\n",
			want: [][]string{
				{"name", "city"},
				{"Bob", "SPb, Russia"},
			},
			wantErr: false,
		},
		{
			name:    "пустой csv",
			input:   "",
			want:    nil,
			wantErr: false,
		},
		{
			name:    "битый csv (незакрытая кавычка)",
			input:   "name,city\nBob,\"SPb\n",
			want:    nil,
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := parseCSV(strings.NewReader(tc.input))

			if tc.wantErr {
				if err == nil {
					t.Fatal("ожидали ошибку, но её нет")
				}
				return
			}

			if err != nil {
				t.Fatalf("неожидали ошибку: %v", err)
			}

			if len(got) != len(tc.want) {
				t.Fatalf("разное число строк: got %d, want %d", len(got), len(tc.want))
			}

			for i := range got {
				if len(got[i]) != len(tc.want[i]) {
					t.Errorf("строка %d: разное число колонок: got %v, want %v", i, got[i], tc.want[i])
					continue
				}
				for j := range got[i] {
					if got[i][j] != tc.want[i][j] {
						t.Errorf("строка %d, колонка %d: got %q, want %q", i, j, got[i][j], tc.want[i][j])
					}
				}
			}
		})
	}
}
