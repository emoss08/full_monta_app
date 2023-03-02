# Generated by Django 4.1.7 on 2023-03-01 21:09

from django.conf import settings
from django.db import migrations, models
import django.db.models.deletion


class Migration(migrations.Migration):

    initial = True

    dependencies = [
        ("dispatch", "0001_initial"),
        ("equipment", "0001_initial"),
        ("location", "0001_initial"),
        ("order", "0001_initial"),
        ("customer", "0002_initial"),
        migrations.swappable_dependency(settings.AUTH_USER_MODEL),
        ("commodities", "0001_initial"),
        ("organization", "0001_initial"),
        ("billing", "0002_initial"),
    ]

    operations = [
        migrations.AddField(
            model_name="ratetable",
            name="destination_location",
            field=models.ForeignKey(
                blank=True,
                help_text="Destination Location for Rate Table",
                null=True,
                on_delete=django.db.models.deletion.PROTECT,
                related_name="destination_rate_tables",
                to="location.location",
                verbose_name="Destination Location",
            ),
        ),
        migrations.AddField(
            model_name="ratetable",
            name="organization",
            field=models.ForeignKey(
                help_text="Organization",
                on_delete=django.db.models.deletion.CASCADE,
                related_name="%(class)ss",
                related_query_name="%(class)s",
                to="organization.organization",
                verbose_name="Organization",
            ),
        ),
        migrations.AddField(
            model_name="ratetable",
            name="origin_location",
            field=models.ForeignKey(
                blank=True,
                help_text="Origin Location for Rate Table",
                null=True,
                on_delete=django.db.models.deletion.PROTECT,
                related_name="origin_rate_tables",
                to="location.location",
                verbose_name="Origin Location",
            ),
        ),
        migrations.AddField(
            model_name="ratetable",
            name="rate",
            field=models.ForeignKey(
                help_text="Rate for Rate Table",
                on_delete=django.db.models.deletion.PROTECT,
                related_name="rate_tables",
                to="dispatch.rate",
                verbose_name="Rate",
            ),
        ),
        migrations.AddField(
            model_name="ratebillingtable",
            name="charge_code",
            field=models.ForeignKey(
                help_text="Charge Code for Rate Billing Table",
                on_delete=django.db.models.deletion.PROTECT,
                related_name="rate_billing_tables",
                to="billing.accessorialcharge",
                verbose_name="Charge Code",
            ),
        ),
        migrations.AddField(
            model_name="ratebillingtable",
            name="organization",
            field=models.ForeignKey(
                help_text="Organization",
                on_delete=django.db.models.deletion.CASCADE,
                related_name="%(class)ss",
                related_query_name="%(class)s",
                to="organization.organization",
                verbose_name="Organization",
            ),
        ),
        migrations.AddField(
            model_name="ratebillingtable",
            name="rate",
            field=models.ForeignKey(
                help_text="Rate for Rate Billing Table",
                on_delete=django.db.models.deletion.PROTECT,
                related_name="rate_billing_tables",
                to="dispatch.rate",
                verbose_name="Rate",
            ),
        ),
        migrations.AddField(
            model_name="rate",
            name="commodity",
            field=models.ForeignKey(
                blank=True,
                help_text="Commodity for Rate",
                null=True,
                on_delete=django.db.models.deletion.SET_NULL,
                related_name="rates",
                to="commodities.commodity",
                verbose_name="Commodity",
            ),
        ),
        migrations.AddField(
            model_name="rate",
            name="customer",
            field=models.ForeignKey(
                blank=True,
                help_text="Customer for Rate",
                null=True,
                on_delete=django.db.models.deletion.SET_NULL,
                related_name="rates",
                to="customer.customer",
                verbose_name="Customer",
            ),
        ),
        migrations.AddField(
            model_name="rate",
            name="equipment_type",
            field=models.ForeignKey(
                blank=True,
                null=True,
                on_delete=django.db.models.deletion.SET_NULL,
                related_name="rates",
                to="equipment.equipmenttype",
                verbose_name="Equipment Type",
            ),
        ),
        migrations.AddField(
            model_name="rate",
            name="order_type",
            field=models.ForeignKey(
                blank=True,
                null=True,
                on_delete=django.db.models.deletion.SET_NULL,
                related_name="rates",
                to="order.ordertype",
                verbose_name="Order Type",
            ),
        ),
        migrations.AddField(
            model_name="rate",
            name="organization",
            field=models.ForeignKey(
                help_text="Organization",
                on_delete=django.db.models.deletion.CASCADE,
                related_name="%(class)ss",
                related_query_name="%(class)s",
                to="organization.organization",
                verbose_name="Organization",
            ),
        ),
        migrations.AddField(
            model_name="fleetcode",
            name="manager",
            field=models.ForeignKey(
                blank=True,
                help_text="Manager for the fleet code.",
                null=True,
                on_delete=django.db.models.deletion.CASCADE,
                related_name="fleet_code_manager",
                to=settings.AUTH_USER_MODEL,
            ),
        ),
        migrations.AddField(
            model_name="fleetcode",
            name="organization",
            field=models.ForeignKey(
                help_text="Organization",
                on_delete=django.db.models.deletion.CASCADE,
                related_name="%(class)ss",
                related_query_name="%(class)s",
                to="organization.organization",
                verbose_name="Organization",
            ),
        ),
        migrations.AddField(
            model_name="dispatchcontrol",
            name="organization",
            field=models.OneToOneField(
                on_delete=django.db.models.deletion.CASCADE,
                related_name="dispatch_control",
                related_query_name="dispatch_controls",
                to="organization.organization",
                verbose_name="Organization",
            ),
        ),
        migrations.AddField(
            model_name="delaycode",
            name="organization",
            field=models.ForeignKey(
                help_text="Organization",
                on_delete=django.db.models.deletion.CASCADE,
                related_name="%(class)ss",
                related_query_name="%(class)s",
                to="organization.organization",
                verbose_name="Organization",
            ),
        ),
        migrations.AddField(
            model_name="commenttype",
            name="organization",
            field=models.ForeignKey(
                help_text="Organization",
                on_delete=django.db.models.deletion.CASCADE,
                related_name="%(class)ss",
                related_query_name="%(class)s",
                to="organization.organization",
                verbose_name="Organization",
            ),
        ),
    ]
