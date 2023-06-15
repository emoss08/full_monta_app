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

# Allowed models and fields for report generation
ALLOWED_MODELS = {
    "User": {
        "app_label": "accounts",
        "allowed_fields": [
            "username",
            "email",
            "date_joined",
            "is_staff",
            "profiles__first_name",
            "profiles__last_name",
            "profiles__address_line_1",
            "profiles__address_line_2",
            "profiles__city",
            "profiles__state",
            "profiles__zip_code",
            "profiles__phone_number",
            "profiles__is_phone_verified",
            "profiles__job_title__name",
            "profiles__job_title__description",
            "department__name",
            "department__description",
            "organization__name",
        ],
    },
    "UserProfile": {
        "app_label": "accounts",
        "allowed_fields": [
            "user",
            "job_title__name",
            "job_title__description",
            "first_name",
            "last_name",
            "address_line_1",
            "address_line_2",
            "city",
            "state",
            "zip_code",
            "phone_number",
            "is_phone_verified",
        ],
    },
    "JobTitle": {
        "app_label": "accounts",
        "allowed_fields": [
            "is_active",
            "name",
            "description",
            "job_function",
        ],
    },
    "Organization": {
        "app_label": "organization",
        "allowed_fields": [
            "name",
            "scac_code",
            "dot_number",
            "address_line_1",
            "address_line_2",
            "city",
            "state",
            "zip_code",
            "phone_number",
            "website",
            "org_type",
            "timezone",
            "language",
            "currency",
            "date_format",
            "time_format",
            "token_expiration_days",
        ],
    },
    "Department": {
        "app_label": "organization",
        "allowed_fields": [
            "organization__name",
            "name",
            "description",
        ],
    },
}
