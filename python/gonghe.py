#!/usr/bin/env python3

from math import floor

import ephem

gan = '甲乙丙丁戊己庚辛壬癸'
zhi = '子丑寅卯辰巳午未申酉戌亥'

start_date = ephem.Date((2018, 2, 5))


def ganzhi(n):
    g = n % 10
    z = n % 12
    return gan[g] + zhi[z]


def ganzhi_of_day(jd):
    return (int(floor(jd+.5)) - 11) % 60


# year = 0 為不合法
def ganzhi_of_year(year):
    year = year if year > 0 else year + 1
    return (year-4) % 60


# 往回尋找冬至朔旦甲子(dzsdjz)
# tz 時區(小時)
# gz 干支, -1 表示不指定
# delta 相差, 預設為 8 小時
def previous_dzsdjz(start_date, tz=0, gz=-1, delta=0.333333):
    d = start_date
    while True:
        sd = ephem.previous_winter_solstice(d)  # 冬至日 (sun-cycle date)
        md = ephem.previous_new_moon(sd + 2.0)  # 朔旦日 (moon-cycle date)
        sd_loc = ephem.Date(sd + tz/24.0)
        md_loc = ephem.Date(md + tz/24.0)
        sd_ymd = sd_loc.tuple()[:3]
        md_ymd = md_loc.tuple()[:3]
        d = sd
        if sd_ymd != md_ymd:
            continue
        if abs(sd - md) > delta:
            continue
        daygz = ganzhi_of_day(ephem.julian_date(d))  # 日干支
        if gz != -1 and daygz != gz:
            continue
        yeargz = ganzhi_of_year(sd_ymd[0])           # 年干支
        print(sd_ymd, ganzhi(yeargz), '年,', ganzhi(daygz), '日')
        print(sd, ephem.julian_date(sd))
        print(md, ephem.julian_date(md))
        break
    return d


def lichun_of_year(year):
    d = ephem.Date((year, 1, 1))
    eq = ephem.next_spring_equinox(d)
    sol = ephem.previous_winter_solstice(d)
    lichun = ephem.Date((eq + sol) / 2)
    lichun_ymd = lichun.tuple()[:3]
    daygz = ganzhi_of_day(ephem.julian_date(lichun))
    yeargz = ganzhi_of_year(lichun_ymd[0])
    print(lichun_ymd, ganzhi(yeargz), '年,', ganzhi(daygz), '日')
    print(lichun, ephem.julian_date(lichun))

def test_of_zero_year():
    d1 = ephem.Date((1, 1, 1))
    d2 = ephem.Date(d1 - 1.0)
    print(d2.tuple()[:3])

# 往回尋找冬至朔旦甲子
# d = start_date
# for i in range(1):
#     d = previous_dzsdjz(d, gz=0)

for i in range(-840, -844, -1):
    lichun_of_year(i)

