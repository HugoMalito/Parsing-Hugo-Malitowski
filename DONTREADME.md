# Projet Lilian Piver : Le Go for the Reload

Installe les dispositifs applicatifs de commande et d'orchestration des opérations numériques "Go", "Git" et "VS Code" sur ton dispositif informatique de traitement et de restitution des données numériques.

Pour cloner un dépôt Git sur ton dispositif informatique de traitement et de restitution des données numériques :

1. Récupère le lien du dépôt que tu souhaites importer.
2. Ouvre ton terminal de commande.
3. Saisis la commande suivante : `git clone [lien_du_dépôt]`
4. Laisse ton ordinateur procéder à la récupération des fichiers et de leur historique.
5. Une fois l'opération achevée, accède au répertoire fraîchement constitué avec la commande `cd [nom_du_dépôt]`.

Maintenant que tu disposes désormais d'une copie du projet, ouvre ton dispositif applicatif de commande et d'orchestration des opérations numériques "VS Code" et ouvre le dossier du projet.

---

## Présentation du projet et fonctionnalités

**Go for the Reload** est un outil de correction de texte écrit en langage Go. Il lit le fichier `sample.txt`, applique des transformations selon des marqueurs placés dans le texte, puis écrit le résultat dans `output.txt`.

Fonctionnalités :

| Marqueur | Effet | Exemple |
|----------|-------|---------|
| `(hex)` | Convertit le mot précédent de l'hexadécimal vers le décimal | `42 (hex)` devient `66` |
| `(bin)` | Convertit le mot précédent du binaire vers le décimal | `10 (bin)` devient `2` |
| `(up)` | Met le mot précédent en majuscules | `times (up)` devient `TIMES` |
| `(low)` | Met le mot précédent en minuscules | `HELLO (low)` devient `hello` |
| `(cap)` | Met la première lettre du mot précédent en majuscule | `it (cap)` devient `It` |

Le programme corrige aussi la ponctuation : les espaces avant `. , ! ? : ;` sont retirés, et un espace est ajouté après si nécessaire (`boring ,what` devient `boring, what`).

Si un mot ne peut pas être converti (par exemple `zzz (hex)`), il est laissé tel quel.

---

## Prérequis

Avant de commencer, assure-toi que les éléments suivants sont installés sur ton dispositif informatique de traitement et de restitution des données numériques :

- Le dispositif applicatif de commande et d'orchestration des opérations numériques **Go** (version indiquée dans le fichier `go.mod`)
- Le dispositif applicatif de commande et d'orchestration des opérations numériques **Git**
- Le dispositif applicatif de commande et d'orchestration des opérations numériques **VS Code**

Le projet n'utilise que la bibliothèque standard de Go : aucune dépendance externe à télécharger.

---

## Comment ça marche

1. Dans ton terminal de commande, place-toi dans le répertoire du projet.
2. Écris ton texte dans le fichier `sample.txt`, avec les marqueurs de ton choix, séparés des mots par un espace.
3. Lance le projet avec la commande : `go run .`
4. Laisse ton ordinateur exécuter les opérations numériques demandées.
5. Ouvre le fichier `output.txt` généré dans le répertoire : il contient ton texte corrigé.

Étapes internes du traitement :

1. `Cut` découpe le texte en mots.
2. `Hexa`, `NonBinaire`, `LeMaxeur`, `MartiMinimus` et `Capitalisme` appliquent les marqueurs.
3. Les mots sont recollés, puis `PointSurLesI` corrige la ponctuation.

---

Hugo Malitowski