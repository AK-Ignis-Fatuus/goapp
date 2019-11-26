FROM golang:latest

# Copy the local package files to the container’s workspace.
COPY main /root/main

RUN go get github.com/gorilla/mux \
    && useradd -mr godo-user \
    && chmod +x /root/main \
    && mkdir -p /home/godo-user/go_do \
    && mv /root/main /home/godo-user/go_do \
    && chown -R godo-user:godo-user /home/godo-user/go_do \
    && touch /root/start.sh \ 
    && echo "#!/bin/sh" >> /root/start.sh \
    && echo "runuser -l godo-user -c '/home/godo-user/go_do/main'" >> /root/start.sh \
    && chmod ugo+x /root/start.sh

EXPOSE 8000

ENTRYPOINT ["/root/start.sh"]
