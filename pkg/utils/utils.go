package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"net"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strconv"

	"github.com/twinj/uuid"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorGray   = "\033[37m"
	ColorWhite  = "\033[97m"
)

func Colorize(colorCode, text string) string {
	if colorCode == "" || text == "" {
		return ""
	}
	return colorCode + text + ColorReset
}

func UUIDv4() string {
	return uuid.NewV4().String()
}

func MkdirAllIfNotExists(pathname string, perm os.FileMode) error {
	dir := path.Dir(pathname)
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(dir, perm); err != nil {
				return err
			}
		}
	}
	return nil
}

func FilExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func WriteToFile(filePath string, content string, filemode os.FileMode) error {
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	if err := os.WriteFile(filePath, []byte(content), filemode); err != nil {
		return err
	}
	return nil
}

func GenHexStr(length int) string {
	bytes := make([]byte, length/2)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}

func FileMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	hashInBytes := hash.Sum(nil)
	md5Hash := hex.EncodeToString(hashInBytes)

	return md5Hash, nil
}
func GetIPv4NetworkIPs() ([]string, error) {
	ips, err := GetNetworkIPs()
	if err != nil {
		return nil, err
	}

	ip4s := make([]string, 0)
	for _, ip := range ips {
		if net.ParseIP(ip).To4() != nil {
			ip4s = append(ip4s, ip)
		}
	}

	return ip4s, nil
}

func GetNetworkIPs() ([]string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	ips := make([]string, 0)
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				ips = append(ips, v.IP.String())
			case *net.IPAddr:
				ips = append(ips, v.IP.String())
			}
		}
	}

	return ips, nil
}

func PrintStructFieldsAndValues(s interface{}, indent string) {
	v := reflect.ValueOf(s)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		fmt.Printf("%s%v is not a struct\n", indent, v.Type())
		return
	}

	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fmt.Printf("%s - %s: ", indent, typeOfS.Field(i).Name)

		if field.Kind() == reflect.Struct {
			fmt.Println()
			PrintStructFieldsAndValues(field.Interface(), indent+"  ")
		} else if field.Kind() == reflect.Ptr && field.Elem().Kind() == reflect.Struct {
			fmt.Println()
			PrintStructFieldsAndValues(field.Interface(), indent+"  ")
		} else {
			if field.CanInterface() {
				fmt.Println(Colorize(ColorGreen, fmt.Sprint(field.Interface())))
			} else {
				fmt.Println()
			}
		}
	}
}

func ParseBool(val string, defVal bool) bool {
	_val, err := strconv.ParseBool(val)
	if err != nil {
		return defVal
	}
	return _val
}

func ParseInt(val string, defVal int) int {
	_val, err := strconv.Atoi(val)
	if err != nil {
		return defVal
	}
	return _val
}

func ParseInt64(val string, defVal int64) int64 {
	_val, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return defVal
	}
	return _val
}

func ParseFloat64(val string, defVal float64) float64 {
	_val, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return defVal
	}
	return _val
}

func ParseUint(val string, defVal uint) uint {
	_val, err := strconv.ParseUint(val, 10, 0)
	if err != nil {
		return defVal
	}
	return uint(_val)
}
