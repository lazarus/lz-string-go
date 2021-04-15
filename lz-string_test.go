package LZString

import "testing"

func TestCompressToBase64(t *testing.T) {
	type args struct {
		uncompressed string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{uncompressed: "123"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompressToBase64(tt.args.uncompressed); got != tt.want {
				t.Errorf("CompressToBase64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompressFromString(t *testing.T) {
	type args struct {
		uncompressed string
		chars        string
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
			if got := CompressFromString(tt.args.uncompressed, tt.args.chars); got != tt.want {
				t.Errorf("CompressFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCompressFromString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompressFromString("1233333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333eeeeeeeeeeeeeeeeee2e23e2qwerqwerqwerqwerqwerqweeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerqwerqwer1233333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333eeeeeeeeeeeeeeeeee2e23e2qwerqwerqwerqwerqwerqweeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerqwerqwer1233333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333eeeeeeeeeeeeeeeeee2e23e2qwerqwerqwerqwerqwerqweeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerqwerqwer1233333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333eeeeeeeeeeeeeeeeee2e23e2qwerqwerqwerqwerqwerqweeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerqwerqwer1233333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333eeeeeeeeeeeeeeeeee2e23e2qwerqwerqwerqwerqwerqweeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerqwerqwer1233333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333eeeeeeeeeeeeeeeeee2e23e2qwerqwerqwerqwerqwerqweeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerqwerqwer1233333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333eeeeeeeeeeeeeeeeee2e23e2qwerqwerqwerqwerqwerqweeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerqwerqwer1233333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333eeeeeeeeeeeeeeeeee2e23e2qwerqwerqwerqwerqwerqweeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerqwerqwer1233333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333eeeeeeeeeeeeeeeeee2e23e2qwerqwerqwerqwerqwerqweeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerqwerqwer1233333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333eeeeeeeeeeeeeeeeee2e23e2qwerqwerqwerqwerqwerqweeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerqwerqwer", keyStrBase64)
	}
}

func BenchmarkCompressFromString2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompressFromString("3123123", keyStrBase64)
	}
}
