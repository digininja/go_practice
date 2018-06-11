package main

import "fmt"
import "regexp"

func main() {
	fmt.Println("vim-go")
	var year_end = regexp.MustCompile(`((19[0-9][0-9])|(20[012][0-9]))$`)

	fmt.Printf("%q\n", year_end.FindString("abc1999"))
	fmt.Printf("%q\n", year_end.FindString("1999abc"))
	fmt.Printf("%q\n", year_end.FindString("1999"))
	fmt.Printf("%q\n", year_end.FindString("asdsdfasdf"))
	fmt.Printf("%q\n", year_end.FindString("2020x"))
	fmt.Printf("%q\n", year_end.FindString("x2020"))

	fmt.Println("first cap last number")
	first_cap_last_num_re := regexp.MustCompile("^[A-Z].*[0-9]$")

	fmt.Printf("%q\n", first_cap_last_num_re.MatchString("Axxx9"))

	fmt.Println("first cap last symbol")

	first_cap_last_symbol_re := regexp.MustCompilePOSIX("^[A-Z].*[[:punct:]]$")

	fmt.Printf("%q\n", first_cap_last_symbol_re.MatchString("Axxx9"))
	fmt.Printf("%q\n", first_cap_last_symbol_re.MatchString("Axxx$"))
	fmt.Printf("%q\n", first_cap_last_symbol_re.MatchString("Axxx%"))
	fmt.Printf("%q\n", first_cap_last_symbol_re.MatchString("Axxx"))

}
