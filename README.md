# ShowMeWhatYouGot

version: v0.1

Utilitaire simple pour afficher du contenu de fichier via http

## Conf

A la racine un fichier config.yaml doit être présent

### Example

```
server: example.host
port : 3112
file:
    puppet: /var/log/puppet/puppet.log
    acces: /var/log/http/acces.log
```

## Spec

Une limite de 2Mo a été fixé (arbitrairement) sur la longueur des fichiers. Dans ce cas la, les derniers 500ko sont affichés.

Le ws /infos liste tous les fichiers exposés.
