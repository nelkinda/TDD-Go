package main

func main() {
	a := App{}
	a.Initialize("docker", "docker", "rest_api_example")

	a.Run(":8080")
}
