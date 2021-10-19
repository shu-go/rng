package rng

import (
	"fmt"
	"strconv"
	"strings"
)

type IPv4 [4]int

func (s IPv4) Next() Sequential {
	ip := s
	if ip[0] == 255 && ip[1] == 255 && ip[2] == 255 && ip[3] == 255 {
		return s
	}

	for i := range ip {
		if ip[3-i] == 255 {
			ip[3-i] = 1
		} else {
			ip[3-i]++
			break
		}
	}

	return IPv4(ip)
}

func (s IPv4) Prev() Sequential {
	ip := s
	if ip[0] == 0 && ip[1] == 0 && ip[2] == 0 && ip[3] == 0 {
		return s
	}

	for i := range ip {
		if ip[3-i] <= 1 {
			ip[3-i] = 255
		} else {
			ip[3-i]--
			break
		}
	}

	return IPv4(ip)
}

func (s IPv4) Less(b Sequential) bool {
	if bb, ok := b.(IPv4); ok {
		for i := range s {
			if s[i] < bb[i] {
				return true
			} else if s[i] > bb[i] {
				return false
			}
		}
	}
	return false
}

func (s IPv4) Equal(b Sequential) bool {
	if bb, ok := b.(IPv4); ok {
		for i := range s {
			if s[i] != bb[i] {
				return false
			}
		}
		return true
	}
	return false
}

func NewIPv4(s string) IPv4 {
	ss := strings.Split(s, ".")
	if len(ss) < 4 {
		ss = append(ss, "0", "0", "0", "0")
	}

	nn := IPv4{0, 0, 0, 0}
	for i, s := range ss {
		ni, err := strconv.Atoi(s)
		if err != nil {
			panic(fmt.Sprintf("%s", ss))
		}
		nn[i] = ni
	}

	return nn
}
