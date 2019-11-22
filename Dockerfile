### test ###

FROM scratch

# Copy the local package files to the container’s workspace.
ADD app.ini godo-app /

EXPOSE 9200

CMD ["/godo-app"]
