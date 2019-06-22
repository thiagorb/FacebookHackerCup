export GOPATH="$(pwd)/"

while true
do
    random_input="$(go run cmd/random_input.go 1 3)"

    if diff <(go run cmd/main.go <<< "$random_input") <(go run cmd/main.go --brute <<< "$random_input")
    then
        echo "same results"
    else
        echo "$random_input"
    fi
done
