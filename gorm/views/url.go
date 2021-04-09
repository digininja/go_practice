package views

import (
	"fmt"
	"github.com/digininja/go_practice/gorm/models"
)

func DumpURL(url models.URL) {
	fmt.Printf("ID: %d\nURL: %s\nStatus: %s\n", url.ID, url.URL, url.Status)
}
