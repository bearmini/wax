package wax

const (
	DefaultMaxMemorySizeInPage = 256 // 256 pages = 1 MB (4kB / page)
)

type RuntimeConfig struct {
	maxMemorySizeInPage uint32
}

func NewRuntimeConfig() *RuntimeConfig {
	return &RuntimeConfig{
		maxMemorySizeInPage: DefaultMaxMemorySizeInPage,
	}
}

func (c *RuntimeConfig) MaxMemorySizeInPage(n uint32) *RuntimeConfig {
	c.maxMemorySizeInPage = n
	return c
}
