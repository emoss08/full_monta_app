# Generated by Django 4.1.7 on 2023-02-18 16:43

from django.db import migrations


class Migration(migrations.Migration):

    dependencies = [
        ("equipment", "0002_alter_equipment_options_and_more"),
    ]

    operations = [
        migrations.AlterModelTable(
            name="equipment",
            table="equipment",
        ),
        migrations.AlterModelTable(
            name="equipmentmaintenanceplan",
            table="equipment_maintenance_plan",
        ),
        migrations.AlterModelTable(
            name="equipmentmanufacturer",
            table="equipment_manufacturer",
        ),
        migrations.AlterModelTable(
            name="equipmenttype",
            table="equipment_type",
        ),
        migrations.AlterModelTable(
            name="equipmenttypedetail",
            table="equipment_type_detail",
        ),
    ]
