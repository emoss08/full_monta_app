# Generated by Django 4.1.5 on 2023-01-21 20:43

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ("order", "0011_alter_order_freight_charge_amount"),
    ]

    operations = [
        migrations.AlterField(
            model_name="order",
            name="mileage",
            field=models.DecimalField(
                blank=True,
                decimal_places=2,
                default=0,
                help_text="Total Mileage",
                max_digits=10,
                null=True,
                verbose_name="Total Mileage",
            ),
        ),
    ]
