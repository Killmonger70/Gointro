# !/bin/bash
# Nicolas ramirez
# Commit script

echo "que va en el commit?: " 
read COMMIT

git commit -am "${COMMIT}"
git pull origin main
git push origin main

git status
