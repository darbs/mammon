
import (
	"fmt"
	"time"
	log "github.com/sirupsen/logrus"

	"astuart.co/go-robinhood"
	"github.com/darbs/mammon/internal/database"
)

var logger *log.Entry

type Collection struct {
	items []interface{}
}