# Generated by Django 4.1.5 on 2023-01-16 03:43

import django.core.validators
from django.db import migrations, models
import order.models


class Migration(migrations.Migration):

    dependencies = [
        ("order", "0006_alter_ordercomment_options_and_more"),
    ]

    operations = [
        migrations.AlterField(
            model_name="orderdocumentation",
            name="document",
            field=models.FileField(
                default=1,
                upload_to=order.models.order_documentation_upload_to,
                validators=[
                    django.core.validators.FileExtensionValidator(
                        allowed_extensions=["pdf"]
                    )
                ],
                verbose_name="Document",
            ),
            preserve_default=False,
        ),
    ]
