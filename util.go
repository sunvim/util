package util

import (
	"bytes"
	"fmt"
	log "github.com/sunvim/slog"
	"math/rand"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

func GenerateStringSlice(element string, size int) []string {
	slice := make([]string, size)
	for size = size - 1; size >= 0; size-- {
		slice[size] = element
	}
	return slice
}

func MonthDays(year, month int) int {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
		if IsLeapYear(year) {
			return 29
		}
		return 28
	default:
		panic(fmt.Sprintf("非法的月份: %d", month))
	}
}

func IsLeapYear(year int) bool {
	if year%100 == 0 {
		return year%400 == 0
	}

	return year%4 == 0
}

func IsLegalMD5Code(code string) bool {
	const CODE_LENGTH = 32

	if len(code) != CODE_LENGTH {
		return false
	}

	legal := true
	for i := 0; i < CODE_LENGTH && legal; i++ {
		c := code[i]

		legal = (c >= '0' && c <= '9') ||
			(c >= 'a' && c <= 'f') ||
			(c >= 'A' && c <= 'Z')
	}

	return legal
}

func JoinStringSlices(slice1, slice2 []string, sep string) string {
	buffer := bytes.NewBuffer(make([]byte, 0, 256))
	buffer.WriteString(strconv.Itoa(len(slice1)))
	for i := 0; i < len(slice1); i++ {
		buffer.WriteString(sep)
		buffer.WriteString(slice1[i])
	}

	for i := 0; i < len(slice2); i++ {
		buffer.WriteString(sep)
		buffer.WriteString(slice2[i])
	}

	return buffer.String()
}

func ParseJoinedStringSlice(s, sep string) (all, slice1, slice2 []string) {
	all = strings.Split(s, sep)
	if len(all) == 0 {
		return nil, nil, nil
	}

	size, err := strconv.Atoi(all[0])
	if err != nil {
		return nil, nil, nil
	}
	all = all[1:]
	if size > len(all) {
		return nil, nil, nil
	}

	return all, all[:size], all[size:]
}

func RandString(slice []string) string {
	len := len(slice)
	if len == 0 {
		return ""
	}
	return slice[rand.Intn(len)]
}

func RemoveDuplicate(slice []string) []string {
	sort.Strings(slice)

	var prevIndex = 0
	size := len(slice)
	for i := 1; i < size; i++ {
		if s := slice[i]; s != slice[prevIndex] {
			prevIndex++
			slice[prevIndex] = s
		}
	}

	return slice[:prevIndex+1]
}

func Recover(desc string) {
	if e := recover(); e != nil {
		buf := make([]byte, 4096)
		i := runtime.Stack(buf, false)
		buf = buf[:i]
		log.Error("desc:", desc, "error out of think=>", e)
		log.Error("Stack:", string(buf))
	}
}

func FindString(eles []string, ele string, skip int) int {
	for i := 0; i < len(eles); i = i + 1 + skip {
		if eles[i] == ele {
			return i
		}
	}

	return -1
}
