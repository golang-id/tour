.PHONY: serve deploy

serve:
	dev_appserver.py app.yaml

deploy:
	gcloud app deploy
