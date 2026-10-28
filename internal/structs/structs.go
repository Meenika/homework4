package structs

import (
	"unsafe"
	"fmt"
)

type User struct {
	ID         int
	Name       string
	Active     bool
	LoginCount int
	Tags       []string
}

func UpdateUserName(users map[int]User, id int, name string) bool {
	// TODO: изменить имя существующего пользователя, хранящегося в map как значение.
	// Если пользователя нет, вернуть false и не добавлять новый ключ.
	user, ok := users[id]
	if ok {
		user.Name = name
		users[id] = user
		return ok
	}
	return ok
}

func ActivateUser(users map[int]User, id int) bool {
	// TODO: сделать существующего пользователя активным.
	// Если пользователя нет, вернуть false.
	user, ok := users[id]
	if ok {
		user.Active = true
		users[id] = user
		return ok
	}
	return false
}

func IncrementLoginCount(users map[int]User, id int) bool {
	// TODO: увеличить LoginCount существующего пользователя на единицу.
	// Если пользователя нет, вернуть false.
	user, ok := users[id]
	if ok {
		user.LoginCount += 1
		users[id] = user
		return ok
	}
	return false
}

func CopyUsers(users map[int]User) map[int]User {
	// TODO: вернуть независимую копию map и вложенных слайсов Tags.
	// Изменения любой части результата не должны менять исходные данные.
	usersCopy := make(map[int]User, len(users))
	for id, user := range users {
		if user.Tags != nil {
			tagsCopy := make([]string, len(user.Tags))
			copy(tagsCopy, user.Tags)
			user.Tags = tagsCopy
		}
		usersCopy[id] = user
	}
	return usersCopy
}

func BuildPointerIndex(users []User) map[int]*User {
	index := make(map[int]*User, len(users))
	for i := range users {
		userCopy := users[i]
		if userCopy.Tags != nil {
			tagsCopy := make([]string, len(userCopy.Tags))
			copy(tagsCopy, userCopy.Tags)
			userCopy.Tags = tagsCopy
		}
		index[userCopy.ID] = &userCopy
	}
	return index
}

func RenameThroughPointer(users map[int]*User, id int, name string) bool {
	user, ok := users[id]
	if !ok || user == nil {
		return false
	}
	user.Name = name
	return true
}

func (u User) DisplayName() string {
	str := fmt.Sprintf("%d:%s", u.ID, u.Name)
	// TODO: вернуть строку "ID:Name".
	return str
}

type Admin struct {
	User
	Permissions []string
}

func (a Admin) DisplayName() string {
	str := fmt.Sprintf("admin:%d:%s", a.User.ID, a.User.Name)
	// TODO: вернуть строку "admin:ID:Name", перекрывая одноимённый метод встроенного User.
	return str
}

func (a Admin) HasPermission(permission string) bool {
	for i := range a.Permissions {
		if permission == a.Permissions[i] {
			return true
		}
	}
	// TODO: проверить точное наличие permission в списке администратора.
	// Регистр, пробелы и повторения не нормализуются.
	return false
}

func RenameEmbeddedUser(admin *Admin, name string) {
	if admin == nil {
		return
	}
	admin.User.Name = name
}

type Account struct {
	ID      int
	Balance int
}

type PremiumAccount struct {
	Account
	Bonus int
}

func (p PremiumAccount) Total() int {
	// TODO: вернуть сумму основного баланса и бонуса.
	sum := p.Account.Balance + p.Bonus
	return sum
}

func PremiumLabel(p PremiumAccount) string {
	// TODO: вернуть строку "ID:Balance total=Total".
	str := fmt.Sprintf("%d:%d total=%d", p.Account.ID, p.Account.Balance, p.Total())
	return str
}

type BadLayout struct {
	Active   bool
	Amount   int64
	Verified bool
	Retries  int32
}

type GoodLayout struct {
	Amount   int64
	Retries  int32
	Active   bool
	Verified bool
}

func UserValueSize(user User) uintptr {
	// TODO: вернуть фактический размер переданного значения User в байтах.
	return unsafe.Sizeof(user)
}

func UserPointerSize(user *User) uintptr {
	// TODO: вернуть фактический размер переменной-указателя на User в байтах.
	// Результат не зависит от того, nil указатель или нет.
	return unsafe.Sizeof(user)
}

func LayoutSizes() (bad, good, saved uintptr) {
	// TODO: вернуть размеры BadLayout и GoodLayout, затем число сэкономленных байт.
	bad = unsafe.Sizeof(BadLayout{})
	good = unsafe.Sizeof(GoodLayout{})
	saved = bad - good
	return bad, good, saved
}


