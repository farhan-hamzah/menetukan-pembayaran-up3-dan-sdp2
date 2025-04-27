package main
import "fmt"

func main(){
	var n int
	fmt.Scan(&n)
	pembayaran(n)
}
func pembayaran(n int){
	var mtk, indo, inggris, rata, up3, sdp2 float64
	var osn, o2sn bool
	for n > 0 {
		fmt.Scan(&mtk, &indo, &inggris, &osn, &o2sn)
		rata = rerata(mtk, indo, inggris)
		if rata >= 75{
			up3 = 10000000
		}else{
			up3 = 15000000
		}
		if osn == true && o2sn == true{
			sdp2 = 5000000
		}else if osn == false && o2sn == true {
			sdp2 = 7000000
		}else{
			sdp2 = 10000000
		}
		fmt.Print(up3, " ", sdp2)
		n-=1
	}
	
}
func rerata(mtk, indo, inggris float64)float64{
	return (mtk + indo + inggris)/3
}