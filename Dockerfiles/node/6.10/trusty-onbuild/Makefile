build:
	cp ../../../../tasque ./
	docker build -t tasque/node:$(VERSION)-$(TASQUE_VERSION)-$(VARIANT) \
							 -t tasque/node:$(VERSION)-$(VARIANT) \
							 -t tasque/node:$(VARIANT) \
							  .
	rm tasque

push:
	docker push tasque/node:$(VERSION)-$(TASQUE_VERSION)-$(VARIANT)
	docker push tasque/node:$(VARIANT)
