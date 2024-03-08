package main

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"

	"changemedaddy/internal/page"
	"changemedaddy/internal/page/attr/title"
	"changemedaddy/internal/page/el/btn"
)

func main() {
	u, _ := uuid.NewRandom()

	p := page.New(
		u,
		page.WithElement(
			btn.New(uuid.New(), 0, "More of My Videos Here <3", "onlyfans.com/yourmom"),
		),
		page.WithAttribute(title.New(uuid.New(), "Your Mom's Website")),
	)

	j, _ := json.Marshal(p)

	fmt.Println(string(j))
}
