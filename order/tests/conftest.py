# --------------------------------------------------------------------------------------------------
#  COPYRIGHT(c) 2023 MONTA                                                                         -
#                                                                                                  -
#  This file is part of Monta.                                                                     -
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

from typing import Any
from collections.abc import Generator

import pytest
from django.utils import timezone
from rest_framework.response import Response
from rest_framework.test import APIClient

from accounting.models import RevenueCode
from accounting.tests.factories import RevenueCodeFactory
from accounts.models import User
from billing.models import AccessorialCharge, DocumentClassification
from billing.tests.factories import (
    AccessorialChargeFactory,
    DocumentClassificationFactory,
)
from customer.factories import CustomerFactory
from customer.models import Customer
from dispatch.factories import CommentTypeFactory
from dispatch.models import CommentType
from equipment.models import EquipmentType
from equipment.tests.factories import EquipmentTypeFactory
from location.factories import LocationFactory
from location.models import Location
from order.models import Order, OrderType
from order.tests.factories import (
    AdditionalChargeFactory,
    OrderCommentFactory,
    OrderDocumentationFactory,
    OrderFactory,
    OrderTypeFactory,
    ReasonCodeFactory,
)
from organization.models import Organization

pytestmark = pytest.mark.django_db


@pytest.fixture
def order_type() -> Generator[Any, Any, None]:
    """
    Pytest Fixture for order Type
    """
    yield OrderTypeFactory()


@pytest.fixture
def order() -> Generator[Any, Any, None]:
    """
    Pytest Fixture for Order
    """
    yield OrderFactory()


@pytest.fixture
def document_classification() -> Generator[Any, Any, None]:
    """
    Pytest Fixture for Document Classification
    """
    yield DocumentClassificationFactory()


@pytest.fixture
def reason_code() -> Generator[Any, Any, None]:
    """
    Pytest Fixture for Reason Code
    """
    yield ReasonCodeFactory()


@pytest.fixture
def order_document() -> Generator[Any, Any, None]:
    """
    Pytest Fixture for Order Documentation
    """
    yield OrderDocumentationFactory()


@pytest.fixture
def additional_charge() -> Generator[Any, Any, None]:
    """
    Pytest Fixture for order Type
    """
    yield AdditionalChargeFactory()


@pytest.fixture
def accessorial_charge() -> Generator[Any, Any, None]:
    """
    Pytest Fixture for Accessorial Charge
    """
    yield AccessorialChargeFactory()


@pytest.fixture
def revenue_code() -> Generator[Any, Any, None]:
    """
    Pytest Fixture for Revenue Code
    """
    yield RevenueCodeFactory()


@pytest.fixture
def customer() -> Generator[Any, Any, None]:
    """
    Pytest Fixture for Customer
    """
    yield CustomerFactory()


@pytest.fixture
def equipment_type() -> Generator[Any, Any, None]:
    """
    Pytest Fixture for Equipment Type
    """
    yield EquipmentTypeFactory()


@pytest.fixture
def order_comment() -> Generator[Any, Any, None]:
    """
    Pytest Fixture for Order Comment
    """
    yield OrderCommentFactory()


@pytest.fixture
def comment_type() -> Generator[Any, Any, None]:
    """
    Pytest Fixture for Comment Type
    """
    yield CommentTypeFactory()


@pytest.fixture
def origin_location() -> Generator[Any, Any, None]:
    """
    Pytest Fixture for Origin Location
    """
    yield LocationFactory()


@pytest.fixture
def destination_location() -> Generator[Any, Any, None]:
    """
    Pytest Fixture for Destination Location
    """
    yield LocationFactory()


@pytest.fixture
def order_api(
    api_client: APIClient,
    organization: Organization,
    order_type: OrderType,
    revenue_code: RevenueCode,
    origin_location: Location,
    destination_location: Location,
    customer: Customer,
    equipment_type: EquipmentType,
    user: User,
) -> Generator[Any, Any, None]:
    """
    Pytest Fixture for Reason Code
    """
    return api_client.post(
        "/api/orders/",
        {
            "organization": f"{organization.id}",
            "order_type": f"{order_type.id}",
            "revenue_code": f"{revenue_code.id}",
            "origin_location": f"{origin_location.id}",
            "origin_appointment": f"{timezone.now()}",
            "destination_location": f"{destination_location.id}",
            "destination_appointment": f"{timezone.now()}",
            "freight_charge_amount": 100.00,
            "customer": f"{customer.id}",
            "equipment_type": f"{equipment_type.id}",
            "entered_by": f"{user.id}",
            "bol_number": "newbol",
        },
        format="json",
    )


@pytest.fixture
def additional_charge_api(
    api_client: APIClient,
    user: User,
    organization: Organization,
    order: Order,
    accessorial_charge: AccessorialCharge,
) -> Generator[Any, Any, None]:
    """
    Additional Charge Factory
    """
    yield api_client.post(
        "/api/additional_charges/",
        {
            "organization": f"{organization.id}",
            "order": f"{order.id}",
            "charge": f"{accessorial_charge.id}",
            "charge_amount": 123.00,
            "unit": 2,
            "entered_by": f"{user.id}",
        },
        format="json",
    )


@pytest.fixture
def order_comment_api(
    order_api: Response, user: User, comment_type: CommentType, api_client: APIClient
) -> Generator[Any, Any, None]:
    """
    Pytest Fixture for Order Comment
    """
    yield api_client.post(
        "/api/order_comments/",
        {
            "order": f"{order_api.data['id']}",
            "comment_type": f"{comment_type.id}",
            "comment": "IM HAPPY YOU'RE HERE",
            "entered_by": f"{user.id}",
        },
        format="json",
    )


@pytest.fixture
def order_documentation_api(
    api_client: APIClient,
    order: Order,
    document_classification: DocumentClassification,
    organization: Organization,
) -> Generator[Any, Any, None]:
    """
    Pytest Fixture for Order Documentation
    """

    with open("order/tests/files/dummy.pdf", "rb") as test_file:
        yield api_client.post(
            "/api/order_documents/",
            {
                "organization": f"{organization}",
                "order": f"{order.id}",
                "document": test_file,
                "document_class": f"{document_classification.id}",
            },
        )


@pytest.fixture
def order_type_api(
    api_client: APIClient, organization: Organization
) -> Generator[Any, Any, None]:
    """
    Order Type Factory
    """
    yield api_client.post(
        "/api/order_types/",
        {
            "organization": organization.id,
            "name": "Foo Bar",
            "description": "Foo Bar",
            "is_active": True,
        },
    )


@pytest.fixture
def reason_code_api(
    api_client: APIClient, organization: Organization
) -> Generator[Any, Any, None]:
    """
    Reason Code Factory
    """
    yield api_client.post(
        "/api/reason_codes/",
        {
            "organization": organization.id,
            "code": "NEWT",
            "description": "Foo Bar",
            "is_active": True,
            "code_type": "VOIDED",
        },
    )
