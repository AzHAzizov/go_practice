package models

var DB []Book

type Book struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	YearPublished int    `json:"year_published"`
	Author        Author `json:"author"`
}

type Author struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	BornYear int    `json:"born_year"`
}

func init() {
	book1 := Book{
		ID:            1,
		Title:         "Lord of the Rings. Vol.1",
		YearPublished: 1978,
		Author: Author{
			Name:     "J.R",
			LastName: "Tolkin",
			BornYear: 1892,
		},
	}

	DB = append(DB, book1)
}

func FindBookKey(id int) (key int, found bool) {

	for k, book := range DB {
		if book.ID == id {
			key = k
			found = true
			break
		}
	}

	return key, found
}

func RemoveFromBooks(key int) []Book {
	DB[key] = DB[len(DB)-1]
	DB[len(DB)-1] = Book{}
	DB = DB[:len(DB)-1]

	return DB
}

func FindBookById(id int) (Book, bool) {
	var book Book
	var found bool

	for _, b := range DB {
		if b.ID == id {
			book = b
			found = true
		}
	}

	return book, found
}
