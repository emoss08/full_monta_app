import { Notifications } from "@/components/layout/notification_menu/notification";
import { Button } from "@/components/ui/button";
import { ComponentLoader } from "@/components/ui/component-loader";
import { InternalLink } from "@/components/ui/link";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { ScrollArea } from "@/components/ui/scroll-area";
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip";
import { useNotifications } from "@/hooks/useQueries";
import axios from "@/lib/axiosConfig";
import { WEB_SOCKET_URL } from "@/lib/constants";
import { createWebsocketManager } from "@/lib/websockets";
import { useUserStore } from "@/stores/AuthStore";
import { useHeaderStore } from "@/stores/HeaderStore";
import { UserNotification } from "@/types/accounts";
import { BellIcon } from "@radix-ui/react-icons";
import { useQueryClient } from "@tanstack/react-query";
import React, { useState } from "react";
import { toast } from "sonner";

function NotificationButton({
  userHasNotifications,
  open,
}: {
  userHasNotifications: boolean;
  open: boolean;
}) {
  return (
    <TooltipProvider delayDuration={100}>
      <Tooltip>
        <TooltipTrigger asChild>
          <Button
            size="icon"
            variant="outline"
            aria-expanded={open}
            className="border-muted-foreground/40 hover:border-muted-foreground/80 group relative size-8"
          >
            <BellIcon className="text-muted-foreground group-hover:text-foreground size-5" />
            <span className="sr-only">Notifications</span>
            {userHasNotifications && (
              <span className="absolute -right-1 -top-1 flex size-2.5">
                <span className="absolute inline-flex size-full animate-ping rounded-full bg-green-400 opacity-100"></span>
                <span className="ring-background relative inline-flex size-2.5 rounded-full bg-green-600 ring-1"></span>
              </span>
            )}
          </Button>
        </TooltipTrigger>
        <TooltipContent side="bottom" sideOffset={5}>
          <span>Notifications</span>
        </TooltipContent>
      </Tooltip>
    </TooltipProvider>
  );
}

function NotificationContent({
  notificationsData,
  notificationsLoading,
  userHasNotifications,
  readAllNotifications,
}: {
  notificationsData: UserNotification | undefined;
  notificationsLoading: boolean;
  userHasNotifications: boolean;
  readAllNotifications: () => void;
}) {
  if (notificationsLoading) {
    return <ComponentLoader className="h-80" />;
  }

  return (
    <>
      <ScrollArea className="h-80 w-full">
        <Notifications
          notification={notificationsData as UserNotification}
          notificationLoading={notificationsLoading}
        />
      </ScrollArea>
      {!userHasNotifications && (
        <div className="select-none items-center justify-center border-t pt-2 text-center text-xs">
          Know when you have new notifications by enabling text notifications in
          your{" "}
          <InternalLink to="/account/settings/">Account Settings</InternalLink>
        </div>
      )}
      {userHasNotifications && (
        <div className="flex items-center justify-center border-t pt-2 text-center">
          <Button
            onClick={readAllNotifications}
            variant="link"
            className="w-full"
          >
            Mark all as read
          </Button>
        </div>
      )}
    </>
  );
}

export function NotificationMenu() {
  const queryClient = useQueryClient();
  const [notificationsMenuOpen, setNotificationMenuOpen] = useHeaderStore.use(
    "notificationMenuOpen",
  );
  const [userHasNotifications, setUserHasNotifications] =
    useState<boolean>(false);
  const { id: userId } = useUserStore.get("user");
  const { notificationsData, notificationsLoading } = useNotifications(userId);
  const webSocketManager = createWebsocketManager();

  const markedAndInvalidate = async () => {
    await axios.get("/user-notifications/?markAsRead=true");
    await queryClient.invalidateQueries({
      queryKey: ["userNotifications", userId],
    });
  };

  const readAllNotifications = async () => {
    const sendNotificationRequest = markedAndInvalidate();

    // Fire Toast
    toast.promise(sendNotificationRequest, {
      loading: "Marking all notifications as read",
      success: "All notifications marked as read",
      error: "Failed to mark all notifications as read",
    });

    setNotificationMenuOpen(false);
  };

  React.useEffect(() => {
    webSocketManager.connect("notifications", `${WEB_SOCKET_URL}/${userId}`, {
      onOpen: () => console.info("Connected to notifications websocket"),
      onMessage: (event: MessageEvent) => {
        const data = JSON.parse(event.data);
        queryClient
          .invalidateQueries({
            queryKey: ["userNotifications", userId],
          })
          .then(() => {
            toast.success(
              <div className="flex flex-col space-y-1">
                <span className="font-semibold">New Report Available!</span>
                <span className="text-xs">{data.content}</span>
              </div>,
            );
          });
      },
      onClose: (event: CloseEvent) => console.info(event),
    });

    if (!userId) {
      return;
    }

    return () => {
      webSocketManager.disconnect("notifications");
    };
  }, [userId]);

  React.useEffect(() => {
    if (
      notificationsData &&
      (notificationsData as UserNotification).unreadList
    ) {
      setUserHasNotifications(
        (notificationsData as UserNotification).unreadList.length > 0,
      );
    }
  }, [notificationsData]);

  return (
    <Popover
      open={notificationsMenuOpen}
      onOpenChange={setNotificationMenuOpen}
    >
      <PopoverTrigger asChild>
        <span>
          <NotificationButton
            userHasNotifications={userHasNotifications}
            open={notificationsMenuOpen}
          />
        </span>
      </PopoverTrigger>
      <PopoverContent
        className="bg-popover w-80 p-3"
        sideOffset={10}
        alignOffset={-40}
        align="end"
      >
        <NotificationContent
          notificationsData={notificationsData as UserNotification}
          notificationsLoading={notificationsLoading}
          userHasNotifications={userHasNotifications}
          readAllNotifications={readAllNotifications}
        />
      </PopoverContent>
    </Popover>
  );
}
