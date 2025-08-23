package actioninfo

import "fmt"

type DataParser interface {
	Parse(string) error
	ActionInfo(string) error
}

func Info(dataset []string, dp DataParser) {
	for _, line := range dataset {
		varible := dp.Parse(line)

		if varible != nil {
			fmt.Errorf(varible.Error())
			continue
		}

		varible2 := dp.ActionInfo(line)
		if varible2 != nil {
			fmt.Errorf(varible2.Error())
		}
	}
}
