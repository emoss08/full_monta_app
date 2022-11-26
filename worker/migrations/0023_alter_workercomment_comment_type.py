# Generated by Django 4.1.3 on 2022-11-25 00:50

import django.db.models.deletion
from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ("dispatch", "0008_commenttype"),
        ("worker", "0022_alter_worker_worker_type_and_more"),
    ]

    operations = [
        migrations.AlterField(
            model_name="workercomment",
            name="comment_type",
            field=models.ForeignKey(
                help_text="Related comment type.",
                on_delete=django.db.models.deletion.CASCADE,
                related_name="comments",
                related_query_name="comments",
                to="dispatch.commenttype",
                verbose_name="Comment Type",
            ),
        ),
    ]
