package main

import "fmt"

type Hit struct {
	source string
	value  string
}

type Hits struct {
	hits []Hit
	Name string
}

/*
If not passing by pointer, a copy of the object is passed.
When this happens, any changes are lost when the function
ends as the copy is no longer needed and so garbage collected.

To confirm it is a different object to the one in main,
check its address.

*/
func (h Hits) AddHitNoStick(source string, value string) {
	fmt.Printf("Address in AddHitNoStick %p\n", &h)
	fmt.Printf("Name: %s\n", h.Name)
	hit := Hit{source, value}
	fmt.Printf("Appending %s\n", value)
	fmt.Printf("Hit %u\n", hit)
	h.hits = append(h.hits, hit)

	fmt.Printf("hits %u\n", h.hits)
	h.Name = value
	fmt.Printf("Name after: %s\n", h.Name)
}

/*
For changes to be made in here, you MUST pass by pointer,
that is what the (h *Hits) means and why everything is referenced
as (*h) rather than just h.

To confirm this is the same object as in main, check the
address.
*/
func (h *Hits) AddHitStick(source string, value string) {
	fmt.Printf("Address in AddHitStick %p\n", h)
	fmt.Printf("Name: %s\n", (*h).Name)
	hit := Hit{source, value}
	fmt.Printf("Appending %s\n", value)
	fmt.Printf("Hit %u\n", hit)
	(*h).hits = append((*h).hits, hit)

	fmt.Printf("hits %u\n", (*h).hits)
	(*h).Name = value
	fmt.Printf("Name after: %s\n", (*h).Name)
}

func (h Hits) GetUniqueValues() []string {
	/*
		It is supposed to be more efficient to do a map
		and check if a key exists than to check if an element
		exists in a slice. So this creates a map of empty
		structs using the value I want to check as the key
	*/
	value_map := map[string]struct{}{}

	for _, hit := range h.hits {
		fmt.Printf("Looking at %s\n", hit.value)
		if _, ok := value_map[hit.value]; !ok {
			value_map[hit.value] = struct{}{}
		}

	}

	/*
		This converts the keys into a slice that can
		then be returned as a slice of strings.
	*/
	keys := make([]string, len(value_map))

	i := 0
	for k := range value_map {
		keys[i] = k
		i++
	}

	return keys
}

func main() {
	hits := Hits{}
	fmt.Printf("Address in main %p\n", &hits)

	hits.AddHitNoStick("ns0", "1.2.3.4")
	fmt.Println()
	hits.AddHitStick("ns1", "1.2.3.4")
	fmt.Println()
	hits.AddHitStick("ns2", "1.2.3.4")
	fmt.Println()
	hits.AddHitStick("nsdefault", "1.2.3.4")
	fmt.Println()
	hits.AddHitStick("ns1", "2.3.4.5")
	fmt.Println()
	hits.AddHitStick("ns2", "3.4.5.6")
	fmt.Println()

	unique_values := hits.GetUniqueValues()

	fmt.Printf("hits: %u\n", unique_values)

	fmt.Println("Unique Values")
	for _, value := range unique_values {
		fmt.Printf("Value: %s\n", value)
	}
	fmt.Println()
}
