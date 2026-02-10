.PHONY: install
install:
	go build -o ${GOBIN}/golang-id-tour .

.PHONY: serve
serve:
	go run .

.PHONY: deploy
deploy:
	gcloud config configurations activate go-tour-id2
	gcloud --project=go-tour-id2 app deploy --promote app.yaml
