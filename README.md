![Build status](https://travis-ci.org/joaopedrosgs/LordOfUltima.svg?branch=master)

# Guidelines
	1. Todo request autorizado requer a struct autor, que contém o login e o cookie de sessão.
	2. Toda variável private tem de ter _ na frente do nome.
---
# Entidades

	1. Login Server: Autoriza login
	2. City Server: Construções e cidades
	3. Action Server: Transporte de recursos e ataques
	4. Map Server: Eventos globais e mapa


---
# Login server

- string **authorize**(string _login, string _password, string _ip)
	1. Recebe o login, senha e ip e retorna um autor (ou vazio caso o login não tenha sido bem sucedido).
	2. O autor gerado é guardado no banco de dados, junto com o login.
- bool **isLogged**(autor _autor)
	1. retorna verdadeiro se o autor existe.
- void **logout**(autor _tor)
	1. remove o autor do banco de dados, causando o logout.
---
# City server

- city **getCity**(autor _autor , int _cityid)
	1. Retorna a cidade onde (_cityId == id) na tabela cidades.
	2. Caso o usuário não seja o dono ou a cidade não exista:
		1. Retorna nulo
		2. Adiciona ao log
- bool **buildEnqueue**(autor _autor, int _cityid, build _build)
	1. Retorna verdadeiro se pôde adicionar a construção à fila.
	2. Se a construção já existe, faz upgrade ou downgrade (dependendo do struct _build passado como referencia)
	
- bool **troopEnqueue**(autor _autor, int _cityid, byte _type, uint _quantity)
	1. retorna verdadeiro se pôde adicionar as tropas na fila.

---

# Action Server

- bool **sendResources**(autor _autor, int _cityid, int _target, byte _resourcetype, uint _quantity)
	1. Retorna verdadeiro se pôde enviar os recursos

- []transportation **getCTransportation**(autor _autor, uint _cityid)
	1. Retorna a lista de transporte de recursos naquela cidade

- []transportation **getTransportation**(autor _autor)
	1. Retorna a lista de transporte de recursos no império.

- bool **enqueueMilitary**(autor _autor, int _cityid, int _target, byte _trooptype, uint _quantity,Time _depart, Duration _duration, byte _type)
	1. Retorna verdadeiro se pôde enviar a ação militar.
	2. Tipos de ação:
		1. Ataque direto = 0
		2. Suporte
		3. Cerco
		4. Fundação
		5. Raid

- []military **getCMilitary**(autor _autor, uint _cityid)
	1. Retorna a lista de ações militares naquela cidade

- []military **getMilitary**(autor _autor)
	1. Retorna a lista de ações militares no império.
