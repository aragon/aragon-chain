set -e
touch ../coverage.txt
for d in $(go list ./... | grep -v vendor); do
    go test -mod=readonly -timeout 10m -coverprofile=profile.out -covermode=atomic $d
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done