# --------------------------------------------------------------------------------------------------
#  COPYRIGHT(c) 2024 Trenova                                                                       -
#                                                                                                  -
#  This file is part of Trenova.                                                                   -
#                                                                                                  -
#  The Monta software is licensed under the Business Source License 1.1. You are granted the right -
#  to copy, modify, and redistribute the software, but only for non-production use or with a total -
#  of less than three server instances. Starting from the Change Date (November 16, 2026), the     -
#  software will be made available under version 2 or later of the GNU General Public License.     -
#  If you use the software in violation of this license, your rights under the license will be     -
#  terminated automatically. The software is provided "as is," and the Licensor disclaims all      -
#  warranties and conditions. If you use this license's text or the "Business Source License" name -
#  and trademark, you must comply with the Licensor's covenants, which include specifying the      -
#  Change License as the GPL Version 2.0 or a compatible license, specifying an Additional Use     -
#  Grant, and not modifying the license in any other way.                                          -
# --------------------------------------------------------------------------------------------------

from __future__ import annotations

import textwrap
import typing
import uuid

from django.conf import settings
from django.core.exceptions import ValidationError
from django.db import models
from django.db.models.functions import Lower
from django.urls import reverse
from django.utils.translation import gettext_lazy as _

from organization.models import BusinessUnit, Organization
from shipment.models import Shipment
from utils.models import ChoiceField, GenericModel, PrimaryStatusChoices, StatusChoices
from utils.types import ModelDelete


@typing.final
class FuelMethodChoices(models.TextChoices):
    """
    A class representing the possible fuel method choices.

    This class inherits from the `models.TextChoices` class and defines three constants:
    - DISTANCE: representing a fuel method based on distance
    - FLAT: representing a flat rate fuel method
    - PERCENTAGE: representing a fuel method based on a percentage
    """

    DISTANCE = "D", _("Distance")
    FLAT = "F", _("Flat")
    PERCENTAGE = "P", _("Percentage")


@typing.final
class BillingExceptionChoices(models.TextChoices):
    """
    A class representing the possible billing exception choices.

    This class inherits from the `models.TextChoices` class and defines five constants:
    - PAPERWORK: representing a billing exception related to paperwork
    - CHARGE: representing a billing exception resulting in a charge
    - CREDIT: representing a billing exception resulting in a credit
    - DEBIT: representing a billing exception resulting in a debit
    - OTHER: representing any other type of billing exception
    """

    PAPERWORK = "PAPERWORK", _("Paperwork")
    CHARGE = "CHARGE", _("Charge")
    CREDIT = "CREDIT", _("Credit")
    DEBIT = "DEBIT", _("Debit")
    OTHER = "OTHER", _("OTHER")


class BillingControl(GenericModel):
    """Stores the billing control information for a related :model: `organization.Organization`

    The BillingControl model stores the billing control information for a related organization.
    It is used to store information such as whether to auto-bill invoices, or if users can or
    cannot delete records from billing history and more.
    """

    @typing.final
    class AutoBillingCriteriaChoices(models.TextChoices):
        """
        A class representing the possible auto billing choices.

        This class inherits from the `models.TextChoices` class and defines three constants:
        - Shipment_DELIVERY: representing a criteria stating to auto bill Shipments on delivery.
        - TRANSFERRED_TO_BILL: representing a criteria stating to auto bill Shipment when
        Shipments are transferred to bill queue.
        - CREDIT: representing a criteria stating to auto bill Shipment when Shipments are
        marked ready to bill in the billing queue.
        """

        SHIPMENT_DELIVERY = "SHIPMENT_DELIVERY", _(
            "Auto Bill when shipment is delivered"
        )
        TRANSFERRED_TO_BILL = "TRANSFERRED_TO_BILL", _(
            "Auto Bill when Shipment are transferred to billing"
        )
        MARKED_READY_TO_BILL = "MARKED_READY", _(
            "Auto Bill when Shipment is marked ready to bill in Billing Queue"
        )

    @typing.final
    class ShipmentTransferCriteriaChoices(models.TextChoices):
        """
        A class representing the possible Shipment transfer choices.

        This class inherits from the `models.TextChoices` class and defines three constants:
        - READY_AND_COMPLETED: representing a criteria stating the Shipment must be `ready_to_bill`
        & `completed` before it can be transferred to billing.
        - COMPLETED: representing a criteria stating the Shipment must be `completed` before it can
        be transferred to billing.
        - READY_TO_BILL: representing a criteria stating the Shipment must be `ready_to_bill` before it
        can be transferred to billing.
        """

        READY_AND_COMPLETED = "READY_AND_COMPLETED", _("Ready to bill & Completed")
        COMPLETED = "COMPLETED", _("Completed")
        READY_TO_BILL = "READY_TO_BILL", _("Ready to bill")

    id = models.UUIDField(
        primary_key=True,
        default=uuid.uuid4,
        editable=False,
        unique=True,
    )
    organization = models.OneToOneField(
        "organization.Organization",
        on_delete=models.CASCADE,
        verbose_name=_("Organization"),
        related_name="billing_control",
    )
    remove_billing_history = models.BooleanField(
        _("Remove Billing History"),
        default=False,
        help_text=_("Whether users can remove records from billing history."),
    )
    auto_bill_shipment = models.BooleanField(
        _("Auto Bill Shipment"),
        default=False,
        help_text=_("Whether to automatically bill shipments directly to customer"),
    )
    auto_mark_ready_to_bill = models.BooleanField(
        _("Auto Mark Ready to Bill"),
        default=False,
        help_text=_(
            "Marks Shipments as ready to bill when they are delivered and meet customer billing requirements."
        ),
    )
    validate_customer_rates = models.BooleanField(
        _("Validate Customer Rates"),
        default=False,
        help_text=_(
            "Validate rates match the customer contract in the billing queue before allowing billing. If the rates"
            " do not match, the Shipment will not be allowed to be billed. If the rates match, the Shipment will be"
            " allowed to be billed. If no contract exists for the customer, the Shipment will be allowed to be billed."
        ),
    )
    auto_bill_criteria = ChoiceField(
        _("Auto Bill Criteria"),
        choices=AutoBillingCriteriaChoices.choices,
        default=AutoBillingCriteriaChoices.MARKED_READY_TO_BILL,
        help_text=_("Define a criteria on when auto billing is to occur."),
        blank=True,
    )
    shipment_transfer_criteria = ChoiceField(
        _("Shipment Transfer Criteria"),
        choices=ShipmentTransferCriteriaChoices.choices,
        default=ShipmentTransferCriteriaChoices.READY_AND_COMPLETED,
        help_text=_("Define when an Shipment can be transferred to billing."),
    )
    enforce_customer_billing = models.BooleanField(
        _("Enforce Customer Billing Requirements"),
        default=True,
        help_text=_(
            "Define if customer billing requirements will be enforced when billing."
        ),
    )

    class Meta:
        """
        Metaclass for BillingControl
        """

        verbose_name = _("Billing Control")
        verbose_name_plural = _("Billing Controls")
        db_table = "billing_control"
        permissions = [
            ("billing.use_billing_client", "Can use the billing client"),
        ]
        db_table_comment = (
            "Stores the billing control information for a related organization."
        )

    def __str__(self) -> str:
        """Billing control string representation

        Returns:
            str: Billing control string representation
        """
        return textwrap.wrap(self.organization.name, width=25, placeholder="...")[0]

    def get_absolute_url(self) -> str:
        """Billing Control absolute url

        Returns:
            Absolute url for the billing control object. For example,
            `/billing_control/edd1e612-cdd4-43d9-b3f3-bc099872088b/'
        """
        return reverse("billing-control-detail", kwargs={"pk": self.pk})

    def clean(self) -> None:
        """Billing control clean method

        Returns:
            None

        Raises:
            ValidationError: If billing control is not valid.
        """
        if self.auto_bill_shipment and not self.auto_bill_criteria:
            raise ValidationError(
                {
                    "auto_bill_criteria": _(
                        "Auto Billing criteria is required when `Auto Bill Shipment` is on. Please try again."
                    ),
                },
                code="invalid",
            )


class ChargeType(GenericModel):
    """Class for storing other charge types.

    Attributes:
        id (models.UUIDField): Primary key for the charge type. It has a default value
        of a new UUID, is not editable, and is unique.
        name (models.CharField): The name of the charge type. It has a max length of
        50 and must be unique.
        description (models.CharField): The description of the charge type. It has a
        max length of 100 and is optional.

    Methods:
        str(self) -> str: Returns the string representation of the charge type, which
        is the first 50 characters of the name.
        get_absolute_url(self) -> str: Returns the absolute URL for the charge type.

    Meta:
        verbose_name (str): The singular form of the name for the charge type model.
        verbose_name_plural (str): The plural form of the name for the charge type model.
        ordering (List[str]): The default ordering for instances of the charge type m
    """

    id = models.UUIDField(
        primary_key=True,
        default=uuid.uuid4,
        editable=False,
        unique=True,
    )
    status = ChoiceField(
        _("Status"),
        choices=PrimaryStatusChoices.choices,
        help_text=_("Status of the charge type."),
        default=PrimaryStatusChoices.ACTIVE,
    )
    name = models.CharField(
        _("Name"),
        max_length=50,
        help_text=_("The name of the charge type."),
    )
    description = models.CharField(
        _("Description"),
        max_length=100,
        blank=True,
        help_text=_("The description of the charge type."),
    )

    class Meta:
        """
        Metaclass for Charge Type
        """

        verbose_name = _("Charge Type")
        verbose_name_plural = _("Charge Types")
        ordering = ["name"]
        db_table = "charge_type"
        db_table_comment = "Stores charge type information for related organization."
        constraints = [
            models.UniqueConstraint(
                Lower("name"),
                "organization",
                name="unique_charge_type_organization",
            )
        ]

    def __str__(self) -> str:
        """Charge Type string representation

        Returns:
            str: Charge Type name
        """
        return textwrap.wrap(self.name, 50)[0]

    def get_absolute_url(self) -> str:
        """Charge Type absolute URL

        Returns:
            str: Charge Type absolute URL
        """
        return reverse("charge-type-detail", kwargs={"pk": self.pk})


class AccessorialCharge(GenericModel):
    """Class for storing information about other charges.

    Attributes:
        code (models.CharField): The code for the other charge. It has a max length of 50
        and must be unique. It is also the primary key for the model.
        is_detention (models.BooleanField): A boolean field indicating whether the other charge is
        a detention charge. It has a default value of False.
        charge_amount (models.DecimalField): The amount of the other charge. It has a max of 10
        digits, with 2 decimal places, and a default value of 1.00.
        method (ChoiceField): The method for calculating the other charge. It has a set of
        choices defined in the FuelMethodChoices class and a default value of
        FuelMethodChoices.DISTANCE.

    Methods:
        str(self) -> str: Returns the string representation of the other charge, which
        is the first 50 characters of the code.
        get_absolute_url(self) -> str: Returns the absolute URL for the other charge.

    Meta:
        verbose_name (str): The singular form of the name for the other charge model.
        verbose_name_plural (str): The plural form of the name for the other charge model.
        ordering (List[str]): The default ordering for instances of the other charge model.
    """

    id = models.UUIDField(
        primary_key=True,
        default=uuid.uuid4,
        editable=False,
        unique=True,
    )
    status = ChoiceField(
        _("Status"),
        choices=PrimaryStatusChoices.choices,
        help_text=_("Status of the charge type."),
        default=PrimaryStatusChoices.ACTIVE,
    )
    code = models.CharField(
        _("Code"),
        max_length=50,
        help_text=_("Code for the other charge."),
    )
    description = models.TextField(
        _("Description"),
        blank=True,
        help_text=_("Description of the other charge."),
    )
    is_detention = models.BooleanField(
        _("Is Detention"),
        default=False,
        help_text=_("Whether the other charge is a detention charge."),
    )
    method = ChoiceField(
        _("Method"),
        choices=FuelMethodChoices.choices,
        default=FuelMethodChoices.DISTANCE,
        help_text=_("Method for calculating the other charge."),
    )
    charge_amount = models.DecimalField(
        _("Additional Charge Amount"),
        max_digits=19,
        decimal_places=4,
        default=0,
        help_text=_("Additional Charge Amount"),
    )

    class Meta:
        """
        Metaclass for the AccessorialCharge model.
        """

        verbose_name = _("Accessorial Charge")
        verbose_name_plural = _("Accessorial Charges")
        ordering = ["code"]
        db_table = "accessorial_charge"
        db_table_comment = (
            "Stores Accessorial charge information for related organization."
        )
        constraints = [
            models.UniqueConstraint(
                Lower("code"),
                "organization",
                name="unique_accesorial_charge_organization",
            )
        ]

    def __str__(self) -> str:
        """Other Charge string representation

        Returns:
            str: Other Charge string representation
        """
        return textwrap.wrap(self.code, 50)[0]

    def get_absolute_url(self) -> str:
        """Other Charge absolute URL

        Returns:
            str: Other Charge absolute URL
        """
        return reverse("accessorial-charges-detail", kwargs={"pk": self.pk})


class DocumentClassification(GenericModel):
    """
    Stores Document Classification information.
    """

    id = models.UUIDField(
        primary_key=True,
        default=uuid.uuid4,
        editable=False,
        unique=True,
    )
    name = models.CharField(
        _("Name"),
        max_length=10,
        help_text=_("Document classification name"),
    )
    description = models.TextField(
        _("Description"),
        blank=True,
        help_text=_("Document classification description"),
    )

    class Meta:
        """
        Metaclass for the DocumentClassification model.
        """

        verbose_name = _("Document Classification")
        verbose_name_plural = _("Document Classifications")
        db_table = "document_classification"
        db_table_comment = (
            "Stores document classification information for related organization."
        )
        constraints = [
            models.UniqueConstraint(
                Lower("name"),
                "organization",
                name="unique_document_class_organization",
            )
        ]

    def __str__(self) -> str:
        """Document classification string representation

        Returns:
            str: Document classification string representation
        """
        return textwrap.wrap(self.name, 50)[0]

    def get_absolute_url(self) -> str:
        """Returns the url to access a particular document classification instance

        Returns:
            str: Document classification url
        """
        return reverse("document-classification-detail", kwargs={"pk": self.pk})

    def delete(self, *args: typing.Any, **kwargs: typing.Any) -> ModelDelete:
        if self.name == "CON":
            raise ValidationError(
                {
                    "name": _(
                        "Document classification with this name cannot be deleted. Please try again."
                    )
                },
                code="invalid",
            )
        return super().delete(*args, **kwargs)

    def clean(self) -> None:
        """DocumentClassification Clean Method

        Returns:
            None

        Raises:
            ValidationError: If Document Classification is not valid.
        """

        super().clean()

        if self.__class__.objects.filter(name=self.name).exclude(pk=self.pk).exists():
            raise ValidationError(
                {
                    "name": _(
                        "Document classification with this name already exists. Please try again."
                    ),
                },
            )

    def update_doc_class(self, **kwargs: typing.Any) -> None:
        """
        Updates the document classification with the given kwargs
        """
        for key, value in kwargs.items():
            setattr(self, key, value)
        self.save()


class BillingQueue(GenericModel):
    """Class for storing information about the billing queue.

    It has several fields, including:
        id (models.UUIDField): primary key and unique identifier for the billing queue.
        shipment_type (models.ForeignKey): foreign key to the `ShipmentType` model, representing the assigned Shipment type
            to the billing queue.
        Shipment (models.ForeignKey): foreign key to the `Shipment` model, representing the assigned Shipment to the billing queue.
        revenue_code (models.ForeignKey): foreign key to the `RevenueCode` model, representing the assigned revenue
            code to the billing queue.
        customer (models.ForeignKey): foreign key to the `Customer` model, representing the assigned customer
            to the billing queue.
        invoice_number (models.CharField): invoice number for the billing queue.
        pieces (models.PositiveIntegerField): total piece count of the Shipment.
        weight (models.DecimalField): total weight of the Shipment.
        bill_type (ChoiceField): bill type for the billing queue, with choices from the `BillTypeChoices` class.
        ready_to_bill (models.BooleanField): Whether Shipment is ready to be billed to the customer.
        bill_date (models.DateField): date the invoice was billed.
        mileage (models.DecimalField): total mileage.
        worker (models.ForeignKey): foreign key to the `Worker` model, representing the assigned worker
            to the billing queue.
        commodity (models.ForeignKey): foreign key to the `Commodity` model, representing the assigned commodity
            to the billing queue.
        commodity_descr (models.CharField): description of the commodity.
        other_charge_total (models.DecimalField): other charge total for the Shipment.
        freight_charge_amount (models.DecimalField): freight charge amount for the Shipment.
        total_amount (models.DecimalField): total amount for the Shipment.
    """

    @typing.final
    class BillTypeChoices(models.TextChoices):
        """
        Status choices for Shipment model
        """

        INVOICE = "INVOICE", _("Invoice")
        CREDIT = "CREDIT", _("Credit")
        DEBIT = "DEBIT", _("Debit")
        PREPAID = "PREPAID", _("Prepaid")
        OTHER = "OTHER", _("Other")

    id = models.UUIDField(
        primary_key=True,
        default=uuid.uuid4,
        editable=False,
        unique=True,
    )
    shipment_type = models.ForeignKey(
        "shipment.ShipmentType",
        on_delete=models.RESTRICT,
        verbose_name=_("Shipment Type"),
        related_name="billing_queue",
        help_text=_("Assigned Shipment type to the billing queue"),
        null=True,
        blank=True,
    )
    shipment = models.ForeignKey(
        "shipment.Shipment",
        on_delete=models.RESTRICT,
        related_name="billing_queue",
        help_text=_("Assigned Shipment to the billing queue"),
        verbose_name=_("Shipment"),
    )
    revenue_code = models.ForeignKey(
        "accounting.RevenueCode",
        on_delete=models.RESTRICT,
        related_name="billing_queue",
        verbose_name=_("Revenue Code"),
        help_text=_("Assigned revenue code to the billing queue"),
        blank=True,
        null=True,
    )
    customer = models.ForeignKey(
        "customer.Customer",
        on_delete=models.RESTRICT,
        related_name="billing_queue",
        help_text=_("Assigned customer to the billing queue"),
        verbose_name=_("Customer"),
    )
    invoice_number = models.CharField(
        _("Invoice Number"),
        max_length=50,
        blank=True,
        help_text=_("Invoice number for the billing queue"),
    )
    pieces = models.PositiveIntegerField(
        _("Pieces"),
        help_text=_("Total Piece Count of the Shipment"),
        default=0,
    )
    weight = models.DecimalField(
        _("Weight"),
        max_digits=10,
        decimal_places=2,
        help_text=_("Total Weight of the Shipment"),
        default=0,
    )
    bill_type = ChoiceField(
        _("Bill Type"),
        choices=BillTypeChoices.choices,
        default=BillTypeChoices.INVOICE,
        help_text=_("Bill type for the billing queue"),
    )
    ready_to_bill = models.BooleanField(
        _("Ready to bill"),
        default=False,
        help_text=_("Shipment is ready to be billed to customer."),
    )
    bill_date = models.DateField(
        _("Billed Date"),
        null=True,
        blank=True,
        help_text=_("Date invoiced was billed."),
    )
    mileage = models.DecimalField(
        _("Total Mileage"),
        max_digits=10,
        decimal_places=2,
        help_text=_("Total Mileage"),
        blank=True,
        null=True,
        default=0,
    )
    worker = models.ForeignKey(
        "worker.Worker",
        on_delete=models.RESTRICT,
        related_name="billing_queue",
        help_text=_("Assigned worker to the billing queue"),
        verbose_name=_("Worker"),
        blank=True,
        null=True,
    )
    commodity = models.ForeignKey(
        "commodities.Commodity",
        on_delete=models.RESTRICT,
        related_name="billing_queue",
        help_text=_("Assigned commodity to the billing queue"),
        verbose_name=_("Commodity"),
        blank=True,
        null=True,
    )
    commodity_descr = models.CharField(
        _("Commodity Description"),
        max_length=255,
        blank=True,
    )
    consignee_ref_number = models.CharField(
        _("Consignee Reference Number"),
        max_length=255,
        blank=True,
        help_text=_("Consignee Reference Number"),
    )
    other_charge_total = models.DecimalField(
        _("Other Charge Total"),
        max_digits=19,
        decimal_places=4,
        default=0,
        help_text=_("Other charge total for Shipment"),
        blank=True,
        null=True,
    )
    freight_charge_amount = models.DecimalField(
        _("Freight Charge Amount"),
        max_digits=19,
        decimal_places=4,
        default=0,
        help_text=_("Freight Charge Amount"),
        blank=True,
        null=True,
    )
    total_amount = models.DecimalField(
        _("Total Amount"),
        max_digits=19,
        decimal_places=4,
        default=0,
        help_text=_("Total amount for Shipment"),
        blank=True,
        null=True,
    )
    is_summary = models.BooleanField(
        _("Is Summary"),
        default=False,
        help_text=_("Is the invoice going to be a summary bill."),
    )
    is_cancelled = models.BooleanField(
        _("Is Cancelled"),
        default=False,
        help_text=_("Is the invoice cancelled."),
    )
    bol_number = models.CharField(
        _("BOL Number"),
        max_length=255,
        help_text=_("BOL Number"),
        blank=True,
    )
    user = models.ForeignKey(
        "accounts.User",
        on_delete=models.RESTRICT,
        related_name="billing_queue",
        help_text=_("Assigned user to the billing queue"),
        verbose_name=_("User"),
        blank=True,
        null=True,
    )
    document_class = models.ForeignKey(
        "billing.DocumentClassification",
        on_delete=models.RESTRICT,
        related_name="billing_queue",
        help_text=_("Assigned document classification to the billing queue"),
        verbose_name=_("Document Classification"),
        blank=True,
        null=True,
    )

    class Meta:
        """
        Metaclass for the BillingQueue model.
        """

        verbose_name = _("Billing Queue")
        verbose_name_plural = _("Billing Queues")
        db_table = "billing_queue"
        db_table_comment = (
            "Stores Temporary Billing Queue information for related organization."
        )
        permissions = [
            ("billing.client", "Has access to the billing client"),
        ]

    def __str__(self) -> str:
        """String Representation of the BillingQueue model

        Returns:
            str: BillingQueue string representation
        """
        return textwrap.wrap(self.shipment.pro_number, 50)[0]

    def get_absolute_url(self) -> str:
        """Billing Queue absolute url

        Returns:
            Absolute url for the billing queue object. For example,
            `/billing_queue/edd1e612-cdd4-43d9-b3f3-bc099872088b/'
        """
        return reverse("billing-queue-detail", kwargs={"pk": self.pk})

    def clean(self) -> None:
        """
        Clean method for the BillingQueue model.

        Raises:
            ValidationError: If Shipment does not meet the requirements for billing.
        """
        super().clean()

        errors = []

        # TODO (WOLFRED): Think about moving this into a validation.py file, and changing it to a dictionary of errors.

        # If Shipment is already billed raise ValidationError
        if self.shipment.billed:
            errors.append(
                _(
                    "Shipment has already been billed. Please try again with a different Shipment."
                )
            )

        # If Shipment is already transferred to billing raise ValidationError
        # if self.shipment.transferred_to_billing:
        #     errors.append(
        #         _(
        #             "Shipment has already been transferred to billing. Please try again with a different Shipment."
        #         )
        #     )

        # If Shipment is voided raise ValidationError
        if self.shipment.status == StatusChoices.VOIDED:
            errors.append(
                _(
                    "Shipment has been voided. Please try again with a different shipment."
                )
            )

        # If billing control `shipment_transfer_criteria` is `READY_AND_COMPLETE` and Shipment `status` is not `COMPLETED`
        # and Shipment `ready_to_bill` is `False` raise ValidationError
        if (
            self.shipment.organization.billing_control.shipment_transfer_criteria
            == BillingControl.ShipmentTransferCriteriaChoices.READY_AND_COMPLETED
            and self.shipment.status != StatusChoices.COMPLETED
            and self.shipment.ready_to_bill is False
        ):
            errors.append(
                _(
                    "Shipment must be `COMPLETED` and `READY_TO_BILL` must be marked before transferring to billing."
                    "Please try again."
                )
            )

        # If billing control `shipment_transfer_criteria` is `COMPLETED` and Shipment `status` is not `COMPLETED`
        # raise ValidationError
        if (
            self.shipment.organization.billing_control.shipment_transfer_criteria
            == BillingControl.ShipmentTransferCriteriaChoices.COMPLETED
            and self.shipment.status != StatusChoices.COMPLETED
        ):
            errors.append(
                _(
                    "Shipment must be `COMPLETED` before transferring to billing. Please try again."
                )
            )

        # if billing control `shipment_transfer_criteria` is `READY_TO_BILL` and Shipment `ready_to_bill` is false
        # raise ValidationError
        if (
            self.shipment.organization.billing_control.shipment_transfer_criteria
            == BillingControl.ShipmentTransferCriteriaChoices.READY_TO_BILL
            and self.shipment.ready_to_bill is False
        ):
            errors.append(
                _(
                    "Shipment must be marked `READY_TO_BILL` before transferring to billing. Please try again."
                )
            )

        if errors:
            raise ValidationError({"shipment": errors})

        # if manually entered invoice number does not start with the organization's invoice prefix
        # raise ValidationError
        if self.invoice_number and not self.invoice_number.startswith(
            self.shipment.organization.invoice_control.invoice_number_prefix
        ):
            raise ValidationError(
                {
                    "invoice_number": _(
                        "Invoice number must start with invoice prefix from Organization's invoice_control. Please "
                        "try again."
                    )
                },
                code="invalid",
            )


@typing.final
class ActionChoices(models.TextChoices):
    """
    Class for storing choices for the `action` field.
    """

    TRANSFERRED = "TRANSFERRED", _("Transferred")
    UN_TRANSFERRED = "UN_TRANSFERRED", _("Un-transferred")
    BILLED = "BILLED", _("Billed")
    VOIDED = "VOIDED", _("Voided")


class BillingLogEntry(GenericModel):
    """
    Class for storing information about the billing transfer log.
    """

    id = models.UUIDField(
        primary_key=True,
        default=uuid.uuid4,
        editable=False,
        unique=True,
        help_text=_("Unique identifier for the Billing Log"),
        verbose_name=_("Billing Log ID"),
    )
    task_id = models.CharField(
        verbose_name=_("Task ID"),
        max_length=255,
        help_text=_("Task ID for the Billing Log"),
        blank=True,
    )
    object_pk = models.CharField(
        verbose_name=_("Object PK"),
        max_length=255,
        help_text=_("Object PK for the Billing Log"),
        db_index=True,
    )
    content_type = models.ForeignKey(
        to="contenttypes.ContentType",
        on_delete=models.CASCADE,
        verbose_name=_("Content Type"),
        related_name="billing_log_entries",
        help_text=_("Content Type for the Billing Log"),
    )
    action = ChoiceField(
        verbose_name=_("Action"),
        choices=ActionChoices.choices,
        help_text=_("Action for the Billing Log"),
        default=ActionChoices.TRANSFERRED,
        db_index=True,
    )
    shipment = models.ForeignKey(
        to="shipment.Shipment",
        on_delete=models.RESTRICT,
        verbose_name=_("Shipment"),
        related_name="billing_log_entries",
        help_text=_("Shipment for the Billing Log"),
        blank=True,
        null=True,
    )
    invoice_number = models.CharField(
        _("Invoice Number"),
        max_length=50,
        blank=True,
        help_text=_("Invoice number for the Billing Log Entry"),
    )
    customer = models.ForeignKey(
        to="customer.Customer",
        on_delete=models.RESTRICT,
        verbose_name=_("Customer"),
        related_name="billing_log_entries",
        help_text=_("Customer for the Billing Log"),
    )
    actor = models.ForeignKey(
        to=settings.AUTH_USER_MODEL,
        on_delete=models.SET_NULL,
        verbose_name=_("Actor"),
        related_name="billing_log_entries",
        help_text=_("Actor for the Billing Log"),
        blank=True,
        null=True,
    )

    class Meta:
        """
        Metaclass for the BillingTransferLog model.
        """

        verbose_name = _("Billing Log Entry")
        verbose_name_plural = _("Billing Log Entries")
        db_table = "billing_log_entry"
        db_table_comment = (
            "Stores Billing Log Entry information for related organization."
        )

    def __str__(self) -> str:
        """
        String representation for the BillingLogEntry model.

        Returns:
            String representation for the BillingLogEntry model.
        """
        return textwrap.shorten(
            f"{self.shipment.pro_number if self.shipment else self.invoice_number} {self.action} by {self.actor}",
            width=100,
            placeholder="...",
        )

    def get_absolute_url(self) -> str:
        """Billing Log Entry absolute url

        Returns:
            Absolute url for the billing transfer log object. For example,
            `/billing_entry_log/edd1e612-cdd4-43d9-b3f3-bc099872088b/'
        """
        return reverse("billing-log-entry-detail", kwargs={"pk": self.pk})


class BillingHistory(GenericModel):
    """
    Class for storing information about the billing history.
    """

    id = models.UUIDField(
        primary_key=True,
        default=uuid.uuid4,
        editable=False,
        unique=True,
        help_text=_("Unique identifier for the billing history"),
    )
    shipment_type = models.ForeignKey(
        "shipment.ShipmentType",
        on_delete=models.RESTRICT,
        verbose_name=_("Shipment Type"),
        related_name="billing_history",
        help_text=_("Assigned Shipment type to the billing history"),
        blank=True,
        null=True,
    )
    shipment = models.ForeignKey(
        "shipment.Shipment",
        on_delete=models.RESTRICT,
        related_name="billing_history",
        help_text=_("Assigned Shipment to the billing history"),
        verbose_name=_("Shipment"),
    )
    revenue_code = models.ForeignKey(
        "accounting.RevenueCode",
        on_delete=models.RESTRICT,
        related_name="billing_history",
        verbose_name=_("Revenue Code"),
        help_text=_("Assigned revenue code to the billing history"),
        blank=True,
        null=True,
    )
    customer = models.ForeignKey(
        "customer.Customer",
        verbose_name=_("Customer"),
        on_delete=models.RESTRICT,
        related_name="billing_history",
        help_text=_("Assigned customer to the billing history"),
    )
    invoice_number = models.CharField(
        _("Invoice Number"),
        max_length=50,
        blank=True,
        help_text=_("Invoice number for the billing history"),
    )
    pieces = models.PositiveIntegerField(
        _("Pieces"),
        help_text=_("Total Piece Count of the Shipment"),
        default=0,
    )
    weight = models.DecimalField(
        _("Weight"),
        max_digits=10,
        decimal_places=2,
        help_text=_("Total Weight of the Shipment"),
        default=0,
    )
    bill_type = ChoiceField(
        _("Bill Type"),
        choices=BillingQueue.BillTypeChoices.choices,
        default=BillingQueue.BillTypeChoices.INVOICE,
        help_text=_("Type of bill"),
    )
    bill_date = models.DateField(
        _("Billed Date"),
        null=True,
        blank=True,
        help_text=_("Date invoiced was billed."),
    )
    mileage = models.DecimalField(
        _("Total Mileage"),
        max_digits=10,
        decimal_places=2,
        help_text=_("Total Mileage"),
        blank=True,
        null=True,
    )
    worker = models.ForeignKey(
        "worker.Worker",
        on_delete=models.RESTRICT,
        related_name="billing_history",
        help_text=_("Assigned worker to the billing history"),
        verbose_name=_("Worker"),
        blank=True,
        null=True,
    )
    commodity = models.ForeignKey(
        "commodities.Commodity",
        on_delete=models.RESTRICT,
        related_name="billing_history",
        blank=True,
        null=True,
        verbose_name=_("Commodity"),
    )
    commodity_descr = models.CharField(
        _("Commodity Description"),
        max_length=255,
        blank=True,
        help_text=_("Description of the commodity"),
    )
    consignee_ref_number = models.CharField(
        _("Consignee Reference Number"),
        max_length=255,
        blank=True,
        help_text=_("Consignee Reference Number"),
    )
    other_charge_total = models.DecimalField(
        _("Other Charge Total"),
        max_digits=19,
        decimal_places=4,
        default=0,
        help_text=_("Other charge total for Shipment"),
        blank=True,
        null=True,
    )
    freight_charge_amount = models.DecimalField(
        _("Freight Charge Amount"),
        max_digits=19,
        decimal_places=4,
        default=0,
        help_text=_("Freight Charge Amount"),
        blank=True,
        null=True,
    )
    total_amount = models.DecimalField(
        _("Total Amount"),
        max_digits=19,
        decimal_places=4,
        default=0,
        help_text=_("Total amount for Shipment"),
        blank=True,
        null=True,
    )
    is_summary = models.BooleanField(
        _("Is Summary"),
        default=False,
        help_text=_("Is the invoice going to be a summary bill."),
    )
    is_cancelled = models.BooleanField(
        _("Is Cancelled"),
        default=False,
        help_text=_("Is the invoice cancelled."),
    )
    bol_number = models.CharField(
        _("BOL Number"),
        max_length=255,
        help_text=_("BOL Number"),
        blank=True,
    )
    user = models.ForeignKey(
        "accounts.User",
        on_delete=models.RESTRICT,
        related_name="billing_history",
        help_text=_("Assigned user to the billing history"),
        verbose_name=_("User"),
        blank=True,
        null=True,
    )
    payment_received = models.BooleanField(
        verbose_name=_("Payment Received"), default=False
    )

    class Meta:
        """
        Metaclass for the BillingHistory model.
        """

        verbose_name = _("Billing History")
        verbose_name_plural = _("Billing Histories")
        ordering = ["shipment"]
        db_table = "billing_history"
        db_table_comment = (
            "Stores Billing History information for related organization."
        )

    def __str__(self) -> str:
        """String Representation of the BillingHistory model

        Returns:
            str: BillingHistory string representation
        """
        return textwrap.shorten(
            f"{self.shipment.pro_number} {self.bill_type} {self.invoice_number}",
            width=50,
            placeholder="...",
        )

    def get_absolute_url(self) -> str:
        """Billing History absolute url

        Returns:
            Absolute url for the billing history object. For example,
            `/billing_control/edd1e612-cdd4-43d9-b3f3-bc099872088b/'
        """
        return reverse("billing-history-detail", kwargs={"pk": self.pk})

    def delete(self, *args: typing.Any, **kwargs: typing.Any) -> ModelDelete:
        if self.organization.billing_control.remove_billing_history is False:
            raise ValidationError(
                {
                    "shipment": _(
                        "Your Organization disallows the deletion of billing history. Please try again."
                    ),
                },
                code="invalid",
            )
        return super().delete(*args, **kwargs)

    def clean(self) -> None:
        """Clean method for the BillingHistory model.

        Returns:
            None

        Raises:
            ValidationError
        """
        if not self.shipment.billed:
            raise ValidationError(
                {
                    "shipment": _(
                        "Shipment has not been billed. Please try again with a different Shipment."
                    ),
                },
            )


class BillingExceptionManager(models.Manager):
    """
    A custom manager for the BillingException model that provides additional functionality for
    managing BillingException instances.
    """

    def get_most_recent_exception(self, *shipment: Shipment) -> BillingException:
        """
        Retrieve the most recent BillingException instance associated with the provided Shipment.

        This method queries the database for the most recent `BillingException` instance associated with
        the provided `Shipment`. It takes one argument:
        - `shipment` (type: `Shipment`): The `Shipment` object to retrieve the most recent `BillingException` for.

        If a `BillingException` instance is found, it is returned. Otherwise, the method returns `None`.

        Args:
            shipment (Shipment): The Shipment object to retrieve the most recent BillingException for.

        Returns:
            BillingException: The most recent BillingException instance associated with the provided Shipment,
            or None if no BillingException instance is found.
        """
        return self.filter(shipment=shipment).latest("created_at")  # type: ignore

    def create_billing_exception(
        self,
        *,
        organization: Organization,
        business_unit: BusinessUnit,
        exception_type: str,
        shipment: Shipment,
        exception_message: str,
    ) -> BillingException:
        """
        Create a new BillingException instance with the provided information.

        Args:
            organization (Organization): The Organization object associated with the BillingException.
            business_unit (BusinessUnit): The BusinessUnit object associated with the BillingException.
            exception_type (str): A string representing the type of the BillingException.
            shipment (Shipment): The Shipment object associated with the BillingException.
            exception_message (str): A message describing the BillingException.

        Returns:
            BillingException: The new BillingException instance.
        """
        return self.create(  # type: ignore
            organization=organization,
            business_unit=business_unit,
            exception_type=exception_type,
            shipment=shipment,
            exception_message=exception_message,
        )


class BillingException(GenericModel):
    """
    Stores billing exception exception information for a related :model:`shipment.Shipment`.
    """

    id = models.UUIDField(
        primary_key=True,
        default=uuid.uuid4,
        editable=False,
        unique=True,
    )
    exception_type = ChoiceField(
        _("Exception Type"),
        choices=BillingExceptionChoices.choices,
        default=BillingExceptionChoices.PAPERWORK,
        help_text=_("Type of billing exception"),
    )
    shipment = models.ForeignKey(
        "shipment.Shipment",
        on_delete=models.RESTRICT,
        related_name="billing_exception",
        help_text=_("Assigned Shipment to the billing exception"),
    )
    exception_message = models.TextField(
        _("Exception Message"),
        help_text=_("Message for the billing exception"),
        blank=True,
    )

    objects = BillingExceptionManager()

    class Meta:
        """
        Metaclass for the BillingException model.
        """

        verbose_name = _("Billing Exception")
        verbose_name_plural = _("Billing Exceptions")
        ordering = ("shipment",)
        db_table = "billing_exception"
        db_table_comment = "Stores Billing exception for related organization."

    def __str__(self) -> str:
        """String Representation of the BillingException model

        Returns:
            str: BillingException string representation
        """
        return textwrap.shorten(
            f"{self.exception_type} - {self.exception_message}",
            width=50,
            placeholder="...",
        )

    def get_absolute_url(self) -> str:
        """Billing Exception absolute url

        Returns:
            str: Absolute url for the billing exception object.
        """
        return reverse("billing-exception-detail", kwargs={"pk": self.pk})


class InvoicePaymentDetail(GenericModel):
    """
    Stores invoice payment detail information for a related :model:`billing.BillingHistory`.
    """

    id = models.UUIDField(
        primary_key=True,
        default=uuid.uuid4,
        editable=False,
        unique=True,
    )
    invoice = models.ForeignKey(
        BillingHistory,
        on_delete=models.RESTRICT,
        help_text=_("Invoice for the payment."),
        verbose_name=_("Invoice"),
    )
    check_number = models.CharField(
        verbose_name=_("Check Number"),
        help_text=_("Check number for the payment."),
        max_length=100,
        blank=True,
    )
    payment_amount = models.DecimalField(
        verbose_name=_("Payment Amount"),
        help_text=_("Payment amount received."),
        max_digits=10,
        decimal_places=2,
    )
    payment_date = models.DateField(
        verbose_name=_("Payment Date"),
        help_text=_("Date payment was received."),
    )
    payment_method = models.CharField(
        verbose_name=_("Payment Method"),
        help_text=_("Payment method used."),
        max_length=100,
    )

    class Meta:
        """
        Metaclass for the InvoicePaymentDetail model.
        """

        verbose_name = _("Invoice Payment Detail")
        verbose_name_plural = _("Invoice Payment Details")
        ordering = ("invoice",)
        db_table = "invoice_payment_detail"
        db_table_comment = "Table for storing invoice payment details."

    def __str__(self) -> str:
        """String Representation of the InvoicePaymentDetail model

        Returns:
            str: InvoicePaymentDetail string representation
        """
        return textwrap.shorten(
            f"{self.invoice} - {self.payment_amount}",
            width=50,
            placeholder="...",
        )

    def get_absolute_url(self) -> str:
        """Invoice Payment Detail absolute url

        Returns:
            str: Absolute url for the invoice payment detail object.
        """
        return reverse("invoice-payment-detail-detail", kwargs={"pk": self.pk})
