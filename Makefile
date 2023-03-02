.PHONY: serve deploy

serve:
	go run .

deploy:
	gcloud config configurations activate personal
	gcloud --project=go-tour-id2 app deploy --promote app.yaml
