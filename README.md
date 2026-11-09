# Push-Swap & Checker

Ce projet consiste à implémenter un algorithme de tri non-comparatif en Go en utilisant deux piles (`a` et `b`) ainsi qu'un ensemble d'instructions prédéfinies. Le projet est composé de deux programmes distincts : **push-swap** et **checker**.

## Présentation des programmes

1. **push-swap** : Prend en argument une liste d'entiers non triés pour initialiser la pile `a`. Le programme calcule et affiche sur la sortie standard (`Stdout`) la suite d'instructions la plus courte possible permettant de trier la pile `a` par ordre croissant.

2. **checker** : Prend la même liste d'entiers en argument pour initialiser la pile `a`, puis lit les instructions reçues sur l'entrée standard (`Stdin`). Le programme exécute les instructions une à une puis affiche :
   * `OK` si la pile `a` est triée et que la pile `b` est vide.
   * `KO` dans tous les autres cas.

## Instructions autorisées

| Instruction | Description |
| --- | --- |
| **pa** | *Push A* : prend le premier élément au sommet de `b` et le place au sommet de `a`. |
| **pb** | *Push B* : prend le premier élément au sommet de `a` et le place au sommet de `b`. |
| **sa** | *Swap A* : échange les 2 premiers éléments au sommet de la pile `a`. |
| **sb** | *Swap B* : échange les 2 premiers éléments au sommet de la pile `b`. |
| **ss** | Exécute `sa` et `sb` en même temps. |
| **ra** | *Rotate A* : décale tous les éléments de la pile `a` d'une position vers le haut (le 1er devient le dernier). |
| **rb** | *Rotate B* : décale tous les éléments de la pile `b` d'une position vers le haut. |
| **rr** | Exécute `ra` et `rb` en même temps. |
| **rra** | *Reverse Rotate A* : décale tous les éléments de la pile `a` d'une position vers le bas (le dernier devient le 1er). |
| **rrb** | *Reverse Rotate B* : décale tous les éléments de la pile `b` d'une position vers le bas. |
| **rrr** | Exécute `rra` et `rrb` en même temps. |

## Structure du projet

```text
.
├── checker/
│   ├── main.go         # Point d'entrée de l'exécutable checker (lecture Stdin, validation finale)
│   └── operations.go   # Types Piles, méthodes de manipulation et validation des arguments
├── push-swap/
│   ├── main.go         # Point d'entrée de l'exécutable push-swap
│   └── logique.go      # Algorithmes de tri, méthodes de manipulation et validation
└── README.md           # Documentation du projet
```

## Compilation

Chaque programme doit être compilé séparément dans son dossier respectif :

```bash
# Compilation de push-swap
cd push-swap
go build -o push-swap

# Compilation de checker
cd ../checker
go build -o checker
```

## Utilisation

### push-swap

Affiche la suite d'opérations pour trier la pile donnée en argument :

```bash
$ ./push-swap "2 1 3 6 5 8"
pb
pb
ra
sa
rrr
pa
pa
```

### checker

Exécute les instructions données sur `Stdin` et vérifie le résultat :

```bash
$ echo -e "rra\npb\nsa\nrra\npa" | ./checker "3 2 1 0"
OK
```

### Combinaison avec un Pipe

Transmettez directement la sortie de `push-swap` dans `checker` pour valider le tri :

```bash
$ ARG="4 67 3 87 23"; ./push-swap "$ARG" | ./checker "$ARG"
OK
```

Pour compter le nombre d'instructions générées par `push-swap` :

```bash
$ ARG="4 67 3 87 23"; ./push-swap "$ARG" | wc -l
6
```

## Gestion des Erreurs

En cas d'erreur de saisie ou de formatage, les programmes écrivent `Error` suivi d'un retour à la ligne sur la **sortie d'erreur standard (`Stderr`)** et s'arrêtent avec un code de retour `1`.

Les cas d'erreur incluent :
* Un argument qui n'est pas un nombre entier valide (ex: `"0 un 2"`).
* La présence de doublons dans la liste fournie (ex: `"1 2 2 3"`).
* Dans le cas de `checker` : une instruction inconnue ou mal formatée lue sur `Stdin`.
* Si aucun argument n'est fourni, le programme s'arrête simplement sans rien afficher.