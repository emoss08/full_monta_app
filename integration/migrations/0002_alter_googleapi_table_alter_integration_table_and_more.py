# Generated by Django 4.1.7 on 2023-02-18 16:43

from django.db import migrations


class Migration(migrations.Migration):

    dependencies = [
        ("integration", "0001_initial"),
    ]

    operations = [
        migrations.AlterModelTable(
            name="googleapi",
            table="google_api",
        ),
        migrations.AlterModelTable(
            name="integration",
            table="integration",
        ),
        migrations.AlterModelTable(
            name="integrationvendor",
            table="integration_vendor",
        ),
    ]
