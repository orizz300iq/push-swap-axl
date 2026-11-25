# Push-Swap & Checker

Implémentation en Go d'un algorithme de tri utilisant deux piles (`a` et `b`) et un jeu d'instructions restreint. Le projet est composé de deux programmes :

- **push-swap** : calcule et affiche la suite d'instructions qui trie la pile `a`.
- **checker** : lit des instructions sur l'entrée standard, les applique, et vérifie si le résultat est trié.

## Sommaire

- [Instructions autorisées](#instructions-autorisées)
- [Structure du projet](#structure-du-projet)
- [Compilation](#compilation)
- [Utilisation](#utilisation)
- [Algorithme de tri](#algorithme-de-tri)
- [Gestion des erreurs](#gestion-des-erreurs)

## Instructions autorisées

| Instruction | Description |
| --- | --- |
| `sa` | *Swap A* : échange les 2 premiers éléments de `a`. |
| `sb` | *Swap B* : échange les 2 premiers éléments de `b`. |
| `ss` | `sa` et `sb` en même temps. |
| `pa` | *Push A* : prend le sommet de `b` et le place au sommet de `a`. |
| `pb` | *Push B* : prend le sommet de `a` et le place au sommet de `b`. |
| `ra` | *Rotate A* : décale `a` vers le haut (le 1er élément devient le dernier). |
| `rb` | *Rotate B* : décale `b` vers le haut. |
| `rr` | `ra` et `rb` en même temps. |
| `rra` | *Reverse Rotate A* : décale `a` vers le bas (le dernier élément devient le 1er). |
| `rrb` | *Reverse Rotate B* : décale `b` vers le bas. |
| `rrr` | `rra` et `rrb` en même temps. |

Une instruction qui n'a pas d'effet (par exemple `sa` sur une pile de moins de 2 éléments) est simplement ignorée.

## Structure du projet

```text
.
├── go.mod
├── README.md
├── checker/
│   ├── main.go         # Lecture de Stdin, exécution des instructions, affichage OK/KO
│   └── operations.go   # Type Piles, instructions, vérification du tri, lecture des arguments
└── push-swap/
    ├── main.go         # Point d'entrée : lecture des arguments, affichage des instructions
    └── logique.go      # Type Piles, instructions, algorithmes de tri, lecture des arguments
```

Les deux programmes sont des `package main` indépendants, d'où la présence d'un type `Piles` dans chaque dossier.

## Compilation

Chaque programme se compile dans son dossier :

```bash
# push-swap
cd push-swap
go build -o push-swap

# checker
cd ../checker
go build -o checker
```

## Utilisation

### push-swap

Prend la liste d'entiers en argument et affiche, une par ligne, les instructions qui trient la pile `a` :

```bash
$ ./push-swap "2 1 0"
sa
rra
```

Les nombres peuvent être passés dans un seul argument entre guillemets ou dans plusieurs arguments :

```bash
$ ./push-swap 2 1 0
sa
rra
```

Si la pile est déjà triée, rien n'est affiché.

### checker

Prend la même liste en argument, lit les instructions sur `Stdin` (une par ligne), les exécute, puis affiche :

- `OK` si `a` est triée **et** `b` est vide ;
- `KO` dans tous les autres cas.

```bash
$ echo -e "rra\npb\nsa\nrra\npa" | ./checker "3 2 1 0"
OK
```

### Utiliser les deux ensemble

```bash
ARG="4 67 3 87 23"
./push-swap "$ARG" | ./checker "$ARG"     # affiche OK
./push-swap "$ARG" | wc -l                # nombre d'instructions générées
```

## Algorithme de tri

`push-swap` choisit la stratégie selon la taille de la pile :

| Taille de `a` | Stratégie |
| --- | --- |
| Déjà triée | Aucune instruction. |
| 2 éléments | Un `sa`. |
| 3 éléments | Analyse des 6 permutations possibles, au plus 2 instructions. |
| 4 à 5 éléments | Tri par sélection : on envoie les plus petits éléments dans `b` (en les amenant au sommet par `ra` ou `rra`, selon le chemin le plus court) jusqu'à ce qu'il en reste 3, on trie ces 3, puis on remet tout dans `a` avec `pa`. |
| 6 éléments ou plus | **Tri radix** (voir ci-dessous). |

### Tri radix (6 éléments ou plus)

1. **Normalisation** : chaque valeur est remplacée par son rang dans la liste triée (0 à n-1). Cela permet de gérer les nombres négatifs et les très grands nombres.
2. **Passes binaires** : pour chaque bit, du poids faible au poids fort :
   - si le bit du sommet de `a` vaut 0, `pb` ;
   - sinon, `ra` ;
   - puis tous les éléments de `b` sont renvoyés dans `a` avec `pa`.
3. **Restauration** : les rangs sont remplacés par les valeurs d'origine.

Le nombre de passes est `⌈log₂(n)⌉`, et chaque passe coûte environ `2n` instructions, soit une complexité en **O(n log n)** instructions.

> Cet algorithme garantit un résultat correct et un nombre d'instructions raisonnable, mais pas la suite d'instructions la plus courte possible.

## Gestion des erreurs

En cas d'erreur, le programme écrit un message d'erreur sur la **sortie d'erreur standard** (`Stderr`) et se termine avec le code de retour `1`.

Cas d'erreur :

- un argument n'est pas un entier valide (ex. `"0 un 2"`) ;
- la liste contient des doublons (ex. `"1 2 2 3"`) ;
- *checker uniquement* : une instruction lue sur `Stdin` est inconnue ou mal formatée.

Si aucun argument n'est fourni, le programme s'arrête sans rien afficher.
