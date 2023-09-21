IFS='.'
read -ra FILENAME <<< $1
rinha ${FILENAME[0]}.rinha > ${FILENAME[0]}.json

cd app && go run main.go ../${FILENAME[0]}.json