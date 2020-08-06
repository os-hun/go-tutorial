package named_results

/*
	Named return values
	@see https://tour.golang.org/basics/7
*/
func Split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
