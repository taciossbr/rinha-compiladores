IFS='.'
read -ra FILENAME <<< $1
rinha ${FILENAME[0]}.rinha > ${FILENAME[0]}.json

cd app && go run cmd/run.go ../${FILENAME[0]}.json