# Generated by Django 4.2 on 2023-04-17 14:03

from django.db import migrations, models


class Migration(migrations.Migration):
    dependencies = [
        ("stops", "0005_alter_stop_appointment_time"),
    ]

    operations = [
        migrations.AlterField(
            model_name="stop",
            name="appointment_time",
            field=models.DateTimeField(
                help_text="The time the equipment is expected to arrive at the stop.",
                verbose_name="Stop Appointment Time",
            ),
        ),
    ]
