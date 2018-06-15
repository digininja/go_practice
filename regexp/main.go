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

	fmt.Printf("%t\n", first_cap_last_num_re.MatchString("Axxx9"))

	fmt.Println("first cap last symbol")

	first_cap_last_symbol_re := regexp.MustCompilePOSIX("^[A-Z].*[[:punct:]]$")

	fmt.Printf("%t\n", first_cap_last_symbol_re.MatchString("Axxx9"))
	fmt.Printf("%t\n", first_cap_last_symbol_re.MatchString("Axxx$"))
	fmt.Printf("%t\n", first_cap_last_symbol_re.MatchString("Axxx%"))
	fmt.Printf("%t\n", first_cap_last_symbol_re.MatchString("Axxx"))

	// https://golang.org/pkg/regexp/syntax/
	const Symbols_Regexp = "=!\"£$%^&*()\\-_\\[\\]#~'@;:<>,./?{}"

	fmt.Println("\n\npunctuation\n\n")
	re := regexp.MustCompile(fmt.Sprintf("^[%s]+$", Symbols_Regexp))
	fmt.Printf("%t\n", re.MatchString("Axxx"))
	fmt.Printf("%t\n", re.MatchString("..."))
	fmt.Printf("%t\n", re.MatchString("!!"))
	fmt.Printf("%t\n", re.MatchString("@:!"))
	fmt.Printf("%t\n", re.MatchString("\"aaa\""))
	fmt.Printf("%t\n", re.MatchString("\"\""))
	fmt.Printf("%t\n", re.MatchString("£$%"))

	fmt.Println("\n\ncombine\n\n")
	re_az := regexp.MustCompile("[a-z]+")
	re_AZ := regexp.MustCompile("[A-Z]+")
	re_09 := regexp.MustCompile("[0-9]+")
	re_symb := regexp.MustCompile(fmt.Sprintf("[%s]+", Symbols_Regexp))
	re_exclude := regexp.MustCompile(fmt.Sprintf("^[A-Za-z0-9%s]+$", Symbols_Regexp))
	s := "abcA$3"
	if re_az.MatchString(s) && re_AZ.MatchString(s) && re_09.MatchString(s) && re_symb.MatchString(s) && re_exclude.MatchString(s) {
		fmt.Printf("Match\n")
	} else {
		fmt.Printf("No Match\n")
	}
	s = "abc$3"
	if re_az.MatchString(s) && re_AZ.MatchString(s) && re_09.MatchString(s) && re_symb.MatchString(s) && re_exclude.MatchString(s) {
		fmt.Printf("Match\n")
	} else {
		fmt.Printf("No Match\n")
	}
	s = "ABC$3"
	if re_az.MatchString(s) && re_AZ.MatchString(s) && re_09.MatchString(s) && re_symb.MatchString(s) && re_exclude.MatchString(s) {
		fmt.Printf("Match\n")
	} else {
		fmt.Printf("No Match\n")
	}
	s = "AB$3"
	if re_az.MatchString(s) && re_AZ.MatchString(s) && re_09.MatchString(s) && re_symb.MatchString(s) && re_exclude.MatchString(s) {
		fmt.Printf("Match\n")
	} else {
		fmt.Printf("No Match\n")
	}
	s = "Aaaaaaaaaaaaaaaaaaaaaaaaaaaa.B$3"
	if re_az.MatchString(s) && re_AZ.MatchString(s) && re_09.MatchString(s) && re_symb.MatchString(s) && re_exclude.MatchString(s) {
		fmt.Printf("Match\n")
	} else {
		fmt.Printf("No Match\n")
	}
	s = "aB4$"
	if re_az.MatchString(s) && re_AZ.MatchString(s) && re_09.MatchString(s) && re_symb.MatchString(s) && re_exclude.MatchString(s) {
		fmt.Printf("Match\n")
	} else {
		fmt.Printf("No Match\n")
	}

	fmt.Println("\n\nnumbers on end\n\n")

	s = "abc3"
	re = regexp.MustCompile(".*([0-9]{1})$")
	hit := re.FindString(s)
	if hit != "" {
		fmt.Printf("hit: %s\n", hit)
	}

	re = regexp.MustCompile(".*([0-9]{1})$")
	fmt.Println(re.FindStringIndex(s))
	loc := (re.FindStringIndex(s))

	fmt.Println(s[loc[0]:loc[1]])

	fmt.Println(re.FindStringSubmatch(s))
	s = "xxx"
	fmt.Println(re.FindStringSubmatch(s))

	s = "xxx44"
	fmt.Println(re.FindStringSubmatch(s))
	re = regexp.MustCompile(".*([0-9]{2})$")
	s = "xxx44"
	fmt.Println(re.FindStringSubmatch(s))
}
