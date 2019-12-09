package v2client

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
)

func (p *Presenter) asJSON() {
	j, err := json.Marshal(p.ServiceReport)
	if err != nil {
		log.Fatalln(err)
	}

	// TODO maybe there's a cleaner way of outputting this than just a println
	fmt.Println(string(j))
}
