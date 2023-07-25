package apirules

import (
	"encoding/json"
	"fmt"
	"github.com/go-gota/gota/dataframe"
	log "github.com/sirupsen/logrus"
	"strings"
)

func (inst *Client) Print(x interface{}) {
	log.Error(x)
}

func (inst *Client) ToString(x interface{}) string {
	return fmt.Sprint(x)
}

func (inst *Client) PrintMany(x ...interface{}) {
	fmt.Printf("%v\n", x)
}

func (inst *Client) JsonToDF(data any) dataframe.DataFrame {
	b, err := json.Marshal(data)
	if err != nil {
		return dataframe.DataFrame{}
	}
	df := dataframe.ReadJSON(strings.NewReader(string(b)))
	return df
}

func (inst *Client) Tags(tag ...string) {
	var includeList []string
	var excludeList []string
	for _, s := range tag {
		exclude := strings.Contains(s, "!")
		if exclude {
			t := strings.Trim(s, "!")
			excludeList = append(excludeList, t)
		} else {
			includeList = append(includeList, s)
		}
	}

	for i, s := range includeList {
		fmt.Println("includeList", i, s)
	}
	for i, s := range excludeList {
		fmt.Println("excludeList", i, s)
	}

}
