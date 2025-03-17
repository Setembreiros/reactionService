package provider

type Provider struct {
	env string
}

func NewProvider(env string) *Provider {
	return &Provider{
		env: env,
	}
}
