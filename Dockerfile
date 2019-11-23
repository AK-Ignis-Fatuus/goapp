### test ###

FROM scratch

# Copy the local package files to the container’s workspace.
ADD app.ini goapp /

EXPOSE 9200

CMD ["/goapp"]
