# -*- coding: utf-8 -*-
from django.apps import AppConfig


class WorkerConfig(AppConfig):
    default_auto_field = "django.db.models.BigAutoField"
    name = "worker"

    def ready(self):
        import worker.signals
