FROM library/golang

# Godep for vendoring
RUN go get github.com/golang/dep/cmd/dep

# Recompile the standard library without CGO
RUN CGO_ENABLED=0 go install -a std

ENV APP_DIR $GOPATH/src/github.com/kanhaiya15/gopf
ENV APP_ENV development
RUN mkdir -p $APP_DIR

# Set the entrypoint
ENTRYPOINT (cd $APP_DIR && ./gopf)
ADD . $APP_DIR

# Compile the binary and statically link
RUN cd $APP_DIR && dep ensure && CGO_ENABLED=0 go build -ldflags '-d -w -s'

EXPOSE 9000