export GOPATH="$(pwd)/"

createPipe () {
    PIPE=$(mktemp -u)
    mkfifo $PIPE
    eval "exec $FD<>\"$PIPE\""
    rm $PIPE
}

generate_test_cases () {
    go run cmd/input_generator.go | head -n 300
}

generate_case_strings () {
    generate_test_cases | while read test_case
    do
        printf "1\n%s\n" "$test_case"
    done
}

FD=3 createPipe
FD=4 createPipe

generate_case_strings | go run cmd/main.go >&3 &
generate_case_strings | go run cmd/main.go --brute >&4 &

generate_test_cases | while read test_case
do
    read opmized_answer <&3
    read brute_answer <&4
    if ! diff <(echo "$opmized_answer") <(echo "$brute_answer") 1>&2
    then
        echo "$test_case" 1>&2
    fi
done