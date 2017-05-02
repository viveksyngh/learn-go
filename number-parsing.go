package main
import "fmt"
import "strconv"

func main() {
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)
	h, _ := strconv.ParseInt("0xf45", 0, 64)
	fmt.Println(h)
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)
	k, _ := strconv.Atoi("123")
	fmt.Println(k)
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}