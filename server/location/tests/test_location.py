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

import pytest
from django.core.exceptions import ValidationError
from rest_framework.test import APIClient

from accounts.models import User
from dispatch.models import CommentType
from location import models
from organization.models import BusinessUnit, Organization

pytestmark = pytest.mark.django_db


def test_location_creation(location: models.Location) -> None:
    """
    Test location creation
    """
    assert location is not None


def test_location_update(location: models.Location) -> None:
    """
    Test location update
    """
    location.name = "New name"
    location.save()
    assert location.name == "New name"


def test_post_location(
    api_client: APIClient,
    organization: Organization,
    business_unit: BusinessUnit,
) -> None:
    response = api_client.post(
        "/api/locations/",
        {
            "organization": organization.id,
            "business_unit": business_unit.id,
            "code": "string",
            "name": "string",
            "address_line_1": "string",
            "city": "string",
            "state": "NC",
            "zip_code": "12345",
        },
        format="json",
    )

    assert response.status_code == 201
    assert response.data["code"] == "string"
    assert response.data["name"] == "string"
    assert response.data["address_line_1"] == "string"
    assert response.data["city"] == "string"
    assert response.data["state"] == "NC"
    assert response.data["zip_code"] == "12345"


def test_post_location_with_contacts(
    api_client: APIClient,
    organization: Organization,
    business_unit: BusinessUnit,
) -> None:
    """Test post location with contacts

    Args:
        api_client (APIClient): Api client
        organization (Organization): Organization object
        business_unit (BusinessUnit): BusinessUnit Object

    Returns:
        None: This function does not return anything.
    """
    response = api_client.post(
        "/api/locations/",
        {
            "organization": organization.id,
            "business_unit": business_unit.id,
            "code": "string",
            "name": "string",
            "address_line_1": "string",
            "city": "string",
            "state": "NC",
            "zip_code": "12345",
            "location_contacts": [
                {
                    "name": "string",
                    "email": "test@test.com",
                    "phone": "(123) 123-4567",
                    "fax": "(123) 123-4567",
                }
            ],
        },
        format="json",
    )

    assert response.status_code == 201
    assert response.data["code"] == "string"
    assert response.data["name"] == "string"
    assert response.data["address_line_1"] == "string"
    assert response.data["city"] == "string"
    assert response.data["state"] == "NC"
    assert response.data["zip_code"] == "12345"
    assert response.data["location_contacts"][0]["name"] == "string"
    assert response.data["location_contacts"][0]["email"] == "test@test.com"
    assert response.data["location_contacts"][0]["phone"] == "(123) 123-4567"
    assert response.data["location_contacts"][0]["fax"] == "(123) 123-4567"


def test_post_location_with_comments(
    api_client: APIClient,
    organization: Organization,
    business_unit: BusinessUnit,
    comment_type: CommentType,
    user: User,
) -> None:
    """Test post location with comments

    Args:
        api_client (APIClient): Api client
        organization (Organization): Organization object
        business_unit (BusinessUnit): BusinessUnit Object
        comment_type (CommentType): CommentType object
        user (User): User object

    Returns:
        None: This function does not return anything.
    """
    response = api_client.post(
        "/api/locations/",
        {
            "organization": organization.id,
            "business_unit": business_unit.id,
            "code": "string",
            "name": "string",
            "address_line_1": "string",
            "city": "string",
            "state": "NC",
            "zip_code": "12345",
            "location_comments": [
                {
                    "comment_type": comment_type.id,
                    "comment": "this is a test comment for now.",
                    "entered_by": user.id,
                }
            ],
        },
        format="json",
    )

    assert response.status_code == 201
    assert response.data["code"] == "string"
    assert response.data["name"] == "string"
    assert response.data["address_line_1"] == "string"
    assert response.data["city"] == "string"
    assert response.data["state"] == "NC"
    assert response.data["zip_code"] == "12345"
    assert response.data["location_comments"][0]["comment_type"] == comment_type.id
    assert (
        response.data["location_comments"][0]["comment"]
        == "this is a test comment for now."
    )
    assert response.data["location_comments"][0]["entered_by"] == user.id


def test_post_location_with_comments_and_contacts(
    api_client: APIClient,
    organization: Organization,
    business_unit: BusinessUnit,
    comment_type: CommentType,
    user: User,
) -> None:
    """Test post location with comments & contacts

    Args:
        api_client (APIClient): Api client
        organization (Organization): Organization object
        business_unit (BusinessUnit): BusinessUnit Object
        comment_type (CommentType): CommentType object
        user (User): User object

    Returns:
        None: This function does not return anything.
    """
    response = api_client.post(
        "/api/locations/",
        {
            "organization": organization.id,
            "business_unit": business_unit.id,
            "code": "string",
            "name": "string",
            "address_line_1": "string",
            "city": "string",
            "state": "NC",
            "zip_code": "12345",
            "location_comments": [
                {
                    "comment_type": comment_type.id,
                    "comment": "this is a test comment for now.",
                    "entered_by": user.id,
                }
            ],
            "location_contacts": [
                {
                    "name": "string",
                    "email": "test@test.com",
                    "phone": "(123) 123-4567",
                    "fax": "(123) 123-4567",
                }
            ],
        },
        format="json",
    )

    assert response.status_code == 201
    assert response.data["code"] == "string"
    assert response.data["name"] == "string"
    assert response.data["address_line_1"] == "string"
    assert response.data["city"] == "string"
    assert response.data["state"] == "NC"
    assert response.data["zip_code"] == "12345"
    assert response.data["location_comments"][0]["comment_type"] == comment_type.id
    assert (
        response.data["location_comments"][0]["comment"]
        == "this is a test comment for now."
    )
    assert response.data["location_contacts"][0]["name"] == "string"
    assert response.data["location_contacts"][0]["email"] == "test@test.com"
    assert response.data["location_contacts"][0]["phone"] == "(123) 123-4567"
    assert response.data["location_contacts"][0]["fax"] == "(123) 123-4567"


def test_put_location_with_comments_and_contacts(
    api_client: APIClient,
    organization: Organization,
    business_unit: BusinessUnit,
    comment_type: CommentType,
    location: models.Location,
    user: User,
) -> None:
    """Test post location with comments & contacts

    Args:
        api_client (APIClient): Api client
        organization (Organization): Organization object
        business_unit (BusinessUnit): BusinessUnit Object
        comment_type (CommentType): CommentType object
        location (Location): Location object
        user (User): User object

    Returns:
        None: This function does not return anything.
    """

    location.refresh_from_db()

    response = api_client.put(
        f"/api/locations/{location.id}/",
        {
            "organization": organization.id,
            "business_unit": business_unit.id,
            "code": "LOCATION",
            "name": "test location",
            "address_line_1": "string",
            "city": "string",
            "state": "NC",
            "zip_code": "12345",
            "location_comments": [
                {
                    "comment_type": comment_type.id,
                    "comment": "this is a test comment for now for the location",
                    "entered_by": user.id,
                }
            ],
            "location_contacts": [
                {
                    "name": "test_contact",
                    "email": "test2@test.com",
                    "phone": "(123) 123-4567",
                    "fax": "(123) 123-4567",
                }
            ],
        },
        format="json",
    )

    assert response.status_code == 200
    assert response.data["code"] == "LOCATION"
    assert response.data["name"] == "test location"
    assert response.data["address_line_1"] == "string"
    assert response.data["city"] == "string"
    assert response.data["state"] == "NC"
    assert response.data["zip_code"] == "12345"
    assert response.data["location_comments"][0]["comment_type"] == comment_type.id
    assert (
        response.data["location_comments"][0]["comment"]
        == "this is a test comment for now for the location"
    )
    assert response.data["location_contacts"][0]["name"] == "test_contact"
    assert response.data["location_contacts"][0]["email"] == "test2@test.com"
    assert response.data["location_contacts"][0]["phone"] == "(123) 123-4567"
    assert response.data["location_contacts"][0]["fax"] == "(123) 123-4567"


def test_cannot_delete_location(
    api_client: APIClient, location: models.Location
) -> None:
    """Test that we cannot delete a location

    Args:
        api_client (APIClient): APIClient
        location (Location): Location object

    Returns:
        None: This function does not return anything.
    """
    response = api_client.delete(f"/api/locations/{location.id}/")

    assert response.status_code == 405
    assert response.data["errors"][0]["detail"] == 'Method "DELETE" not allowed.'


def test_location_contact_phone_number(
    location_contact: models.LocationContact,
) -> None:
    """Test Location contact phone number throws validation error when the
    phone number input is invalid

    Args:
        location_contact (models.LocationContact): LocationContact Object

    Returns:
        None: This function does not return anything.
    """

    location_contact.phone = "1234"

    with pytest.raises(ValidationError) as excinfo:
        location_contact.full_clean()

    assert (
        excinfo.value.message_dict["phone"][0]
        == "Phone number must be in the format (xxx) xxx-xxxx"
    )
