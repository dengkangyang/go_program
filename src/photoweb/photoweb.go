package main
/*
import(
	"io"
	"log"
	"path"
	"net/http"
	"io/ioutil"
	"html/template"
	"runtime/debug"
)

const(
	ListDir = 0x0001
	UPLOAD_DIR = "./uploads"
	TEMPLATE_DIR = "./views"
)

templates := make(map[string]*template.Template)

func init(){
	fileInfoArr, err :=  ioutil.ReadDir(TEMPLATE_DIR)
	check(err)

	var templateName, templatePath string
	for _, fileInfo := range fileInfoArr{
		templateName = fileInfo.Name
		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}

		templatePath = TEMPLATE_DIR + "/" + templateName
		log.Println("Loading template:", templatePath)
		t := template.Must(template.ParseFiles(templatePath))
		templates[temp1] = t
	}
}
*/

func main(){
	
}