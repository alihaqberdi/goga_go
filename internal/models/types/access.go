package types

const (
	Username Access = 1 + iota
	Origin
	//InitStr
)

type Access int16

func (p Access) String() string {
	return access2str[p]
}

var (
	access2str = map[Access]string{}
	str2access = map[string]Access{}
)

func init() {
	access2str[Username] = "username"
	access2str[Origin] = "origin"
	//access2str[InitStr] = "init_str"

	for p, s := range access2str {
		str2access[s] = p
	}

}
