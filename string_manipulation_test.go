package strman

import (
	"reflect"
	"testing"
)

func Test_split(t *testing.T) {
	tests := []struct {
		name   string
		source string
		want   []string
	}{
		{
			name:   "split kebab",
			source: "split-kebab-case",
			want:   []string{"split", "kebab", "case"},
		},
		{
			name:   "split snake",
			source: "split_snake_case",
			want:   []string{"split", "snake", "case"},
		},
		{
			name:   "split screaming snake",
			source: "SPLIT_SNAKE_CASE",
			want:   []string{"split", "snake", "case"},
		},
		{
			name:   "split mixed",
			source: "split_mixed-case",
			want:   []string{"split", "mixed", "case"},
		},
		{
			name:   "split camel",
			source: "splitCamelCase",
			want:   []string{"split", "camel", "case"},
		},
		{
			name:   "split pascal",
			source: "SplitPascalCase",
			want:   []string{"split", "pascal", "case"},
		},
		{
			name:   "split snake numbers",
			source: "split_12_snake",
			want:   []string{"split", "12", "snake"},
		},
		{
			name:   "split kebab numbers",
			source: "split-12-kebab",
			want:   []string{"split", "12", "kebab"},
		},
		{
			name:   "split camel numbers",
			source: "split12Camel",
			want:   []string{"split", "12", "camel"},
		},
		{
			name:   "split pascal numbers",
			source: "Split12Pascal",
			want:   []string{"split", "12", "pascal"},
		},
		{
			name:   "split screaming snake numbers",
			source: "SPLIT_12_SCREAMING",
			want:   []string{"split", "12", "screaming"},
		},
		{
			name:   "split screaming kebab numbers",
			source: "SPLIT-12-SCREAMING",
			want:   []string{"split", "12", "screaming"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := split(tt.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("split() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToDelimited(t *testing.T) {
	tests := []struct {
		name      string
		source    string
		delimiter string
		want      string
	}{
		{
			name:      "from kebab to delimited",
			source:    "from-kebab",
			delimiter: ".",
			want:      "from.kebab",
		},
		{
			name:      "from snake to delimited",
			source:    "from_snake",
			delimiter: ".",
			want:      "from.snake",
		},
		{
			name:      "from screaming snake to delimited",
			source:    "FROM_SCREAMING_SNAKE",
			delimiter: ".",
			want:      "from.screaming.snake",
		},
		{
			name:      "from camel to delimited",
			source:    "fromCamel",
			delimiter: ".",
			want:      "from.camel",
		},
		{
			name:      "from Pascal to delimited",
			source:    "FromPascal",
			delimiter: ".",
			want:      "from.pascal",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToDelimited(tt.source, tt.delimiter); got != tt.want {
				t.Errorf("ToDelimited() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToScreamingDelimited(t *testing.T) {
	tests := []struct {
		name      string
		source    string
		delimiter string
		want      string
	}{
		{
			name:      "from kebab to screaming delimited",
			source:    "from-kebab",
			delimiter: ".",
			want:      "FROM.KEBAB",
		},
		{
			name:      "from snake to screaming delimited",
			source:    "from_snake",
			delimiter: ".",
			want:      "FROM.SNAKE",
		},
		{
			name:      "from screaming snake to screaming delimited",
			source:    "FROM_SCREAMING_SNAKE",
			delimiter: ".",
			want:      "FROM.SCREAMING.SNAKE",
		},
		{
			name:      "from camel to screaming delimited",
			source:    "fromCamel",
			delimiter: ".",
			want:      "FROM.CAMEL",
		},
		{
			name:      "from pascal to screaming delimited",
			source:    "FromPascal",
			delimiter: ".",
			want:      "FROM.PASCAL",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToScreamingDelimited(tt.source, tt.delimiter); got != tt.want {
				t.Errorf("ToScreamingDelimited() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToCamel(t *testing.T) {
	tests := []struct {
		name      string
		source    string
		delimiter string
		want      string
	}{
		{
			name:      "from kebab to camel",
			source:    "from-kebab",
			delimiter: ".",
			want:      "fromKebab",
		},
		{
			name:      "from snake to camel",
			source:    "from_snake",
			delimiter: ".",
			want:      "fromSnake",
		},
		{
			name:      "from screaming snake to camel",
			source:    "FROM_SCREAMING_SNAKE",
			delimiter: ".",
			want:      "fromScreamingSnake",
		},
		{
			name:      "from camel to camel",
			source:    "fromCamel",
			delimiter: ".",
			want:      "fromCamel",
		},
		{
			name:      "from pascal to camel",
			source:    "FromPascal",
			delimiter: ".",
			want:      "fromPascal",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToCamel(tt.source); got != tt.want {
				t.Errorf("ToCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToPascal(t *testing.T) {
	tests := []struct {
		name      string
		source    string
		delimiter string
		want      string
	}{
		{
			name:      "from kebab to pascal",
			source:    "from-kebab",
			delimiter: ".",
			want:      "FromKebab",
		},
		{
			name:      "from snake to pascal",
			source:    "from_snake",
			delimiter: ".",
			want:      "FromSnake",
		},
		{
			name:      "from screaming snake to pascal",
			source:    "FROM_SCREAMING_SNAKE",
			delimiter: ".",
			want:      "FromScreamingSnake",
		},
		{
			name:      "from camel to pascal",
			source:    "fromCamel",
			delimiter: ".",
			want:      "FromCamel",
		},
		{
			name:      "from pascal to pascal",
			source:    "FromPascal",
			delimiter: ".",
			want:      "FromPascal",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToPascal(tt.source); got != tt.want {
				t.Errorf("ToPascal() = %v, want %v", got, tt.want)
			}
		})
	}
}
