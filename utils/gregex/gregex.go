/**
 *
 * @author 半城风雨
 * @since 2021/11/19
 * @File : gregex
 */
package gregex

import (
	"regexp"
	"sync"
)

var (
	regexMu = sync.RWMutex{}
	// Cache for regex object.
	// Note that:
	// 1. It uses sync.RWMutex ensuring the concurrent safety.
	// 2. There's no expiring logic for this map.
	regexMap = make(map[string]*regexp.Regexp)
)

func MatchString(pattern string, src string) ([]string, error) {
	if r, err := getRegexp(pattern); err == nil {
		return r.FindStringSubmatch(src), nil
	} else {
		return nil, err
	}
}

func getRegexp(pattern string) (regex *regexp.Regexp, err error) {
	// Retrieve the regular expression object using reading lock.
	regexMu.RLock()
	regex = regexMap[pattern]
	regexMu.RUnlock()
	if regex != nil {
		return
	}
	// If it does not exist in the cache,
	// it compiles the pattern and creates one.
	regex, err = regexp.Compile(pattern)
	if err != nil {
		return
	}
	// Cache the result object using writing lock.
	regexMu.Lock()
	regexMap[pattern] = regex
	regexMu.Unlock()
	return
}
