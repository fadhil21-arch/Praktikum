package main
import "fmt"

func main(){
	var r, d1, d2, d3 int
	var membesar, mengecil, status bool
	fmt.Scan(&r)
	d1 = r / 100
	d2 = (r / 10) % 10
	d3 = r % 10

	membesar = d1 < d2 && d2 < d3
	mengecil = d1 > d2 && d2 > d3
	status = membesar || mengecil
	fmt.Println(status)


}