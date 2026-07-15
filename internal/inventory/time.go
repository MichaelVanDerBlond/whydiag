package inventory

import "time"

func Timezone() string {
	return time.Now().Location().String()
}
