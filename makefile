buildi:
ifndef n
	@n = <none>
endif
ifndef t
	@t = latest
endif
	@docker build -t $(n):$(t) .


rmi:
ifndef n
	@echo parameter n [name] is required
	@exit 1
endif
ifndef t
	@echo parameter t [tag] is required
	@exit 1
endif
	@docker rmi $(n):$(t)