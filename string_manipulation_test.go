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
				t.Errorf("split() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestConversions(t *testing.T) {
	type want struct {
		delimited          string
		screamingDelimited string
		kebab              string
		screamingKebab     string
		snake              string
		screamingSnake     string
		camel              string
		pascal             string
	}
	tests := []struct {
		name   string
		source string
		want   want
	}{
		{
			name:   "from period",
			source: "lorem.ipsum.$.sat.12",
			want: want{
				delimited:          "lorem.ipsum.$.sat.12",
				screamingDelimited: "LOREM.IPSUM.$.SAT.12",
				kebab:              "lorem-ipsum-$-sat-12",
				screamingKebab:     "LOREM-IPSUM-$-SAT-12",
				snake:              "lorem_ipsum_$_sat_12",
				screamingSnake:     "LOREM_IPSUM_$_SAT_12",
				camel:              "loremIpsum$Sat12",
				pascal:             "LoremIpsum$Sat12",
			},
		},
		{
			name:   "from double under",
			source: "lorem__ipsum__$__sat__12",
			want: want{
				delimited:          "lorem.ipsum.$.sat.12",
				screamingDelimited: "LOREM.IPSUM.$.SAT.12",
				kebab:              "lorem-ipsum-$-sat-12",
				screamingKebab:     "LOREM-IPSUM-$-SAT-12",
				snake:              "lorem_ipsum_$_sat_12",
				screamingSnake:     "LOREM_IPSUM_$_SAT_12",
				camel:              "loremIpsum$Sat12",
				pascal:             "LoremIpsum$Sat12",
			},
		},
		{
			name:   "from screaming period",
			source: "LOREM.IPSUM.$.SAT.12",
			want: want{
				delimited:          "lorem.ipsum.$.sat.12",
				screamingDelimited: "LOREM.IPSUM.$.SAT.12",
				kebab:              "lorem-ipsum-$-sat-12",
				screamingKebab:     "LOREM-IPSUM-$-SAT-12",
				snake:              "lorem_ipsum_$_sat_12",
				screamingSnake:     "LOREM_IPSUM_$_SAT_12",
				camel:              "loremIpsum$Sat12",
				pascal:             "LoremIpsum$Sat12",
			},
		},
		{
			name:   "from kebab",
			source: "lorem-ipsum-$-sat-12",
			want: want{
				delimited:          "lorem.ipsum.$.sat.12",
				screamingDelimited: "LOREM.IPSUM.$.SAT.12",
				kebab:              "lorem-ipsum-$-sat-12",
				screamingKebab:     "LOREM-IPSUM-$-SAT-12",
				snake:              "lorem_ipsum_$_sat_12",
				screamingSnake:     "LOREM_IPSUM_$_SAT_12",
				camel:              "loremIpsum$Sat12",
				pascal:             "LoremIpsum$Sat12",
			},
		},
		{
			name:   "from screaming kebab",
			source: "LOREM-IPSUM-$-SAT-12",
			want: want{
				delimited:          "lorem.ipsum.$.sat.12",
				screamingDelimited: "LOREM.IPSUM.$.SAT.12",
				kebab:              "lorem-ipsum-$-sat-12",
				screamingKebab:     "LOREM-IPSUM-$-SAT-12",
				snake:              "lorem_ipsum_$_sat_12",
				screamingSnake:     "LOREM_IPSUM_$_SAT_12",
				camel:              "loremIpsum$Sat12",
				pascal:             "LoremIpsum$Sat12",
			},
		},
		{
			name:   "from snake",
			source: "lorem_ipsum_$_sat_12",
			want: want{
				delimited:          "lorem.ipsum.$.sat.12",
				screamingDelimited: "LOREM.IPSUM.$.SAT.12",
				kebab:              "lorem-ipsum-$-sat-12",
				screamingKebab:     "LOREM-IPSUM-$-SAT-12",
				snake:              "lorem_ipsum_$_sat_12",
				screamingSnake:     "LOREM_IPSUM_$_SAT_12",
				camel:              "loremIpsum$Sat12",
				pascal:             "LoremIpsum$Sat12",
			},
		},
		{
			name:   "from screaming snake",
			source: "LOREM_IPSUM_$_SAT_12",
			want: want{
				delimited:          "lorem.ipsum.$.sat.12",
				screamingDelimited: "LOREM.IPSUM.$.SAT.12",
				kebab:              "lorem-ipsum-$-sat-12",
				screamingKebab:     "LOREM-IPSUM-$-SAT-12",
				snake:              "lorem_ipsum_$_sat_12",
				screamingSnake:     "LOREM_IPSUM_$_SAT_12",
				camel:              "loremIpsum$Sat12",
				pascal:             "LoremIpsum$Sat12",
			},
		},
		{
			name:   "from camel",
			source: "loremIpsum$Sat12",
			want: want{
				delimited:          "lorem.ipsum.$.sat.12",
				screamingDelimited: "LOREM.IPSUM.$.SAT.12",
				kebab:              "lorem-ipsum-$-sat-12",
				screamingKebab:     "LOREM-IPSUM-$-SAT-12",
				snake:              "lorem_ipsum_$_sat_12",
				screamingSnake:     "LOREM_IPSUM_$_SAT_12",
				camel:              "loremIpsum$Sat12",
				pascal:             "LoremIpsum$Sat12",
			},
		},
		{
			name:   "from pascal",
			source: "LoremIpsum$Sat12",
			want: want{
				delimited:          "lorem.ipsum.$.sat.12",
				screamingDelimited: "LOREM.IPSUM.$.SAT.12",
				kebab:              "lorem-ipsum-$-sat-12",
				screamingKebab:     "LOREM-IPSUM-$-SAT-12",
				snake:              "lorem_ipsum_$_sat_12",
				screamingSnake:     "LOREM_IPSUM_$_SAT_12",
				camel:              "loremIpsum$Sat12",
				pascal:             "LoremIpsum$Sat12",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToDelimited(tt.source, "."); got != tt.want.delimited {
				t.Errorf("invalid result for delimited: got %s - want %s", got, tt.want.delimited)
			}
			if got := ToScreamingDelimited(tt.source, "."); got != tt.want.screamingDelimited {
				t.Errorf("invalid result for screaming delimited: got %s - want %s", got, tt.want.screamingDelimited)
			}
			if got := ToKebab(tt.source); got != tt.want.kebab {
				t.Errorf("invalid result for kebab: got %s - want %s", got, tt.want.screamingDelimited)
			}
			if got := ToScreamingKebab(tt.source); got != tt.want.screamingKebab {
				t.Errorf("invalid result for screaming kebab: got %s - want %s", got, tt.want.screamingKebab)
			}
			if got := ToSnake(tt.source); got != tt.want.snake {
				t.Errorf("invalid result for snake: got %s - want %s", got, tt.want.snake)
			}
			if got := ToScreamingSnake(tt.source); got != tt.want.screamingSnake {
				t.Errorf("invalid result for screaming snake: got %s - want %s", got, tt.want.screamingSnake)
			}
			if got := ToCamel(tt.source); got != tt.want.camel {
				t.Errorf("invalid result for camel: got %s - want %s", got, tt.want.camel)
			}
			if got := ToPascal(tt.source); got != tt.want.pascal {
				t.Errorf("invalid result for pascal: got %s - want %s", got, tt.want.pascal)
			}
		})
	}
}
