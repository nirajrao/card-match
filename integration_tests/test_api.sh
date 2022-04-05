#!/bin/bash

go install card-match/cmd/api

$(api --port 8080) &

# Maintain an $ATTEMPTS variable in case the server takes longer to serve.
ATTEMPTS=0
check=$(curl -s -w "%{http_code}\n" -L "localhost:8080/ping" -o /dev/null)
until [ $ATTEMPTS -eq 100 ]
do
        if [[ $check == 200 || $check == 403 ]]
        then
                echo "/ping endpoint responded successfully."
                exit 0
        fi
        check=$(curl -s -w "%{http_code}\n" -L "localhost:8080/ping" -o /dev/null)
        ATTEMPTS=$(( ATTEMPTS + 1))
        echo $ATTEMPTS
done

# We tried 100 times and didn't receive a successful response.
# Exit with an error code.
echo "/ping endpoint did not respond successfully."
exit 1
