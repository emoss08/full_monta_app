# Generated by Django 4.2 on 2023-04-24 02:52

from django.db import migrations, models


class Migration(migrations.Migration):
    dependencies = [
        ("integration", "0005_delete_googleapilog"),
    ]

    operations = [
        migrations.AddField(
            model_name="googleapi",
            name="auto_geocode",
            field=models.BooleanField(
                default=False,
                help_text="This determines if locations will automatically be geocoded, once they are created/updated.",
                verbose_name="Auto Geocode",
            ),
        ),
    ]
