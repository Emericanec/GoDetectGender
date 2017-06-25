# О библиотеке:

Определение пола по ФИО для русского языка на Golang

# Установка:
```go
    import "github.com/emericanec/GoDetectGender"
```

# Использование:
```go
    maleFullName := GoDetectGender.FullName{"Иванов", "Иван", "Иванович"}
	result := GoDetectGender.GetGender(maleFullName)

	if result == GoDetectGender.MALE {
	    fmt.Println("Это мужчина")
	} else if result == GoDetectGender.FEMALE {
	    fmt.Println("Это женщина")
	} else if result == GoDetectGender.UNDEFINED {
	    fmt.Println("Не удалось определить пол")
	}
```