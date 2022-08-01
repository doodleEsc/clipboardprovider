package clipboardprovider

type Provider interface {
	Copy([]byte)
	Paste() bool
}

type ProviderSet []Provider

func (ps ProviderSet) Copy(data []byte) {
    for _, p := range ps {
        p.Copy(data)
    }
}

func (ps ProviderSet) Paste() bool {
    for _, p := range ps {
        if ok := p.Paste(); ok {
            return true
        }
    }
    return false
}
