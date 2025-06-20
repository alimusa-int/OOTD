package main

import "fmt"

type clothes struct {
	name       string
	categories string
	colors     string
	formality  int
	lastWorn   string
}

const NMAX = 1000

type clothing [NMAX]clothes

func main() {
	var choice, n int
	var wardrobe clothing

	option()
	fmt.Scan(&choice)
	for choice != 11 {
		switch choice {
		case 1:
			addItem(&wardrobe, &n)
		case 2:
			listItem(wardrobe, n)
			modifyItem(&wardrobe, n)
		case 3:
			listItem(wardrobe, n)
			deleteItem(&wardrobe, &n)
		case 4:
			listItem(wardrobe, n)
		case 5:
			searchCategories(wardrobe, n)
		case 6:
			searchColors(wardrobe, n)
		case 7:
			sortLastWorn(&wardrobe, n)
			searchLastWorn(wardrobe, n)
		case 8:
			sortFormality(&wardrobe, n)
			listItem(wardrobe, n)
		case 9:
			sortLastWorn(&wardrobe, n)
			listItem(wardrobe, n)
		case 10:
			outfitRecommendation(wardrobe, n)
		case 11:
			return
		default:
			fmt.Println("Invalid number")
		}
		option()
		fmt.Scan(&choice)
	}
}

func option() {
	fmt.Println("============================================================================================")
	fmt.Println("Let's Manage your OOTD")
	fmt.Println("1. Add Item")
	fmt.Println("2. Modify Item")
	fmt.Println("3. Delete Item")
	fmt.Println("4. List Item")
	fmt.Println("5. Search Item by Categories")
	fmt.Println("6. Search Item by Colors")
	fmt.Println("7. Search Item by Last Worn")
	fmt.Println("8. Sort by formality")
	fmt.Println("9. Sort by Last Worn")
	fmt.Println("10. Outfit Recommendation")
	fmt.Println("11. Exit")
	fmt.Println("============================================================================================")
	fmt.Print("Choose number: ")

}

func addItem(wardrobe *clothing, n *int) {
	if *n < NMAX {
		fmt.Print("Item name: ")
		fmt.Scan(&wardrobe[*n].name)
		choooseCategories(wardrobe, n)
		fmt.Print("Colors: ")
		fmt.Scan(&wardrobe[*n].colors)
		fmt.Print("Formality Level (1-3): ")
		fmt.Scan(&wardrobe[*n].formality)
		fmt.Print("Last Worn (YYYY-MM-DD): ")
		fmt.Scan(&wardrobe[*n].lastWorn)
		(*n)++
	} else {
		fmt.Println("Wardrobe is full.")
	}
}

func choooseCategories(wardrobe *clothing, n *int) {
	var categoriesChoose int
	var otherCategories string
	fmt.Println("Categories: ")
	fmt.Println("1. shirt")
	fmt.Println("2. trouser")
	fmt.Println("3. hat")
	fmt.Println("4. jeans")
	fmt.Println("5. jacket")
	fmt.Println("6. sweater")
	fmt.Println("7. hoodie")
	fmt.Println("8. coat")
	fmt.Println("9. t-shirt")
	fmt.Println("10. short")
	fmt.Println("11. dresses")
	fmt.Println("12. skirt")
	fmt.Println("13. sneaker")
	fmt.Println("14. sandal")
	fmt.Println("15. loafer")
	fmt.Println("16. other categories")
	fmt.Print("Choose number: ")
	fmt.Scan(&categoriesChoose)
	switch categoriesChoose {
	case 1:
		wardrobe[*n].categories = "shirt"
	case 2:
		wardrobe[*n].categories = "trouser"
	case 3:
		wardrobe[*n].categories = "hat"
	case 4:
		wardrobe[*n].categories = "jeans"
	case 5:
		wardrobe[*n].categories = "jacket"
	case 6:
		wardrobe[*n].categories = "sweater"
	case 7:
		wardrobe[*n].categories = "hoodie"
	case 8:
		wardrobe[*n].categories = "coat"
	case 9:
		wardrobe[*n].categories = "t-shirt"
	case 10:
		wardrobe[*n].categories = "short"
	case 11:
		wardrobe[*n].categories = "dresses"
	case 12:
		wardrobe[*n].categories = "skirt"
	case 13:
		wardrobe[*n].categories = "sneaker"
	case 14:
		wardrobe[*n].categories = "sandal"
	case 15:
		wardrobe[*n].categories = "loafer"
	case 16:
		fmt.Print("Enter Categories: ")
		fmt.Scan(&otherCategories)
		wardrobe[*n].categories = otherCategories
	}

}

func modifyItem(wardrobe *clothing, n int) {
	var idx int
	fmt.Print("Enter item number to modify: ")
	fmt.Scan(&idx)
	if idx <= n {
		fmt.Print("New Item name: ")
		fmt.Scan(&wardrobe[idx-1].name)
		fmt.Print("New Categories: ")
		fmt.Scan(&wardrobe[idx-1].categories)
		fmt.Print("New Colors: ")
		fmt.Scan(&wardrobe[idx-1].colors)
		fmt.Print("New Formality Level: ")
		fmt.Scan(&wardrobe[idx-1].formality)
		fmt.Print("New Last Worn Date: ")
		fmt.Scan(&wardrobe[idx-1].lastWorn)
	} else {
		fmt.Println("Item not found")
	}

}

func deleteItem(wardrobe *clothing, n *int) {
	var delete int
	fmt.Print("Which number you want to delete ? ")
	fmt.Scan(&delete)
	if delete <= *n {
		for delete-1 < *n {
			wardrobe[delete-1] = wardrobe[delete]
			delete++
		}
		(*n)--
	} else {
		fmt.Println("Item not found")
	}
}

func listItem(wardrobe clothing, n int) {
	fmt.Printf("%-4s %-20s %-15s %-10s %-15s %-12s\n", "No", "Name", "Category", "Color", "Formality lvl", "Last Worn")
	for i := 0; i < n; i++ {
		fmt.Printf("%-4d %-20s %-15s %-10s %-15d %-12s\n", i+1, wardrobe[i].name, wardrobe[i].categories, wardrobe[i].colors, wardrobe[i].formality, wardrobe[i].lastWorn)
	}
}

func searchCategories(wardrobe clothing, n int) {
	var search string
	var found int = 0
	fmt.Print("Enter Categories: ")
	fmt.Scan(&search)
	fmt.Printf("%-4s %-20s %-15s %-10s %-15s %-12s\n", "No", "Name", "Category", "Color", "Formality lvl", "Last Worn")
	for i := 0; i < n; i++ {
		if wardrobe[i].categories == search {
			fmt.Printf("%-4d %-20s %-15s %-10s %-15d %-12s\n", i+1, wardrobe[i].name, wardrobe[i].categories, wardrobe[i].colors, wardrobe[i].formality, wardrobe[i].lastWorn)
			found++
		}
	}

	if found == 0 {
		fmt.Println("Categories not found")
	}
}

func searchColors(wardrobe clothing, n int) {
	var search string
	var found int = 0
	fmt.Print("Enter Colors: ")
	fmt.Scan(&search)
	fmt.Printf("%-4s %-20s %-15s %-10s %-15s %-12s\n", "No", "Name", "Category", "Color", "Formality lvl", "Last Worn")
	for i := 0; i < n; i++ {
		if wardrobe[i].colors == search {
			fmt.Printf("%-4d %-20s %-15s %-10s %-15d %-12s\n", i+1, wardrobe[i].name, wardrobe[i].categories, wardrobe[i].colors, wardrobe[i].formality, wardrobe[i].lastWorn)
			found++
		}
	}

	if found == 0 {
		fmt.Println("Colors not found")
	}

}

func searchLastWorn(wardrobe clothing, n int) {
	var left, mid, right, found int
	var x string
	fmt.Print("Enter date you want to search: ")
	fmt.Scan(&x)
	found = -1
	left = 0
	right = n - 1

	for left <= right && found == -1 {
		mid = (left + right) / 2
		if x < wardrobe[mid].lastWorn {
			right = mid - 1
		} else if x > wardrobe[mid].lastWorn {
			left = mid + 1
		} else {
			fmt.Printf("%-4s %-20s %-15s %-10s %-15s %-12s\n", "No", "Name", "Category", "Color", "Formality lvl", "Last Worn")
			fmt.Printf("%-4d %-20s %-15s %-10s %-15d %-12s\n", mid+1, wardrobe[mid].name, wardrobe[mid].categories, wardrobe[mid].colors, wardrobe[mid].formality, wardrobe[mid].lastWorn)
			found++
		}
	}

	if found == -1 {
		fmt.Println("No item found with that last worn date.")
	}
}

func sortFormality(wardrobe *clothing, n int) {
	var pass, idx, i int
	var temp clothes
	pass = 1

	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i <= n-1 {
			if wardrobe[i].formality > wardrobe[idx].formality {
				idx = i
			}
			i++
		}
		temp = wardrobe[pass-1]
		wardrobe[pass-1] = wardrobe[idx]
		wardrobe[idx] = temp
		pass++
	}

}

func sortLastWorn(wardrobe *clothing, n int) {
	var pass, i int
	var temp clothes
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = wardrobe[pass]
		for i > 0 && (temp.lastWorn < wardrobe[i-1].lastWorn) {
			wardrobe[i] = wardrobe[i-1]
			i--
		}
		wardrobe[i] = temp
		pass++
	}

}
func outfitRecommendation(wardrobe clothing, n int) {
	fmt.Println("Outfit recommendation for rainy days:")
	darkColors := [7]string{"black", "navy", "gray", "brown", "dark green", "dark blue", "dark red"}
	rainCategories := [4]string{"jacket", "coat", "hoodie", "sweater"}
	var i, j int
	var found bool
	found = false
	fmt.Printf("%-4s %-20s %-15s %-10s %-15s %-12s\n", "No", "Name", "Category", "Color", "Formality lvl", "Last Worn")
	for i = 0; i < n; i++ {
		colorMatch := false
		categoryMatch := false

		for j = 0; j < 7; j++ {
			if wardrobe[i].colors == darkColors[j] {
				colorMatch = true
			}
		}

		for j = 0; j < 4; j++ {
			if wardrobe[i].categories == rainCategories[j] {
				categoryMatch = true
			}
		}

		if colorMatch && categoryMatch {
			fmt.Printf("%-4d %-20s %-15s %-10s %-15d %-12s\n", i+1, wardrobe[i].name, wardrobe[i].categories, wardrobe[i].colors, wardrobe[i].formality, wardrobe[i].lastWorn)
			found = true
		}
	}
	if !found {
		fmt.Println("No suitable outfits found for a rainy day.")
	}

	found = false
	fmt.Printf("\nOutfit recommendation for formal meeting:\n")
	fmt.Printf("%-4s %-20s %-15s %-10s %-15s %-12s\n", "No", "Name", "Category", "Color", "Formality lvl", "Last Worn")
	for i = 0; i < n; i++ {
		if wardrobe[i].formality == 3 {
			fmt.Printf("%-4d %-20s %-15s %-10s %-15d %-12s\n", i+1, wardrobe[i].name, wardrobe[i].categories, wardrobe[i].colors, wardrobe[i].formality, wardrobe[i].lastWorn)
			found = true
		}
	}
	if !found {
		fmt.Println("No suitable outfits found for formal meeting.")
	}
}
