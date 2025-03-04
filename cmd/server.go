package cmd

import (
	"github.com/afa7789/tallypound/internal/cache"
	"github.com/afa7789/tallypound/internal/compound"
	"github.com/afa7789/tallypound/internal/domain"
	"github.com/afa7789/tallypound/internal/server"
)

// execute the server package to start the server running
func ServerExecute(f domain.Flags) error {
	c := cache.NewCache()
	cc := compound.NewCacheCompound(c)
	s := server.NewServer(cc)

	s.Start(*f.Port)
	return nil
}
