package config

import (
	"encoding/json"
	"net/url"
	"os"
)

type VaultClients []VaultClient

type VaultClient struct {
	Vault            *url.URL
	AccountDirectory *url.URL
	Unlock           []string
	Authentication   VaultClientAuthentication
	TLS              vaultClientTLS
}

type environmentVariable url.URL

func (e environmentVariable) Get() string {
	u := url.URL(e)
	return os.Getenv(u.Host)
}

func (e environmentVariable) IsSet() bool {
	u := url.URL(e)
	if u.Host == "" {
		return false
	}
	_, b := os.LookupEnv(u.Host)
	return b
}

type VaultClientAuthentication struct {
	Token       *environmentVariable
	RoleId      *environmentVariable
	SecretId    *environmentVariable
	ApprolePath string
}

type vaultClientTLS struct {
	CaCert     *url.URL
	ClientCert *url.URL
	ClientKey  *url.URL
}

type vaultClientJSON struct {
	Vault            string
	AccountDirectory string
	Unlock           []string
	Authentication   vaultClientAuthenticationJSON
	Tls              vaultClientTLSJSON
}

type vaultClientAuthenticationJSON struct {
	Token       string
	RoleId      string
	SecretId    string
	ApprolePath string
}

type vaultClientTLSJSON struct {
	CaCert     string
	ClientCert string
	ClientKey  string
}

func (c *VaultClient) UnmarshalJSON(b []byte) error {
	j := new(vaultClientJSON)
	if err := json.Unmarshal(b, j); err != nil {
		return err
	}
	vc, err := j.vaultClient()
	if err != nil {
		return err
	}
	*c = vc
	return nil
}

func (c vaultClientJSON) vaultClient() (VaultClient, error) {
	vault, err := url.Parse(c.Vault)
	if err != nil {
		return VaultClient{}, err
	}

	accountDirectory, err := url.Parse(c.AccountDirectory)
	if err != nil {
		return VaultClient{}, err
	}

	authentication, err := c.Authentication.vaultClientAuthentication()
	if err != nil {
		return VaultClient{}, err
	}

	tls, err := c.Tls.vaultClientTls()
	if err != nil {
		return VaultClient{}, err
	}

	return VaultClient{
		Vault:            vault,
		AccountDirectory: accountDirectory,
		Unlock:           c.Unlock,
		Authentication:   authentication,
		TLS:              tls,
	}, nil
}

func (c vaultClientAuthenticationJSON) vaultClientAuthentication() (VaultClientAuthentication, error) {
	token, err := url.Parse(c.Token)
	if err != nil {
		return VaultClientAuthentication{}, err
	}

	roleId, err := url.Parse(c.RoleId)
	if err != nil {
		return VaultClientAuthentication{}, err
	}

	secretId, err := url.Parse(c.SecretId)
	if err != nil {
		return VaultClientAuthentication{}, err
	}

	var (
		tEnv = environmentVariable(*token)
		rEnv = environmentVariable(*roleId)
		sEnv = environmentVariable(*secretId)
	)

	return VaultClientAuthentication{
		Token:       &tEnv,
		RoleId:      &rEnv,
		SecretId:    &sEnv,
		ApprolePath: c.ApprolePath,
	}, nil
}

func (c vaultClientTLSJSON) vaultClientTls() (vaultClientTLS, error) {
	caCert, err := url.Parse(c.CaCert)
	if err != nil {
		return vaultClientTLS{}, err
	}

	clientCert, err := url.Parse(c.ClientCert)
	if err != nil {
		return vaultClientTLS{}, err
	}

	clientKey, err := url.Parse(c.ClientKey)
	if err != nil {
		return vaultClientTLS{}, err
	}

	return vaultClientTLS{
		CaCert:     caCert,
		ClientCert: clientCert,
		ClientKey:  clientKey,
	}, nil
}