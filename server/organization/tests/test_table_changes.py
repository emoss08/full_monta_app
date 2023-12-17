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

from io import StringIO
from unittest.mock import patch

import pytest
from django.core.exceptions import ValidationError
from django.core.management import call_command

from kafka.managers import KafkaManager
from organization import factories, models
from organization.models import TableChangeAlert
from organization.services.psql_triggers import (
    check_function_exists,
    check_trigger_exists,
)
from organization.services.table_choices import TABLE_NAME_CHOICES

pytestmark = pytest.mark.django_db


def test_create_table_charge_alert(organization: models.Organization) -> None:
    """Tests the creation a table charge alert.

    Returns:
        None: This function does not return anything.
    """
    table_charge = models.TableChangeAlert.objects.create(
        business_unit=organization.business_unit,
        organization=organization,
        is_active=True,
        name="Test",
        database_action="INSERT",
        table=TABLE_NAME_CHOICES[0][0],
    )

    assert table_charge.organization == organization
    assert table_charge.is_active is True
    assert table_charge.name == "Test"
    assert (
        table_charge.database_action
        == models.TableChangeAlert.DatabaseActionChoices.INSERT
    )
    assert table_charge.table == TABLE_NAME_CHOICES[0][0]


def test_table_change_insert_database_action_save() -> None:
    """Tests the creation of a table change alert with INSERT Action adds the proper function,
    trigger, and listener name.

    Returns:
        None: This function does not return anything.
    """
    table_change = factories.TableChangeAlertFactory(database_action="INSERT")

    assert table_change.function_name == f"notify_new_{table_change.table}"
    assert table_change.trigger_name == f"after_insert_{table_change.table}"
    assert table_change.listener_name == f"new_added_{table_change.table}"


def test_table_change_insert_adds_insert_trigger() -> None:
    """Tests that the insert trigger is added to the database when a table change alert is created
    with INSERT action.

    Returns:
        None: This function does not return anything.
    """
    table_change = factories.TableChangeAlertFactory(database_action="INSERT")

    trigger_check = check_trigger_exists(
        table_name=table_change.table, trigger_name=table_change.trigger_name
    )
    function_check = check_function_exists(function_name=table_change.function_name)

    assert trigger_check is True
    assert function_check is True


def test_delete_table_change_removes_trigger() -> None:
    """
    Tests that the trigger is removed from the database when a table change alert is deleted.

    Returns:
        None: This function does not return anything.
    """
    table_change = factories.TableChangeAlertFactory(database_action="INSERT")

    trigger_check = check_trigger_exists(
        table_name=table_change.table, trigger_name=table_change.trigger_name
    )
    function_check = check_function_exists(function_name=table_change.function_name)
    assert trigger_check is True
    assert function_check is True

    table_change.delete()

    trigger_check_2 = check_trigger_exists(
        table_name=table_change.table, trigger_name=table_change.trigger_name
    )
    function_check_2 = check_function_exists(function_name=table_change.function_name)

    assert trigger_check_2 is False
    assert function_check_2 is False


def test_table_change_database_action_update() -> None:
    """Test changing the database action removes and adds the proper function, trigger, and listener
    names.

    Returns:
        None: This function does not return anything.
    """
    table_change = factories.TableChangeAlertFactory(database_action="INSERT")

    assert table_change.function_name == f"notify_new_{table_change.table}"
    assert table_change.trigger_name == f"after_insert_{table_change.table}"
    assert table_change.listener_name == f"new_added_{table_change.table}"

    table_change.database_action = "UPDATE"
    table_change.save()

    assert table_change.function_name == f"notify_updated_{table_change.table}"
    assert table_change.trigger_name == f"after_update_{table_change.table}"
    assert table_change.listener_name == f"updated_{table_change.table}"


def test_command() -> None:
    """Tests that the psql_listener command runs successfully.

    Returns:
        None: This function does not return anything.
    """
    with patch("psycopg.connect"), patch(
        "django.core.management.color.supports_color", return_value=False
    ):
        out = StringIO()
        call_command("psql_listener", stdout=out)
        assert "Starting PostgreSQL listener..." in out.getvalue()


def test_save_table_change_alert_kafka_without_topic(
    organization: models.Organization,
) -> None:
    """Tests that a ValidationError is raised when trying to save a TableChangeAlert with source as
    ``Kafka`` but no topic.

    Returns:
        None: This function does not return anything.
    """
    # Create a TableChangeAlert instance with source as Kafka but no topic
    kafka_alert = TableChangeAlert(source=TableChangeAlert.SourceChoices.KAFKA)

    # Expect a ValidationError when trying to save
    with pytest.raises(ValidationError) as excinfo:
        kafka_alert.clean()

    # Check if the error message is correct
    assert excinfo.value.message_dict["topic"] == [
        "Topic is required when source is Kafka."
    ]


def test_save_table_change_alert_postgres_without_table(
    organization: models.Organization,
) -> None:
    """Tests that a ValidationError is raised when trying to save a TableChangeAlert with source as
    Postgres but no table.

    Args:
        organization (models.Organization); Organization instance.

    Returns:
        None: This function does not return anything.
    """
    # Create a TableChangeAlert instance with source as Postgres but no table
    alert = TableChangeAlert(
        organization=organization,
        source=TableChangeAlert.SourceChoices.POSTGRES,
    )

    # Expect a ValidationError when trying to save
    with pytest.raises(ValidationError) as excinfo:
        alert.clean()

    # Check if the error message is correct
    assert excinfo.value.message_dict["table"] == [
        "Table is required when source is Postgres."
    ]


def test_cannot_save_if_kafka_offline(organization: models.Organization) -> None:
    """Test validationError is thrown if source is ``KAFKA`` and Kafka is offline.

    Args:
        organization (models.Organization): Organization instance.

    Returns:
        None: This function does not return anything.

    Notes:
        This test requires Kafka to be offline. If Kafka is online, this test will fail.
    """

    manager = KafkaManager()

    if manager.is_kafka_available():
        pytest.skip("Kafka is online. Skipping test.")

    alert = TableChangeAlert(
        organization=organization,
        source=TableChangeAlert.SourceChoices.KAFKA,
        topic="test",
    )

    with pytest.raises(ValidationError) as excinfo:
        alert.clean()

    assert excinfo.value.message_dict["source"] == [
        f"Unable to connect to Kafka at {manager.kafka_host}:{manager.kafka_port}."
        " Please check your connection and try again."
    ]


def test_cannot_save_delete_if_source_not_kafka(
    organization: models.Organization,
) -> None:
    """Test ValidationError is thrown if ``database_action`` is ``delete`` and source is not ``KAFKA``.

    Args:
        organization (models.Organization): Organization instance.

    Returns:
        None: This function does not return anything.
    """

    alert = TableChangeAlert(
        organization=organization,
        source=TableChangeAlert.SourceChoices.POSTGRES,
        table=TABLE_NAME_CHOICES[0][0],
        database_action=TableChangeAlert.DatabaseActionChoices.DELETE,
    )
    with pytest.raises(ValidationError) as excinfo:
        alert.clean()

    assert excinfo.value.message_dict["database_action"] == [
        "Database action can only be delete when source is Kafka."
        " Please change the source to Kafka and try again."
    ]
