# -*- coding: utf-8 -*-
# Generated by Django 4.1.2 on 2022-11-02 16:17

import django.db.models.deletion
import django_extensions.db.fields
from django.db import migrations, models


class Migration(migrations.Migration):

    initial = True

    dependencies = [
        ("organization", "0006_alter_depotdetail_address_line_1_and_more"),
    ]

    operations = [
        migrations.CreateModel(
            name="DispatchControl",
            fields=[
                (
                    "id",
                    models.BigAutoField(
                        auto_created=True,
                        primary_key=True,
                        serialize=False,
                        verbose_name="ID",
                    ),
                ),
                (
                    "created",
                    django_extensions.db.fields.CreationDateTimeField(
                        auto_now_add=True, verbose_name="created"
                    ),
                ),
                (
                    "modified",
                    django_extensions.db.fields.ModificationDateTimeField(
                        auto_now=True, verbose_name="modified"
                    ),
                ),
                (
                    "record_service_incident",
                    models.CharField(
                        choices=[
                            ("Never", "Never"),
                            ("Pickup", "Pickup"),
                            ("Delivery", "Delivery"),
                            ("Pickup and Delivery", "Pickup and Delivery"),
                            ("All except shipper", "All except shipper"),
                        ],
                        default="Never",
                        max_length=19,
                        verbose_name="Record Service Incident",
                    ),
                ),
                (
                    "grace_period",
                    models.PositiveIntegerField(
                        default=0,
                        help_text="Grace period for the service incident in minutes.",
                        verbose_name="Grace Period",
                    ),
                ),
                (
                    "deadhead_target",
                    models.DecimalField(
                        decimal_places=2,
                        default=0.0,
                        help_text="Deadhead Mileage target for the company.",
                        max_digits=5,
                        verbose_name="Deadhead Target",
                    ),
                ),
                (
                    "driver_assign",
                    models.BooleanField(
                        default=False,
                        help_text="Enforce driver assign to orders for the company.",
                        verbose_name="Enforce Driver Assign",
                    ),
                ),
                (
                    "trailer_continuity",
                    models.BooleanField(
                        default=False,
                        help_text="Enforce trailer continuity for the company.",
                        verbose_name="Enforce Trailer Continuity",
                    ),
                ),
                (
                    "distance_method",
                    models.CharField(
                        choices=[("Google", "Google"), ("Monta", "Monta")],
                        default="Monta",
                        help_text="Distance method for the company.",
                        max_length=20,
                        verbose_name="Distance Method",
                    ),
                ),
                (
                    "dupe_trailer_check",
                    models.BooleanField(
                        default=False,
                        help_text="Enforce duplicate trailer check for the company.",
                        verbose_name="Enforce Duplicate Trailer Check",
                    ),
                ),
                (
                    "regulatory_check",
                    models.BooleanField(
                        default=False,
                        help_text="Enforce regulatory check for the company.",
                        verbose_name="Enforce Regulatory Check",
                    ),
                ),
                (
                    "prev_orders_on_hold",
                    models.BooleanField(
                        default=False,
                        help_text="Prevent dispatch of orders on hold for the company.",
                        verbose_name="Prevent Orders On Hold",
                    ),
                ),
                (
                    "organization",
                    models.OneToOneField(
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="dispatch_control",
                        related_query_name="dispatch_controls",
                        to="organization.organization",
                        verbose_name="Organization",
                    ),
                ),
            ],
            options={
                "verbose_name": "Dispatch Control",
                "verbose_name_plural": "Dispatch Controls",
                "ordering": ["organization"],
            },
        ),
    ]
