# Generated by Django 4.1.3 on 2022-11-26 02:54

from django.db import migrations

import utils.models


class Migration(migrations.Migration):

    dependencies = [
        ("order", "0056_rename_hazmat_id_order_hazmat"),
    ]

    operations = [
        migrations.AlterField(
            model_name="order",
            name="status",
            field=utils.models.ChoiceField(
                choices=[
                    ("N", "New"),
                    ("P", "In Progress"),
                    ("C", "Completed"),
                    ("H", "Hold"),
                    ("B", "Billed"),
                    ("V", "Voided"),
                ],
                default="N",
                max_length=1,
                verbose_name="Status",
            ),
        ),
    ]
