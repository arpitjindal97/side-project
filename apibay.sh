
#data=$(curl --silent 'https://apibay.org/precompiled/data_top100_200.json' | jq '.[].info_hash')
#data=$(curl --silent 'https://apibay.org/q.php?q=category%3A400' | jq '.[].info_hash')
data=$(curl --silent 'https://apibay.org/q.php?q=category%3A200' | jq '.[].info_hash')
#data=$(curl --silent 'https://apibay.org/precompiled/data_top100_300.json' | jq '.[].info_hash')

for info in $data
do

  curl --location --request POST 'http://188.166.76.22:30001/api/torrents' \
  --header 'Content-Type: application/json' \
  --data-raw '{ "infohash":'$info' }'
  echo ""
done

