package lib

import (
	"fmt"
	"html/template"
	"os"
)

func Basic() {
	templateString := `This will be replaced by a template`
	t, err := template.New("title").Parse(templateString)
	if err != nil {
		fmt.Printf("Error in creating a template %v", err)
	}
	err = t.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Printf("Error in creating a template %v", err)
	}
}
