package configs

const (
	jsonType = "json"
	yamlType = "yaml"
	tomlType = "toml"
)

var client Client

// InitConfigClient is used to initialise and get the instance of a config client
func InitConfigClient(options Options) (err error) {

	switch options.Provider {
	case FileBased:
		client, err = newFileBasedClient(options.Params)
	default:
		err = ErrProviderNotSupported
	}

	return
}

func Get() Client {
	return client
}

func Close() error {
	return client.Close()
}
