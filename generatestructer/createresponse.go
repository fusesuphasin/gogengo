package generatestructer

import (
	"fmt"
	"gogenerate/generatestructer/gotemplate"
	"os"
	"time"
)

func CreateResponseTemplate(path string) {
	//write respone struct file
	controllerFileName := fmt.Sprintf("%v/app/utils/response/response.go", path)
	res, err := os.Create(controllerFileName)
	if(err!=nil){
		fmt.Println(err)
	}
	
	gotemplate.RespackageTemplate.Execute(res, struct {
		Timestamp time.Time
		DoubleQuote string
	}{
		Timestamp: time.Now(),
		DoubleQuote: "`",
	})
}