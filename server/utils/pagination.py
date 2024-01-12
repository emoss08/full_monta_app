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
from rest_framework.pagination import LimitOffsetPagination
from rest_framework.response import Response


class TrenovaPagination(LimitOffsetPagination):
    default_limit = 10
    max_limit = None

    def paginate_queryset(self, queryset, request, view=None):
        self.limit = self.get_limit(request)
        if self.limit is None:
            # When limit is 'all', count without pagination
            self.count = queryset.count()
            return list(queryset)
        else:
            # Standard limit-offset pagination
            self.offset = self.get_offset(request)
            self.count = self.get_count(queryset)
            return super().paginate_queryset(queryset, request, view=view)

    def get_paginated_response(self, data):
        if self.limit is None:
            # Return total count without next/previous links for 'all'
            return Response(
                {"count": self.count, "next": None, "previous": None, "results": data}
            )
        else:
            # Standard paginated response
            return super().get_paginated_response(data)

    def get_limit(self, request):
        limit_param = request.query_params.get(self.limit_query_param)
        if limit_param == "all":
            return None
        try:
            # Ensuring limit is a positive integer
            limit = int(limit_param)
            if limit < 1:
                raise ValueError
            return limit
        except (TypeError, ValueError):
            # Default to the standard implementation if not valid
            return super().get_limit(request)
