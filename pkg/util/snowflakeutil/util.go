package snowflakeutil

import (
	"net"
	"os"
	"strconv"

	"github.com/sony/sonyflake"
)

var _sf *sonyflake.Sonyflake

func init() {
	saltStr, ok := os.LookupEnv("SNOW_FLAKE_SALT")
	var salt uint16
	if ok {
		i, err := strconv.ParseUint(saltStr, 10, 16)
		if err != nil {
			panic(err)
		}
		salt = uint16(i)
	}
	ips, err := getLocalIPs()
	if err != nil {
		panic(err)
	}
	_sf = sonyflake.NewSonyflake(sonyflake.Settings{
		MachineID: func() (u uint16, e error) {
			return sumIPs(ips) + salt, nil
		},
	})
	if _sf == nil {
		panic("sonyflake init failed")
	}
}

func getLocalIPs() ([]net.IP, error) {
	var ips []net.IP
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ips, err
	}
	for _, a := range addrs {
		if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			ips = append(ips, ipNet.IP)
		}
	}
	return ips, nil
}

func sumIPs(ips []net.IP) uint16 {
	total := 0
	for _, ip := range ips {
		for i := range ip {
			total += int(ip[i])
		}
	}
	return uint16(total)
}

// 生成新的 ID
func NewId() uint64 {
	id, err := _sf.NextID()
	if err != nil {
		panic("get sony flake uid failed:" + err.Error())
	}
	return id
}

func NewIdString() string {
	return strconv.FormatUint(NewId(), 10)
}
