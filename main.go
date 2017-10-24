package main

func main() {
	app := App{}

	app.Initialize()
	app.Router.Run()
}
