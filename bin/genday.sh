#!/usr/bin/env sh

DAY=$1

mkdir -p days/day${DAY} data/${DAY}


echo "Generate day ${DAY}"
sed -e "s/DY/${DAY}/g" < tpl/day.go.tpl > "days/day${DAY}/day.go"
sed -e "s/DY/${DAY}/g" < tpl/day_test.go.tpl > "days/day${DAY}/day_test.go"
touch "data/${DAY}/01.twitter"
echo "Done"
