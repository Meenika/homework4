package maps


import (
	"slices"
	"sort"
)


type User struct {
	ID     int
	Name   string
	City   string
	Active bool
	Tags   []string
}

type Product struct {
	SKU      string
	Quantity int
	Price    int
}

func CountWords(words []string) map[string]int {
	wordsCopy := make(map[string]int, len(words))
	for _, word := range words{
		wordsCopy[word]++
	}

	return wordsCopy
}

func GetOrDefault(values map[string]int, key string, fallback int) int {
	// TODO: вернуть значение существующего ключа или fallback, если ключ отсутствует.
	// Нулевое значение по существующему ключу должно возвращаться как обычное значение.
	for k, value := range values {
		if k == key {
			return value
		}
	}
	return fallback
}

func DeleteAndReport(values map[string]int, key string) bool {
	for k, _ := range values {
		if k == key {
			delete(values, key)
			return true
		}
	}
	// TODO: удалить ключ и сообщить, существовал ли он до вызова.
	// Для отсутствующего ключа, пустой и nil map вернуть false без panic.
	return false
}

func MergeCounters(left, right map[string]int) map[string]int {
	result := make(map[string]int)
	for k, v := range left {
		result[k] = v
	}
	for k, v := range right {
		_, ok := result[k]
		if ok {
			result[k] += v
		} else {
			result[k] = v
		}
	}
	// TODO: объединить два счётчика в новой map, складывая значения общих ключей.
	// Исходные map после вызова должны остаться без изменений.
	return result
}

func KeysByValueAtLeast(values map[string]int, min int) []string {
	// TODO: вернуть ключи, значения которых не меньше min.
	// Результат должен иметь стабильный алфавитный порядок.
	keys := make([]string, 0)
	for k, value := range values {
		if value >= min {
			keys = append(keys, k)
		}
	}
	slices.Sort(keys)
	return keys
}

func MaxKeyByValue(values map[string]int) (string, bool) {
	found := false
	maxNum := 0
	saveKey := ""
	if len(values) == 0 {
		return "", false
	}
	for k, v := range values {
	    if !found || v > maxNum || (v == maxNum && k < saveKey) {
        maxNum = v
        saveKey = k
        found = true
    }
	}
	// TODO: найти ключ с максимальным значением.
	// Для пустого входа вернуть "", false; при равенстве выбрать меньший ключ по алфавиту.
	return saveKey, true
}

func BuildUserIndex(users []User) map[int]User {
	index := make(map[int]User)
	for _, user := range users {
		index[user.ID] = user
	}
	// TODO: построить индекс пользователей по ID.
	// Если один ID встречается несколько раз, в результате должен остаться последний пользователь.
	return index
}

func GroupActiveUsersByCity(users []User) map[string][]User {
	activeUsers := make(map[string][]User)
	for _, user := range users {
		if user.Active {
			activeUsers[user.City] = append(activeUsers[user.City], user)
		}
	}
	// TODO: сгруппировать только активных пользователей по городу.
	// Порядок пользователей внутри каждого города должен совпадать с исходным слайсом.
	return activeUsers
}

func CountTags(users []User) map[string]int {
	countTags := make(map[string]int)
    for _, user := range users {
        for _, tag := range user.Tags {
            countTags[tag]++
        }
    }

	// TODO: посчитать все появления тегов у всех пользователей.
	// Повторяющийся тег у одного пользователя также считается отдельным появлением.
	return countTags
}

func BuildInventory(products []Product) map[string]Product {
	storage := make(map[string]Product)
	for _, p := range products {
		storage[p.SKU] = p
	}
	// TODO: построить склад по SKU.
	// При повторном SKU оставить последнюю запись из входного слайса.
	return storage
}

func ReserveStock(inventory map[string]Product, sku string, count int) bool {
	if count <= 0 {
		return false
	}
	if inventory[sku].Quantity < count {
		return false
	}
	product, ok := inventory[sku]
	if !ok {
		return false
	}
	product.Quantity -= count
	inventory[sku] = product
	// TODO: уменьшить остаток существующего товара на положительное count, если товара хватает.
	// При неуспехе вернуть false и не менять склад.
	return true
}

func Restock(inventory map[string]*Product, sku string, count int) bool {
	if count <= 0 {
		return false
	}
	product, ok := inventory[sku]
	if !ok || product == nil{
		return false
	}
	product.Quantity += count
	// TODO: увеличить остаток товара, хранящегося в map как указатель.
	// Ключ должен существовать, указатель не должен быть nil, count должен быть положительным.
	return true
}

func LowStockSKUs(inventory map[string]Product, limit int) []string {
    result := make([]string, 0)
    for sku, product := range inventory {
        if product.Quantity <= limit {
            result = append(result, sku)
        }
    }
    sort.Strings(result)
	return result
}

func InventoryValue(inventory map[string]Product) int {
    total := 0
    for _, product := range inventory {
        total += product.Quantity * product.Price
    }
    return total
}

func ApplyPriceUpdates(inventory map[string]Product, updates map[string]int) int {
    count := 0
    for sku, newPrice := range updates {
        if newPrice <= 0 {
            continue
        }
        product, ok := inventory[sku]
        if !ok {
            continue
        }
        product.Price = newPrice
        inventory[sku] = product
        count++
    }
    return count
}
