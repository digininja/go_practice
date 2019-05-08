package main

import "fmt"
import "strings"

func main() {
	ree := strings.NewReplacer("<registration>", "<subitem name=\"registration\">", "</registration>", "</subitem>", "<postcode>", "<subitem name=\"postcode\">", "</postcode>", "</subitem>", "<first_line>", "<subitem name=\"first_line\">", "</first_line>", "</subitem>", "<second_line>", "<subitem name=\"second_line\">", "</second_line>", "</subitem>", "<town>", "<subitem name=\"town\">", "</town>", "</subitem>", "<county>", "<subitem name=\"county\">", "</county>", "</subitem>")
	fmt.Printf(ree.Replace("<registration>hello"))
}
