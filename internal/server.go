package internal

import (
	"errors"
	"fmt"
	common "github.com/chenshijian73-qq/Doraemon/pkg"
	"github.com/olekukonko/tablewriter"
	"os"
	"sort"
)

// ListServers merge basic context servers and current context servers
func ListServers(serverSort bool) Servers {
	var servers Servers
	if serverSort {
		sort.Sort(basicConfig.Servers)
	}
	servers = append(servers, basicConfig.Servers...)
	if currentConfig.configPath != basicConfig.configPath {
		if serverSort {
			sort.Sort(currentConfig.Servers)
		}
		servers = append(servers, currentConfig.Servers...)
	}
	return servers
}

// PrintServers print server list
func PrintServers(serverSort bool) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "User", "Tags", "Address"})
	ss := ListServers(serverSort)
	for _, s := range ss {
		table.Append([]string{s.Name, s.User, fmt.Sprint(s.Tags), fmt.Sprintf("%s:%d", s.Address, s.Port)})
	}
	table.Render()
}

// PrintServerDetail print single server detail
func PrintServerDetail(serverName string) {
	s, err := findServerByName(serverName)
	common.CheckAndExit(err)

	table := tablewriter.NewWriter(os.Stdout)
	table.Append([]string{"NAME", s.Name})
	table.Append([]string{"USER", s.User})
	table.Append([]string{"ADDR", fmt.Sprintf("%s:%d", s.Address, s.Port)})
	table.Append([]string{"PROXY", s.Proxy})
	table.Append([]string{"CONFIG", s.ConfigPath})
	table.Render()
}

// findServerByName find server from config by server name
func findServerByName(name string) (*Server, error) {
	for _, s := range ListServers(false) {
		if s.Name == name {
			return s, nil
		}
	}
	return nil, errors.New("server not found")
}

// findServersByTag find servers from config by server tag
func findServersByTag(tag string) (Servers, error) {
	var servers Servers
	for _, s := range ListServers(false) {
		tmpServer := s
		for _, t := range tmpServer.Tags {
			if tag == t {
				servers = append(servers, tmpServer)
			}
		}
	}
	if len(servers) == 0 {
		return nil, errors.New("server not found")
	}
	return servers, nil
}
