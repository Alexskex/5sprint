package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, item := range dataset {
		if err := dp.Parse(item); err != nil {
			log.Print(err)
			continue
		}

		output, err := dp.ActionInfo()
		if err != nil {
			log.Print(err)
			continue
		}

		fmt.Print(output)
	}
}

