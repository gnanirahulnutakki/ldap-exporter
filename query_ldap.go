package main

import (
    "fmt"
    "log"
    "os"

    "github.com/go-ldap/ldap/v3"
)

func queryLDAP(server string, bindDN string, bindPassword string, baseDN string, filter string) (*ldap.SearchResult, error) {
    l, err := ldap.Dial("tcp", server)
    if err != nil {
        return nil, err
    }
    defer l.Close()

    // Bind with the given bindDN and bindPassword
    err = l.Bind(bindDN, bindPassword)
    if err != nil {
        return nil, err
    }

    // Search for entries matching the given filter in the given baseDN
    searchRequest := ldap.NewSearchRequest(
        baseDN,
        ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
        filter,
        []string{"*"},
        nil,
    )
    searchResult, err := l.Search(searchRequest)
    if err != nil {
        return nil, err
    }

    return searchResult, nil
}

func main() {
    server := os.Getenv("LDAP_URL")
    bindDN := os.Getenv("LDAP_BIND_DN")
    bindPassword := os.Getenv("LDAP_BIND_PASSWORD")
    baseDN := os.Getenv("LDAP_BASE_DN")
    filter := os.Getenv("LDAP_FILTER")

    searchResult, err := queryLDAP(server, bindDN, bindPassword, baseDN, filter)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("LDAP search returned %d entries\n", len(searchResult.Entries))
    for _, entry := range searchResult.Entries {
        fmt.Println(entry)
    }
}
