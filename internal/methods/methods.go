package methods

import (
	"errors"
	"fmt"
)

var (
	ErrNilAccount        = errors.New("nil account")
	ErrNilDestination    = errors.New("nil destination")
	ErrSameAccount       = errors.New("source and destination are the same")
	ErrInactive          = errors.New("inactive account")
	ErrInvalidAmount     = errors.New("invalid amount")
	ErrInsufficientFunds = errors.New("insufficient funds")
	ErrNilCart           = errors.New("nil cart")
	ErrInvalidSKU        = errors.New("invalid sku")
	ErrCartLimit         = errors.New("cart limit exceeded")
)

type Account struct {
	ID       int
	Owner    string
	Balance  int
	Active   bool
	Currency string
}

func (a Account) Label() string {
	// TODO: вернуть строку "#ID Owner: Balance Currency" с одним пробелом после двоеточия.
	str := fmt.Sprintf("#%d %s: %d %s", a.ID, a.Owner, a.Balance, a.Currency)
	return str
}

func (a Account) CanWithdraw(amount int) bool {
	// TODO: сообщить, можно ли списать amount без нарушения состояния счёта.
	// Счёт должен быть активным, сумма — положительной, денег должно хватать.
	if a.Active == true && amount > 0 && amount <= a.Balance {
		return true
	}
	return false
}

func (a Account) Renamed(owner string) Account {
	// TODO: вернуть копию счёта с новым владельцем.
	// Исходное значение Account не должно изменяться.
	a.Owner = owner
	return a
}

func (a *Account) Rename(owner string) error {
	if a == nil {
		return ErrNilAccount
	}
	// TODO: изменить владельца исходного счёта.
	// Для nil receiver вернуть ErrNilAccount и не паниковать.
	a.Owner = owner
	return nil
}

func (a *Account) Deposit(amount int) error {
	if a == nil {
		return ErrNilAccount
	}
	if !a.Active {
		return ErrInactive
	}
	if amount <= 0 {
		return ErrInvalidAmount
	}
	a.Balance += amount
	// TODO: пополнить активный счёт на положительную сумму.
	// При ошибке вернуть подходящую sentinel-ошибку и сохранить прежний баланс.
	return nil
}

func (a *Account) Withdraw(amount int) error {
	if a == nil {
		return ErrNilAccount
	}
	if !a.Active {
		return ErrInactive
	}
	if amount <= 0 {
		return ErrInvalidAmount
	}
	if amount > a.Balance {
		return ErrInsufficientFunds
	}
	a.Balance -= amount
	// TODO: списать положительную сумму с активного счёта, если денег достаточно.
	// При ошибке вернуть подходящую sentinel-ошибку и сохранить прежний баланс.
	return nil
}

func (a *Account) TransferTo(destination *Account, amount int) error {
	if a == nil {
		return ErrNilAccount
	}
	if destination == nil {
		return ErrNilDestination
	}
	if !a.Active || !destination.Active{
		return ErrInactive
	}
	if amount <= 0 {
		return ErrInvalidAmount
	}
	if amount > a.Balance {
		return ErrInsufficientFunds
	}
	if a == destination {
		return ErrSameAccount
	}
	a.Balance -= amount
	destination.Balance += amount
	// TODO: атомарно перевести сумму между двумя разными активными счетами.
	// При любой ошибке оба баланса должны остаться прежними.
	return nil
}

func (a *Account) Activate() error {
	if a == nil {
		return ErrNilAccount
	}
	if !a.Active {
		a.Active = true
	}
	// TODO: сделать счёт активным.
	// Для nil receiver вернуть ErrNilAccount.
	return nil
}

func (a *Account) Deactivate() error {
	if a == nil {
		return ErrNilAccount
	}
	if a.Active {
		a.Active = false
	}
	// TODO: сделать счёт активным.
	// Для nil receiver вернуть ErrNilAccount.
	return nil
	// TODO: сделать счёт неактивным.
	// Для nil receiver вернуть ErrNilAccount.
	return nil
}

func (a Account) Snapshot() Account {
	copyA := a
	// TODO: вернуть независимую копию текущего значения счёта.
	return copyA
}

func (a *Account) Reset() {
	if a == nil {
		return
	}
	a.ID = 0
	a.Owner = ""
	a.Balance = 0
	a.Active = false
	a.Currency = ""
	// TODO: сбросить все поля существующего счёта в zero value.
	// Вызов для nil receiver должен быть безопасным.
}

func (a *Account) ApplyBonus(percent int) error {
	if a == nil {
		return ErrNilAccount
	}
	if !a.Active {
		return ErrInactive
	}
	if percent < 0 {
		return ErrInvalidAmount
	}
	bonus := a.Balance * percent / 100
	a.Balance += bonus
	// TODO: увеличить баланс активного счёта на указанный неотрицательный процент.
	// Дробная часть от целочисленного расчёта отбрасывается; при ошибке баланс не менять.
	return nil
}

func (a Account) SameOwner(other Account) bool {
	if a.Owner == other.Owner {
		return true
	}
	// TODO: сравнить владельцев двух счетов точным сравнением строк.
	return false
}

type Cart struct {
	Items    map[string]int
	MaxItems int
}

func (c *Cart) Add(sku string, count int) error {
	if c == nil {
		return ErrNilCart
	}
	if sku == "" {
		return ErrInvalidSKU
	}
	if count <= 0 {
		return ErrInvalidAmount
	}
	total := 0
    for _, amount := range c.Items {
        total += amount
    }
    if total + count > c.MaxItems {
        return ErrCartLimit
    }
    if c.Items == nil {
        c.Items = make(map[string]int)
    }
    c.Items[sku] += count
	return nil
}

func (c Cart) TotalItems() int {
	count := 0
	for _, v := range c.Items {
		count += v
	}
	// TODO: вернуть сумму количеств всех товаров корзины.
	return count
}
