package utility

import(
	"log"
	"os"
	"strings"
	"github.com/google/uuid"
)

//generate a random uuid
func IdGenerator() string {
	id := uuid.New().String()
	return id
}

//CreateDirectory creates a directory at specified location
func CreateDirectory(pathToDirectory string,DirectoryName string){
	fullPath :=pathToDirectory+"/"+DirectoryName
	err :=os.MkdirAll(fullPath,os.ModePerm)
	if err!=nil{
		log.Fatal(err)
	}
}

//FetchFileNameUrl return the fileName of the Url
func FetchFileNameFromUrl(url string)string{
	s := strings.Split(url, "/")
	return s[len(s)-1]
}

func FindExtensionOfFile(file string)string{
	s := strings.Split(file,".")
	return s[len(s)-1]
}


