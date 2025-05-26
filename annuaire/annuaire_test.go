package annuaire

import "testing"

func TestAjouterEtRechercher(t *testing.T) {
	Ajouter("Alice", "0102030405")
	c, ok := Rechercher("Alice")
	if !ok || c.Tel != "0102030405" {
		t.Errorf("Échec de la recherche du contact ajouté")
	}
}

func TestSupprimer(t *testing.T) {
	Ajouter("Bob", "0607080910")
	Supprimer("Bob")
	_, ok := Rechercher("Bob")
	if ok {
		t.Errorf("Le contact n'a pas été supprimé correctement")
	}
}
