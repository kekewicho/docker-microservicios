# Docker Microservicios

<!-- [TODO] agregar descripción --> 

## 🛑 Información importante

Este repositorio forma parte de una practica de la materia Arquitectura de Software de Ingenieria de Software UAZ. El codigo fue tomado del (Repositorio original)[https://gitlab.com/tareas-arquitectura-de-software-curso/microservicios/docker-microservicios.git].

Ademas de seguir las instrucciones, se debe asegurar contar con lo siguiente:

- En el directorio del notificador se debera agregar un archivo .env con los siguientes datos:
    - TOKEN
    - CHAT_ID

## Estructura del proyecto

Este repositorio contiene los siguientes directorios y archivos:

```bash
    ├── cliente                              # componente GUI que funciona como cliente
    ├── gestor-de-clientes                   # microservicio gestor de clientes
    ├── docs                                 # carpeta de documentación
    │  ├── context-view.png                  # vista del contexto del sistema
    │  ├── global-assurance.drawio           # archivo editable de daiagramas del sistema 
    ├── notificador                          # microservicio notificador 
    ├── pagos                                # microservicio pagos
    ├── reporteador                          # microservicio reporteador
    ├── simulador                            # componente que simula los pagos realizados
    ├── tyk                                  # archivos compartidos con el gateway
    │  ├── apps                              # definición de APIs en el gateway
    │  |  ├── keyless-gestor-clientes.json   # definición de microservicio API
    │  |  ├── keyless-notificador.json       # definición de microservicio Notifier
    │  |  ├── keyless-pagos.json             # definición de microservicio Payment
    │  |  ├── keyless-reporteador.json       # definición de microservicio Reporter
    │  ├── tyk.standalone.conf               # archivo de configuración de tyk
    ├── .gitignore                           # exclusiones de git
    ├── docker-compose.yml                   # definición de contenedores para ambiente docker
    ├── README.md                            # este archivo
```

## Instalación

Descarga el código del repositorio utilizando el siguiente comando:

`git clone https://github.com/kekewicho/docker-microservicios.git`

Ingresa a la carpeta del proyecto:

`cd docker-microservicios`

## Ejecución

Para ejecutar el sistema utiliza el siguiente comando:

`docker-compose up -d`

Para detener el sistema utiliza el siguiente comando:

`docker-compose down`

> al ejecutar el sistema se creará una carpeta `volume` donde se guardará la información que se genere en el sistema
> si deseas reiniciar los datos basta con eliminar dicha carpeta

## Preguntas Frecuentes

### ¿Necesito instalar Docker?

Por supuesto, la herramienta Docker es vital para la ejecución de este sistema. Para conocer más acerca de Docker puedes visitar el siguiente [enlace](https://medium.com/@javiervivanco/que-es-docker-79d506f7b2fc).

> Para realizar la instalación de Docker en Windows puedes consultar el siguiente [enlace](https://medium.com/@tushar0618/installing-docker-desktop-on-window-10-501e594fc5eb)


> Para realizar la instalación de Docker en Linux puedes consultar el siguiente [enlace](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-on-ubuntu-20-04-es)

## Versión

3.0.0 - Febrero 2022

## Autores

- Perla Velasco
- Yonathan Martinez
- Jorge Solis