package main

func main() {
	settings := new(Config)
	settings.Configure()
	runServer(settings)
}
