package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("test", "test@test.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "test", client.Name)
	assert.Equal(t, "test@test.com", client.Email)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("test", "test@test.com")
	err := client.Update("test2", "test@test.com")
	assert.Nil(t, err)
	assert.Equal(t, "test2", client.Name)
	assert.Equal(t, "test@test.com", client.Email)
}

func TestUpdateClientWithInvalidArgs(t *testing.T) {
	client, _ := NewClient("test", "test@test.com")
	err := client.Update("", "test@test.com")
	assert.Error(t, err, "name is required")
}

func TestAddAccountToClient(t *testing.T) {
	client, _ := NewClient("test", "test@test.com")
	account := NewAccount(client)
	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
}
