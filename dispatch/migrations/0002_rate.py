# Generated by Django 4.1.6 on 2023-02-09 20:48

import uuid

import django.db.models.deletion
import django.utils.timezone
import django_extensions.db.fields
import django_lifecycle.mixins
from django.db import migrations, models


class Migration(migrations.Migration):
    dependencies = [
        ("organization", "0009_organization_address_line_1_and_more"),
        ("commodities", "0003_alter_hazardousmaterial_name"),
        ("equipment", "0006_alter_equipment_vin_number"),
        ("customer", "0010_alter_customerbillingprofile_email_profile_and_more"),
        ("order", "0017_remove_ordercontrol_enforce_customer_and_more"),
        ("dispatch", "0001_initial"),
    ]

    operations = [
        migrations.CreateModel(
            name="Rate",
            fields=[
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
                    "id",
                    models.UUIDField(
                        default=uuid.uuid4,
                        editable=False,
                        primary_key=True,
                        serialize=False,
                        unique=True,
                    ),
                ),
                (
                    "rate_number",
                    models.CharField(
                        editable=False,
                        help_text="Rate Number for Rate",
                        max_length=10,
                        unique=True,
                        verbose_name="Rate Number",
                    ),
                ),
                (
                    "effective_date",
                    models.DateField(
                        default=django.utils.timezone.now,
                        help_text="Effective Date for Rate",
                        verbose_name="Effective Date",
                    ),
                ),
                (
                    "expiration_date",
                    models.DateField(
                        default=django.utils.timezone.now,
                        help_text="Expiration Date for Rate",
                        verbose_name="Expiration Date",
                    ),
                ),
                (
                    "commodity",
                    models.ForeignKey(
                        blank=True,
                        help_text="Commodity for Rate",
                        null=True,
                        on_delete=django.db.models.deletion.SET_NULL,
                        related_name="rates",
                        to="commodities.commodity",
                        verbose_name="Commodity",
                    ),
                ),
                (
                    "customer",
                    models.ForeignKey(
                        blank=True,
                        help_text="Customer for Rate",
                        null=True,
                        on_delete=django.db.models.deletion.SET_NULL,
                        related_name="rates",
                        to="customer.customer",
                        verbose_name="Customer",
                    ),
                ),
                (
                    "equipment_type",
                    models.ForeignKey(
                        blank=True,
                        null=True,
                        on_delete=django.db.models.deletion.SET_NULL,
                        related_name="rates",
                        to="equipment.equipmenttype",
                        verbose_name="Equipment Type",
                    ),
                ),
                (
                    "order_type",
                    models.ForeignKey(
                        blank=True,
                        null=True,
                        on_delete=django.db.models.deletion.SET_NULL,
                        related_name="rates",
                        to="order.ordertype",
                        verbose_name="Order Type",
                    ),
                ),
                (
                    "organization",
                    models.ForeignKey(
                        help_text="Organization",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="%(class)ss",
                        related_query_name="%(class)s",
                        to="organization.organization",
                        verbose_name="Organization",
                    ),
                ),
            ],
            options={
                "verbose_name": "Rate",
                "verbose_name_plural": "Rates",
                "ordering": ["rate_number"],
            },
            bases=(django_lifecycle.mixins.LifecycleModelMixin, models.Model),
        ),
    ]
