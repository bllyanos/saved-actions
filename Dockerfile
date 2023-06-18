FROM alpine:3.18.2

WORKDIR /src/job

RUN apk add curl git

RUN echo "#!/bin/sh \necho 'hello'" > ./script.sh

RUN chmod +x ./script.sh

ENTRYPOINT [ "./script.sh" ]
