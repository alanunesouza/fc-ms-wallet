package gateway

import "github.com.br/alanunesouza/fc-ms-wallet/internal/entity"

type TransactionGatway interface {
	Create(transaction *entity.Transaction) error
}
