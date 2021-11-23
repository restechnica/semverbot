package versions

import "fmt"

func AddPrefix(version string, prefix string) string {
	return fmt.Sprintf("%s%s", prefix, version)
}
