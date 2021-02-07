# Desafio 1 - Ambiente de desenvolvimento e validação de models

## Informações do desafio

- Nesse desafio você deverá montar seu ambiente de desenvolvimento com Docker utilizando o Dockerfile e docker-compose de nosso projeto prático.


- Crie uma struct de User com ID (uuid), Name e Email. Além disso, implemente uma função NewUser que seja capaz de validar que ID, Name e Email são obrigatórios, caso o contrário, ela deve retornar um erro.

- Suba seu desafio em um repositório Github chamado de "imersao-fullstack-fullcycle" e nos envie através do campo abaixo.

# Desafio 2 Consumo e publicação de mensagens no Kafka

## Informações do desafio

- Nesse desafio você deverá criar uma simples aplicação utilizando Golang que seja capaz de publicar uma mensagem em um tópico do Apache Kafka e também consumir mensagens desse tópico.
- Para deixar mais claro a separação das responsabilidades, crie uma pasta chamada producer e outra consumer. Em cada uma delas crie um arquivo main.go que será responsável por produzir e consumir as mensagens respectivamente
- Utilize o mesmo Dockerfile e docker-compose.yaml utilizados no projeto do CodePix. Fique na liberdade para escolher o nome do tópico que será utilizado
- Realize o commit em seu repositório do Github e entre com o endereço nos campos abaixo