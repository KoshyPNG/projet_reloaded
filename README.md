# Projet Reloaded

Ce projet est un outil Go qui lit un fichier texte, applique plusieurs transformations linguistiques et de formatage, puis écrit le résultat dans un autre fichier.

Il est conçu pour corriger et normaliser des phrases selon des règles spécifiques comme :

- correction des articles `a` / `an`
- conversion en majuscules, minuscules ou capitalisation
- conversion hexadécimale et binaire
- gestion de la ponctuation, apostrophes et guillemets
- suppression des mots vides et nettoyage de l’édition de texte

## Structure du projet

- `main/main.go` : point d’entrée du programme
- `format.go` : nettoyage et formatage des mots
- `erreur.go` : application des règles de transformation
- `up_low_cap.go` : gestion des majuscules/minuscules/capitales
- `hexa_bin.go` : conversion hexadécimale et binaire
- `point.go` / `apo_guil.go` : ponctuation, apostrophes et guillemets
- `sample.txt` : exemple d’entrée
- `sortie.txt` : exemple de sortie obtenue

## Prérequis

- Go 1.20 ou plus

## Utilisation

Depuis la racine du projet, exécutez :

```bash
go run ./main fichier_entree.txt fichier_sortie.txt
```

Exemple :

```bash
go run ./main sample.txt sortie.txt
```

Le programme lit le contenu du fichier d’entrée, applique les transformations, puis écrit le résultat dans le fichier de sortie.

## Exemple de sample.txt

Entrée :

```text
it (cap) was the best of times, it was the worst of times (up) ,
```

Sortie :

```text
It was the best of times, it was the worst of TIMES,
```

## Projet

Projet réalisé dans le cadre de l’exercice de programmation en Go.
