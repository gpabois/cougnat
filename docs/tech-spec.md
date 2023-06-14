# Document technique - Cougnat

## Préambule
Ce document formalise les spécifications et principes techniques employées dans la conception du logiciel de signalement de nuisance (Cougnat)

## Signalement (Reporting)
Chaque signalement est enregistré et tracé par le service signalement.

Il notifie via un bus évenementiel (comme un courtier de message: RabbitMQ) la création/suppression de signalements.

## Surveillance des Nuisances (Monitoring)
Chaque signalement est transféré au service Surveillance de la Pollution (Monitoring).

A des fins d'optimisation, le signalement est stocké par échantillonnage suivant deux espaces :
- Sur une période de temps fixe  dit "Période de temps Plancher" (par défaut 1 minute)
- Sur une tuile de carte à zoom fixe dit "Zoom Plancher" (par défaut 16)

La combinaison tuile et période de temps est appelée tuile temporisée (Timed Tile) et forme un espace discret à 4D.

A temps constant, le signalement est également stocké dans une série de tuiles parentes, du zoom plancher au zoom plafond. 

Ainsi, tout signalement est regroupé par tuile temporisé pour assurer :
- la génération des tuiles de pollution pour être affiché sur une carte glissante (Slippy Map) ;
- l'affichage de données statistiques (courbes, répartition...) par séries temporelles ;

Une fois définie, la période de temps ne doit pas être modifiée pour maintenir la cohérence des données de pollution.

### Recommandation pour la configuration de la surveillance des nuisances
- Période de temps plancher : 1 mn
- Zoom plancher : 18
- Zoom plafond : 11

### Limites et considérations

L'hypothèse prise est la zone géographique de la France.

La France est largement couverte par 4 tuiles à Z=11. 

Pour un zoom plancher de 16, le pire scénario comprend 4 096 tuiles/mn, soit environ 2,152 milliards de tuiles/an.

La précision d'une tuile Z=16, est de 19"², soit une surface d'environ 0,24 km² au niveau de Bordeaux. Elle apparaît dès lors largement satisfaisante pour les besoins d'une surveillance de la pollution.

La génération d'une série de tuiles du zoom plancher au plafond (4) génère dans le pire des scénarios 5 460 tuiles/mn, soit un surcoût de 33 %. 

Ce surcoût permet d'économiser du temps de génération pour les tuiles de pollution à temps constant.

Sur une année, le nombre de tuiles maximales à stockée serait de 2 869 776 000 tuiles, soit pour une hypothèse d'un taille de données par tuile de 20 ko une taille de stockage de 57 To/an. 

L'hypothèse ne prend pas en compte l'indexation de la tuile suivant ses coordonnées (X, Y, Z, T).

La problématique de stockage de données apparaît une limite véritable du système. 

Un sharding serait préférable sur une échelle de temps annuel pour limiter d'éventuelles dégradations de performance pour les requêtes à fonctionnement nominal (surveillance en temps réel).