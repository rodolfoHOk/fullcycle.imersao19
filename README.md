# Imersão FullStack & FullCycle 19 - FCTube

> Evento realizado pela FullCycle / CodeEdu

## Tecnologias

- Python
- Django Framework
- Golang
- RabbitMQ
- PostgreSQL
- Nginx
- Typescript / Javascript
- Next.js / React.js
- Tailwind CSS
- Docker e Docker Compose

## Sobre o Projeto

![Sobre o projeto 01](/files/ifc19-01.png)

![Sobre o projeto 02](/files/ifc19-02.png)

![Sobre o projeto 03](/files/ifc19-04.png)

![Sobre o projeto 04](/files/ifc19-05.png)

![Sobre o projeto 05](/files/ifc19-08.png)

![Sobre o projeto 06](/files/ifc19-09.png)

![Sobre o projeto 07](/files/ifc19-10.png)

## Imagens

<img src="https://raw.githubusercontent.com/rodolfoHOk/portfolio-img/main/images/fullcycle-fctube-01.png" alt="FC Tube 01" height="400" />

<img src="https://raw.githubusercontent.com/rodolfoHOk/portfolio-img/main/images/fullcycle-fctube-02.png" alt="FC Tube 02" height="400" />

<img src="https://raw.githubusercontent.com/rodolfoHOk/portfolio-img/main/images/fullcycle-fctube-03.png" alt="FC Tube 03" height="400" />

## Repositórios

- [Microsserviço Django - Admin dos vídeos](/django/README.md)
- [Microsserviço Golang - Conversor de vídeos](/go_transcoder/README.md)
- [Microsserviço Nextjs - Frontend](/nextjs/README.md)

## Rodar

### Requisitos

- Docker
- Docker Compose

### Comandos

- terminal 1: docker compose up
- terminal 2: docker compose exec go_app_dev bash
- terminal 2: go run cmd/videoconverter/main.go
- terminal 3: docker compose exec django bash
- terminal 3: pipenv shell
- terminal 3: python manage.py migrate
- terminal 3: python manage.py createsuperuser
- terminal 3: python manage.py loaddata initial_data
- terminal 3: ./django/copy-media.sh
- terminal 3: python manage.py runserver 0.0.0.0:8000
- terminal 4: docker compose exec django bash
- terminal 4: pipenv shell
- terminal 4: python manage.py consumer_upload_chunks_to_external_storage
- terminal 5: docker compose exec django bash
- terminal 5: pipenv shell
- terminal 5: python manage.py consumer_register_processed_video_path
- terminal 6: docker compose exec nextjs bash
- terminal 6: npm run dev

### Acessos

- web page: http://localhost:3000
- admin: http://localhost:8000/admin
