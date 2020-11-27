# API de Receitas
API em GoLang para busca de receitas a partir de ingredientes. 
A busca é feita na API [RecipePuppy](http://www.recipepuppy.com/about/api/) e para cada receita
um há um GIF, vindo da API do [Giphy](https://developers.giphy.com/docs/).

## Execução da API
Para executar a API siga os seguintes passos:
1) Clone o repositório na sua máquina `git clone git@github.com:matheusmhmelo/api-recipe.git`
2) Acesse a pasta raiz do projeto
3) Execute o comando `make start` para inicializar todos os serviços necessários 
4) Já é possível acessar as rotas da API a partir do endereço `127.0.0.1:8082`
5) Para finalizar a execução execute o comando `make stop`

## Rotas

Rota responsável pela busca da receita: \
`GET - /recipes/?i={ingredient_1},{ingredient_2},{ingredient_3}` 

**Resposta:**
```{ 
   "keywords": [
       "onions",
       "garlic"
   ],
   "recipes": [
       {
           "title": "Crock Pot Caramelized Onions",
           "ingredients": [
               "beef broth",
               "butter",
               "garlic",
               "onions"
           ],
           "link": "http://www.recipezaar.com/Crock-Pot-Caramelized-Onions-102934",
           "gif": "https://giphy.com/gifs/crock-pot-7IYPkpk7fiXaRt4yEp"
       }
   ]
 }
```

## Configurações
As configurações necessárias para a execução da API são:
 - Porta de execução (padrão: 8080)
 - Rota para requisição da API RecipePuppy
 - Rota para requisição da API Giphy
 - Chave de acesso para a API Giphy
 - Configurações de Redis (cache)
 
Os arquivos de configurações são `config.yaml` (configurações para execução local) e 
`docker-config.yaml` (configurações para execução no Docker).

Ao alterar qualquer informação do arquivo `docker-config.yaml` é necessário atualizar o 
serviço executando o comando `docker service update dm_recipe --force`.

## Infraestrutura
### Docker
O projeto é executado em Docker Swarm e os serviços necessários para execução são:
- API (dm_recipe)
- Redis (dm_recipe_redis)

### Cache
A API possui um sistema de cache para as requisições ao Giphy, tal requisição retorna um
GIF a partir do título da receita, visando um desempenho melhor, o resultado da busca por esse
título é salvo num Cache para futuras requisições.

### Makefile
Para facilitar a execução algumas ações estão disponíveis em Scripts:
- `make start`: Inicia os serviços do Docker
- `make stop`: Finaliza os serviços do Docker
- `make build tag={tag}`: Atualiza a imagem da API salva no Docker de acordo com a `tag`informada

O último comando deve ser executado apenas quando houver alterações no código e tais alterações 
devam ser aplicadas na imagem do docker.

### Testes Unitários
A API possui testes unitários para três pacotes:
- `format`: Responsável pela formatação dos ingredientes recebidos na requisição.
- `recipepuppy`: Responsável pela busca das receitas na API do RecipePuppy.
- `giphy`: Responsável por buscar o GIF na API do Giphy

Para testar o projeto basta executar o comando `go test ./...`. \
Para testar apenas um pacote específico execute o comando `go test ./internal/services/{packageName}/...`,
onde `packageName` pode ser **format**, **recipepuppy** ou **giphy**.


