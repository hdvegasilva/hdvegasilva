package drum

import (
	"fmt"
	"path"
	"testing"
)

/*
00 00 00 00 04

6b69636b(instrument-kick)
| 01-00-00-00 | 01-00-00-00 | 01-00-00-00 | 01-00-00-00 |

01 00 00 00 05

hex(instrument-snare)
| 00-00-00-00 | 01-00-00-00 | 00-00-00-00 | 01-00-00-00 |

02 00 00 00 04

hex(instrument-clap)
| 00-00-00-00 | 01-00-01-00 | 00-00-00-00 | 00-00-00-00 |

03 00 00 00 07

hex(instrument-hh-open)
| 00-00-01-00 | 00-00-01-00 | 01-00-01-00 | 00-00-01-00 |

04 00 00 00 08

hex(instrument-hh-close)
| 01-00-00-00 | 01-00-00-00 | 00-00-00-00 | 01-00-00-01 |

05 00 00 00 07

hex(instrument-cowbell)
| 00-00-00-00 | 00-00-00-00 | 00-00-01-00 | 00-00-00-00 |
*/

func TestDecodeFile(t *testing.T) {
	tData := []struct {
		path   string
		output string
	}{
		{"pattern_1.splice",
			`Saved with HW Version: 0.808-alpha
Tempo: 120
(0) kick	|x---|x---|x---|x---|
(1) snare	|----|x---|----|x---|
(2) clap	|----|x-x-|----|----|
(3) hh-open	|--x-|--x-|x-x-|--x-|
(4) hh-close	|x---|x---|----|x--x|
(5) cowbell	|----|----|--x-|----|
`,
		},
		{"pattern_2.splice",
			`Saved with HW Version: 0.808-alpha
		Tempo: 98.4
		(0) kick	|x---|----|x---|----|
		(1) snare	|----|x---|----|x---|
		(3) hh-open	|--x-|--x-|x-x-|--x-|
		(5) cowbell	|----|----|x---|----|
		`,
		},
		/*{"pattern_3.splice",
					`Saved with HW Version: 0.808-alpha
		Tempo: 118
		(40) kick	|x---|----|x---|----|
		(1) clap	|----|x---|----|x---|
		(3) hh-open	|--x-|--x-|x-x-|--x-|
		(5) low-tom	|----|---x|----|----|
		(12) mid-tom	|----|----|x---|----|
		(9) hi-tom	|----|----|-x--|----|
		`,
				},
				{"pattern_4.splice",
					`Saved with HW Version: 0.909
		Tempo: 240
		(0) SubKick	|----|----|----|----|
		(1) Kick	|x---|----|x---|----|
		(99) Maracas	|x-x-|x-x-|x-x-|x-x-|
		(255) Low Conga	|----|x---|----|x---|
		`,
				},
				{"pattern_5.splice",
					`Saved with HW Version: 0.708-alpha
		Tempo: 999
		(1) Kick	|x---|----|x---|----|
		(2) HiHat	|x-x-|x-x-|x-x-|x-x-|
		`,
				},*/
	}

	for _, exp := range tData {
		decoded, err := DecodeFile(path.Join("fixtures", exp.path))
		if err != nil {
			t.Fatalf("something went wrong decoding %s - %v", exp.path, err)
		} else {
			fmt.Print(decoded)
		}
		/*if fmt.Sprint(decoded) != exp.output {
			t.Logf("decoded:\n%#v\n", fmt.Sprint(decoded))
			t.Logf("expected:\n%#v\n", exp.output)
			t.Fatalf("%s wasn't decoded as expect.\nGot:\n%s\nExpected:\n%s",
				exp.path, decoded, exp.output)
		}
		*/
	}
}
