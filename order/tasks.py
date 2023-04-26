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
from typing import List

from django.db.models import QuerySet

from backend.celery import app

from accounts.models import User
from billing.services import mass_order_billing, single_order_billing
from billing.services.transfer_to_billing import transfer_to_billing_queue_service
from core.exceptions import ServiceException
from order import selectors, services
from organization.models import Organization
from utils.types import ModelUUID


@app.task(
    name="consolidate_order_documentation",
    bind=True,
    max_retries=3,
    default_retry_delay=60,
)
def consolidate_order_documentation(self, *, order_id: ModelUUID) -> None:  # type: ignore
    """Consolidate Order

    Query the database for the Order and call the consolidate_pdf
    service to combine the PDFs into a single PDF.

    Args:
        self (celery.app.task.Task): The task object
        order_id (str): Order ID

    Returns:
        None: None

    Raises:
        ObjectDoesNotExist: If the Order does not exist in the database.
    """

    try:
        if order := selectors.get_order_by_id(order_id=order_id):
            services.combine_pdfs_service(order=order)
        else:
            return None

    except ServiceException as exc:
        raise self.retry(exc=exc) from exc


@app.task(name="bill_order_task", bind=True, max_retries=3, default_retry_delay=60)
def bill_order_task(self, user_id: ModelUUID, order_id: ModelUUID) -> None:    # type: ignore
    """Bill Order

    Query the database for the Order and call the bill_order
    service to bill the order.

    Args:
        self (celery.app.task.Task): The task object
        user_id (str): User ID
        order_id (str): Order ID

    Returns:
        None: None

    Raises:
        ObjectDoesNotExist: If the Order does not exist in the database.
    """

    try:
        if order := selectors.get_order_by_id(order_id=order_id):
            single_order_billing.bill_order(order=order, user_id=user_id)
        else:
            return None
    except ServiceException as exc:
        raise self.retry(exc=exc) from exc


@app.task(name="mass_order_bill_task", bind=True, max_retries=3, default_retry_delay=60)
def mass_order_bill_task(self, *, user_id: ModelUUID) -> None:  # type: ignore
    """Bill Order

    Args:
        self (celery.app.task.Task): The task object
        user_id (str): User ID

    Returns:
        None: None

    Raises:
        ObjectDoesNotExist: If the Order does not exist in the database.
    """
    try:
        mass_order_billing.mass_order_billing_service(
            user_id=user_id, task_id=self.request.id
        )
    except ServiceException as exc:
        raise self.retry(exc=exc) from exc


@app.task(name="transfer_to_billing_task", bind=True)
def transfer_to_billing_task(  # type: ignore
    self, *, user_id: ModelUUID, order_pros: List[str]
) -> None:
    """
    Starts a Celery task to transfer the specified order(s) to billing for the logged in user.

    Args:
        self: The Celery task instance.
        user_id: A string representing the ID of the user who initiated the transfer.
        order_pros: A list of strings representing the order IDs to transfer.

    Returns:
        None.

    Raises:
        self.retry: If an ObjectDoesNotExist exception is raised while processing the task.

    This Celery task function calls the `transfer_to_billing_queue_service` function to create BillingQueue objects
    for each order in the provided list of order IDs, and updates the transfer status and transfer date of each order.

    If an ObjectDoesNotExist exception is raised while processing the task, the Celery task will automatically retry
    the task until it succeeds, with an exponential backoff strategy.

    The function expects the following arguments:
    - self: The Celery task instance.
    - user_id: A string representing the ID of the user who initiated the transfer.
    - order_pros: A list of strings representing the order IDs to transfer.

    The `transfer_to_billing_queue_service` function is called to perform the actual transfer of the specified order(s).
    If this operation raises an ObjectDoesNotExist exception, the function will retry the task with an exponential
    backoff strategy until it succeeds.

    Finally, the function returns None.
    """

    try:
        transfer_to_billing_queue_service(
            user_id=user_id, order_pros=order_pros, task_id=self.request.id
        )
    except ServiceException as exc:
        raise self.retry(exc=exc) from exc


@app.task(
    name="automate_mass_order_billing", bind=True, max_retries=3, default_retry_delay=60
)
def automate_mass_order_billing(self) -> str:  # type: ignore
    """Automated Mass Billing Tasks that uses system user to bill orders.

    Filter the database for the Organizations that have auto bill orders enabled and call the mass_order_billing_service
    service to bill the orders.

    Args:
        self (celery.app.task.Task): The task object

    Returns:
        str: A string containing the results of the automated mass billing tasks.

    Raises:
        ServiceException: If the Order does not exist in the database.
    """
    system_user = User.objects.get(username="sys")
    organizations: QuerySet[Organization] = Organization.objects.filter(
        billing_control__auto_bill_orders=True
    )
    results = []
    for organization in organizations:
        try:
            mass_order_billing.mass_order_billing_service(
                user_id=system_user.id, task_id=self.request.id
            )
            results.append(
                f"Automated Mass Billing Task for {organization.name} was successful."
            )
        except ServiceException as exc:
            raise self.retry(exc=exc) from exc
    if results:
        return "\n".join(results)
    return "No organizations have auto billing enabled."
