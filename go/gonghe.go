package main

import (
	"fmt"
	"math"

	"github.com/soniakeys/meeus/julian"
	pp "github.com/soniakeys/meeus/planetposition"
	"github.com/soniakeys/meeus/solar"
	"github.com/soniakeys/meeus/solstice"
	"github.com/soniakeys/sexagesimal"
	"github.com/soniakeys/unit"
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

func pingqiToDingqi(e *pp.V87Planet, jd, degree float64) float64 {
	year, _, _ := julian.JDToCalendar(jd)
	angle := unit.AngleFromDeg(degree)

	return calcSolarTerms(e, year-1, angle)
}

func calcSolarTerms(e *pp.V87Planet, year int, angle unit.Angle) float64 {
	var jd0, jd1 float64
	var stDegree, stDegreep unit.Angle

	jd1 = solstice.March2(year, e) + 365*float64(angle)/360.0
	for ok := true; ok; ok = (math.Abs(jd1-jd0) > 0.0000001) {
		jd0 = jd1
		stDegree, _, _ = solar.ApparentVSOP87(e, jd0)
		stDegree -= angle
		stDegreep1, _, _ := solar.ApparentVSOP87(e, jd0+0.000005)
		stDegreep2, _, _ := solar.ApparentVSOP87(e, jd0-0.000005)
		stDegreep = (stDegreep1 - stDegreep2) / 0.00001
		jd1 = jd0 - float64(stDegree/stDegreep)
	}

	return jd1
}

func ganzhiOfDay(jd float64) string {
	d := int(math.Floor(jd+.5)) - 11
	return ganzhi(d)
}

func ganzhiOfYear(year int) string {
	if year <= 0 {
		year++
	}
	y := (year - 4)
	return ganzhi(y)
}

func lichunOfYear(year int) {
	e, err := pp.LoadPlanetPath(pp.Earth, "./vendor/github.com/ctdk/vsop87")
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
	lichun = pingqiToDingqi(e, lichun, 315.0)
	y, m, dt := julian.JDToCalendar(lichun)
	d, t := depart(dt)
	dayGZ := ganzhiOfDay(lichun)
	yearGZ := ganzhiOfYear(year)
	fmt.Printf("%f,%04d-%02d-%02d%02s,%s年,%s日\n", lichun, year, m, d, sexa.FmtTime(unit.TimeFromDay(t)), yearGZ, dayGZ)
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
