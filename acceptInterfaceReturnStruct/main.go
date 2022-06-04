package main

// usecaseはidから本を検索する

type BookFetcher interface {
	GetById()
}

func main() {
	featcher := NewBookDb()
	SearchBook(featcher)
}

func SearchBook(fetcher BookFetcher) {
	fetcher.GetById()
}
