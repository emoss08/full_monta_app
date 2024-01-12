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

import textwrap
import uuid
from typing import Any, final

from django.core.exceptions import ValidationError
from django.core.validators import EmailValidator, RegexValidator
from django.db import models
from django.db.models.functions import Lower
from django.urls import reverse
from django.utils.functional import cached_property
from django.utils.translation import gettext_lazy as _
from localflavor.us.models import USZipCodeField

from billing.models import AccessorialCharge, DocumentClassification
from utils.models import ChoiceField, GenericModel, PrimaryStatusChoices, Weekdays


@final
class FuelMethodChoices(models.TextChoices):
    """
    Fuel Method Choices
    """

    DISTANCE = "D", _("Distance")
    FLAT = "F", _("Flat")
    PERCENTAGE = "P", _("Percentage")


class Customer(GenericModel):
    """
    Stores customer information for billing and invoicing
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
        help_text=_("Status of the Customer."),
        default=PrimaryStatusChoices.ACTIVE,
    )
    code = models.CharField(
        _("Code"),
        max_length=10,
        editable=False,
        help_text=_("Customer code"),
    )
    name = models.CharField(
        _("Name"),
        max_length=150,
        help_text=_("Customer name"),
        db_index=True,
    )
    address_line_1 = models.CharField(
        _("Address Line 1"),
        max_length=150,
        help_text=_("Address line 1"),
    )
    address_line_2 = models.CharField(
        _("Address Line 2"),
        max_length=150,
        blank=True,
        help_text=_("Address line 2"),
    )
    city = models.CharField(
        _("City"),
        max_length=150,
        help_text=_("City"),
    )
    state = models.CharField(
        _("State"),
        max_length=5,
        help_text=_("State"),
    )
    zip_code = USZipCodeField(
        _("Zip Code"),
        help_text=_("Zip code"),
    )
    has_customer_portal = models.BooleanField(
        _("Has Customer Portal?"),
        help_text=_(
            "Designates whether this customer has the customer portal. "
            "active or not."
        ),
        default=False,
    )
    auto_mark_ready_to_bill = models.BooleanField(
        _("Automatic Billing Readiness?"),
        help_text=_("Enable automatic marking of billing readiness"),
        default=False,
    )
    advocate = models.ForeignKey(
        to="accounts.User",
        on_delete=models.CASCADE,
        related_name="customers",
        related_query_name="customer",
        help_text=_("Customer Advocate assigned to this customer."),
        verbose_name=_("Customer Advocate"),
        blank=True,
        null=True,
    )

    class Meta:
        """
        Customer Metaclass
        """

        verbose_name = _("Customer")
        verbose_name_plural = _("Customers")
        db_table = "customer"
        constraints = [
            models.UniqueConstraint(
                Lower("code"),
                "organization",
                name="unique_customer_organization",
            )
        ]

    def __str__(self) -> str:
        """Customer string representation

        Returns:
            str: Customer string representation
        """
        return textwrap.shorten(
            f"Customer {self.code}: {self.name}", width=40, placeholder="..."
        )

    def save(self, *args: Any, **kwargs: Any) -> None:
        """Customer save method

        Args:
            *args (Any): Variable length argument list.
            **kwargs (Any): Arbitrary keyword arguments

        Returns:
            None: This function does return anything.
        """

        if not self.code:
            self.code = self.generate_customer_code()

        super().save(*args, **kwargs)

    def get_absolute_url(self) -> str:
        """Returns the url to access a particular customer instance

        Returns:
            str: Customer url
        """
        return reverse("customer-detail", kwargs={"pk": self.pk})

    def generate_customer_code(self) -> str:
        """Generates a unique customer code.

        The customer code is created by using the first 5 characters of the customer's
        name and concatenating it with a generated 4-digit numeric sequence. This sequence
        is incremented by 1 with each new code generated.

        Returns:
            str: A unique customer code.
        """
        base_code = self.name.replace(" ", "")[:5].upper()
        sequence = 1

        while True:
            new_code = f"{base_code}{sequence:04d}"
            if not self.__class__.objects.filter(code=new_code).exists():
                return new_code.upper()  # Convert to uppercase before returning
            sequence += 1
            if len(f"{base_code}{sequence:04d}") > 10:
                # If we've reached the max length, reset the base code and sequence
                base_code = base_code[:-1]
                sequence = 1

    @cached_property
    def get_address_combination(self) -> str:
        """
        Returns:
            str: String representation of the customer address.
        """
        return f"{self.address_line_1} {self.address_line_2} {self.city} {self.state} {self.zip_code}"

    @cached_property
    def get_address(self) -> str:
        """
        Returns:
            str: String representation of the customer address.
        """
        return f"{self.address_line_1} {self.address_line_2}"

    @cached_property
    def get_city_state_zip(self) -> str:
        """
        Returns:
            str: String representation of the customer address.
        """
        return f"{self.city}, {self.state} {self.zip_code}"


class CustomerEmailProfile(GenericModel):
    """
    Stores Customer Email Profile related to the :model:`customer.Customer`. model.
    """

    id = models.UUIDField(
        primary_key=True,
        default=uuid.uuid4,
        editable=False,
        unique=True,
    )
    customer = models.OneToOneField(
        to=Customer,
        on_delete=models.CASCADE,
        related_name="email_profile",
        help_text=_("Customer assigned to Email Profile"),
        verbose_name=_("Customer"),
        blank=True,
        null=True,
    )
    subject = models.CharField(
        _("Subject"),
        max_length=100,
        help_text=_("Subject"),
        blank=True,
    )
    comment = models.CharField(
        _("Comment"),
        max_length=100,
        help_text=_("Comment"),
        blank=True,
    )
    from_address = models.CharField(
        _("From Address"),
        max_length=255,
        help_text=_("From Address"),
        blank=True,
    )
    blind_copy = models.CharField(
        _("Blind Copy"),
        max_length=255,
        help_text=_("Comma separated list of email addresses"),
        blank=True,
    )
    read_receipt = models.BooleanField(
        _("Read Receipt"),
        help_text=_("Read Receipt"),
        default=False,
    )
    read_receipt_to = models.EmailField(
        _("Read Receipt To"),
        max_length=255,
        help_text=_("Read Receipt To"),
        blank=True,
    )
    attachment_name = models.CharField(
        _("Attachment Name"),
        max_length=255,
        help_text=_("Attachment Name"),
        blank=True,
    )

    class Meta:
        """
        Metaclass for CustomerEmailProfile
        """

        verbose_name = _("Customer Email Profile")
        verbose_name_plural = _("Customer Email Profiles")
        db_table = "customer_email_profile"

    def __str__(self) -> str:
        """CustomerEmailProfile string representation

        Returns:
            str: Customer Email Profile string representation
        """
        return textwrap.shorten(
            f"Customer Email Profile for Customer {self.customer.name if self.customer else 'None'}",
            width=60,
            placeholder="...",
        )

    def clean(self) -> None:
        """CustomerEmailProfile clean method

        Returns:
            None: This function does not return anything.
        """
        # Call the parent class's clean method first
        super().clean()

        # Split the blind_copy field by comma to get individual emails
        emails = [
            email.strip() for email in self.blind_copy.split(",") if email.strip()
        ]

        # Use Django's EmailValidator to validate each email
        validator = EmailValidator()
        for email in emails:
            try:
                validator(email)
            except ValidationError as e:
                raise ValidationError(
                    {
                        "blind_copy": _(
                            f"{email} is not a valid email address. Please try again."
                        ),
                    },
                    code="invalid",
                ) from e

    def get_absolute_url(self) -> str:
        """Returns the url to access a particular customer email profile instance

        Returns:
            str: Customer email profile url
        """
        return reverse("customer-email-profile-detail", kwargs={"pk": self.pk})


class CustomerRuleProfile(GenericModel):
    """
    Stores Customer Rule Profile information related to :model:`customer.Customer`. model.
    """

    id = models.UUIDField(
        primary_key=True,
        default=uuid.uuid4,
        editable=False,
        unique=True,
    )
    name = models.CharField(
        _("Name"),
        max_length=50,
        help_text=_("Name"),
    )
    customer = models.OneToOneField(
        to=Customer,
        on_delete=models.CASCADE,
        related_name="rule_profile",
        help_text=_("Customer assigned to Rule Profile"),
        verbose_name=_("Customer"),
        blank=True,
        null=True,
    )
    document_class = models.ManyToManyField(
        DocumentClassification,
        related_name="customer_rule_profile",
        verbose_name=_("Document Class"),
        help_text=_("Document class"),
        blank=True,
    )

    class Meta:
        """
        Metaclass for CustomerRuleProfile
        """

        verbose_name = _("Customer Rule Profile")
        verbose_name_plural = _("Customer Rule Profiles")
        db_table = "customer_rule_profile"
        constraints = [
            models.UniqueConstraint(
                Lower("name"),
                "organization",
                name="unique_customer_rule_profile_organization",
            )
        ]

    def __str__(self) -> str:
        """CustomerRuleProfile string representation

        Returns:
            str: Customer Rule Profile string representation
        """
        return textwrap.shorten(
            f"Customer Rule profile {self.name}",
            width=40,
            placeholder="...",
        )

    def get_absolute_url(self) -> str:
        """Returns the url to access a particular customer rule profile instance

        Returns:
            str: Customer rule profile url
        """
        return reverse("customer-rule-profile-detail", kwargs={"pk": self.pk})


class CustomerContact(GenericModel):
    """
    Stores contract information related to :model:`billing.Customer`.
    """

    id = models.UUIDField(
        primary_key=True,
        default=uuid.uuid4,
        editable=False,
        unique=True,
    )
    customer = models.ForeignKey(
        Customer,
        on_delete=models.CASCADE,
        related_name="contacts",
        related_query_name="contact",
        verbose_name=_("Customer"),
        help_text=_("Customer"),
    )
    is_active = models.BooleanField(
        _("Active"),
        default=True,
        help_text=_(
            "Designates whether this customer contact should be treated as active. "
            "Unselect this instead of deleting customer contacts."
        ),
    )
    name = models.CharField(
        _("Name"),
        max_length=150,
        help_text=_("Contact name"),
        db_index=True,
    )
    email = models.EmailField(
        _("Email"),
        max_length=150,
        help_text=_("Contact email"),
        blank=True,
    )
    title = models.CharField(
        _("Title"),
        max_length=100,
        help_text=_("Contact title"),
        blank=True,
    )
    phone = models.CharField(
        _("Phone Number"),
        max_length=15,
        blank=True,
        help_text=_("The phone number of the user"),
        validators=[
            RegexValidator(
                regex=r"^\(?([0-9]{3})\)?[-. ]?([0-9]{3})[-. ]?([0-9]{4})$",
                message=_("Phone number must be in the format (xxx) xxx-xxxx"),
            )
        ],
    )
    is_payable_contact = models.BooleanField(
        _("Payable Contact"),
        default=False,
        help_text=_("Designates whether this contact is the payable contact"),
    )

    class Meta:
        """
        Metaclass for CustomerContact
        """

        verbose_name = _("Customer Contact")
        verbose_name_plural = _("Customer Contacts")
        db_table = "customer_contact"

    def __str__(self) -> str:
        """Customer Contact string representation

        Returns:
            str: Customer Contact string representation
        """
        return textwrap.shorten(
            f"Contact {self.name} for Customer {self.customer.name}",
            width=50,
            placeholder="...",
        )

    def get_absolute_url(self) -> str:
        """Returns the url to access a particular customer contact instance

        Returns:
            str: Customer contact url
        """
        return reverse("customer-contact-detail", kwargs={"pk": self.pk})

    def clean(self) -> None:
        """Customer Contact clean method

        Returns:
            None

        Raises:
            ValidationError: If the customer contact is not valid.
        """
        super().clean()

        if self.is_payable_contact and not self.email:
            raise ValidationError(
                {
                    "email": _(
                        "Payable contact must have an email address. Please Try Again."
                    ),
                }
            )


class CustomerFuelProfile(GenericModel):
    """
    Stores Customer Fuel Profile information related to the :model:`billing.Customer`. model.
    """

    @final
    class DaysToUseChoices(models.TextChoices):
        """
        Days to use choices
        """

        DELIVERY = "D", _("Delivery Date")
        SHIPMENT = "S", _("Actual Shipment Date")
        SCHED_SHIPMENT = "C", _("Scheduled Shipment Date")
        ENTERED = "E", _("Date Entered")

    @final
    class FuelRegionChoices(models.TextChoices):
        """
        Fuel Region choices
        """

        USA = "USA", _("US National Average")
        EAST = "EAST", _("East Coast")
        NEW_ENGLAND = "NE", _("New England")
        GENERAL_ATL = "GA", _("General Atlantic")
        LOWER_ATL = "LA", _("Lower Atlantic")
        MIDWEST = "MW", _("Midwest")
        GULF_COAST = "GC", _("Gulf Coast")
        ROCKY_MOUNTAIN = "RM", _("Rocky Mountain")
        WEST_COAST = "WC", _("West Coast")
        CALIFORNIA = "CA", _("California")
        WEST_COAST_NO_LA = "WCL", _("West Coast (No LA)")

    @final
    class TableChoices(models.TextChoices):
        """
        Table choices
        """

        PERCENTAGE = "P", _("Percentage")
        FLAT = "F", _("Flat")
        DISTANCE = "D", _("Distance")
        TABLE = "T", _("Table")

    id = models.UUIDField(
        primary_key=True,
        default=uuid.uuid4,
        editable=False,
        unique=True,
    )
    customer = models.ForeignKey(
        Customer,
        on_delete=models.CASCADE,
        related_name="customer_fuel_profiles",
        related_query_name="customer_fuel_profile",
        help_text=_("Customer"),
        verbose_name=_("Customer"),
    )
    fsc_code = models.ForeignKey(
        AccessorialCharge,
        on_delete=models.CASCADE,
        related_name="customer_fuel_profiles",
        related_query_name="customer_fuel_profile",
        help_text=_("FSC Code"),
        verbose_name=_("FSC Code"),
    )
    start_date = models.DateField(
        _("Start Date"),
        help_text=_("Start Date"),
    )
    end_date = models.DateField(
        _("End Date"),
        help_text=_("End Date"),
        null=True,
        blank=True,
    )
    days_to_use = ChoiceField(
        _("Days to Use"),
        choices=DaysToUseChoices.choices,
        help_text=_("Days to Use"),
    )
    shipment_type = models.ForeignKey(
        "shipment.ShipmentType",
        on_delete=models.CASCADE,
        related_name="customer_fuel_profiles",
        related_query_name="customer_fuel_profile",
        help_text=_("shipment Type"),
        verbose_name=_("shipment Type"),
    )
    fuel_region = ChoiceField(
        _("Fuel Region"),
        choices=FuelRegionChoices.choices,
        help_text=_("Fuel Region"),
    )
    fsc_method = ChoiceField(
        _("FSC Method"),
        choices=TableChoices.choices,
        help_text=_("FSC Method"),
    )
    customer_fuel_table = models.ForeignKey(
        "CustomerFuelTable",
        on_delete=models.CASCADE,
        related_name="customer_fuel_profiles",
        related_query_name="customer_fuel_profile",
        help_text=_("Customer Fuel Table"),
        verbose_name=_("Customer Fuel Table"),
        blank=True,
        null=True,
    )
    base_price = models.DecimalField(
        _("Base Price"),
        max_digits=16,
        decimal_places=3,
        help_text=_("Base Price"),
        blank=True,
        null=True,
    )
    fuel_variance = models.DecimalField(
        _("Fuel Variance"),
        max_digits=16,
        decimal_places=3,
        help_text=_("Fuel Variance ex: 0.02"),
        blank=True,
        null=True,
    )
    amount = models.DecimalField(
        _("Amount"),
        max_digits=16,
        decimal_places=5,
        help_text=_("Amount"),
        blank=True,
        null=True,
    )
    percentage = models.DecimalField(
        _("Percentage"),
        max_digits=6,
        decimal_places=2,
        help_text=_("Percentage"),
        blank=True,
        null=True,
    )

    class Meta:
        """
        Meta Class for Customer Fuel Profile
        """

        verbose_name = _("Customer Fuel Profile")
        verbose_name_plural = _("Customer Fuel Profiles")
        db_table = "customer_fuel_profile"

    def __str__(self) -> str:
        """Customer Fuel Profile string representation

        Returns:
            str: Customer Fuel Profile string representation
        """
        return textwrap.shorten(
            f"Customer Fuel Profile for Customer {self.customer.name}",
            width=40,
            placeholder="...",
        )

    def get_absolute_url(self) -> str:
        """Returns the url to access a particular customer fuel profile instance

        Returns:
            str: Customer fuel profile url
        """
        return reverse("customer-fuel-profile-details", kwargs={"pk": self.pk})

    def clean(self) -> None:
        """CustomerFuelProfile clean method

        Returns:
            None

        Raises:
            ValidationError: If the Customer Fuel Profile is not valid.

        """
        super().clean()

        if self.fsc_method == CustomerFuelProfile.TableChoices.TABLE:
            raise ValidationError(
                {
                    "customer_fuel_table": _(
                        "Customer Fuel Table is required if the FSC Method is Table. Please try again."
                    )
                },
                code="required",
            )


class CustomerFuelTable(GenericModel):
    """
    Stores Customer Fuel Profile Information related to the :model:`billing.Customer` model.
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
        help_text=_("Customer Fuel Profile Name"),
    )
    description = models.CharField(
        _("Description"),
        max_length=150,
        help_text=_("Customer Fuel Profile Description"),
        blank=True,
    )

    class Meta:
        """
        Meta Class for Customer Fuel Table
        """

        verbose_name = _("Customer Fuel Table")
        verbose_name_plural = _("Customer Fuel Table")
        db_table = "customer_fuel_table"
        constraints = [
            models.UniqueConstraint(
                Lower("name"),
                "organization",
                name="unique_customer_fuel_table_name_organization",
            )
        ]

    def __str__(self) -> str:
        """Customer Fuel Table string representation

        Returns:
            str: Customer Fuel Table string representation
        """
        return textwrap.shorten(
            f"Customer Fuel Table {self.name}",
            width=30,
            placeholder="...",
        )

    def get_absolute_url(self) -> str:
        """Returns the url to access a particular customer fuel profile instance

        Returns:
            str: Customer fuel profile url
        """
        return reverse("billing:customer-fuel-profile-detail", kwargs={"pk": self.pk})


class CustomerFuelTableDetail(GenericModel):
    """
    Stores detailed information related to the `customer.CustomerFuelTable` model.
    """

    id = models.UUIDField(
        primary_key=True,
        default=uuid.uuid4,
        editable=False,
        unique=True,
    )
    customer_fuel_table = models.ForeignKey(
        CustomerFuelTable,
        on_delete=models.CASCADE,
        related_name="customer_fuel_table_details",
        related_query_name="customer_fuel_table_detail",
        help_text=_("Customer Fuel Profile"),
        verbose_name=_("Customer Fuel Profile"),
    )
    amount = models.DecimalField(
        _("Amount"),
        max_digits=16,
        decimal_places=5,
        help_text=_("Amount"),
        blank=True,
        null=True,
    )
    method = ChoiceField(
        _("Method"),
        choices=FuelMethodChoices.choices,
        help_text=_("Method"),
    )
    start_price = models.DecimalField(
        _("Start Price"),
        max_digits=5,
        decimal_places=3,
        help_text=_("Start Price"),
        blank=True,
        null=True,
    )
    percentage = models.DecimalField(
        _("Percentage"),
        max_digits=6,
        decimal_places=2,
        help_text=_("Percentage"),
        blank=True,
        null=True,
    )

    class Meta:
        """
        Meta class for Customer Fuel Profile Details
        """

        verbose_name = _("Customer Fuel Profile Detail")
        verbose_name_plural = _("Customer Fuel Profile Details")
        db_table = "customer_fuel_profile_detail"

    def __str__(self) -> str:
        """Customer Fuel Profile Detail string representation

        Returns:
            str: Customer Fuel Profile Detail string representation
        """
        return textwrap.shorten(
            f"Details for Customer fuel table {self.customer_fuel_table.name}",
            width=50,
            placeholder="...",
        )

    def get_absolute_url(self) -> str:
        """Returns the url to access a particular customer fuel profile detail instance

        Returns:
            str: Customer fuel profile detail url
        """
        return reverse("billing:customer-fuel-profile-detail", kwargs={"pk": self.pk})


class DeliverySlot(GenericModel):
    """
    Stores delivery slot information related to the :model:`billing.Customer` model.
    """

    id = models.UUIDField(
        primary_key=True,
        default=uuid.uuid4,
        editable=False,
        unique=True,
    )
    customer = models.ForeignKey(
        Customer,
        on_delete=models.CASCADE,
        related_name="delivery_slots",
        help_text=_("Customer"),
        verbose_name=_("Customer"),
    )
    day_of_week = models.PositiveSmallIntegerField(
        _("Day of Week"),
        choices=Weekdays.choices,
        default=Weekdays.MONDAY,
        help_text=_("Day of the week associated with the delivery slot."),
    )
    start_time = models.TimeField(_("Start Time"), help_text=_("Start Time"))
    end_time = models.TimeField(_("End Time"), help_text=_("End Time"))
    location = models.ForeignKey(
        "location.Location",
        on_delete=models.CASCADE,
        help_text=_("Location"),
        verbose_name=_("Location"),
    )

    class Meta:
        """
        Metaclass for the Delivery Slot model.
        """

        verbose_name = _("Delivery Slot")
        verbose_name_plural = _("Delivery Slots")
        db_table = "delivery_slot"
        constraints = [
            models.UniqueConstraint(
                fields=[
                    "customer",
                    "day_of_week",
                    "start_time",
                    "end_time",
                    "location",
                ],
                name="unique_ds_customer_day_start_end_loc",
            ),
            models.CheckConstraint(
                check=models.Q(start_time__lt=models.F("end_time")),
                name="start_time_lt_end_time",
            ),
        ]

    def __str__(self) -> str:
        """String representation of Delivery Slot

        Returns:
            str: Delivery Slot string representation
        """
        return textwrap.shorten(
            f"Delivery Slot for {self.customer.name} on {self.get_day_of_week_display()} from {self.start_time} to"
            f" {self.end_time}",
            width=50,
            placeholder="...",
        )

    def get_absolute_url(self) -> str:
        """Returns the url to access a particular delivery slot instance

        Returns:
            str: Delivery slot url
        """
        return reverse("delivery-slot-detail", kwargs={"pk": self.pk})

    def clean(self) -> None:
        super().clean()

        errors = {}

        # Check if the start time is less than the end time
        if self.start_time >= self.end_time:
            errors["start_time"] = _(
                "Start time must be less than end time. Please try again."
            )

        # Check if the delivery slot overlaps with another delivery slot
        if (
            self.__class__.objects.filter(
                customer=self.customer,
                day_of_week=self.day_of_week,
                location=self.location,
            )
            .exclude(pk=self.pk)
            .filter(
                models.Q(start_time__lte=self.start_time, end_time__gt=self.start_time)
                | models.Q(start_time__lt=self.end_time, end_time__gte=self.end_time)
                | models.Q(start_time__gte=self.start_time, end_time__lte=self.end_time)
            )
            .exists()
        ):
            errors["start_time"] = _(
                "Delivery slot overlaps with another delivery slot. Please try again."
            )

        if errors:
            raise ValidationError(errors)
