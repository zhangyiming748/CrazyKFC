package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func main() {
	a := day()      //今天
	b := subDay()   //今天是这个月的第几天
	c := subYear()  //今天是今年的第几天
	d := thisYear() //今年
	e := week()     //今天是今年的第几周
	ret := communicate(a, b, c, d, e)
	//
	if runtime.GOOS == "android" {
		writeAll("/data/data/com.termux/files/home/storage/documents/report.txt", ret)
	} else {
		writeAll("report.txt", ret)
	}
	fmt.Println(ret)
}

func day() string {
	ms := time.Now().Format("1月2号")
	return ms
}
func subDay() string {
	day := time.Now().Format("2") //今天
	return day
}
func thisYear() string {
	t := time.Now().Year()
	return strconv.Itoa(t)
}
func subYear() string {
	day := strconv.Itoa(time.Now().YearDay())
	return day
}
func week() string {
	_, w := getWeek()
	s := strconv.Itoa(w)
	return s
}

func getWeek() (y, w int) {

	t := time.Now().Format("20060102")
	timeLayout := "20060102"
	loc, _ := time.LoadLocation("Local")
	tmp, _ := time.ParseInLocation(timeLayout, t, loc)
	return tmp.ISOWeek()
}

func communicate(a, b, c, d, e string) string {
	day := isThursday()
	ie, _ := strconv.Atoi(e)
	iepp := ie + 1
	epp := strconv.Itoa(iepp)
	switch day {
	case "before":
		return strings.Join([]string{"今天是", a, ",是我这个月零收入的第", b, "天,也是我今年零收入的第", c, "天,依旧没人带我赚大钱,外面又潮又冷,我很饿,身上的钱只够买一个馒头了,吃了这顿饭以后该怎么办?我很迷茫,以前电子厂一个月800非常的满足,每月还可以下两三次馆子,除去花销,有时候还能结余十多块,自从认识你们,就开始怀疑人生了,你们一谈就是一个月赚几万,看到你们晒的图,经常晚上睡不着,有时候甚至怀疑我们不是在同一个世界,大家在一个群里就是缘分,莫让错过成为你的遗憾,这星期四是KFC", d, "年第", e, "次疯狂星期四,V我50,拒绝遗憾\n"}, "")
	case "bingo":
		return strings.Join([]string{"今天是", a, ",是我这个月零收入的第", b, "天,也是我今年零收入的第", c, "天,依旧没人带我赚大钱,外面又潮又冷,我很饿,身上的钱只够买一个馒头了,吃了这顿饭以后该怎么办?我很迷茫,以前电子厂一个月800非常的满足,每月还可以下两三次馆子,除去花销,有时候还能结余十多块,自从认识你们,就开始怀疑人生了,你们一谈就是一个月赚几万,看到你们晒的图,经常晚上睡不着,有时候甚至怀疑我们不是在同一个世界,大家在一个群里就是缘分,莫让错过成为你的遗憾,今天就是KFC", d, "年第", e, "次疯狂星期四,择日不如撞日,V我50,拒绝遗憾\n"}, "")
	case "after":
		return strings.Join([]string{"今天是", a, ",是我这个月零收入的第", b, "天,也是我今年零收入的第", c, "天,依旧没人带我赚大钱,外面又潮又冷,我很饿,身上的钱只够买一个馒头了,吃了这顿饭以后该怎么办?我很迷茫,以前电子厂一个月800非常的满足,每月还可以下两三次馆子,除去花销,有时候还能结余十多块,自从认识你们,就开始怀疑人生了,你们一谈就是一个月赚几万,看到你们晒的图,经常晚上睡不着,有时候甚至怀疑我们不是在同一个世界,大家在一个群里就是缘分,莫让错过成为你的遗憾,下星期四是KFC", d, "年第", epp, "次疯狂星期四,V我50,拒绝遗憾\n"}, "")
	default:
		return ""
	}
}
func isThursday() string {
	f := time.Now().Weekday().String()
	if f == "Thursday" {
		return "bingo"
	} else if f == "Sunday" || f == "Monday" || f == "Tuesday" || f == "Wednesday" {
		return "before"
	} else {
		return "after"
	}
}
func writeAll(fname, content string) {
	f, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE, 0776)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	n, err := f.WriteString(content)
	if err != nil {
		log.Println("写文件出错")
	} else {
		log.Printf("写入%d个字节", n)
	}
}
