/*
 * COPYRIGHT(c) 2023 MONTA
 *
 * This file is part of Monta.
 *
 * The Monta software is licensed under the Business Source License 1.1. You are granted the right
 * to copy, modify, and redistribute the software, but only for non-production use or with a total
 * of less than three server instances. Starting from the Change Date (November 16, 2026), the
 * software will be made available under version 2 or later of the GNU General Public License.
 * If you use the software in violation of this license, your rights under the license will be
 * terminated automatically. The software is provided "as is," and the Licensor disclaims all
 * warranties and conditions. If you use this license's text or the "Business Source License" name
 * and trademark, you must comply with the Licensor's covenants, which include specifying the
 * Change License as the GPL Version 2.0 or a compatible license, specifying an Additional Use
 * Grant, and not modifying the license in any other way.
 */

import {
  Box,
  Button,
  createStyles,
  Group,
  rem,
  SimpleGrid,
} from "@mantine/core";
import { SelectInput } from "@/components/ui/fields/SelectInput";
import React from "react";
import { useForm, yupResolver } from "@mantine/form";
import { SwitchInput } from "@/components/ui/fields/SwitchInput";
import { useMutation, useQueryClient } from "react-query";
import axios from "@/lib/AxiosConfig";
import { notifications } from "@mantine/notifications";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faCheck, faXmark } from "@fortawesome/pro-solid-svg-icons";
import { APIError } from "@/types/server";
import { ValidatedNumberInput } from "@/components/ui/fields/NumberInput";
import { ValidatedTextInput } from "@/components/ui/fields/TextInput";
import {
  InvoiceControl,
  InvoiceControlFormValues,
} from "@/types/apps/invoicing";
import { ValidatedTextArea } from "@/components/ui/fields/TextArea";
import { dateFormatChoices } from "@/utils/apps/invoicing";
import { ValidatedFileInput } from "@/components/ui/fields/FileInput";
import { invoiceControlSchema } from "@/utils/apps/invoicing/schema";
import { useFormStyles } from "@/styles/FormStyles";

interface Props {
  invoiceControl: InvoiceControl;
}

export const InvoiceControlForm: React.FC<Props> = ({ invoiceControl }) => {
  const { classes } = useFormStyles();
  const [loading, setLoading] = React.useState<boolean>(false);
  const queryClient = useQueryClient();

  const mutation = useMutation(
    (values: InvoiceControlFormValues | FormData) =>
      axios.put(`/invoice_control/${invoiceControl.id}/`, values, {
        headers: {
          "Content-Type": "multipart/form-data",
        },
      }),
    {
      onSuccess: () => {
        queryClient
          .invalidateQueries({
            queryKey: ["invoiceControl"],
          })
          .then(() => {
            notifications.show({
              title: "Success",
              message: "Invoice Control updated successfully",
              color: "green",
              withCloseButton: true,
              icon: <FontAwesomeIcon icon={faCheck} />,
            });
          });
      },
      onError: (error: any) => {
        const { data } = error.response;
        if (data.type === "validation_error") {
          data.errors.forEach((error: APIError) => {
            form.setFieldError(error.attr, error.detail);
            if (error.attr === "non_field_errors") {
              notifications.show({
                title: "Error",
                message: error.detail,
                color: "red",
                withCloseButton: true,
                icon: <FontAwesomeIcon icon={faXmark} />,
                autoClose: 10_000, // 10 seconds
              });
            }
          });
        }
      },
      onSettled: () => {
        setLoading(false);
      },
    }
  );

  const form = useForm<InvoiceControlFormValues>({
    validate: yupResolver(invoiceControlSchema),
    initialValues: {
      invoice_number_prefix: invoiceControl.invoice_number_prefix,
      credit_memo_number_prefix: invoiceControl.credit_memo_number_prefix,
      invoice_due_after_days: invoiceControl.invoice_due_after_days,
      invoice_terms: invoiceControl.invoice_terms || "",
      invoice_footer: invoiceControl.invoice_footer || "",
      invoice_logo: invoiceControl.invoice_logo || "",
      invoice_logo_width: invoiceControl.invoice_logo_width,
      show_invoice_due_date: invoiceControl.show_invoice_due_date,
      invoice_date_format: invoiceControl.invoice_date_format,
      show_amount_due: invoiceControl.show_amount_due,
      attach_pdf: invoiceControl.attach_pdf,
    },
  });

  const handleSubmit = (values: InvoiceControlFormValues) => {
    setLoading(true);
    const formData = new FormData();
    for (const key in values) {
      if (Object.prototype.hasOwnProperty.call(values, key)) {
        let element = values[key as keyof InvoiceControlFormValues];
        if (element instanceof File || typeof element === "string") {
          formData.append(key, element);
        } else if (
          typeof element === "boolean" ||
          typeof element === "number"
        ) {
          formData.append(key, element.toString());
        }
      }
    }
    mutation.mutate(formData);
  };

  return (
    <form onSubmit={form.onSubmit((values) => handleSubmit(values))}>
      <Box className={classes.div}>
        <Box>
          <SimpleGrid cols={3} breakpoints={[{ maxWidth: "sm", cols: 1 }]}>
            <ValidatedTextInput
              form={form}
              className={classes.fields}
              name="invoice_number_prefix"
              label="Invoice Number Prefix"
              placeholder="Invoice Number Prefix"
              variant="filled"
              description="Define a prefix for invoice numbers."
              withAsterisk
            />
            <ValidatedTextInput
              form={form}
              className={classes.fields}
              name="credit_memo_number_prefix"
              label="Credit Memo Number Prefix"
              placeholder="Credit Memo Number Prefix"
              description="Define a prefix for credit note numbers."
              variant="filled"
              withAsterisk
            />
            <ValidatedNumberInput
              form={form}
              className={classes.fields}
              name="invoice_due_after_days"
              label="Invoice Due After Days"
              placeholder="Invoice Due After Days"
              description="Define the number of days after which an invoice is due."
              variant="filled"
              withAsterisk
            />
          </SimpleGrid>
          <ValidatedTextArea
            form={form}
            className={classes.fields}
            name="invoice_terms"
            label="Invoice Terms"
            placeholder="Invoice Terms"
            description="Define the terms and conditions for invoices."
            variant="filled"
          />
          <ValidatedTextArea
            form={form}
            className={classes.fields}
            name="invoice_footer"
            label="Invoice Footer"
            description="Define the footer for invoices."
            placeholder="Invoice Footer"
            variant="filled"
          />
          <SimpleGrid cols={3} breakpoints={[{ maxWidth: "sm", cols: 1 }]}>
            <SelectInput
              form={form}
              data={dateFormatChoices}
              className={classes.fields}
              name="invoice_date_format"
              label="Invoice Date Format"
              placeholder="Invoice Date Format"
              description="Define the date format for invoices."
              variant="filled"
              withAsterisk
            />
            <ValidatedNumberInput
              form={form}
              className={classes.fields}
              name="invoice_logo_width"
              label="Invoice Logo Width"
              placeholder="Invoice Logo Width"
              description="Define the width of the invoice logo."
              variant="filled"
              withAsterisk
            />
            <ValidatedFileInput
              form={form}
              className={classes.fields}
              name="invoice_logo"
              label="Invoice Logo"
              placeholder="Invoice Logo"
              description="Define the logo for invoices."
              variant="filled"
              value={form.values.invoice_logo}
              accept="image/png,image/jpeg"
            />
            <SwitchInput
              form={form}
              className={classes.fields}
              name="show_invoice_due_date"
              label="Show Invoice Due Date"
              description="Show the invoice due date on the invoice."
            />
            <SwitchInput
              form={form}
              className={classes.fields}
              name="show_amount_due"
              label="Show Amount Due"
              description="Show the amount due on the invoice."
            />
            <SwitchInput
              form={form}
              className={classes.fields}
              name="show_amount_due"
              label="Show Amount Due"
              description="Attach the invoice PDF to the invoice email."
            />
          </SimpleGrid>
          <Group position="right" mt="md">
            <Button
              color="white"
              type="submit"
              className={classes.control}
              loading={loading}
            >
              Submit
            </Button>
          </Group>
        </Box>
      </Box>
    </form>
  );
};