FROM lightningnetwork/golang-alpine:latest

COPY . /go/src/github.com/evgeniy-scherbina/hashicorp_infrastructure

WORKDIR /go/src/github.com/evgeniy-scherbina/hashicorp_infrastructure
RUN go install ./services/addition/cmd/addition-service/
RUN go install ./services/subtraction/cmd/subtraction-service/
RUN go install ./services/multiplication/cmd/multiplication-service/
RUN go install ./services/division/cmd/division-service/

RUN go install ./cmd/...