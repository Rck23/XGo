# Proyecto X en Go

Este proyecto es una aplicación AWS Lambda escrita en Go que maneja eventos de API Gateway. Utiliza el SDK de AWS para Go v2 y el framework de AWS Lambda para Go. La aplicación inicializa la configuración de AWS, valida parámetros de entorno, obtiene un secreto de AWS Secrets Manager, se conecta a una base de datos, y procesa la solicitud según el método y la ruta especificados.

## Requisitos

- Go versión 1.22.3
- AWS CLI configurado con acceso a AWS
- Variables de entorno configuradas correctamente(`SecretName`, `BucketName`, `UrlPrefix`)

## Instalación

1. Clona el repositorio:
