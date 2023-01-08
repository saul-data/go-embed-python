package pip

import (
	"fmt"
	"github.com/saul-data/go-embed-python/embed_util"
	"github.com/saul-data/go-embed-python/pip/internal/data"
)

func NewPipLib(name string) (*embed_util.EmbeddedFiles, error) {
	return embed_util.NewEmbeddedFiles(data.Data, fmt.Sprintf("pip-%s", name))
}
