package main

func main() {
	DB = NewDB("test.db")
	DB.InitDB()

	ImportNewEntries()
}
