curl -Ss https://learn.zone01kisumu.ke/assets/superhero/all.json | jq ' .[] | select( .id == 170 ) | .name,.gender,.power '
curl -Ss https://learn.zone01kisumu.ke/assets/superhero/all.json | jq ' .[] | select( .id == 170 ) | .powerstats["power"] '
curl -Ss https://learn.zone01kisumu.ke/assets/superhero/all.json | jq ' .[] | select( .id == 170 ) | .appearance["gender"] ' 

