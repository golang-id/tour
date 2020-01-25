.PHONY: serve deploy

serve:
	go run .

deploy:
	gcloud --project=go-tour-id2 app deploy --promote app.yaml
