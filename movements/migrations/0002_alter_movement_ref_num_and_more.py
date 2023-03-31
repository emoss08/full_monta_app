# Generated by Django 4.1.7 on 2023-03-11 09:23

from django.db import migrations, models


class Migration(migrations.Migration):
    dependencies = [
        ("movements", "0001_initial"),
    ]

    operations = [
        migrations.AlterField(
            model_name="movement",
            name="ref_num",
            field=models.CharField(
                editable=False,
                help_text="Movement Reference Number",
                max_length=10,
                verbose_name="Movement Reference Number",
            ),
        ),
        migrations.AddConstraint(
            model_name="movement",
            constraint=models.UniqueConstraint(
                fields=("ref_num", "organization"),
                name="unique_movement_ref_num_organization",
            ),
        ),
    ]
