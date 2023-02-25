# Generated by Django 4.1.7 on 2023-02-22 22:37

from django.db import migrations

import utils.models


class Migration(migrations.Migration):
    dependencies = [
        ("accounts", "0005_rename_phone_userprofile_phone_number_and_more"),
    ]

    operations = [
        migrations.AddField(
            model_name="jobtitle",
            name="job_function",
            field=utils.models.ChoiceField(
                choices=[
                    ("MANAGER", "Manager"),
                    ("MANAGEMENT_TRAINEE", "Management Trainee"),
                    ("SUPERVISOR", "Supervisor"),
                    ("DISPATCHER", "Dispatcher"),
                    ("BILLING", "Billing"),
                    ("FINANCE", "Finance"),
                    ("SAFETY", "Safety"),
                    ("SYS_ADMIN", "System Administrator"),
                ],
                default=1,
                help_text="Relevant job function of the job title.",
                max_length=18,
                verbose_name="Job Function",
            ),
            preserve_default=False,
        ),
    ]
