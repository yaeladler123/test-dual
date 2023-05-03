

#####################
# File in progress
#####################


FROM alpine:latest AS certs

# HACK to allow copy to work even if the certs directory is missing, we have to copy at least one file, using Dockerfile
# also we use a character range on the certs diretcory name so we are guaranteed a match
COPY Dockerfile cred[s] ./
# need to rename the cert from .pem to .crt
RUN if [ -f CalculiCA-cert.pem ] ; then mv -v CalculiCA-cert.pem CalculiCA-cert.crt ; fi
# now create an empty file we won't mind copying so we can repeat the trick
RUN touch .empty-file

FROM alpine:latest AS image
# HACK to allow safe multi-arch binaries
ARG binaryDir=.

LABEL maintainer=Calculi \
    email=engineering@guide-rails.io

RUN apk -f -q update && apk -f -q --no-cache add ca-certificates
COPY --from=certs .empty-file CalculiCA-cert.cr[t] /usr/local/share/ca-certificates/
RUN update-ca-certificates && rm /usr/local/share/ca-certificates/.empty-file

COPY ${binaryDir}/build/potato-service /usr/local/bin/potato-service

RUN set -e ; \
    adduser -G root -S guiderails -u 100 ; \
    mkdir -p /var/vcap/sys/log ; \
    chmod -R g+rwX /var/vcap/sys/log ; \
    chgrp -R root /var/vcap/sys/log ;
USER 100
ENV PATH="${PATH}:/usr/local/bin"

ENTRYPOINT /usr/local/bin/potato-service