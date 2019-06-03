package main

func main() {
	for i := 12; i <= 100; i++ {
		x := false 
		if i%3 == 0 {
			x = true
			print("Co")
		} 
		if i%5 == 0 {
			x = true
			print("Blue")
		} 
		if(x == false) {
			print(i)
		}
		println(" ")
	}
}
