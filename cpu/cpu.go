package cpu

import (
	"os/exec"
	"strconv"
	"strings"
)

//GetSummary returns a summary of the cpu stats
// TODO

//GetTemp returns the current CPU temperature in C
func GetTemp() (temp float64, err error) {
	cmd := exec.Command("bash", "-c", "/usr/bin/vcgencmd measure_temp | egrep -o '[0-9]*\\.[0-9]*'")
	out, err := cmd.Output()
	if err != nil {
		return -1, err
	}
	res := strings.TrimSuffix(string(out), "\n")
	temp, _ = strconv.ParseFloat(res, 64)
	return temp, nil
}

//GetFreq returns the current CPU clock speed in Hz
func GetFreq() (freq float64, err error) {
	cmd := exec.Command("bash", "-c", "sudo cat /sys/devices/system/cpu/cpu0/cpufreq/cpuinfo_cur_freq")
	out, err := cmd.Output()
	if err != nil {
		return -1, err
	}
	res := strings.TrimSuffix(string(out), "\n")
	freq, _ = strconv.ParseFloat(res, 64)
	return freq, nil
}
