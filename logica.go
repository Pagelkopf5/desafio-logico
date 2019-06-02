package main

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			println("CoBlue")
		} else if i%3 == 0 {
			println("Co")
		} else if i%5 == 0 {
			println("Blue")
		} else {
			println(i)
		}
	}
}
