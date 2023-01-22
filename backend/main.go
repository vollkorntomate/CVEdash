package main

func main() {
	Config = LoadConfig("config.json")
	DB = NewDB(Config.DBFilePath)
	DB.InitDB()

	ImportNewEntries()
}
