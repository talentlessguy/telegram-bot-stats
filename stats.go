package stats

import (
	"io/ioutil"
	"log"
	"os"

	"encoding/json"
)

// statFile - stat filename
const statFile = "stat.json"

const defaultJSON = `
{
	"count": 0,
	"ids": []
}
`

func statfileExists() bool {
	info, err := os.Stat(statFile)

	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Stat - stat.json struct
type Stat struct {
	Count int   `json:"count"`
	IDs   []int `json:"ids"`
}

// ParseJSON - parse stat.json file
func ParseStatJSON() Stat {
	file, err := os.Open(statFile)

	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)

	if err != nil {
		log.Println(err)
	}

	var v Stat

	json.Unmarshal(b, &v)

	return v
}

// ReadStatJSON - read json file and return parsed count and ids
func ReadStatJSON() Stat {
	if statfileExists() {
		v := ParseStatJSON()

		return Stat{
			Count: v.Count,
			IDs:   v.IDs,
		}
	} else {
		err := ioutil.WriteFile("stat.json", []byte(defaultJSON), 0644)

		if err != nil {
			log.Println(err)
		}

		return Stat{}
	}
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// AddUserToStat - add user ID to stat.json file
func AddUserToStat(id int) {

	stat := ReadStatJSON()

	if !contains(stat.IDs, id) {

		stat.Count++
		stat.IDs = append(stat.IDs, id)

		outputJSON, _ := json.Marshal(stat)

		err := ioutil.WriteFile("stat.json", outputJSON, 0644)

		if err != nil {
			log.Println(err)
		}
	}
}
