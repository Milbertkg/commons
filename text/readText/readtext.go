package readtext

//Imports programs and functions
import (
	"log"
	"os"
)

func ReadText(text_route string) string {

	//read file
	content, err := os.ReadFile(text_route)

	if err != nil {
		log.Fatal(err)
	}

	return (string(content))

}
