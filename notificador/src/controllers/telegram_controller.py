##!/usr/bin/env python
# -*- coding: utf-8 -*-
#-------------------------------------------------------------------------
# Archivo: telegram_controller.py
# Capitulo: Estilo Microservicios
# Autor(es): Perla Velasco & Yonathan Mtz. & Jorge Solís
# Version: 3.0.0 Febrero 2022
# Descripción:
#
#   Ésta clase define el controlador del microservicio API. 
#   Implementa la funcionalidad y lógica de negocio del Microservicio.
#
#   A continuación se describen los métodos que se implementaron en ésta clase:
#
#                                             Métodos:
#           +------------------------+--------------------------+-----------------------+
#           |         Nombre         |        Parámetros        |        Función        |
#           +------------------------+--------------------------+-----------------------+
#           |     send_message()     |         Ninguno          |  - Procesa el mensaje |
#           |                        |                          |    recibido en la     |
#           |                        |                          |    petición y ejecuta |
#           |                        |                          |    el envío a         |
#           |                        |                          |    Telegram.          |
#           +------------------------+--------------------------+-----------------------+
#
#-------------------------------------------------------------------------
from flask import request, jsonify
import json
import requests
from src.helpers.config import load_config



class TelegramController:

    @staticmethod
    def send_message():
        data = json.loads(request.data)
        if not data:
            return jsonify({"msg": "invalid request"}), 400
        
        config = load_config()

        url = f"https://api.telegram.org/bot{config['TOKEN']}/sendMessage?chat_id={config['CHAT_ID']}&text=Ha%20habido%20un%20cambio%20en%20tu%20poliza"

        response = requests.get(url)
        
        if response.status_code == 200:

            return jsonify({"msg": "message sent"}), 200
        
        print(response.status_code, response.json())
        return jsonify({"msg": str(response.json())}), 500