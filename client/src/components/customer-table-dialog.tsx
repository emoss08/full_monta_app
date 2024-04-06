import { Button } from "@/components/ui/button";
import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetFooter,
  SheetHeader,
  SheetTitle,
} from "@/components/ui/sheet";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { useCustomMutation } from "@/hooks/useCustomMutation";
import { cn } from "@/lib/utils";
import { customerSchema } from "@/lib/validations/CustomerSchema";
import { useCustomerFormStore } from "@/stores/CustomerStore";
import { type CustomerFormValues as FormValues } from "@/types/customer";
import { type TableSheetProps } from "@/types/tables";
import { yupResolver } from "@hookform/resolvers/yup";
import { useState } from "react";
import { FormProvider, useForm } from "react-hook-form";
import { CustomerContactForm } from "./customer-contacts-form";
import { CustomerEmailProfileForm } from "./customer-email-profile-form";
import { CustomerInfoForm } from "./customer-info-form";
import { CustomerRuleProfileForm } from "./customer-rule-profile-form";
import { DeliverySlotForm } from "./delivery-slots-form";

export function CustomerForm({ open }: { open: boolean }) {
  const [activeTab, setActiveTab] = useCustomerFormStore.use("activeTab");

  return (
    <Tabs
      defaultValue="info"
      value={activeTab}
      className="w-full flex-1"
      onValueChange={setActiveTab}
    >
      <TabsList>
        <TabsTrigger value="info">Information</TabsTrigger>
        <TabsTrigger value="emailProfile">Email Profile</TabsTrigger>
        <TabsTrigger value="ruleProfile">Rule Profile</TabsTrigger>
        <TabsTrigger value="deliverySlots">Delivery Slots</TabsTrigger>
        <TabsTrigger value="contacts">Contacts</TabsTrigger>
        <TabsTrigger value="detentionPolicy">Detention Policy</TabsTrigger>
      </TabsList>
      <TabsContent value="info">
        <CustomerInfoForm open={open} />
      </TabsContent>
      <TabsContent value="emailProfile">
        <CustomerEmailProfileForm />
      </TabsContent>
      <TabsContent value="ruleProfile">
        <CustomerRuleProfileForm open={open} />
      </TabsContent>
      <TabsContent value="deliverySlots">
        <DeliverySlotForm open={open} />
      </TabsContent>
      <TabsContent value="contacts">
        <CustomerContactForm />
      </TabsContent>
      <TabsContent value="detentionPolicy">
        <p>Coming Soon...</p>
      </TabsContent>
    </Tabs>
  );
}

export function CustomerTableSheet({ onOpenChange, open }: TableSheetProps) {
  const [isSubmitting, setIsSubmitting] = useState<boolean>(false);

  const customerForm = useForm<FormValues>({
    resolver: yupResolver(customerSchema),
    defaultValues: {
      status: "A",
      code: "",
      name: "",
      addressLine1: "",
      addressLine2: "",
      city: "",
      state: "",
      zipCode: "",
      hasCustomerPortal: false,
      autoMarkReadyToBill: false,
      deliverySlots: [],
      contacts: [],
      emailProfile: {
        subject: "",
        comment: "",
        fromAddress: "",
        blindCopy: "",
        readReceipt: false,
        readReceiptTo: "",
        attachmentName: "",
      },
      ruleProfile: {
        name: "",
        documentClass: [],
      },
    },
  });

  const { control, handleSubmit, reset } = customerForm;

  const mutation = useCustomMutation<FormValues>(
    control,
    {
      method: "POST",
      path: "/customers/",
      successMessage: "Customer created successfully.",
      queryKeysToInvalidate: ["customers-table-data"],
      closeModal: true,
      errorMessage: "Failed to create new customer.",
    },
    () => setIsSubmitting(false),
    reset,
  );

  function onSubmit(values: FormValues) {
    setIsSubmitting(true);
    mutation.mutate(values);
  }

  return (
    <Sheet open={open} onOpenChange={onOpenChange}>
      <SheetContent className={cn("w-full xl:w-1/2")}>
        <SheetHeader>
          <SheetTitle>Add New Customer</SheetTitle>
          <SheetDescription>
            Use this form to add a new customer to the system.
          </SheetDescription>
        </SheetHeader>
        <FormProvider {...customerForm}>
          <form
            onSubmit={handleSubmit(onSubmit)}
            className="flex h-full flex-col overflow-y-auto"
          >
            <CustomerForm open={open} />
            <SheetFooter className="mb-12">
              <Button
                type="reset"
                variant="secondary"
                onClick={() => onOpenChange(false)}
                className="w-full"
              >
                Cancel
              </Button>
              <Button type="submit" isLoading={isSubmitting} className="w-full">
                Save
              </Button>
            </SheetFooter>
          </form>
        </FormProvider>
      </SheetContent>
    </Sheet>
  );
}
