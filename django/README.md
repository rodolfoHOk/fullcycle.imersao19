# Imersão FullCycle 19 - Plataforma de streaming de vídeos

## Descrição

Repositório do microsserviço Django (Admin dos vídeos)

## Tecnologias utilizadas

- Python
- Django
- PostgreSQL
- Docker
- HTML
- CSS
- Javascript

## Rodar a aplicação

### Requerimentos

- Python 3.12.6

### Comandos

Coloque a variável `PIPENV_VENV_IN_PROJECT` no seu `.bashrc` ou `.zshrc`:

```bash
export PIPENV_VENV_IN_PROJECT=1
```

Levante os containers do PostgreSQL e PGAdmin:

```bash
docker-compose up -d
```

Instale o Pipenv:

```bash
pip install pipenv
```

Instale as dependências:

```bash
pipenv install
```

A partir daqui, precisamos sempre rodar os comandos dentro do ambiente virtual do Pipenv:

```bash
pipenv shell
```

Rode as migrações do Django:

```bash
python manage.py migrate
```

Crie um superusuário:

```bash
python manage.py createsuperuser
```

Rode o servidor:

```bash
python manage.py runserver
```

Acesse o admin em [http://localhost:8000/admin]().

## Guia de comandos utilizados

- Fedora SilverBlue: rpm-ostree install python3-pip
- terminal: sudo pip install pipenv
- .zshrc: export PIPENV_VENV_IN_PROJECT=1
- mkdir django
- cd django
- pipenv install django
- pipenv shell
- django-admin startproject videos
- python manage.py runserver
- python manage.py migrate
- python manage.py createsuperuser
- pipenv install psycopg2-binary
- django-admin startapp core
- python manage.py makemigrations

## Readme principal

- [README.md](../README.md)
