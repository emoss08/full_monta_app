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

import factory


class GeneralLedgerAccountFactory(factory.django.DjangoModelFactory):
    """
    General Ledger Account factory
    """

    class Meta:
        """
        Metaclass for GeneralLedger Account Factory
        """

        model = "accounting.GeneralLedgerAccount"

    business_unit = factory.SubFactory("organization.factories.BusinessUnitFactory")
    organization = factory.SubFactory("organization.factories.OrganizationFactory")
    account_number = factory.Faker(
        "numerify",
        text="####-##",
    )
    account_type = factory.Faker(
        "random_element",
        elements=("ASSET", "LIABILITY", "EQUITY", "REVENUE", "EXPENSE"),
    )


class RevenueCodeFactory(factory.django.DjangoModelFactory):
    """
    Revenue Code factory
    """

    class Meta:
        """
        Metaclass for RevenueCode Factory
        """

        model = "accounting.RevenueCode"

    business_unit = factory.SubFactory("organization.factories.BusinessUnitFactory")
    organization = factory.SubFactory("organization.factories.OrganizationFactory")
    code = factory.Faker("pystr", max_chars=4)
    description = factory.Faker("text", locale="en_US", max_nb_chars=100)


class DivisionCodeFactory(factory.django.DjangoModelFactory):
    """
    Division Code Factory
    """

    class Meta:
        """
        Metaclass for DivisionCodeFactory
        """

        model = "accounting.DivisionCode"

    business_unit = factory.SubFactory("organization.factories.BusinessUnitFactory")
    organization = factory.SubFactory("organization.factories.OrganizationFactory")
    status = "A"
    code = factory.Faker("pystr", max_chars=4)
    description = factory.Faker("text", locale="en_US", max_nb_chars=100)
