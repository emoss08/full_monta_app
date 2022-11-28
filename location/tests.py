# Create your tests here.
"""
COPYRIGHT 2022 MONTA

This file is part of Monta.

Monta is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

Monta is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with Monta.  If not, see <https://www.gnu.org/licenses/>.
"""

import pytest

from location.factories import (
    LocationCategoryFactory,
    LocationContactFactory,
    LocationFactory,
)


@pytest.fixture()
def location_category():
    """
    Location category fixture
    """
    return LocationCategoryFactory()


@pytest.fixture()
def location():
    """
    Location fixture
    """
    return LocationFactory()


@pytest.fixture()
def location_contact():
    """
    Location contact fixture
    """
    return LocationContactFactory()


@pytest.mark.django_db
def test_location_category_creation(location_category):
    """
    Test location category creation
    """
    assert location_category is not None


@pytest.mark.django_db
def test_location_category_update(location_category):
    """
    Test location category update
    """
    location_category.name = "New name"
    location_category.save()
    assert location_category.name == "New name"


@pytest.mark.django_db
def test_location_creation(location):
    """
    Test location creation
    """
    assert location is not None


@pytest.mark.django_db
def test_location_update(location):
    """
    Test location update
    """
    location.name = "New name"
    location.save()
    assert location.name == "New name"


@pytest.mark.django_db
def test_location_category_assigned(location):
    """
    Test location category assigned
    """
    assert location.location_category is not None


@pytest.mark.django_db
def test_location_contact_creation(location):
    """
    Test location contact creation
    """
    assert location.location_contact is not None
