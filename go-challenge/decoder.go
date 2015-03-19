package drum

import (
	"encoding/hex"
	//"encoding/json"
	"fmt"
	"io/ioutil"
	//"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// DecodeFile decodes the drum machine file found at the provided path
// and returns a pointer to a parsed pattern which is the entry point to the
// rest of the data.
// TODO: implement
//func DecodeFile(path string) (*Pattern, error) {
func DecodeFile(path string) (*Pattern, error) {
	data, err := ioutil.ReadFile(path)
	//check(err)
	//fmt.Print(string(dat))

	//dec := json.NewDecoder(strings.NewReader(string(dat)))
	//fmt.Print("Dec: ")
	//fmt.Print(dec)
	//fmt.Print([]byte(hex.EncodeToString(data)))

	//s := hex.EncodeToString(data)
	//fmt.Println(s)

	/*decodedData, dDataError := hex.DecodeString(s)

	if dDataError != nil {
		err = dDataError
	} else {
		fmt.Println(decodedData)
	}*/
	fmt.Println(hex.EncodeToString(data))
	//p := &Pattern{s}
	//p := data
	p := &Pattern{hex.Dump(data)}
	//p := hex.EncodeToString(data)
	//fmt.Print(p)

	return p, err
}

// Pattern is the high level representation of the
// drum pattern contained in a .splice file.
// TODO: implement
// Pattern is an exported type.
type Pattern struct {
	fileData string
}

/*Version string
Tempo   int
Tracks  MusicTracks*/

// Track is an exported type.
type Track struct {
	ID    string
	Name  string
	Steps string
}

// MusicTracks is an exported type.
type MusicTracks []Track
