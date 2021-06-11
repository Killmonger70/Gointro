# !/bin/bash
# Nicolas ramirez
# Commit script

read -p "que va en el commit?" COMMIT

git commit -am "${COMMIT}"
git pull origin main
git push origin main
