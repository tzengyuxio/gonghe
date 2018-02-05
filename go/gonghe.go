package main

import (
	"fmt"
	"math"

	"github.com/soniakeys/meeus/julian"
	pp "github.com/soniakeys/meeus/planetposition"
	"github.com/soniakeys/meeus/solstice"
)

var gan = []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
var zhi = []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}

func ganzhi(n int) string {
	g := n % 10
	if g < 0 {
		g += 10
	}
	z := n % 12
	if z < 0 {
		z += 12
	}
	return gan[g] + zhi[z]
}

func depart(n float64) (int, float64) {
	i := math.Floor(n) // integer part
	f := n - i         // fractional part
	return int(i), f
}

func ganzhiOfDay(jd float64) string {
	d := int(math.Floor(jd+.5)) - 11
	return ganzhi(d)
}

func ganzhiOfYear(year int) string {
	y := (year - 4)
	return ganzhi(y)
}

func lichunOfYear(year int) {
	e, err := pp.LoadPlanetPath(pp.Earth, "/Users/tzengyuxio/SDK/VI_81")
	if err != nil {
		fmt.Println(err)
		return
	}
	y := year
	if year < 0 {
		y++
	}
	eq := solstice.March2(y, e)
	sol := solstice.December2(y-1, e)
	lichun := (eq + sol) / 2
	y, m, dt := julian.JDToCalendar(lichun)
	d, _ := depart(dt)
	fmt.Printf("%d-%d-%d\n", year, m, d)
}

func testOfZeroYear() {
	d1 := julian.CalendarGregorianToJD(1, 1, 1)
	d2 := julian.CalendarJulianToJD(1, 1, 1)
	d3 := julian.CalendarJulianToJD(0, 1, 1)
	d1p := d1 - 1.0 // day 1 previous day
	d2p := d2 - 1.0 // day 2 previous day
	d3p := d3 - 1.0 // day 3 previous day
	y, m, dt := julian.JDToCalendar(d1p)
	fmt.Printf("%d-%d-%f\n", y, m, dt) // 1-1-2
	y, m, dt = julian.JDToCalendar(d2p)
	fmt.Printf("%d-%d-%f\n", y, m, dt) // 0-12-31
	y, m, dt = julian.JDToCalendar(d3p)
	fmt.Printf("%d-%d-%f\n", y, m, dt) // -1-12-31
}

func main() {
	// testOfZeroYear()
	lichunOfYear(-841)
	lichunOfYear(-842)
}
