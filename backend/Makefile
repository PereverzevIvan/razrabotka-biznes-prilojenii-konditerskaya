# SERVICES := gateway hello tools components orders
SERVICES := gateway tools components productsorders

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

test:
	echo "hello"
