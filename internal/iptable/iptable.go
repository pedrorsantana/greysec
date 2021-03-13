package iptable

import (
	"github.com/coreos/go-iptables/iptables"
)

const (
	tableName = "filter"
)

func BlockIP(ip string) (bool, error) {
	ipt, err := iptables.New()
	if err != nil {
		return false, err
	}

	if exists, _ := ipt.Exists(tableName, "OUTPUT", "-s", ip, "-j", "DROP"); exists == true {
		return true, nil
	}
	if exists, _ := ipt.Exists(tableName, "INPUT", "-s", ip, "-j", "DROP"); exists == true {
		return true, nil
	}

	err = ipt.Insert(tableName, "OUTPUT", 1, "-s", ip, "-j", "DROP")
	if err != nil {
		return false, err
	}

	err = ipt.Insert(tableName, "INPUT", 1, "-s", ip, "-j", "DROP")
	if err != nil {
		return false, err
	}

	return true, nil
}
