package dancong

import "strings"

// Context
type Context struct {
	config map[string]interface{}
}

func (ctx *Context) GetConfig(key string) (value interface{}, ok bool) {
	keys := strings.Split(key, ".")
	if value, ok = ctx.config[keys[0]]; !ok {
		return
	}
	for _, key := range keys[1:] {
		value, ok = value.(map[interface{}]interface{})[key]
		if !ok {
			return
		}
	}
	return
}
