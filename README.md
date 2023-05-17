# Analyse dynamique de modules de kernel Windows


## Objectifs

- Démocratiser l'utilisation du débogueur fourni par Microsoft, Windbg.  Le logiciel peut paraitre intimidant à première vue mais après avoir appris quelques commandes et en le combinant avec Ghidra, on arrive à y voir plus clair.

- Expérimenter en développant notre propre fuzzer.  Voir comment on peut influencer le contenu de variable clé ainsi que le flow d'exécution du pilote depuis notre application en "user space".


## Fuzzer

Base de fuzzer dans go-kernel-fuzz.  Ajustez le code au besoin après analyse statique, compilez avec `go build`

## Exploit

Squelette d'exploit dans go-kernel-exploit.  Inspiré d'un exploit d'un pilote développé par Lenovo ici: [https://github.com/alfarom256/CVE-2022-3699/tree/main/CVE-2022-3699](https://github.com/alfarom256/CVE-2022-3699/tree/main/CVE-2022-3699)

Par contre, pas possible de reproduire la technique par laquelle ils arrivent à lire le contenu d'une cellule mémoire spécifiée avec une adresse virtuelle.


