SERVICES := gateway tools components orders

up:
	@for service in $(SERVICES); do \
		(cd $$service && make up); \
	done

restart:
	@for service in $(SERVICES); do \
		(cd $$service && make restart); \
	done

down:
	@for service in $(SERVICES); do \
		(cd $$service && make down); \
	done

