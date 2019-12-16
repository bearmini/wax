package wax

const (
	DefaultMaxMemorySizeInPage = 256 // 256 pages = 1 MB (4kB / page)
)

type RuntimeConfig struct {
	maxMemorySizeInPage uint32
	importFunc          map[Name]map[Name]FuncInst
	importTable         map[Name]map[Name]TableInst
	importMemory        map[Name]map[Name]MemInst
	importGlobal        map[Name]map[Name]GlobalInst
}

func NewRuntimeConfig() *RuntimeConfig {
	return &RuntimeConfig{
		maxMemorySizeInPage: DefaultMaxMemorySizeInPage,
		importFunc:          make(map[Name]map[Name]FuncInst),
		importTable:         make(map[Name]map[Name]TableInst),
		importMemory:        make(map[Name]map[Name]MemInst),
		importGlobal:        make(map[Name]map[Name]GlobalInst),
	}
}

func (c *RuntimeConfig) MaxMemorySizeInPage(n uint32) *RuntimeConfig {
	c.maxMemorySizeInPage = n
	return c
}

func (c *RuntimeConfig) AddImportFunc(module, name Name) *RuntimeConfig {
	if _, ok := c.importFunc[module]; !ok {
		c.importFunc[module] = make(map[Name]FuncInst)
	}

	c.importFunc[module][name] = FuncInst{
		HostCode: NewHostFunc(Name(module), Name(name)),
	}

	return c
}

func (c *RuntimeConfig) AddImportTable(module, name Name) *RuntimeConfig {
	/*
		c.importTable[module][name] = TableInst{
			Elem:
			Max:
		}
	*/

	return c
}

func (c *RuntimeConfig) AddImportMemory(module, name Name, b []byte, max *uint32) *RuntimeConfig {
	if _, ok := c.importMemory[module]; !ok {
		c.importMemory[module] = make(map[Name]MemInst)
	}
	c.importMemory[module][name] = MemInst{
		Data: b,
		Max:  max,
	}

	return c
}

func (c *RuntimeConfig) AddImportGlobal(module, name Name, val Val, mut Mut) *RuntimeConfig {
	if _, ok := c.importGlobal[module]; !ok {
		c.importGlobal[module] = make(map[Name]GlobalInst)
	}
	c.importGlobal[module][name] = GlobalInst{
		Value: val,
		Mut:   mut,
	}

	return c
}

func (c *RuntimeConfig) getImportFunc(im *Import) *FuncInst {
	m, ok := c.importFunc[im.Mod]
	if !ok {
		return nil
	}

	fi, ok := m[im.Nm]
	if !ok {
		return nil
	}

	return &fi
}

func (c *RuntimeConfig) getImportTable(im *Import) *TableInst {
	m, ok := c.importTable[im.Mod]
	if !ok {
		return nil
	}

	ti, ok := m[im.Nm]
	if !ok {
		return nil
	}

	return &ti
}

func (c *RuntimeConfig) getImportMemory(im *Import) *MemInst {
	m, ok := c.importMemory[im.Mod]
	if !ok {
		return nil
	}

	mi, ok := m[im.Nm]
	if !ok {
		return nil
	}

	return &mi
}

func (c *RuntimeConfig) getImportGlobal(im *Import) *GlobalInst {
	m, ok := c.importGlobal[im.Mod]
	if !ok {
		return nil
	}

	gi, ok := m[im.Nm]
	if !ok {
		return nil
	}

	return &gi
}
