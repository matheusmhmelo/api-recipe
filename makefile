build:
	bash scripts/build.sh $(tag)

start:
	docker stack deploy -c docker-compose.yaml dm

stop:
	docker service rm dm_recipe dm_recipe_redis
