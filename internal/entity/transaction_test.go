package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	client, _ := NewClient("test", "test@test.com")
	account := NewAccount(client)
	client2, _ := NewClient("test2", "test2@test.com")
	account2 := NewAccount(client2)

	account.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(account, account2, 100)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, 1100.0, account2.Balance)
	assert.Equal(t, 900.0, account.Balance)
}

func TestCreateTransactionWithInsuficientBalance(t *testing.T) {
	client, _ := NewClient("test", "test@test.com")
	account := NewAccount(client)
	client2, _ := NewClient("test2", "test2@test.com")
	account2 := NewAccount(client2)

	account.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(account, account2, 2000)
	assert.NotNil(t, err)
	assert.Error(t, err, "insufficient funds")
	assert.Nil(t, transaction)
	assert.Equal(t, 1000.0, account2.Balance)
	assert.Equal(t, 1000.0, account.Balance)
}

func TestCreateTransactionWithAmountZero(t *testing.T) {
	client, _ := NewClient("test", "test@test.com")
	account := NewAccount(client)
	client2, _ := NewClient("test2", "test2@test.com")
	account2 := NewAccount(client2)

	account.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(account, account2, 0)
	assert.NotNil(t, err)
	assert.Error(t, err, "amount must be greater than zero")
	assert.Nil(t, transaction)
	assert.Equal(t, 1000.0, account2.Balance)
	assert.Equal(t, 1000.0, account.Balance)
}
