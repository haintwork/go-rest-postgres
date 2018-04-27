package main

func main() {
	a := App{}
	// a.Initialize(
	// 	os.Getenv("APP_DB_USERNAME"),
	// 	os.Getenv("APP_DB_PASSWORD"),
	// 	os.Getenv("APP_DB_NAME"))
	a.Initialize(
		"hainguyen",
		"",
		"test")

	a.Run(":8910")
}
