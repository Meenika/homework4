package common

import (
	"github.com/rinatkh/homework_go_4/internal/methods"
	"fmt"
)

type Operation struct {
	AccountID int
	Kind      string
	Amount    int
}

type LedgerResult struct {
	Balances map[int]int
	Errors   []string
}

func ApplyOperations(accounts map[int]*methods.Account, operations []Operation) LedgerResult {
	// TODO: последовательно применить deposit/withdraw к счетам и продолжать после ошибок.
	// В Balances вернуть итоговые балансы ненулевых счетов, в Errors — сообщения в порядке операций.
	result := LedgerResult{
		Balances: make(map[int]int),
		Errors:   make([]string, 0),
	}
	for _, operation := range operations {
		account, ok := accounts[operation.AccountID]
		if !ok || account == nil {
			result.Errors = append(result.Errors, fmt.Sprintf("account %d: not found", operation.AccountID),
			)
			continue
		}
		var err error
		switch operation.Kind {
		case "deposit":
			err = account.Deposit(operation.Amount)
		case "withdraw":
			err = account.Withdraw(operation.Amount)
		default:
			result.Errors = append(result.Errors, fmt.Sprintf("account %d: unknown operation %q", operation.AccountID, operation.Kind),
			)
			continue
		}
		if err != nil {
			result.Errors = append(result.Errors, fmt.Sprintf("account %d: %s", operation.AccountID, err.Error()),
			)
		}
	}
	for id, account := range accounts {
		if account != nil {
			result.Balances[id] = account.Balance
		}
	}
	return result
}

func BuildAccountIndex(accounts []methods.Account) map[int]*methods.Account {
	// TODO: построить индекс по ID из независимых копий счетов.
	// Разные ключи не должны вести к одной переменной, изменения индекса не должны менять входной слайс.
	index := make(map[int]*methods.Account, len(accounts))
	for i := range accounts {
		accountCopy := accounts[i]
		index[accountCopy.ID] = &accountCopy
	}
	return index
}

type Event struct {
	UserID int
	Name   string
	Active *bool
}

func ApplyUserEvents(
	users map[int]string,
	active map[int]bool,
	events []Event,
) (map[int]string, map[int]bool) {
	// TODO: применить события к состоянию пользователей и вернуть обе map.
	// Непустое Name обновляет имя, ненулевой Active обновляет статус; nil map должны поддерживаться.
	if users == nil {
		users = make(map[int]string)
	}
	if active == nil {
		active = make(map[int]bool)
	}
	for _, event := range events {
		if event.Name != "" {
			users[event.UserID] = event.Name
		}
		if event.Active != nil {
			active[event.UserID] = *event.Active
		}
	}
	return users, active
}
