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
        daygz = ganzhi_of_day(ephem.julian_date(d)) # 日干支
        if gz != -1 and daygz != gz:
            continue
        print(sd_ymd, ganzhi(daygz))
        print(sd, ephem.julian_date(sd))
        print(md, ephem.julian_date(md))
        break
    return d


d = start_date
for i in range(1):
    d = previous_dzsdjz(d, gz=0)
