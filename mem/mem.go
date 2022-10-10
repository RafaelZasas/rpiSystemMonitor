package mem

import (
  "os/exec"
)

//GetFree returns the percentage of memory that is free
func GetFree() (res string, err error) {
  cmd := exec.Command("free | grep Mem | awk '{print $4/$2 * 100.0}'")
  out, err := cmd.Output()
  if err != nil {
    return "", err
  }
  return string(out), nil
}


//GetUsed returns the percentage of memory that is in use
func GetUsed()(res string, err error) {
  cmd := exec.Command("free | grep Mem | awk '{print $3/$2 * 100.0}'")
  out, err := cmd.Output()
  if err != nil {
    return "", err
  }
  return string(out), err
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
