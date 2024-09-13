package versions

import "fmt"

func AddSuffix(version string, suffix string) string {
	return fmt.Sprintf("%s%s", version, suffix)
}
