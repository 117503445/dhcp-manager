package store

import (
	"sync"

	"github.com/117503445/landns/pkg/rpcgen"
	// "github.com/rs/zerolog/log"
)

type LeasesStore struct {
	sync.RWMutex
	leasesByTag   map[string][]*rpcgen.Lease
	HostnameIpMap map[string]string
}

func NewLeasesStore() *LeasesStore {
	return &LeasesStore{
		leasesByTag:   make(map[string][]*rpcgen.Lease),
		HostnameIpMap: make(map[string]string),
	}
}

func (s *LeasesStore) GetIpByHostname(hostname string) string {
	s.RLock()
	defer s.RUnlock()
	return s.HostnameIpMap[hostname]
}

// 把 leasesByTag 聚合成 HostnameIpMap
// 并且把重复的 hostname-ip 对应的 tag 也记录下来
func MergeLeases(leasesByTag map[string][]*rpcgen.Lease) (hostnameIpMap map[string]string, repeatLeaseTagMap map[*rpcgen.Lease]string) {
	hostnameIpMap = make(map[string]string)

	repeatLeaseTagMap = make(map[*rpcgen.Lease]string)
	repeatHostnameMap := make(map[string]interface{})
	hostNameLeaseMap := make(map[string]*rpcgen.Lease)

	for tag, leases := range leasesByTag {
		for _, lease := range leases {
			if _, ok := repeatHostnameMap[lease.Hostname]; ok {
				repeatLeaseTagMap[lease] = tag
			} else {
				if _, ok := hostNameLeaseMap[lease.Hostname]; ok {
					if hostNameLeaseMap[lease.Hostname].Ip != lease.Ip {
						repeatHostnameMap[lease.Hostname] = nil
						repeatLeaseTagMap[lease] = tag
						repeatLeaseTagMap[hostNameLeaseMap[lease.Hostname]] = tag
						delete(hostNameLeaseMap, lease.Hostname)
					}
				} else {
					hostNameLeaseMap[lease.Hostname] = lease
				}
			}
		}
	}

	for hostname, lease := range hostNameLeaseMap {
		hostnameIpMap[hostname] = lease.Ip
	}

	return
}

func (s *LeasesStore) SetLeases(tag string, leases []*rpcgen.Lease) {
	s.Lock()
	defer s.Unlock()
	s.leasesByTag[tag] = leases

	s.HostnameIpMap, _ = MergeLeases(s.leasesByTag)
}

// // hostname -> tags
// repeatHostnameTagMap := make(map[string][]string)
// // hostname -> leases
// repeatHostnameLeaseMap := make(map[string][]*rpcgen.Lease)

// // reset HostnameIpMap by leasesByTag
// for tag, leases := range s.leasesByTag {
// 	for _, lease := range leases {
// 		if _, ok := repeatHostnameLeaseMap[lease.Hostname]; ok {
// 			found := false
// 			for _, lease := range repeatHostnameLeaseMap[lease.Hostname] {
// 				if lease.Ip == lease.Ip {
// 					found = true
// 					break
// 				}
// 			}
// 			if !found {
// 				repeatHostnameLeaseMap[lease.Hostname] = append(repeatHostnameLeaseMap[lease.Hostname], lease)
// 				repeatHostnameTagMap[lease.Hostname] = append(repeatHostnameTagMap[lease.Hostname], tag)
// 			}
// 		}else{

// 		}

// 		// if _, ok := s.HostnameIpMap[lease.Hostname]; ok {
// 		// 	if s.HostnameIpMap[lease.Hostname] != lease.Ip {
// 		// 		log.Warn().Str("tag", tag).Str("hostname", lease.Hostname).Str("ip", lease.Ip).Msg("hostname has multiple ips")
// 		// 	}
// 		// 	continue
// 		// }
// 		// s.HostnameIpMap[lease.Hostname] = lease.Ip
// 	}
// 	}
// }
