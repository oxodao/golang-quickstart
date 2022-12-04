package services

import "github.com/oxodao/go-quickstart/orm"

var GET *Provider

type Provider struct {
}

func Load() error {
	err := orm.Load()
	if err != nil {
		return err
	}

	prv := Provider{}

	GET = &prv

	return nil
}
