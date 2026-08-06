package actioninfo

import "fmt"

type DataParser interface {
	Parse(dataString string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, trainingInfo := range dataset {
		err := dp.Parse(trainingInfo)
		if err != nil {
			fmt.Println(err)
			continue
		}

		trainingData, err := dp.ActionInfo()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(trainingData)
	}
}
