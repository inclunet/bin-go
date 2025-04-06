# IncluGames

IncluGames é uma plataforma de jogos acessíveis e inclusivos criada para oferecer a possibilidade de pessoas com deficiência interagirem com jogos tradicionais como bingo, jogo da forca, Jogo da velha e outros jogos que tradicionalmente são inacessíveis para pessoas com deficiência. A plataforma é inteiramente desenvolvida para o ambiente web e este é o backend que dá suporte para toda a plataforma.

# Desenvolvido em SvelteKit

O Inclugames é inteiramente desenvolvido com SvelteKit que é derivado do Svelte.

<em>"SvelteKit é um framework para desenvolvimento rápido de aplicativos web robustos e de alto desempenho usando Svelte . Se você vem do React, SvelteKit é similar ao Next. Se você vem do Vue, SvelteKit é similar ao Nuxt."</em>

<strong>Documentação do SvelteKit</strong> https://svelte.dev/docs/kit/introduction

<em>"Svelte é um framework de UI que usa um compilador para permitir que você escreva componentes incrivelmente concisos que fazem o mínimo de trabalho no navegador, usando linguagens que você já conhece — HTML, CSS e JavaScript. É uma carta de amor ao desenvolvimento web."</em>

<strong>Documentação do Svelt</strong> https://svelte.dev/

## Preparando seu ambiente de desenvolvimento

Este projeto requer [Go 1.22](https://golang.org/dl/) e [Node.js 20.17.0](https://nodejs.org/en/download/). Siga os passos abaixo para clonar o repositório e rodar o servidor da API:

### Clonando o Repositório

```sh
git clone git@github.com:inclunet/bin-go
cd bin-go
```

### Rodando o Servidor da API

Certifique-se de ter o Go 1.22 instalado. Para rodar o servidor, execute:

```sh
go run main.go --port=8080
```

A API estará disponível em `http://localhost:8080`.

## Preparando o Front-end

A pasta `/bingo` contém o front-end do sistema. Ele já está configurado para usar a API localmente, que deve estar rodando na porta 8080. O front-end é desenvolvido em SvelteKit. Siga os passos abaixo para preparar e iniciar o front-end:

### Instalando Dependências

Certifique-se de ter o Node.js 20.17.0 instalado. Navegue até a pasta `/bingo` e instale as dependências:

```sh
cd bingo
npm install
```

### Rodando o Servidor de Desenvolvimento

Para iniciar o servidor de desenvolvimento do SvelteKit, execute:

```sh
npm run dev
```

Isso iniciará o servidor de desenvolvimento, que estará disponível em `http://localhost:5173`.

### Rodando o Servidor de Desenvolvimento na Rede Local

Para iniciar o servidor de desenvolvimento do SvelteKit para acessá-lo dentro da rede local por outros dispositivos que estejam na mesma rede, execute:

```sh
npm run dev -- --host
```

Para acessá-lo, utilize as seguintes URLs `http://localhost:5173` na máquina local e `http://<IP do server de desenvolvimento>:5173` para acessar de outro dispositivo que esteja na mesma rede.


## Como Contribuir

1. Crie uma branch para sua feature (`git checkout -b feature/nome-da-feature`)
2. Faça commit das suas alterações (`git commit -m 'Adiciona nova feature'`)
3. Faça push para a branch (`git push origin feature/nome-da-feature`)
4. Abra um Pull Request

