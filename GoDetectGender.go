package GoDetectGender

import "strings"

type FullName struct {
	LastName   string
	FirstName  string
	Patronymic string
}

type Gender int

type NameType int

const (
	UNDEFINED Gender = 0
	MALE      Gender = 1
	FEMALE    Gender = 2
)

const (
	LAST_NAME  NameType = 0
	FIRST_NAME NameType = 1
	PATRONYMIC NameType = 2
)

var last_name_list = map[Gender][]string{
	FEMALE: {"ова", "ева", "ина", "ая", "яя", "екая", "цкая"},
	MALE:   {"ов", "ев", "ин", "ын", "ой", "цкий", "ский", "цкой", "ской"},
}

var first_name_list = map[Gender][]string{
	FEMALE: {"авдотья", "аврора", "агата", "агния", "агриппина", "ада", "аксинья", "алевтина", "александра", "алёна", "алена", "алина", "алиса", "алла", "альбина", "амалия", "анастасия", "ангелина", "анжела", "анжелика", "анна", "антонина", "анфиса", "арина", "белла", "божена", "валентина", "валерия", "ванда", "варвара", "василина", "василиса", "вера", "вероника", "виктория", "виола", "виолетта", "вита", "виталия", "владислава", "власта", "галина", "глафира", "дарья", "диана", "дина", "ева", "евгения", "евдокия", "евлампия", "екатерина", "елена", "елизавета", "ефросиния", "ефросинья", "жанна", "зиновия", "злата", "зоя", "ивонна", "изольда", "илона", "инга", "инесса", "инна", "ирина", "ия", "капитолина", "карина", "каролина", "кира", "клавдия", "клара", "клеопатра", "кристина", "ксения", "лада", "лариса", "лиана", "лидия", "лилия", "лина", "лия", "лора", "любава", "любовь", "людмила", "майя", "маргарита", "марианна", "мариетта", "марина", "мария", "марья", "марта", "марфа", "марьяна", "матрёна", "матрена", "матрона", "милена", "милослава", "мирослава", "муза", "надежда", "настасия", "настасья", "наталия", "наталья", "нелли", "ника", "нина", "нинель", "нонна", "оксана", "олимпиада", "ольга", "пелагея", "полина", "прасковья", "раиса", "рената", "римма", "роза", "роксана", "руфь", "сарра", "светлана", "серафима", "снежана", "софья", "софия", "стелла", "степанида", "стефания", "таисия", "таисья", "тамара", "татьяна", "ульяна", "устиния", "устинья", "фаина", "фёкла", "фекла", "феодора", "хаврония", "христина", "эвелина", "эдита", "элеонора", "элла", "эльвира", "эмилия", "эмма", "юдифь", "юлиана", "юлия", "ядвига", "яна", "ярослава"},
	MALE:   {"абрам", "аверьян", "авраам", "агафон", "адам", "азар", "акакий", "аким", "аксён", "александр", "алексей", "альберт", "анатолий", "андрей", "андрон", "антип", "антон", "аполлон", "аристарх", "аркадий", "арнольд", "арсений", "арсентий", "артем", "артём", "артемий", "артур", "аскольд", "афанасий", "богдан", "борис", "борислав", "бронислав", "вадим", "валентин", "валерий", "варлам", "василий", "венедикт", "вениамин", "веньямин", "венцеслав", "виктор", "вилен", "виталий", "владилен", "владимир", "владислав", "владлен", "всеволод", "всеслав", "вячеслав", "гавриил", "геннадий", "георгий", "герман", "глеб", "григорий", "давид", "даниил", "данил", "данила", "демьян", "денис", "димитрий", "дмитрий", "добрыня", "евгений", "евдоким", "евсей", "егор", "емельян", "еремей", "ермолай", "ерофей", "ефим", "захар", "иван", "игнат", "игорь", "илларион", "иларион", "илья", "иосиф", "казимир", "касьян", "кирилл", "кондрат", "константин", "кузьма", "лавр", "лаврентий", "лазарь", "ларион", "лев", "леонард", "леонид", "лука", "максим", "марат", "мартын", "матвей", "мефодий", "мирон", "михаил", "моисей", "назар", "никита", "николай", "олег", "осип", "остап", "павел", "панкрат", "пантелей", "парамон", "пётр", "петр", "платон", "потап", "прохор", "роберт", "ростислав", "савва", "савелий", "семён", "семен", "сергей", "сидор", "спартак", "тарас", "терентий", "тимофей", "тимур", "тихон", "ульян", "фёдор", "федор", "федот", "феликс", "фирс", "фома", "харитон", "харлам", "эдуард", "эммануил", "эраст", "юлиан", "юлий", "юрий", "яков", "ян", "ярослав"},
}

var patronymic_list = map[Gender][]string{
	FEMALE: {"овна", "евна", "ична"},
	MALE:   {"ович", "евич", "ич"},
}

func GetGender(fullName FullName) Gender {
	genderByFirstName := genderBy(FIRST_NAME, fullName.FirstName)
	genderByLastName := genderBy(LAST_NAME, fullName.LastName)
	genderByPatronymic := genderBy(PATRONYMIC, fullName.Patronymic)

	gendersOnNames := []Gender{genderByFirstName, genderByLastName, genderByPatronymic}

	return determineGender(gendersOnNames)
}

func determineGender(genders []Gender) Gender {
	male := false
	female := false
	gender := UNDEFINED

	for _, value := range genders {
		if value == MALE {
			male = true
		}

		if value == FEMALE {
			female = true
		}
	}

	if male && !female {
		gender = MALE
	}

	if !male && female {
		gender = FEMALE
	}

	return gender
}

func genderBy(nameType NameType, name string) Gender {
	name = normalize(name)

	gender := UNDEFINED
	resultFemale := false
	resultMale := false

	switch nameType {
	case LAST_NAME:
		resultFemale = hasSuffixArray(name, last_name_list[FEMALE])
		resultMale = hasSuffixArray(name, last_name_list[MALE])
	case FIRST_NAME:
		resultFemale = in_array(name, first_name_list[FEMALE])
		resultMale = in_array(name, first_name_list[MALE])
	case PATRONYMIC:
		resultFemale = hasSuffixArray(name, patronymic_list[FEMALE])
		resultMale = hasSuffixArray(name, patronymic_list[MALE])
	}

	if resultMale {
		gender = MALE
	} else if resultFemale {
		gender = FEMALE
	}

	return gender
}

func normalize(name string) string {
	return strings.TrimSpace(strings.ToLower(name))
}

func in_array(val string, array []string) bool {
	exists := false

	for _, v := range array {
		if val == v {
			exists = true
			return exists
		}
	}

	return exists
}

func hasSuffixArray(val string, array []string) bool {
	exists := false

	for _, v := range array {
		if strings.HasSuffix(val, v) {
			exists = true
			return exists
		}
	}

	return exists
}
