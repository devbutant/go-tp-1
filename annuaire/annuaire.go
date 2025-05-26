package annuaire

type Contact struct {
	Nom  string
	Tel  string
}

var Contacts = make(map[string]Contact)

func Init() {
	Contacts["Charlie"] = Contact{Nom: "Charlie", Tel: "0811223344"}
	Contacts["Marie Martin"] = Contact{Nom: "Marie Martin", Tel: "0987654321"}
}

func Ajouter(nom, tel string) bool {
	if _, exists := Contacts[nom]; exists {
		return false // déjà présent
	}
	Contacts[nom] = Contact{Nom: nom, Tel: tel}
	return true
}

func Modifier(nom, tel string) bool {
	if _, exists := Contacts[nom]; !exists {
		return false
	}
	Contacts[nom] = Contact{Nom: nom, Tel: tel}
	return true
}

func Supprimer(nom string) bool {
	if _, exists := Contacts[nom]; !exists {
		return false
	}
	delete(Contacts, nom)
	return true
}

func Rechercher(nom string) (Contact, bool) {
	contact, exists := Contacts[nom]
	return contact, exists
}

func Lister() []Contact {
	liste := []Contact{}
	for _, c := range Contacts {
		liste = append(liste, c)
	}
	return liste
}
