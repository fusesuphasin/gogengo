package generatestructer

import (
	"fmt"
	"gogenerate/generate/generatestructer/gotemplate"
	"log"
	"os"
	"time"
)

type Response struct{
	
}

func (rp *Response) CreateResponseTemplate(path string) {
	//write respone struct file
	controllerFileName := fmt.Sprintf("%v/app/utils/response/response.go", path)
	res, err := os.Create(controllerFileName)
	if(err!=nil){
		log.Println(err)
	}
	
	gotemplate.RespackageTemplate.Execute(res, struct {
		Timestamp time.Time
		DoubleQuote string
	}{
		Timestamp: time.Now(),
		DoubleQuote: "`",
	})
}