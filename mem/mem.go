package mem

import (
	"os/exec"
	"strconv"
	"strings"
)

//GetFree returns the percentage of memory that is free
func GetFree() (free float64, err error) {
	cmd := "free | grep Mem | awk '{print $4/$2 * 100.0}'"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return -1, err
	}
	strOut := strings.TrimSuffix(string(out), "\n")
	free, err = strconv.ParseFloat(strOut, 64)
	return free, err
}

//GetUsed returns the percentage of memory that is in use
func GetUsed() (used float64, err error) {
	cmd := "free | grep Mem | awk '{print $3/$2 * 100.0}'"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return -1, err
	}
	strOut := strings.TrimSuffix(string(out), "\n")
	used, err = strconv.ParseFloat(strOut, 64)
	return used, err
}

//GetSwpFree returns the percentage of swap file memory that is free
func GetSwpFree() (res string, err error) {
	cmd := exec.Command("free | grep Swap | awk '{print $4/$2 * 100.0}'")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), err
}

//GetSwpUsed returns the percentage of swap file memory that is in use
func GetSwpUsed() (res string, err error) {
	cmd := exec.Command("free | grep Swap | awk '{print $3/$2 * 100.0}'")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), err
}
