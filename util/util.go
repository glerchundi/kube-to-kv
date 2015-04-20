package util

import (
    "fmt"
    "net"
    "reflect"
    "strconv"
    "strings"

    "github.com/glerchundi/kube2kv/log"
)

// Dump object
func Dump(v interface{}) {
    if v == nil {
        return
    }
    s := reflect.ValueOf(v).Elem()
    typeOfT := s.Type()

    log.Debug(typeOfT.String())
    for i := 0; i < s.NumField(); i++ {
        f := s.Field(i)
        log.Debug(fmt.Sprintf("%d: %s %s = '%v'", i, typeOfT.Field(i).Name, f.Type(), f.Interface()))
    }
}

// Get Backend Node From SRV
func GetBackendNodesFromSRV(backend, srv string) ([]string, error) {
    nodes := make([]string, 0)

    // Split, if domain was found
    srvScheme := "http"
    srvTarget := srv
    parts := strings.SplitN(srv, ",", 2)
    if len(parts) == 2 {
        srvScheme = parts[0]
        srvTarget = parts[1]
    }

    // Ignore the CNAME as we don't need it.
    _, addrs, err := net.LookupSRV(backend, "tcp", srvTarget)
    if err != nil {
        return nodes, err
    }

    // Generate node from SRV output
    for _, srv := range addrs {
        host := strings.TrimRight(srv.Target, ".")
        port := strconv.FormatUint(uint64(srv.Port), 10)
        nodes = append(nodes, fmt.Sprintf("%s://%s", srvScheme, net.JoinHostPort(host, port)))
    }

    return nodes, nil
}

// Default to values provided by functor
func GetBackendNodesFromSRVOrElse(backend, srv string, orElse func() []string) []string {
    if nodes, err := GetBackendNodesFromSRV(backend, srv); err == nil && len(nodes) > 0 {
        return nodes
    }
    return orElse()
}
