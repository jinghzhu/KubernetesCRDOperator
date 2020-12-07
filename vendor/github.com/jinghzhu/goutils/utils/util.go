package utils

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// IntMax returns the max value of int32, which is 2147483647.
func IntMax() int {
    i := 0
    for j := 0; j < 31; j++ {
        i = (i << 1) | 1
    }
    
    return i
}

// IntMin returns the min value of int32, which is -2147483648.
func IntMin() int {
    return -1 * (1 << 31)
}

// IsIPv6 checks whether the input is a valid IPv6 address.
func IsIPv6(ip string) bool {
    ips := strings.Split(ip, ":")
    l := len(ips)
    if l != 8 {
        return false
    }
    validStr := "0123456789ABCDEFabcdef"
    for _, v := range ips {
        if len(v) < 1 || len(v) > 4 {
            return false
        }
        for _, x := range v {
            if !strings.Contains(validStr, string(x)) {
                return false
            }
        }
    }
    
    return true
}

// IsIPv4 checks whether the stirng is a valid IPv4 address.
func IsIPv4(ip string) bool {
	ips := strings.Split(ip, ".")
	l := len(ips)
	if l != 4 {
		return false
	}
	for _, v := range ips {
		lv, tmp := len(v), 0
		if lv > 3 || lv < 1 {
			return false
		}
		if v[0] == '0' && lv > 1 {
			return false
		}
		for _, x := range v {
			if x < '0' || x > '9' {
				return false
			}
			tmp = 10*tmp + int(x-'0')
		}
		if tmp > 255 {
			return false
		}
	}

	return true
}

// Struct2String accepts any interface{} and return to JSON based string.
func Struct2String(v interface{}) string {
	result, err := json.Marshal(v)
	if err != nil {
		errMsg := "Fail to translate to json"
		fmt.Println(errMsg)
		return fmt.Sprintf("%v", v)
	}
	return string(result)
}

// PanicHandler catches a panic and logs an error. Suppose to be called via defer.
func PanicHandler() (caller string, fileName string, lineNum int, stackTrace string, rec interface{}) {
	buf := make([]byte, stackBuffer)
	runtime.Stack(buf, false)
	name, file, line := GetCallerInfo(2)
	if r := recover(); r != nil {
		caller, fileName, stackTrace = name, file, string(buf)
		lineNum = line
		rec = r
		fmt.Printf("%s %s ln%d: PANIC Defered : %v\n", name, file, line, r)
		fmt.Printf("%s %s ln%d: Stack Trace : %s", name, file, line, string(buf))
	}

	return caller, fileName, lineNum, stackTrace, rec
}

// GetCallerInfo returns the name of method caller and file name. It also returns the line number.
func GetCallerInfo(level int) (caller, fileName string, lineNum int) {
	if level < 1 || level > maxCallerLevel {
		level = defaultCallerLevel
	}

	pc, file, line, ok := runtime.Caller(level)
	fileDefault := ""
	lineDefault := -1
	nameDefault := ""
	if ok {
		fileDefault = file
		lineDefault = line
	}
	details := runtime.FuncForPC(pc)
	if details != nil {
		nameDefault = details.Name()
	}

	return nameDefault, fileDefault, lineDefault
}

// GetMountPoints returns all moutpoins in a string array
func GetMountPoints(server string) ([]string, error) {
	b, err := exec.Command("showmount", "-e", server).Output()
	if err != nil {
		fmt.Println("error in showmount: " + err.Error())
		return nil, err
	}
	s := strings.TrimSpace(string(b))
	// The fist line of showmount -e <server> is Exports list on <server>
	firstLine := strings.Index(s, "\n")
	sArr := strings.Split(s[firstLine+1:], "\n")
	for i := 0; i < len(sArr); i++ {
		index := strings.Index(sArr[i], " ")
		temp := sArr[i]
		sArr[i] = temp[:index]
	}
	return sArr, nil
}

// Locate returns the line number and file name in the current goroutine statck trace. The argument skip is the number of stack frames to ascend, with 0 identifying the caller of Caller.
func Locate(skip int) (filename string, line int) {
	if skip < 0 {
		skip = 2
	}
	_, path, line, ok := runtime.Caller(skip)
	file := ""
	if ok {
		_, file = filepath.Split(path)
	} else {
		fmt.Println("Fail to get method caller")
		line = -1
	}
	return file, line
}

// Retry will retry the given condition function with specific time interval and retry round.
// It will return true if the condition is met. If it is timeout, it will return false. Otherwise,
// it will return the error encountered in the retry round.
func Retry(interval time.Duration, round int, retry func() (bool, error)) (bool, error) {
	if round < 1 {
		round = 1
	}
	if interval > time.Hour {
		interval = time.Hour
	}
	var err error
	done := false
	for i := 0; i < round; i++ {
		done, err = retry()
		if done {
			break
		}
		time.Sleep(interval)
	}

	if done {
		return true, nil
	}
	if err != nil {
		return false, err
	}

	return false, nil
}
