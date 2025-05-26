# TP Annuaire de Contacts

Un programme simple en Go pour gérer un annuaire de contacts.

## Commandes

```bash
# Installation
go mod init github.com/devbutant/tp-note
go mod tidy

# Exécution
go run main.go -action lister                    # Lister tous les contacts
go run main.go -action ajouter -nom "Jean" -tel "0123456789"    # Ajouter un contact
go run main.go -action rechercher -nom "Jean"    # Rechercher un contact
go run main.go -action modifier -nom "Jean" -tel "0987654321"   # Modifier un contact
go run main.go -action supprimer -nom "Jean"     # Supprimer un contact
```

## Fonctionnalités

- Ajouter un contact
- Modifier un contact
- Supprimer un contact
- Rechercher un contact
- Lister tous les contacts

## Structure du projet

```
tp-note/
├── annuaire/
│   ├── annuaire.go      # Logique de gestion des contacts
│   └── annuaire_test.go # Tests unitaires
├── main.go              # Point d'entrée du programme
└── README.md           # Documentation
```

## Contacts par défaut

Le programme est initialisé avec deux contacts par défaut :
- Charlie (0811223344)
- Marie Martin (0987654321)

## Développement

Pour ajouter de nouvelles fonctionnalités ou modifier le code existant :

1. Modifiez les fichiers dans le dossier `annuaire/` pour la logique métier
2. Modifiez `main.go` pour ajouter de nouvelles commandes
3. Exécutez les tests avec `go test ./...` 