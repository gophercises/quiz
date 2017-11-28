// returns a map of csv file having two string columns
package reader

import (
	"bufio"
	"os"
	"strings"
)

func ReadCsv(file string) (ans map[string]string) {
	f, err := os.Open(file)
	defer f.Close()
	Check(err)
	reader := bufio.NewReader(f)
	ans = make(map[string]string)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		//fmt.Println(line)
		arr := strings.Split(line, ",")
		ans[arr[0]] = arr[1]
	}
	return
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
